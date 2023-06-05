package model

// 定义分页查询参数结构体
type PageReq struct {
	PageNo   int `json:'pageNo'`   // 页码
	PageSize int `json:'pageSize'` // 每页记录数
}

// 定义分页查询返回结构体
type PageRes struct {
	PageNo     int         `json:'pageNo'`     // 页码
	PageSize   int         `json:'pageSize'`   // 每页记录数
	TotalCount int         `json:'totalCount'` // 总数
	List       interface{} // 查询结果
}

// 初始化分页参数，设置默认值
func InitPageReq(req *PageReq, defaultPageNo, defaultPageSize int) {
	if req.PageNo == 0 || req.PageSize == 0 {
		req.PageNo = defaultPageNo
		req.PageSize = defaultPageSize
	}
}
