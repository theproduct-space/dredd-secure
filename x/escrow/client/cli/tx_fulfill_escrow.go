package cli

import (
	"dredd-secure/x/escrow/types"
	"strconv"

	"encoding/json"

	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdFulfillEscrow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fulfill-escrow [id] [denom-mapping]",
		// ex: fulfill-escrow 1 'base_denom1=ibcHash1,base_denom2=ibcHash2'
		Short: "Broadcast message fulfill_escrow",
		Args:  cobra.ExactArgs(2), // Expect exactly 2 arguments
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Println("FULFILL ESCROW CLI")
			fmt.Println("FULFILL ESCROW CLI")
			fmt.Println("FULFILL ESCROW CLI")
			fmt.Println("FULFILL ESCROW CLI")
			argID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Parse the JSON or comma-separated list of key-value pairs from the argument
			denomMappingsJSON := args[1]

			var denomMapping map[string]string
			err = json.Unmarshal([]byte(denomMappingsJSON), &denomMapping)
			if err != nil {
				return err
			}

			
			fmt.Println("denomMapping in CLI", denomMapping)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fmt.Println("clientCtx", clientCtx)

			// Create the MsgFulfillEscrow message with the map
			msg := types.NewMsgFulfillEscrow(
				clientCtx.GetFromAddress().String(),
				argID,
				denomMapping, // Pass the map here
			)

			fmt.Println("msg", msg)
			if err := msg.ValidateBasic(); err != nil {
				fmt.Println("err ValidateBasic", err)
				return err
			}

			fmt.Println("validated basic !")
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}