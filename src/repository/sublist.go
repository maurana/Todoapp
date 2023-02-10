package repository

import (
	"fmt"
	"todoapp/src/abstraction"
	"todoapp/src/model"

	"gorm.io/gorm"
)

type Sublist interface {
	Find(ctx *abstraction.Context, m *model.SublistFilterModel, p *abstraction.Pagination) (*[]model.SublistEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.SublistEntityModel, error)
	Create(ctx *abstraction.Context, e *model.SublistEntityModel) (*model.SublistEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.SublistEntityModel) (*model.SublistEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.SublistEntityModel) (*model.SublistEntityModel, error)
}

type sublist struct {
	abstraction.Repository
}

func NewSublist(db *gorm.DB) *sublist {
	return &sublist{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *sublist) Find(ctx *abstraction.Context, m *model.SublistFilterModel, p *abstraction.Pagination) (*[]model.SublistEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.SublistEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.SublistEntityModel{})

	// filter
	query = r.Filter(ctx, query, m)

	// sort
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}
	if p.SortBy == nil {
		sortBy := "title"
		p.SortBy = &sortBy
	}
	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	// pagination
	if p.Page == nil {
		page := 1
		p.Page = &page
	}
	if p.PageSize == nil {
		pageSize := 10
		p.PageSize = &pageSize
	}
	info = abstraction.PaginationInfo{
		Pagination: p,
	}
	limit := *p.PageSize + 1
	offset := (*p.Page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	err := query.Find(&datas).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, &info, err
	}

	info.Count = len(datas)
	info.MoreRecords = false
	if len(datas) > *p.PageSize {
		info.MoreRecords = true
		info.Count -= 1
		datas = datas[:len(datas)-1]
	}

	return &datas, &info, nil
}

func (r *sublist) FindByID(ctx *abstraction.Context, id *int) (*model.SublistEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.SublistEntityModel
	err := conn.Where("sublist_id = ?", id).First(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *sublist) Create(ctx *abstraction.Context, e *model.SublistEntityModel) (*model.SublistEntityModel, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Create(e).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(e).First(e).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}


func (r *sublist) Update(ctx *abstraction.Context, id *int, e *model.SublistEntityModel) (*model.SublistEntityModel, error) {
	conn := r.CheckTrx(ctx)

	// err := conn.Where("sublist_id = ?", id).First(e).WithContext(ctx.Request().Context()).Error
	// if err != nil {
	// 	return nil, err
	// }

	err := conn.Model(e).Updates(e).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *sublist) Delete(ctx *abstraction.Context, id *int, e *model.SublistEntityModel) (*model.SublistEntityModel, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("sublist_id = ?", id).Delete(e).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return e, nil
}