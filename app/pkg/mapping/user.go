package mapping

import (
	"default_ddd/app/internal/adapters/repository/models"
	"default_ddd/app/pkg/dbo"
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/dto"
)

// UserDTOtoDDO --
func UserDTOtoDDO(req *dto.CreateUserRequest) *ddo.UserDDO {
	return &ddo.UserDDO{
		Name: req.UserName,
	}
}

// UserDDOtoDBO --
func UserDDOtoDBO(ddo *ddo.UserDDO) *dbo.UserDBO {
	return &dbo.UserDBO{
		UUID:      ddo.UUID,
		Name:      ddo.Name,
		Balance:   ddo.Balance,
		UpdatedAt: ddo.UpdatedAt,
		CreatedAt: ddo.CreatedAt,
	}
}

// UserDBOtoDDO --
func UserDBOtoDDO(userDBO *dbo.UserDBO) *ddo.UserDDO {
	return &ddo.UserDDO{
		UUID:      userDBO.UUID,
		Name:      userDBO.Name,
		Balance:   userDBO.Balance,
		UpdatedAt: userDBO.UpdatedAt,
		CreatedAt: userDBO.CreatedAt,
	}
}

// UserDBOtoModel --
func UserDBOtoModel(dbo *dbo.UserDBO) *models.UserModel {
	return &models.UserModel{
		UUID:      dbo.UUID,
		Name:      dbo.Name,
		Balance:   dbo.Balance,
		UpdatedAt: dbo.UpdatedAt,
		CreatedAt: dbo.CreatedAt,
	}
}

// UserModelToDBO --
func UserModelToDBO(model *models.UserModel) *dbo.UserDBO {
	return &dbo.UserDBO{
		UUID:      model.UUID,
		Name:      model.Name,
		Balance:   model.Balance,
		UpdatedAt: model.UpdatedAt,
		CreatedAt: model.CreatedAt,
	}
}
