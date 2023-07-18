//引入仓库
import pinia from '@/store'
import useUserStore from '@/store/module/use'
let useStore = useUserStore(pinia)

export const isHasPermission = (app:any)=>{
app.directive('has',{
  mounted(el:any,options:any) {
    if(!useStore.permissionList.includes(options.value)){
      el.parentNode.removeChild(el)
      // el.style.display = 'none';
    }
  }
})
}

/*
自定义指令说明：
app.directive('has',{
    mounted(el,options) {
    console.log(el);
    console.log(options);
  }
})
表示注册一个名为has的自定义指令，mounted是组件挂载完毕会执行一次
el是使用全局指令的DOM元素，options中的value是可以拿到自定义指令的值
判断options的value值在数组中是否存在，有就留着，没有就从DOM树上干掉
*/