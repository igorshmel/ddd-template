package mapping

import (
	"default_ddd/app/internal/adapters/repository/models"
	"default_ddd/app/pkg/dbo"
	"default_ddd/app/pkg/ddo"
	"default_ddd/app/pkg/dto"
	status "default_ddd/app/pkg/vars/statuses"
)

// OrderDDOtoDBO --
func OrderDDOtoDBO(order *ddo.OrderDDO) *dbo.OrderDBO {
	return &dbo.OrderDBO{
		UUID:          order.UUID,
		UserUUID:      order.UserUUID,
		TotalPrice:    order.TotalPrice,
		CartItemUUIDs: order.CartItemUUIDs,
		Status:        status.OrderStatus(order.Status),
		UpdatedAt:     order.UpdatedAt,
		CreatedAt:     order.CreatedAt,
	}
}

// OrderDBOtoModel --
func OrderDBOtoModel(order *dbo.OrderDBO) *models.OrderModel {
	return &models.OrderModel{
		UUID:          order.UUID,
		UserUUID:      order.UserUUID,
		TotalPrice:    order.TotalPrice,
		CartItemUUIDs: order.CartItemUUIDs,
		Status:        string(order.Status),
		UpdatedAt:     order.UpdatedAt,
		CreatedAt:     order.CreatedAt,
	}
}

// OrderDTOtoDDO --
func OrderDTOtoDDO(dtoCreateOrder *dto.CreateOrderRequest) *ddo.OrderDDO {
	return &ddo.OrderDDO{
		UserUUID:      dtoCreateOrder.UserUUID,
		CartItemUUIDs: dtoCreateOrder.CartItemUUIDs,
	}
}
