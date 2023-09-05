let BASE_URL;
//开发环境中
if (process.env.NODE_ENV === 'development') {
	// 开发环境
	BASE_URL = 'http://127.0.0.1:8000/';
} else {
	// 生产环境
	BASE_URL = 'https://7483f88q46.goho.co';
}
 
export default BASE_URL;