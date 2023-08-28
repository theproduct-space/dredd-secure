package cli

import (
	"dredd-secure/x/escrow/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateEscrow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-escrow [initiator-coins] [fulfiller-coins] [start-date] [end-date] [api-conditions]",
		Short: "Broadcast message create-escrow",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argInitiatorCoins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}
			argFulfillerCoins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}
			argStartDate := args[2]
			argEndDate := args[3]
			argOracleConditions := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateEscrow(
				clientCtx.GetFromAddress().String(),
				argInitiatorCoins,
				argFulfillerCoins,
				argStartDate,
				argEndDate,
				argOracleConditions,
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
