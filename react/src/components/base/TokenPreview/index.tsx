import { Coin } from 'dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin'
import React from 'react'

interface TokenPreviewProps {
    token: Coin;
}

function TokenPreview(props: TokenPreviewProps) {
    const {token} = props;
    return (
        <div>
            <div className="token-img">IMG</div>
            <div className="token-amount">{token?.amount}</div>
            <div className="token-denom">{token?.denom}</div>
        </div>
    )
}

export default TokenPreview