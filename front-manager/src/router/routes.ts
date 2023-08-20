//对外暴露配置路由(常量路由):全部用户都可以访问到的路由
export const constantRoute = [
  {
    //登录
    path: '/login',
    component: () => import('@/view/login/index.vue'),
    name: 'login',
    meta: {
      title: '登录',
      hidden: true, //代表路由标题在菜单中是否隐藏  true:隐藏 false:不隐藏
      icon: 'Promotion' //菜单文字左侧的图标,支持element-plus全部图标
    }
  },
  {
    //注册
    path: '/register',
    component: () => import('@/view/register/index.vue'),
    name: 'register',
    meta:{
      title: '注册',
      hidden: true,
      icon: 'Promotion'
    }
  },
  {
    //首页，登录成功以后展示数据的路由
    path: '/',
    component: () => import('@/layout/index.vue'),
    name: 'layout',
    meta: {
      title: '',
      hidden: false,
      icon: ''
    },
    redirect: '/home',
    children: [
      {
        path: '/home',
        component: () => import('@/view/home/index.vue'),
        meta: {
          title: '首页',
          hidden: false,
          icon: 'HomeFilled'
        }
      },
    ]
  }, 
  {
    //404
    path: '/404',
    component: () => import('@/view/404/index.vue'),
    name: '404',
    meta: {
      title: '404',
      hidden: true,
      icon: 'DocumentDelete'
    }
  }, 
]

//异步路由
export const asyncRoute = [
  {
    //数据大屏
    path: '/screen',
    component: () => import('@/layout/index.vue'),
    redirect: '/screen',
    children: [
      {
        path: '/screen',
        component: () => import('@/view/screen/index.vue'),
        name: 'ScreenView',
        meta: {
          hidden: false,
          title: '数据大屏',
          icon: 'Platform',
        }
      },
    ]
  },
  {
    //设备管理----充电站/充电桩
    path: '/device',
    component: () => import('@/layout/index.vue'),
    name: 'Device',
    meta: {
      title: '设备管理',
      hidden: false,
      icon: 'Suitcase'
    },
    redirect: '/device/station',
    children: [
      {
        path: '/device/station',
        component: () => import('@/view/device/station/index.vue'),
        name: 'ChargeStationView',
        meta: {
          title: '充电站管理',
          icon: 'Flag',
          has:true
        }
      },
      {
        path: '/device/pile',
        component: () => import('@/view/device/pile/index.vue'),
        name: 'ChargePileView',
        meta: {
          title: '充电桩管理',
          icon: 'Stamp'
        }
      }
    ]
  },
  {
    //系统设置---用户管理/租户管理/角色管理/权限管理/充值管理
    path: '/acl',
    component: () => import('@/layout/index.vue'),
    name: 'Acl',
    meta: {
      title: '系统设置',
      hidden: false,
      icon: 'List'
    },
    redirect: '/acl/user',
    children: [
      {
        path: '/acl/user',
        component: () => import('@/view/acl/user/index.vue'),
        name: 'UserView',
        meta: {
          title: '用户管理',
          icon: 'User'
        }
      },
      // 管理台不需要租户管理，路由不展示了
      // {
      //   path: '/acl/tenant',
      //   component: () => import('@/view/acl/tenant/index.vue'),
      //   name: 'TenantView',
      //   meta: {
      //     title: '租户管理',
      //     icon: 'Grid'
      //   }
      // },
      {
        path: '/acl/role',
        component: () => import('@/view/acl/role/index.vue'),
        name: 'RoleView',
        meta: {
          title: '角色管理',
          icon: 'Avatar'
        }
      },
      {
        path: '/acl/authority',
        component: () => import('@/view/acl/authority/index.vue'),
        name: 'PermissionView',
        meta: {
          title: '权限管理',
          icon: 'Monitor',
          hidden: true,
        }
      },
      {
        path: '/acl/record',
        component: () => import('@/view/acl/record/index.vue'),
        name: 'RechargeView',
        meta: {
          title: '充值管理',
          icon: 'Money'
        }
      },
    ]
  },
  {
    //运营管理----订单管理/价格规则
    path: '/opration',
    component: () => import('@/layout/index.vue'),
    name: 'Opration',
    meta: {
      title: '运营管理',
      hidden: false,
      icon: 'Management'
    },
    redirect: '/opration/order',
    children: [
      {
        path: '/opration/order',
        component: () => import('@/view/opration/order/index.vue'),
        name: 'ChargeOrderView',
        meta: {
          title: '充电订单',
          hidden: false,
          icon: 'Lock'
        }
      },
      {
        path: '/opration/priceRange',
        component: () => import('@/view/opration/priceRange/index.vue'),
        name: 'PriceView',
        meta: {
          title: '充电价格',
          hidden: false,
          icon: 'PriceTag'
        }
      }
    ]
  },
  {
    //审计
    path: '/audit',
    component: () => import('@/layout/index.vue'),
    redirect: '/audit',
    children: [
      {
        path: '/audit',
        component: () => import('@/view/audit/index.vue'),
        name: 'AtuditLogView',
        meta: {
          title: '审计日志',
          hidden: false,
          icon: 'Edit'
        }
      },
    ]
  },
]

//任意路由
export const anyRoute = {
  //任意路由
  path: '/:pathMatch(.*)*',
  redirect: '/404',
  name: 'Any',
  meta: {
    title: '任意路由',
    hidden: true
  }
}