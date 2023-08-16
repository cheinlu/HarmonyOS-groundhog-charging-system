//存储TOKEN
export const SET_TOKEN = (token) => {
  uni.setStorageSync('TOKEN', token)
}
//本地存储获取数据
export const GET_TOKEN = () => {
  return uni.getStorageSync('TOKEN')
}
//本地存储删除数据方法
export const REMOVE_TOKEN = () => {
  uni.removeStorageSync('TOKEN')
}
//存储用户信息
export const SET_INFO = (info) => {
  uni.setStorageSync('info', info)
}
//本地存储获取用户信息
export const GET_INFO = () => {
  return uni.getStorageSync('info')
}
//本地存储删除用户信息
export const REMOVE_INFO = () => {
  uni.removeStorageSync('info')
}
