package keeper

import (
	"dredd-secure/x/escrow/constants"
	"dredd-secure/x/escrow/types"
	"fmt"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"

	"strconv"
	"strings"

	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
)

type (
	Keeper struct {
		bank          types.BankKeeper
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
		channelKeeper types.ChannelKeeper
		portKeeper    types.PortKeeper
		scopedKeeper  exported.ScopedKeeper
		govKeeper	  *govkeeper.Keeper
	}
)

func NewKeeper(
	bank types.BankKeeper,
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper exported.ScopedKeeper,
	govKeeper *govkeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		bank:       bank,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
		govKeeper: 	   govKeeper,
	}
}

// ----------------------------------------------------------------------------
// IBC Keeper Logic
// ----------------------------------------------------------------------------

// ChanCloseInit defines a wrapper function for the channel Keeper's function.
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IsBound checks if the IBC app module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the IBC app module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the IBC app module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the IBC app module to claim a capability that core IBC
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// StoreOracleResponsePacket is a function that receives an OracleResponsePacketData from BandChain.
func (k Keeper) StoreOracleResponsePacket(ctx sdk.Context, res types.OracleResponsePacketDataPacketData) error {

	// Find the oracleId from the clientID
	oracleId := strings.Split(res.ClientId, "_")[0]

	switch oracleId {
		case constants.OracleCryptoCurrencyPriceScriptId:
			// Decode the result from the response packet.
			result, err := bandtypes.DecodeResult(res.Result)
			if err != nil {
				return err
			}

			for _, r := range result {
				oraclePrice := types.OraclePrice{
					Symbol:      r.Symbol,
					ResolveTime: res.RequestTime,
					Price:       strconv.FormatUint(r.Rate, 10),
				}
				changed := k.UpdateOraclePrice(ctx, oraclePrice)
				if changed {
					ctx.EventManager().EmitEvent(
						sdk.NewEvent(
							types.EventTypePriceUpdate,
							sdk.NewAttribute(types.AttributeKeySymbol, r.Symbol),
							sdk.NewAttribute(types.AttributeKeyPrice, fmt.Sprintf("%d", r.Rate)),
							sdk.NewAttribute(types.AttributeKeyTimestamp, res.ResolveStatus),
						),
					)
				}
			}
		default: 
			return types.ErrOracleScriptNotConfigured
	}

	return nil
}

// RequestBandChainData is a function that sends an OracleRequestPacketData to BandChain via IBC.
func (k Keeper) RequestBandChainData(
	ctx sdk.Context,
	sourceChannel string,
	oracleRequestPacket bandtypes.OracleRequestPacketData,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {
	portID := k.GetPort(ctx)
	channel, found := k.channelKeeper.GetChannel(ctx, portID, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrChannelNotFound,
			"port ID (%s) channel ID (%s)",
			portID,
			sourceChannel,
		)
	}

	destinationPort := channel.GetCounterparty().GetPortID()
	destinationChannel := channel.GetCounterparty().GetChannelID()

	// Get the capability associated with the given channel.
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// Get the next sequence number for the given channel and port.
	sequence, found := k.channelKeeper.GetNextSequenceSend(
		ctx, portID, sourceChannel,
	)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port %s",
			sourceChannel,
			portID,
		)
	}

	// Create a new packet with the oracle request packet data and the sequence number.
	packet := channeltypes.NewPacket(
		oracleRequestPacket.GetBytes(),
		sequence,
		portID,
		sourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.ZeroHeight(),
		uint64(ctx.BlockTime().UnixNano()+int64(20*time.Minute)),
	)

	// Send the packet via the channel and capability associated with the given channel.
	_, err := k.channelKeeper.SendPacket(ctx, channelCap, portID, sourceChannel, timeoutHeight, timeoutTimestamp, packet.Data)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetSrcChannel (ctx sdk.Context) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SourceChannelKey))
	b := store.Get(types.KeyPrefix(types.SourceChannelKey))
	if b == nil {
		return types.NotSet
	}

	srcChannel := string(b)

	if len(srcChannel) == 0 {
		return types.NotSet
	}

	return srcChannel
}

func (k Keeper) SetSrcChannel (ctx sdk.Context, channel string) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SourceChannelKey))
	b := []byte(channel)
	if b == nil {
		fmt.Println(b)
		return
	}
	store.Set(types.KeyPrefix(types.SourceChannelKey), b)
}
