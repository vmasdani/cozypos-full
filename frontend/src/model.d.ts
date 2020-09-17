export interface ApiKey {
  id: number,
  apiKey: string,
  created_at: string | undefined,
  updated_at: string | undefined
} 
export interface Item {
  id: number,
  uid: string,
  name: string,
  description: string,
  price: number,
  manufacturingPrice: number,
  created_at: string | undefined,
  updated_at: string | undefined
}

export interface Project {
  id: number,
  uid: string,
  name: string,
  startDate: string,
  created_at: string | undefined,
  updated_at: string | undefined
}

export interface Transaction {
  id: number,
  uid: string,
  cashier: string,
  priceIsCustom: boolean,
  customPrice: number,
  projectId: number,
  created_at: string | undefined,
  updated_at: string | undefined
}

export interface StockIn {
  id: number,
  uid: string,
  pic: string,
  itemId: number,
  projectId: number,
  qty: number,
  created_at: string | undefined,
  updated_at: string | undefined
}

export interface ItemTransaction {
  id: number,
  uid: string,
  itemId: number,
  transactionId: number,
  qty: number,
  created_at: string | undefined,
  updated_at: string | undefined
}

export interface ItemStockIn {
  id: number,
  uid: string,
  itemId: number,
  stockInId: number,
  qty: number,
  created_at: string | undefined,
  updated_at: string | undefined
}

export interface ItemProject {
  id: number,
  uid: string,
  itemId: number,
  projectId: number,
  qty: number,
  created_at: string | undefined,
  updated_at: string | undefined
}