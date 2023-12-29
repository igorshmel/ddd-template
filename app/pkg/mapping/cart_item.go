package mapping

import (
	"default_ddd/app/internal/adapters/repository/models"
	"default_ddd/app/pkg/dbo"
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/dto"
	status "default_ddd/app/pkg/vars/statuses"
)

// CartItemDTOtoDDO --
func CartItemDTOtoDDO(req *dto.CreateCartItemRequest) *ddo.CartItemDDO {
	return &ddo.CartItemDDO{
		ProductUUID: req.ProductUUID,
		Quantity:    req.Quantity,
	}
}

// CartItemDDOtoDBO --
func CartItemDDOtoDBO(ddo *ddo.CartItemDDO) *dbo.CartItemDBO {
	return &dbo.CartItemDBO{
		UUID:        ddo.UUID,
		ProductUUID: ddo.ProductUUID,
		Quantity:    ddo.Quantity,
		Status:      ddo.Status.String(),
		Price:       ddo.Price,
		UpdatedAt:   ddo.UpdatedAt,
		CreatedAt:   ddo.CreatedAt,
	}
}

// CartItemDBOtoModel --
func CartItemDBOtoModel(dbo *dbo.CartItemDBO) *models.CartItemModel {
	return &models.CartItemModel{
		UUID:        dbo.UUID,
		ProductUUID: dbo.ProductUUID,
		Quantity:    dbo.Quantity,
		Status:      dbo.Status,
		Price:       dbo.Price,
		UpdatedAt:   dbo.UpdatedAt,
		CreatedAt:   dbo.CreatedAt,
	}
}

func CartItemsDBOtoDDO(carts []dbo.CartItemDBO) []ddo.CartItemDDO {
	var ddoCartItem []ddo.CartItemDDO

	for _, cart := range carts {
		ddoCartItem = append(ddoCartItem, ddo.CartItemDDO{
			UUID:        cart.UUID,
			ProductUUID: cart.ProductUUID,
			Quantity:    cart.Quantity,
			Status:      status.CartItemStatus(cart.Status),
			Price:       cart.Price,
			UpdatedAt:   cart.UpdatedAt,
			CreatedAt:   cart.CreatedAt,
		})
	}
	return ddoCartItem
}

// CartItemsModelToDBO --
func CartItemsModelToDBO(cartItems []models.CartItemModel) []dbo.CartItemDBO {

	var dboProducts []dbo.CartItemDBO

	for _, item := range cartItems {
		dboProducts = append(dboProducts, dbo.CartItemDBO{
			UUID:        item.UUID,
			ProductUUID: item.ProductUUID,
			Quantity:    item.Quantity,
			Price:       item.Price,
			Status:      item.Status,
			UpdatedAt:   item.UpdatedAt,
			CreatedAt:   item.CreatedAt,
		})
	}

	return dboProducts
}
