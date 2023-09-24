package consts

type ContextKeyStruct string

const (
	TenantIDKey         ContextKeyStruct = "tenantID" //存储在 ctx 中租户id的 key
	DefaultTenantValue  int              = 1          // 默认的租户id的值（测试用户）
	UserContextKey      ContextKeyStruct = "userCtx"
	AdminRoleName       string           = "系统管理员"
	AdminRolePermission string           = "ChargeStationView,ChargeStationAdd,ChargeStationUpdate,ChargeStationDel,ChargePileView,ChargePileAdd,ChargePileUpdate,ChargePileDel,PriceView,PriceAdd,PriceDel,PriceUpdate,ChargeOrderView,AddPrice,AtuditLogView,PermissionView,PermissionViewBind,RoleView,RoleAdd,RoleDel,RoleUpdate,UserView,UserAdd,UserDel,UserUpdate"
)

const (
// ErrCap

)
