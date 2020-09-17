import { Transaction, ItemTransaction, Item, StockIn, Project } from "./model";
import { ItemTransactionView } from "./view";

export interface TransactionPostBody {
  transaction: Transaction,
  itemTransactions: ItemTransactionView[],
  itemTransactionDeleteIds: number[]
}

export interface ItemPostBody {
  item: Item,
  withInitialStock: boolean,
  initialStockQty: number,
  project: Project | null
}

export interface StockInPostBody {
  item: Item, 
  stockIns: StockIn[]
}