package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
)

type (
	UpdateChannelRequestProposalJSON struct {
		Title          string             `json:"title"           yaml:"title"`
		Description    string             `json:"description"     yaml:"description"`
		ChannelRequest string			  `json:"channel_request" yaml:"channel_request"`
		Deposit        string             `json:"deposit"         yaml:"deposit"`
	}
)

func ParseUpdateChannelRequestProposalJSON(
	cdc *codec.LegacyAmino,
	proposalFile string,
) (UpdateChannelRequestProposalJSON, error) {
	proposal := UpdateChannelRequestProposalJSON{}
	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}