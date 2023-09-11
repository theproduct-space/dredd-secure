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
		Use:   "create-escrow [initiator-coins] [fulfiller-coins] [tips] [start-date] [end-date] [oracle-conditions]",
		Short: "Broadcast message create-escrow",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argInitiatorCoins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}
			argFulfillerCoins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}
			argTips, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}
			argStartDate := args[3]
			argEndDate := args[4]
			argOracleConditions := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateEscrow(
				clientCtx.GetFromAddress().String(),
				argInitiatorCoins,
				argFulfillerCoins,
				argTips,
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
