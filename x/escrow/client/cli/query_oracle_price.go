package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"dredd-secure/x/escrow/types"
)

func CmdListOraclePrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-oracle-price",
		Short: "list all OraclePrice",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOraclePriceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OraclePriceAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowOraclePrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-oracle-price [symbol]",
		Short: "shows a OraclePrice",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argSymbol := args[0]

			params := &types.QueryGetOraclePriceRequest{
				Symbol: argSymbol,
			}

			res, err := queryClient.OraclePrice(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
