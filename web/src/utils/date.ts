export const formatDate = (date: string) => {
	const d = new Date(date);
	return `${d.getFullYear()}-${(d.getMonth() + 1) < 10 ? '0' + (d.getMonth() + 1) : (d.getMonth() + 1)}-${d.getDate() < 10 ? '0' + d.getDate() : d.getDate()} ${d.getHours() < 10 ? '0' + d.getHours() : d.getHours()}:${d.getMinutes() < 10 ? '0' + d.getMinutes() : d.getMinutes()}:${d.getSeconds() < 10 ? '0' + d.getSeconds() : d.getSeconds()}`;
}
