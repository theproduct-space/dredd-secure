import { V1Beta1Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/rest";
import { Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin";
import { IToken } from "~baseComponents/TokenSelector";
import { IContract } from "~sections/CreateContract";
import { ICondition } from "~sections/CreateContract/AddConditions";
import assets from "~src/tokens.json";


export const CoinToIToken = (c: V1Beta1Coin | undefined): IToken | undefined => {
    if (c) {
        const token = assets.tokens.find((t) => t.denom === c.denom);
        if (token) {
            console.log("token", token);
            return {
                name: token.name,
                denom: token.denom,
                amount: Number(c.amount),
                logos: token.logos,
                display: token.display,
                chain_name: token.chain_name,
                selectedAmount: Number(c.amount)
            };
        }
    }

    return;
};

export const ContractToEscrow = (c: IContract): {
    initiatorCoins: Coin[],
    fulfillerCoins: Coin[],
    tips: Coin[],
    startDate: string,
    endDate: string,
    apiConditions: string,
} => {
    const initiatorCoins: Coin[] = [
        {
            denom: c.initiatorCoins.denom,
            amount: c.initiatorCoins.selectedAmount?.toString() ?? "0",
        },
    ];
    const fulfillerCoins: Coin[] = [
        {
            denom: c.fulfillerCoins.denom,
            amount: c.fulfillerCoins.selectedAmount?.toString() ?? "1", // TODO set back to 0
        },
    ];

    let tips: Coin[] = [];
    if (c.tips) {
        tips = [
            {
                denom: c.tips.denom,
                amount: c.tips.selectedAmount?.toString() ?? "1", // TODO set back to 0
            },
        ];
    }

    // Conditions message preparation
    let endDate = String((new Date("9999-12-31").getTime() / 1000).toFixed());
    let startDate = String((Date.now() / 1000).toFixed());
    const apiConditionsArray: ICondition[] = [];
    c.conditions?.map((condition) => {
        switch (condition.type) {
            case "startDate":
                startDate = String(condition.value);
                return;
            case "endDate":
                endDate = String(condition.value);
                return;
            case "apiCondition":
                apiConditionsArray.push(condition);
                return;
        }
    });

    // the sendMsgCreateEscrow will accept a apiConditions array stringified
    const apiConditions: string = JSON.stringify(apiConditionsArray);

    return {
        initiatorCoins: initiatorCoins,
        fulfillerCoins: fulfillerCoins,
        tips: tips,
        startDate: startDate,
        endDate: endDate,
        apiConditions: apiConditions,
    }
}