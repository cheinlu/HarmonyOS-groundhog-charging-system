const http = uni.$u.http

//发起登录请求    /wx-api/login    post    参数：{ "code":"xxxxxx"}
export const requestLogin = (data)=>http.post('/wx-api/login',data)

//获取个人中心信息   /wx-api/me/info   get
export const requestUserInfo = () => http.get('/wx-api/me/info')

//请求验证码  /wx-api/validcode   get   参数:{"phone":"xxx"}
export const  requestVerificationCode = (params = {}) => http.get(`/wx-api/validcode`, params)

//发起注册请求   /wx-api/register   post   
/*
参数：{"code": "xxx","phone": "xxx","verifyCode": "xxx","nickname": "xxx","avatarUrl": "xxx","gender":"0"}
*/
export const requestRegister = (data)=>http.post('/wx-api/register',data)