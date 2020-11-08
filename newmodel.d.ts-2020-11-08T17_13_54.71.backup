/* Do not change, this code is generated from Golang structs */


export interface ApiKey {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt?: string;
    apiKey: string;
}
export interface Item {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt?: string;
    uid: string;
    name: string;
    description: string;
    price: number;
    manufacturingPrice: number;
    itemTransactions: ItemTransaction[] | null;
    stockIns: StockIn[] | null;
}
export interface Project {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt?: string;
    uid: string;
    name: string;
    startDate: string;
    transactions: Transaction[] | null;
}
export interface Transaction {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt?: string;
    uid: string;
    cashier: string;
    priceIsCustom: boolean;
    customPrice: number;
    projectId: number;
    itemTransactions: ItemTransaction[] | null;
}
export interface StockIn {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt?: string;
    uid: string;
    pic: string;
    itemId: number;
    projectId: number;
    qty: number;
}
export interface ItemTransaction {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt?: string;
    uid: string;
    itemId: number;
    transactionId: number;
    qty: number;
    item: Item;
}