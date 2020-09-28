/* Do not change, this code is generated from Golang structs */


export class ApiKey {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    apiKey: string;

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new ApiKey();
        result.id = source["id"];
        result.createdAt = source["createdAt"];
        result.updatedAt = source["updatedAt"];
        result.deletedAt = source["deletedAt"];
        result.apiKey = source["apiKey"];
        return result;
    }

}
export class ItemTransaction {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    uid: string;
    itemId: number;
    transactionId: number;
    qty: number;
    item: Item;

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new ItemTransaction();
        result.id = source["id"];
        result.createdAt = source["createdAt"];
        result.updatedAt = source["updatedAt"];
        result.deletedAt = source["deletedAt"];
        result.uid = source["uid"];
        result.itemId = source["itemId"];
        result.transactionId = source["transactionId"];
        result.qty = source["qty"];
        result.item = source["item"] ? Item.createFrom(source["item"]) : null;
        return result;
    }

}
export class Item {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    uid: string;
    name: string;
    description: string;
    price: number;
    manufacturingPrice: number;
    itemTransactions: ItemTransaction[];

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new Item();
        result.id = source["id"];
        result.createdAt = source["createdAt"];
        result.updatedAt = source["updatedAt"];
        result.deletedAt = source["deletedAt"];
        result.uid = source["uid"];
        result.name = source["name"];
        result.description = source["description"];
        result.price = source["price"];
        result.manufacturingPrice = source["manufacturingPrice"];
        result.itemTransactions = source["itemTransactions"] ? source["itemTransactions"].map(function(element: any) { return ItemTransaction.createFrom(element); }) : null;
        return result;
    }

}
export class Transaction {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    uid: string;
    cashier: string;
    priceIsCustom: boolean;
    customPrice: number;
    projectId: number;
    itemTransactions: ItemTransaction[];

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new Transaction();
        result.id = source["id"];
        result.createdAt = source["createdAt"];
        result.updatedAt = source["updatedAt"];
        result.deletedAt = source["deletedAt"];
        result.uid = source["uid"];
        result.cashier = source["cashier"];
        result.priceIsCustom = source["priceIsCustom"];
        result.customPrice = source["customPrice"];
        result.projectId = source["projectId"];
        result.itemTransactions = source["itemTransactions"] ? source["itemTransactions"].map(function(element: any) { return ItemTransaction.createFrom(element); }) : null;
        return result;
    }

}
export class Project {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    uid: string;
    name: string;
    startDate: string;
    transactions: Transaction[];

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new Project();
        result.id = source["id"];
        result.createdAt = source["createdAt"];
        result.updatedAt = source["updatedAt"];
        result.deletedAt = source["deletedAt"];
        result.uid = source["uid"];
        result.name = source["name"];
        result.startDate = source["startDate"];
        result.transactions = source["transactions"] ? source["transactions"].map(function(element: any) { return Transaction.createFrom(element); }) : null;
        return result;
    }

}

export class StockIn {
    id: number;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
    uid: string;
    pic: string;
    itemId: number;
    projectId: number;
    qty: number;

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new StockIn();
        result.id = source["id"];
        result.createdAt = source["createdAt"];
        result.updatedAt = source["updatedAt"];
        result.deletedAt = source["deletedAt"];
        result.uid = source["uid"];
        result.pic = source["pic"];
        result.itemId = source["itemId"];
        result.projectId = source["projectId"];
        result.qty = source["qty"];
        return result;
    }

}
