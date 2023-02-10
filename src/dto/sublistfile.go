package dto

import (
	"todoapp/src/abstraction"
	"todoapp/src/model"
	res "todoapp/pkg/util/response"
)

// Get
type SublistfileGetRequest struct {
	abstraction.Pagination
}

type SublistfileGetResponse struct {
	Datas          []model.SublistfileEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type SublistfileGetResponseDoc struct {
	Body struct {
		Meta res.Meta                  `json:"meta"`
		Data []model.SublistfileEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type SublistfileGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type SublistfileGetByIDResponse struct {
	model.SublistfileEntityModel
}
type SublistfileGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta              `json:"meta"`
		Data SublistfileGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type SublistfileCreateRequest struct {
	model.SublistfileEntity
}
type SublistfileCreateResponse struct {
	model.SublistfileEntityModel
}
type SublistfileCreateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SublistfileCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type SublistfileUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.SublistfileEntity
}
type SublistfileUpdateResponse struct {
	model.SublistfileEntityModel
}
type SublistfileUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SublistfileUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type SublistfileDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type SublistfileDeleteResponse struct {
	model.SublistfileEntityModel
}
type SublistfileDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta             `json:"meta"`
		Data SublistfileDeleteResponse `json:"data"`
	} `json:"body"`
}