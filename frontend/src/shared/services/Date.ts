export function formatDate(date_string:any) {
  const date = new Date(date_string)
  return date.toLocaleString('ja-JP')
}
