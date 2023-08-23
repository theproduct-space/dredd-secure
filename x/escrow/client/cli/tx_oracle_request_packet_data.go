package cli

import (
	"strconv"

	"dredd-secure/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channelutils "github.com/cosmos/ibc-go/v7/modules/core/04-channel/client/utils"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendOracleRequestPacketData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-oracle-request-packet-data [src-port] [src-channel] [client-id] [oracle-script-id] [calldata] [ask-count] [min-count] [fee-limit] [prepare-gas] [execute-gas]",
		Short: "Send a OracleRequestPacketData over IBC",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			argClientId := args[2]
			argOracleScriptId, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argCalldata := []byte(args[4])
			argAskCount, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argMinCount, err := cast.ToUint64E(args[6])
			if err != nil {
				return err
			}
			argFeeLimit, err := sdk.ParseCoinsNormalized(args[7])
			if err != nil {
				return err
			}
			argPrepareGas, err := cast.ToUint64E(args[8])
			if err != nil {
				return err
			}
			argExecuteGas, err := cast.ToUint64E(args[9])
			if err != nil {
				return err
			}

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendOracleRequestPacketData(creator, srcPort, srcChannel, timeoutTimestamp, argClientId, argOracleScriptId, argCalldata, argAskCount, argMinCount, argFeeLimit, argPrepareGas, argExecuteGas)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
