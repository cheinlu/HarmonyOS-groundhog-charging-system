//定义全部接口返回数据都拥有ts类型
export interface ResponseData{
code :number
message:string
}

//用户登录接口携带参数的ts类型
export interface loginFormData {
  username: string
  password: string
}

//定义登录接口返回数据类型
export interface loginResponseData extends ResponseData{
  data:{
    role:string
    token:string
  }
}