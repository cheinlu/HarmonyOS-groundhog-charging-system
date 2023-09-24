package controller_saas

import (
	"context"
	v2 "login-demo/api/saas"
	"login-demo/internal/logic"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
)

type StationController struct {
}

// 充电站列表
func (StationController) List(ctx context.Context, req *v2.StationListReq) (pageRes model.PageRes, err error) {
	model.InitPageReq(&req.PageReq, 1, 10)
	query := entity.Station{
		Name:    req.Name,
		Address: req.Address,
	}
	stations, count, err := logic.Station.StationList(ctx, query, req.PageReq)
	res := make([]*v2.StationListRes, len(stations))
	for i, station := range stations {
		res[i] = &v2.StationListRes{
			Id:         station.Id,
			Name:       station.Name,
			Address:    station.Address,
			Coordinate: station.Coordinate,
			TenantName: station.TenantName,
			ImageUrl:   station.ImageUrl,
			CreateAt:   station.CreateAt.Time,
			UpdateAt:   station.UpdateAt.Time,
		}
	}
	pageRes.List, pageRes.PageNo, pageRes.PageSize, pageRes.TotalCount = res, req.PageNo, req.PageSize, count
	return
}

// 充电站添加
func (StationController) Add(ctx context.Context, req *v2.StationAddReq) (res *v2.StationAddRes, err error) {
	var station do.Station = do.Station{
		Name:       req.Name,
		Address:    req.Address,
		Coordinate: req.Coordinate,
		ImageUrl:   req.ImageUrl,
	}
	err = logic.Station.Add(ctx, station)
	return
}

// 删除充电站
func (StationController) Del(ctx context.Context, req *v2.StationDelReq) (res *v2.StationDelRes, err error) {
	err = logic.Station.Del(ctx, req.Id)
	return
}

func (StationController) Update(ctx context.Context, req *v2.StationUpdateReq) (res *v2.StationUpdateRes, err error) {

	station := do.Station{
		Id:         req.Id,
		Name:       req.Name,
		Address:    req.Address,
		Coordinate: req.Coordinate,
		ImageUrl:   req.ImageUrl,
	}
	err = logic.Station.Update(ctx, station)
	return
}

func (StationController) Upload(ctx context.Context, req *v2.StationUploadReq) (res *v2.StationUploadRes, err error) {
	err, fileName, fileUrl := logic.File.FileUpload(ctx, req.File)
	if err != nil {
		return
	}
	res = &v2.StationUploadRes{
		Name: fileName,
		Url:  fileUrl,
	}
	return
}
