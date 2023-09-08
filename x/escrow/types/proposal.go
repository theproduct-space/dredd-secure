package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	UpdateChannelRequest = "UpdateChannelRequest"
)

var _ govtypes.Content = &UpdateChannelRequestProposal{}

func init() {
	govtypes.RegisterProposalType(UpdateChannelRequest)
}

func NewUpdateChannelRequestProposal(
	title, description string, channel string,
) *UpdateChannelRequestProposal {
	return &UpdateChannelRequestProposal{title, description, channel}
}

// GetTitle returns the title of a update symbol request proposal.
func (p *UpdateChannelRequestProposal) GetTitle() string { return p.Title }

// GetDescription returns the description of a update symbol request proposal.
func (p *UpdateChannelRequestProposal) GetDescription() string { return p.Description }

// ProposalRoute returns the routing key of a update symbol request proposal.
func (*UpdateChannelRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a update symbol request proposal.
func (*UpdateChannelRequestProposal) ProposalType() string { return UpdateChannelRequest }

// ValidateBasic validates the update symbol request proposal.
func (p *UpdateChannelRequestProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return ValidateChannelRequest(p.ChannelRequest)
}

func NewChannelRequest(channel string) string {
	return channel
}

func ValidateChannelRequest(channel string) error {
	if len(channel) == 0 {
		return ErrEmptyChannelRequest
	}

	return nil
}