import { Item, Transaction, Project, ItemTransaction, StockIn } from "./model";
import { TransactionView, ItemTransactionView } from './view';
import { makeDateString } from './helpers';

export const initialItem: Item = {
  id: 0,
  uid: "",
  name: "",
  description: "",
  price: 0,
  manufacturingPrice: 0,
  created_at: undefined,
  updated_at: undefined
}

export const initialTransaction: Transaction = {
  id: 0,
  uid: "",
  cashier: "",
  priceIsCustom: false,
  customPrice: 0,
  projectId: 0,
  created_at: undefined,
  updated_at: undefined
}

export const initialProject: Project = {
  id: 0,
  uid: "",
  name: "",
  startDate: makeDateString(new Date()),
  created_at: undefined,
  updated_at: undefined
}

export const initialTransactionView: TransactionView = {
  transaction: { ...initialTransaction },
  itemTransactions: [],
  totalPrice: 0
}

export const initialItemTransaction: ItemTransaction = {
  id: 0,
  uid: "",
  itemId: 0,
  transactionId: 0,
  qty: 0,
  created_at: undefined,
  updated_at: undefined
}

export const initialItemTransactionView: ItemTransactionView = {
  itemTransaction: {...initialItemTransaction},
  item: {...initialItem}
}

export const initialStockIn: StockIn = {
  id: 0,
  uid: "",
  pic: "",
  itemId: 0,
  projectId: 0,
  qty: 0,
  created_at: undefined,
  updated_at: undefined
}