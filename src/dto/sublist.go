package dto

import (
	"todoapp/src/abstraction"
	"todoapp/src/model"
	res "todoapp/pkg/util/response"
)

// Get
type SublistGetRequest struct {
	abstraction.Pagination
	model.SublistFilterModel
}
type SublistGetResponse struct {
	Datas          []model.SublistEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type SublistGetResponseDoc struct {
	Body struct {
		Meta res.Meta                  `json:"meta"`
		Data []model.SublistEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type SublistGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type SublistGetByIDResponse struct {
	model.SublistEntityModel
}
type SublistGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data SublistGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type SublistCreateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.SublistEntity
}
type SublistCreateResponse struct {
	model.SublistEntityModel
}
type SublistCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SublistCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type SublistUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.SublistEntity
}
type SublistUpdateResponse struct {
	model.SublistEntityModel
}
type SublistUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SublistUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type SublistDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type SublistDeleteResponse struct {
	model.SublistEntityModel
}
type SublistDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SublistDeleteResponse `json:"data"`
	} `json:"body"`
}