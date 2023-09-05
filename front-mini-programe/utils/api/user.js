import { setRequestConfig } from '@/utils/request.js';

// 调用setRequestConfig函数进行请求配置
setRequestConfig();
const http = uni.$u.http
// 发起登录请求
export const requestLogin = (data) => http.post('/wx-api/login', data);
//请求验证码  /wx-api/validcode   get   参数:{"phone":"xxx"}
export const  requestVerificationCode = (params = {}) => http.get(`/wx-api/validcode`, params)
//发起注册请求   /wx-api/register   post   
/*
参数：{"code": "xxx","phone": "xxx","verifyCode": "xxx","nickname": "xxx","avatarUrl": "xxx","gender":"0"}
*/
export const requestRegister = (data)=>http.post('/wx-api/register',data)
//获取个人中心信息   /wx-api/me/info   get
export const requestUserInfo = () => http.get('/wx-api/me/info')
//修改个人昵称  /wx-api/update/nickname post  参数：newNickname
export const requestNickname = (data)=>http.post('/wx-api/update/nickname',data)