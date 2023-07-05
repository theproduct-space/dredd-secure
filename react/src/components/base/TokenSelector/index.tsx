import React, { useState } from 'react'
import TokenItem from '~baseComponents/TokenItem';

export interface IToken {
    name: string;
    amount?: number;
    denom: string;
}

export interface TokenSelectorProps {
    onSave: (token: IToken | undefined) => void;
    address?: string;
    selectedToken?: IToken;
}

export function TokenSelector(props: TokenSelectorProps) {
    const [selectedToken, setSelectedToken] = useState(props.selectedToken);
    const { onSave, address } = props;

    const displayOwnedToken = () => {
    }

    const displayAllToken = () => {
        const tokens: IToken[] = [{
            name: "token",
            denom: "tok",
        }]

        return (
            <TokenItem {...tokens[0]} />
        )
    }

    return (
        <div className="modal">
            <div className="card">
                <div className="card-headers">
                    Select a token
                    <div className="search-bar">Search token</div>
                </div>
                <div className="card-body">
                    {
                        displayAllToken()
                    }
                </div>
            </div>
        </div>
    )
}
