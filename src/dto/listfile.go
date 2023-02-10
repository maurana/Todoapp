package dto

import (
	"todoapp/src/abstraction"
	"todoapp/src/model"
	res "todoapp/pkg/util/response"
)

// Get
type ListfileGetRequest struct {
	abstraction.Pagination
}

type ListfileGetResponse struct {
	Datas          []model.ListfileEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type ListfileGetResponseDoc struct {
	Body struct {
		Meta res.Meta                  `json:"meta"`
		Data []model.ListfileEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type ListfileGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type ListfileGetByIDResponse struct {
	model.ListfileEntityModel
}
type ListfileGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data ListfileGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type ListfileCreateRequest struct {
	model.ListfileEntity
}
type ListfileCreateResponse struct {
	model.ListfileEntityModel
}
type ListfileCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ListfileCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type ListfileUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.ListfileEntity
}
type ListfileUpdateResponse struct {
	model.ListfileEntityModel
}
type ListfileUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ListfileUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type ListfileDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type ListfileDeleteResponse struct {
	model.ListfileEntityModel
}
type ListfileDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data ListfileDeleteResponse `json:"data"`
	} `json:"body"`
}