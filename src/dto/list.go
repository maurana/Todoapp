package dto

import (
	"todoapp/src/abstraction"
	"todoapp/src/model"
	res "todoapp/pkg/util/response"
)

// Get
type ListGetRequest struct {
	abstraction.Pagination
	model.ListFilterModel
}
type ListGetResponse struct {
	Datas          []model.ListEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type ListGetResponseDoc struct {
	Body struct {
		Meta res.Meta                  `json:"meta"`
		Data []model.ListEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type ListGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type ListGetByIDResponse struct {
	model.ListEntityModel
}
type ListGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data ListGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type ListCreateRequest struct {
	model.ListEntity
}
type ListCreateResponse struct {
	model.ListEntityModel
}
type ListCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ListCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type ListUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.ListEntity
}
type ListUpdateResponse struct {
	model.ListEntityModel
}
type ListUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ListUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type ListDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type ListDeleteResponse struct {
	model.ListEntityModel
}
type ListDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ListDeleteResponse `json:"data"`
	} `json:"body"`
}