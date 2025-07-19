export interface Investment {
    id: string;
    name: string;
    amount: number;
    assetClass: AssetClass;
}

export enum AssetClass {
    EQUITY = "Equity",
    BOND = "Bond",
    REAL_ESTATE = "Real Estate",
    COMMODITY = "Commodity",
    CASH = "Cash"
}