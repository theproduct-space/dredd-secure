package cli

import (
	"dredd-secure/x/escrow/types"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateEscrow())
	cmd.AddCommand(CmdCancelEscrow())
	cmd.AddCommand(CmdFulfillEscrow())
	cmd.AddCommand(CmdOptOutEscrow())
	cmd.AddCommand(CmdSendOracleRequestPacketData())
	cmd.AddCommand(CmdSetSourceChannel())
	// this line is used by starport scaffolding # 1

	return cmd
}
