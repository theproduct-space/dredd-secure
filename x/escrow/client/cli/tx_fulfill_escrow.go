package cli

import (
	"dredd-secure/x/escrow/types"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdFulfillEscrow() *cobra.Command {
	cmd := &cobra.Command{
		Use: "fulfill-escrow [id] [denom-mapping]",
		// ex: fulfill-escrow 1 'base_denom1=ibcHash1,base_denom2=ibcHash2'
		Short: "Broadcast message fulfill_escrow",
		Args:  cobra.ExactArgs(2), // Expect exactly 2 arguments
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Parse the key-value pairs from the argument
			denomMappingsStr := args[1]
			denomMappings := strings.Split(denomMappingsStr, ",")

			var denomMapping []*types.KeyVal
			for _, kvStr := range denomMappings {
				kv := strings.Split(kvStr, "=")
				if len(kv) != 2 {
					return fmt.Errorf("invalid key-value pair: %s", kvStr)
				}
				kvPair := &types.KeyVal{Key: kv[0], Value: kv[1]}
				denomMapping = append(denomMapping, kvPair)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Create the MsgFulfillEscrow message with the map
			msg := types.NewMsgFulfillEscrow(
				clientCtx.GetFromAddress().String(),
				argID,
				denomMapping, // Pass the map here
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
