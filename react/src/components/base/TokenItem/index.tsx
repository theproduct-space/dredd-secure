import React from 'react'
import { IToken } from '~baseComponents/TokenSelector'

function TokenItem(props: IToken) {
    return (
        <div className="single-token">
            <div className="token-info">
                <div className="token-img">IMG</div>
                <div className="token">
                    <div className="token-name">{props.name}</div>
                    <div className="token-denom">{props.denom}</div>
                </div>

                {
                    props.amount && props.amount > 0 &&
                    <div className="token-amount">
                        {props.amount}
                    </div>
                }
            </div>
        </div>
    )
}

export default TokenItem