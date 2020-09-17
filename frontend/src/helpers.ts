export type RequestStatus 
  = 'Success'
  | 'Error'
  | 'Loading'
  | 'NotAsked'

const months = [ 'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec' ]

export const formatIdr = (moneyAmount: number) => {
  return Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(moneyAmount)
}

export const makeDateString = (date: Date) => {
  const y = date.getFullYear()
  const m = date.getMonth() + 1 < 10 ? `0${date.getMonth() + 1}` : `${date.getMonth()}`
  const d = date.getDate() < 10 ? `0${date.getDate()}` : `${date.getDate()}`

  return `${y}-${m}-${d}`
}

export const makeReadableDateString = (date: Date) => {
  return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`
}