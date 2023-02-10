package list

import (
	"os"
	"io"
	"fmt"
	"todoapp/src/abstraction"
	"todoapp/src/dto"
	"todoapp/src/factory"
	"todoapp/src/model"
	res "todoapp/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

// Get
// @Summary Get list
// @Description Get list
// @Tags list
// @Accept json
// @Produce json
// @param request query dto.ListGetRequest true "request query"
// @Success 200 {object} dto.ListGetResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /list [get]
func (h *handler) Get(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ListGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Find(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get datas success", &result.PaginationInfo).Send(c)
}

// Get By ID
// @Summary Get list by id
// @Description Get list by id
// @Tags list
// @Accept json
// @Produce json
// @Param id path int true "id path"
// @Success 200 {object} dto.ListGetByIDResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /list/{id} [get]
func (h *handler) GetByID(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ListGetByIDRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	fmt.Printf("%+v", payload)

	result, err := h.service.FindByID(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Create list
// @Summary Create list
// @Description Create list
// @Tags list
// @Accept  json
// @Produce  json
// @Param request body dto.ListCreateRequest true "request body"
// @Success 200 {object} dto.ListCreateResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /list [post]
func (h *handler) Create(c echo.Context) error {
	cc := c.(*abstraction.Context)
	payload := new(dto.ListCreateRequest)
	
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	
	path := "./data_upload/listfile/"
	var datafile []*model.ListfileEntity
	
	form, err := cc.MultipartForm()
	if err != nil {
		res.ErrorResponse(err).Send(c)
	}
	files := form.File["listfile"]

	if len(files) > 0 {
		for _, file := range files {
			fmt.Println("Filename :", file.Filename)
			fmt.Println("Size :", file.Size)
			fmt.Println("Type :", file.Header.Get("Content-Type"))

			if file.Header.Get("Content-Type") == "text/plain" || file.Header.Get("Content-Type") == "application/pdf" {
				datax := &model.ListfileEntity{Filename: file.Filename}
				datafile = append(datafile, datax)

				src, err := file.Open()
				if err != nil {
					res.ErrorResponse(err).Send(c)
				}
				defer src.Close()

				dst, err := os.OpenFile(path+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					res.ErrorResponse(err).Send(c)
				}
				defer dst.Close()

				if _, err = io.Copy(dst, src); err != nil {
					res.ErrorResponse(err).Send(c)
				}

			} else {
				res.ErrorResponse(err).Send(c)
			}
		}

		result, err := h.service.CreateWithFile(cc, payload, datafile)
		if err != nil {
			return res.ErrorResponse(err).Send(c)
		}
	    return res.SuccessResponse(result).Send(c)
	} else {
		result, err := h.service.Create(cc, payload)
		if err != nil {
			return res.ErrorResponse(err).Send(c)
		}
	 
	 return res.SuccessResponse(result).Send(c)
	}
}

// Update list
// @Summary Update list
// @Description Update list
// @Tags list
// @Accept  json
// @Produce  json
// @Param id path int true "id path"
// @Param request body dto.ListUpdateRequest true "request body"
// @Success 200 {object} dto.ListUpdateResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /list/{id} [put]
func (h *handler) Update(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ListUpdateRequest)
	if err := c.Bind(&payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Update(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete list
// @Summary Delete list
// @Description Delete list
// @Tags list
// @Accept  json
// @Produce  json
// @Param id path int true "id path"
// @Success 200 {object}  dto.ListDeleteResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /list/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ListDeleteRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Delete(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}