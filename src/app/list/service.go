package list

import (
	"errors"
	"todoapp/src/abstraction"
	"todoapp/src/dto"
	"todoapp/src/factory"
	"todoapp/src/model"
	"todoapp/src/repository"
	"todoapp/pkg/util/trxmanager"
	res "todoapp/pkg/util/response"

	"gorm.io/gorm"
)

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.ListGetRequest) (*dto.ListGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.ListGetByIDRequest) (*dto.ListGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.ListCreateRequest) (*dto.ListCreateResponse, error)
	CreateWithFile(ctx *abstraction.Context, payload *dto.ListCreateRequest, mf *[]model.ListfileEntity) (*dto.ListCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.ListUpdateRequest) (*dto.ListUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.ListDeleteRequest) (*dto.ListDeleteResponse, error)
}

type service struct {
	Repository repository.List
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.ListRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.ListGetRequest) (*dto.ListGetResponse, error) {
	var result *dto.ListGetResponse
	var datas *[]model.ListEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.ListFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.ListGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.ListGetByIDRequest) (*dto.ListGetByIDResponse, error) {
	var result *dto.ListGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.ListGetByIDResponse{
		ListEntityModel: *data,
	}

	return result, nil
}

func (s *service) CreateWithFile(ctx *abstraction.Context, payload *dto.ListCreateRequest, mf []*model.ListfileEntity) (*dto.ListCreateResponse, error) {
	var result *dto.ListCreateResponse
	var data *model.ListEntityModel
	var datafile []model.ListfileEntityModel

	data = &model.ListEntityModel{}
	for i := 0; i < len(mf); i++ {
		datax := model.ListfileEntityModel{Context: ctx, ListfileEntity: *mf[i]}
		datafile = append(datafile, datax)
	}

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.ListEntity = payload.ListEntity
		data.Listfile = datafile
		data, err = s.Repository.Create(ctx, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.ListCreateResponse{
		ListEntityModel: *data,
	}


	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.ListCreateRequest) (*dto.ListCreateResponse, error) {
	var result *dto.ListCreateResponse
	var data *model.ListEntityModel

	data = &model.ListEntityModel{}	
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.ListEntity = payload.ListEntity
		data, err = s.Repository.Create(ctx, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.ListCreateResponse{
		ListEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.ListUpdateRequest) (*dto.ListUpdateResponse, error) {
	var result *dto.ListUpdateResponse
	var data *model.ListEntityModel

	data = &model.ListEntityModel{}
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}
        
		data.Context = ctx
		data.ListEntity = payload.ListEntity
		data, err = s.Repository.Update(ctx, &payload.ID, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ListUpdateResponse{
		ListEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.ListDeleteRequest) (*dto.ListDeleteResponse, error) {
	var result *dto.ListDeleteResponse
	var data *model.ListEntityModel

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}

		data.Context = ctx
		data, err = s.Repository.Delete(ctx, &payload.ID, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ListDeleteResponse{
		ListEntityModel: *data,
	}

	return result, nil
}