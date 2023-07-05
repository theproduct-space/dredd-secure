import React from 'react'
import { IToken } from '~baseComponents/TokenSelector';

export interface TokenElementProps {
    onClick?: () => void;
    selectedToken?: IToken;
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
                                <div className="token-name">{selectedToken.name}</div>
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