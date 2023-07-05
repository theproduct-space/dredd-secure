import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import { TokenElement } from '~baseComponents/TokenElement';
import { IToken, TokenSelector } from '~baseComponents/TokenSelector';

function CreateContract() {
    enum Modals {
        Own,
        Wanted,
        Tips
    };

    const [modalToOpen, setModalToOpen] = useState<Modals | undefined>();
    const [selectedOwnToken, setSelectedOwnToken] = useState<IToken | undefined>();
    const [selectedWantedToken, setSelectedWantedToken] = useState<IToken | undefined>();
    const [selectedTokenTips, setSelectedTokenTips] = useState<IToken | undefined>();

    const handleSaving = (t: IToken | undefined) => {
        let func;
        switch (modalToOpen) {
            case Modals.Own:
                func = setSelectedOwnToken;
                break;
            case Modals.Wanted:
                func = setSelectedWantedToken;
                break;
            default:
                func = setSelectedTokenTips;
                break;
        }

        func(t);
        setModalToOpen(undefined);
    }

    const displayModal = () => {
        let modal;
        switch (modalToOpen) {
            case Modals.Own:
                modal = selectedOwnToken;
                break;
            case Modals.Wanted:
                modal = selectedWantedToken;
                break;
            default:
                modal = selectedTokenTips;
                break;
        }
        return (
            <TokenSelector selectedToken={modal} onSave={handleSaving} />
        )
    }


    const displayConditionTypes = () => {
        const conditionTypes = ["Deadline", "Milestone"];

        return conditionTypes.map((type) => { return <option value={type}>{type}</option> });
    }
    // This is for testing purposes
    const conditions = [
        {
            type: "Deadline",
            value: "07/26/2025"
        },
        {
            type: "Milestone",
            value: "08/02/2024"
        },
    ]

    return (
        <div>
            <Link to="/">GO BACK</Link>
            <div>
                <span className="overheader">STEP 1</span>
                <div className="title-2">Create Contract</div>
                <div className="card">
                    <div className="conditions-management">
                        <div className="subtitle">Add Conditions</div>
                        {
                            conditions.map((condition, index) => {
                                return (
                                    <div className="condition"> {/* Might be a component for a condition and maybe a section for condition-list */}
                                        <div className="condition-number">Condition #{index + 1}</div>
                                        <div className="condition-value">
                                            <select value={condition.type}>
                                                {displayConditionTypes()}
                                            </select>
                                            <input value={condition.value}></input>
                                            <span>-</span>
                                        </div>
                                    </div>
                                )
                            })
                        }
                        <div className="add-condition">Add Another Condition</div>
                    </div>
                    <div className="assets-management">
                        <div className="subtitle">Choose Assets for Exchange</div>
                        <div className="small-text">To complete this escrow, you must choose an asset you want to give and an asset to receive</div>
                        <div className="assets">
                            <div className="assets-selection">
                                <div className="sub-subtitle">Select Your Assets:</div>
                                {/* Will take as a prop another component for the base display. Here, it will be a "Select Token" button */}
                                <TokenElement selectedToken={selectedOwnToken} onClick={() => setModalToOpen(Modals.Own)} baseButton={<span>Select Token</span>} />
                            </div>
                            <div className="assets-selection">
                                <div className="sub-subtitle">Asset you want to receive:</div>
                                <TokenElement selectedToken={selectedWantedToken} onClick={() => setModalToOpen(Modals.Wanted)} baseButton={<span>Select Token</span>} />
                            </div>
                        </div>
                    </div>
                    <div className="tips-section">
                        <span>Tips and donations go a long way.</span>
                        <div>
                            <span>We are a free service. Lorem ipsum</span>
                            {/* Will take as a prop another component for the base display. Here, it will be a "Add Tip" link or button */}
                            <TokenElement selectedToken={selectedTokenTips} onClick={() => setModalToOpen(Modals.Tips)} baseButton={<span>Add Tip</span>} />
                        </div>
                    </div>
                </div>
            </div>
            {
                modalToOpen != undefined && displayModal()
            }
        </div>
    )
}

export default CreateContract
