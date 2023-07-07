import React, { useState } from 'react'
import { TokenElement } from '~baseComponents/TokenElement';
import TokenPreview from '~baseComponents/TokenPreview';
import { TokenSelector } from '~baseComponents/TokenSelector';
import { IContract } from '~sections/CreateContract/CreateContract'

interface PaymentSectionProps {
    contract: IContract;
}

function PaymentSection(props: PaymentSectionProps) {
    const { contract } = props;
    const [ selectedTokenTips, setSelectedTokenTips ] = useState(contract.tips);
    const [ displayTips, setDisplayTips ] = useState(false);

    const displayTipsSelection = () => {
        return (<TokenSelector selectedToken={selectedTokenTips} onSave={setSelectedTokenTips} />)
    }
    console.log(contract)
    return (
        <div>
            <div className="Title">Review and Confirm Exchange</div>
            <div className="card">
                <div className="card-subtitle">Conditions</div>

                {
                    contract.conditions.map((condition) => {
                        return (
                            <div>
                                <div className="condition-name">{condition.type}</div>
                                <div className="condition-value">{condition.value}</div>
                            </div>
                        )
                    })
                }

                <div className="card-subtitle">You are exchanging</div>
                <div>
                    <TokenPreview token={contract.initiatorCoins} />
                    <div className="exchange-icon"></div>
                    <TokenPreview token={contract.fulfillerCoins} />
                </div>

                <div className="tips-section">
                    <span>Tips and donations go a long way.</span>
                    <div>
                        <span>We are a free service. Lorem ipsum</span>
                        {/* Will take as a prop another component for the base display. Here, it will be a "Add Tip" link or button */}
                        <TokenElement selectedToken={selectedTokenTips} onClick={() => setDisplayTips(true)} baseButton={<span>Add Tip</span>} />
                    </div>
                </div>
            </div>
            {
                displayTips && displayTipsSelection()
            }
        </div>
    )
}

export default PaymentSection