//封装一个函数:获取一个结果:当前早上|上午|下午|晚上
export const getTime = () => {
  let message = ''
  //通过内置构造函数Date
  const hours = new Date().getHours()
  //情况的判断
  if (hours <= 9) {
    message = '早上'
  } else if (hours <= 12) {
    message = '上午'
  } else if (hours <= 18) {
    message = '下午'
  } else {
    message = '晚上'
  }
  return message
}

//封装一个函数：简化时间格式
export const formatDate = (dateStr:any)=>{
  const date = new Date(dateStr)
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}
