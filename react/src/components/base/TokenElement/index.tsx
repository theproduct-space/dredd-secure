import { Coin } from 'dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin';
import React from 'react'

export interface TokenElementProps {
    onClick?: () => void;
    selectedToken?: Coin;
    baseButton?: JSX.Element;
}

export function TokenElement(props: TokenElementProps) {
    const { onClick, selectedToken, baseButton } = props;
    const handleOnClick = () => {
        if (onClick) onClick();
        console.log("clicked");
    }

    return (
        <>
            {
                selectedToken ?
                    <div className="token-display">
                        <div className="token-amount">{selectedToken.amount}</div>
                        <div className="token-info" onClick={() => handleOnClick()}>
                            <div className="token-img"></div>
                            <div className="token">
                                <div className="token-name">{selectedToken.denom}</div>
                                <div className="token-denom">{selectedToken.denom}</div>
                            </div>
                        </div>
                    </div>
                    :
                <span onClick={() => handleOnClick()}>
                    {baseButton}
                </span>
            }
        </>
    )
}