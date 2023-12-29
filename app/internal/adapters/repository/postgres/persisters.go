package postgres

import (
	"default_ddd/app/internal/adapters/repository/models"
	"default_ddd/app/pkg/dbo"
	"default_ddd/app/pkg/mapping"
	status "default_ddd/app/pkg/vars/statuses"
	"gorm.io/gorm"
)

// CreateUser --
func (ths *SQLStore) CreateUser(userDBO *dbo.UserDBO) error {
	if err := ths.verifyPointer(userDBO); err != nil {
		return err
	}

	userModel := mapping.UserDBOtoModel(userDBO)
	return ths.db.Table(userModel.TableName()).Create(&userModel).Error
}

// CreateCartItem --
func (ths *SQLStore) CreateCartItem(cartItemDBO *dbo.CartItemDBO) error {
	if err := ths.verifyPointer(cartItemDBO); err != nil {
		return err
	}

	cartItemModel := mapping.CartItemDBOtoModel(cartItemDBO)
	return ths.db.Table(cartItemModel.TableName()).Create(&cartItemModel).Error
}

// CreateProduct --
func (ths *SQLStore) CreateProduct(productDBO *dbo.ProductDBO) error {
	if err := ths.verifyPointer(productDBO); err != nil {
		return err
	}

	productModel := mapping.ProductDBOtoModel(productDBO)
	return ths.db.Table(productModel.TableName()).Create(&productModel).Error
}

// UpdateProduct --
func (ths *SQLStore) UpdateProduct(product *dbo.ProductDBO, cartItem *dbo.CartItemDBO) error {
	if err := ths.verifyPointer(product); err != nil {
		return err
	}

	if err := ths.db.Model(models.ProductModel{}).
		Where("uuid = ?", product.UUID).
		Where("quantity >= ?", cartItem.Quantity).
		Updates(map[string]interface{}{
			"quantity":   gorm.Expr("quantity - ?", cartItem.Quantity),
			"updated_at": product.UpdatedAt}).
		Error; err != nil {
		return err
	}

	return nil
}

// CreateOrder --
func (ths *SQLStore) CreateOrder(orderDBO *dbo.OrderDBO) error {
	if err := ths.verifyPointer(orderDBO); err != nil {
		return err
	}

	orderModel := mapping.OrderDBOtoModel(orderDBO)
	return ths.db.Table(orderModel.TableName()).Create(orderModel).Error
}

// UpdateUser --
func (ths *SQLStore) UpdateUser(userDBO *dbo.UserDBO) error {
	if err := ths.verifyPointer(userDBO); err != nil {
		return err
	}

	if err := ths.db.Model(models.UserModel{}).
		Where("uuid = ?", userDBO.UUID).
		Updates(map[string]interface{}{
			"balance":    userDBO.Balance,
			"updated_at": userDBO.UpdatedAt}).
		Error; err != nil {
		return err
	}

	return nil
}

// UpdateCartItems --
func (ths *SQLStore) UpdateCartItems(orderDBO *dbo.OrderDBO) error {
	if err := ths.verifyPointer(orderDBO); err != nil {
		return err
	}

	if err := ths.db.Model(models.CartItemModel{}).
		Where("status = ?", status.CartItemProcessing).
		Where("cart_items.uuid IN (?)", orderDBO.CartItemUUIDs).
		Updates(map[string]interface{}{
			"status":     status.CartItemCompleted,
			"updated_at": orderDBO.UpdatedAt}).
		Error; err != nil {
		return err
	}

	return nil
}
