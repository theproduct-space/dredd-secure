package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"dredd-secure/x/escrow/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.CmdProposalUpdateSourceChannel)