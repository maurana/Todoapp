package sublist

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
	Find(ctx *abstraction.Context, payload *dto.SublistGetRequest) (*dto.SublistGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.SublistGetByIDRequest) (*dto.SublistGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.SublistCreateRequest) (*dto.SublistCreateResponse, error)
	CreateWithFile(ctx *abstraction.Context, payload *dto.SublistCreateRequest, mf *[]model.SublistfileEntity) (*dto.SublistCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.SublistUpdateRequest) (*dto.SublistUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.SublistDeleteRequest) (*dto.SublistDeleteResponse, error)
}

type service struct {
	Repository repository.Sublist
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.SublistRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.SublistGetRequest) (*dto.SublistGetResponse, error) {
	var result *dto.SublistGetResponse
	var datas *[]model.SublistEntityModel

	datas, info, err := s.Repository.Find(ctx, &payload.SublistFilterModel, &payload.Pagination)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.SublistGetResponse{
		Datas:          *datas,
		PaginationInfo: *info,
	}

	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.SublistGetByIDRequest) (*dto.SublistGetByIDResponse, error) {
	var result *dto.SublistGetByIDResponse

	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.SublistGetByIDResponse{
		SublistEntityModel: *data,
	}

	return result, nil
}

func (s *service) CreateWithFile(ctx *abstraction.Context, payload *dto.SublistCreateRequest, mf []*model.SublistfileEntity) (*dto.SublistCreateResponse, error) {
	var result *dto.SublistCreateResponse
	var data *model.SublistEntityModel
	var datafile []model.SublistfileEntityModel
	var list_id int = payload.ID

	data = &model.SublistEntityModel{}
	for i := 0; i < len(mf); i++ {
		datax := model.SublistfileEntityModel{Context: ctx, SublistfileEntity: *mf[i]}
		datafile = append(datafile, datax)
	}

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.SublistEntity = payload.SublistEntity
		data.ListId = list_id
		data.Sublistfile = datafile
		data, err = s.Repository.Create(ctx, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.SublistCreateResponse{
		SublistEntityModel: *data,
	}


	return result, nil
}

func (s *service) Create(ctx *abstraction.Context, payload *dto.SublistCreateRequest) (*dto.SublistCreateResponse, error) {
	var result *dto.SublistCreateResponse
	var data *model.SublistEntityModel
	var list_id int = payload.ID

	data = &model.SublistEntityModel{}	
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data.Context = ctx
		data.SublistEntity = payload.SublistEntity
		data.ListId = list_id
		data, err = s.Repository.Create(ctx, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err

	}

	result = &dto.SublistCreateResponse{
		SublistEntityModel: *data,
	}

	return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.SublistUpdateRequest) (*dto.SublistUpdateResponse, error) {
	var result *dto.SublistUpdateResponse
	var data *model.SublistEntityModel

	data = &model.SublistEntityModel{}
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		_, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err)
		}
        
		data.Context = ctx
		data.SublistEntity = payload.SublistEntity
		data, err = s.Repository.Update(ctx, &payload.ID, data)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.SublistUpdateResponse{
		SublistEntityModel: *data,
	}

	return result, nil
}

func (s *service) Delete(ctx *abstraction.Context, payload *dto.SublistDeleteRequest) (*dto.SublistDeleteResponse, error) {
	var result *dto.SublistDeleteResponse
	var data *model.SublistEntityModel

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

	result = &dto.SublistDeleteResponse{
		SublistEntityModel: *data,
	}

	return result, nil
}