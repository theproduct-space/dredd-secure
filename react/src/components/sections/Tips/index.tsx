import { Coin } from 'dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin'
import React from 'react'
import { TokenElement } from '~baseComponents/TokenElement'

interface TipsSectionProps {
    selectedToken: Coin | undefined;
    onClick: () => void;
}

function TipsSection(props: TipsSectionProps) {
    const { selectedToken, onClick } = props;

    return (
        <div className="tips-section">
            <span>Tips and donations go a long way.</span>
            <div>
                <span>We are a free service. Lorem ipsum</span>
                {/* Will take as a prop another component for the base display. Here, it will be a "Add Tip" link or button */}
                <TokenElement selectedToken={selectedToken} onClick={onClick} baseButton={<span>Add Tip</span>} />
            </div>
        </div>
    )
}

export default TipsSection