package mapping

import (
	"default_ddd/app/internal/adapters/repository/models"
	"default_ddd/app/pkg/dbo"
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/dto"
)

// ProductDTOtoDDO --
func ProductDTOtoDDO(req *dto.CreateProductRequest) *ddo.ProductDDO {
	return &ddo.ProductDDO{
		Title:    req.Title,
		Price:    req.Price,
		Quantity: req.Quantity,
	}
}

// ProductDDOtoDBO --
func ProductDDOtoDBO(ddo *ddo.ProductDDO) *dbo.ProductDBO {
	return &dbo.ProductDBO{
		UUID:      ddo.UUID,
		Title:     ddo.Title,
		Price:     ddo.Price,
		Quantity:  ddo.Quantity,
		UpdatedAt: ddo.UpdatedAt,
		CreatedAt: ddo.CreatedAt,
	}
}

// ProductDBOtoModel --
func ProductDBOtoModel(dbo *dbo.ProductDBO) *models.ProductModel {
	return &models.ProductModel{
		UUID:      dbo.UUID,
		Title:     dbo.Title,
		Price:     dbo.Price,
		Quantity:  dbo.Quantity,
		UpdatedAt: dbo.UpdatedAt,
		CreatedAt: dbo.CreatedAt,
	}
}

// ProductModelToDBO --
func ProductModelToDBO(model *models.ProductModel) *dbo.ProductDBO {
	return &dbo.ProductDBO{
		UUID:      model.UUID,
		Title:     model.Title,
		Price:     model.Price,
		Quantity:  model.Quantity,
		UpdatedAt: model.UpdatedAt,
		CreatedAt: model.CreatedAt,
	}
}

// ProductDBOtoDDO --
func ProductDBOtoDDO(dbo *dbo.ProductDBO) *ddo.ProductDDO {
	return &ddo.ProductDDO{
		UUID:      dbo.UUID,
		Title:     dbo.Title,
		Price:     dbo.Price,
		Quantity:  dbo.Quantity,
		UpdatedAt: dbo.UpdatedAt,
		CreatedAt: dbo.CreatedAt,
	}
}
