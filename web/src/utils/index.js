export default { 
	formatDate: (date) => {
 		const d = new Date(date);
		const year = d.getFullYear();
 		const month = d.getMonth() + 1;
 		const day = d.getDate();
 		const hour = d.getHours();
 		const minute = d.getMinutes();
 		const second = d.getSeconds();
 		return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day} ${hour < 10 ? '0' + hour : hour}:${minute < 10 ? '0' + minute : minute}:${second < 10 ? '0' + second : second}`;
	}
};
