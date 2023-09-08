package cli

import (
	"dredd-secure/x/escrow/types"
	"dredd-secure/x/escrow/utils"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = strconv.Itoa(0)

func CmdProposalUpdateSourceChannel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-source-channel-request [proposal-file]",
		Short: "Submit update source channel request proposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			proposal, err := utils.ParseUpdateChannelRequestProposalJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			content := types.NewUpdateChannelRequestProposal(
				proposal.Title, proposal.Description, proposal.ChannelRequest,
			)

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)

			if err != nil {
				return err
			}

			msg, err := govv1beta1.NewMsgSubmitProposal(content, deposit, from)

			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}