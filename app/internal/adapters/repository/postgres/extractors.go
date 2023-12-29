package postgres

import (
	"default_ddd/app/internal/adapters/repository/models"
	"default_ddd/app/pkg/dbo"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/mapping"
	"strings"
)

// GetUserByUUID возвращает сущность User из БД--
func (ths *SQLStore) GetUserByUUID(userUUID string) (*dbo.UserDBO, error) {
	if len(strings.TrimSpace(userUUID)) == 0 {
		return nil, errs.ErrEmptyData
	}

	userModel := models.UserModel{}
	resDB := ths.db.Find(&userModel, "uuid", userUUID)
	if err := resDB.Error; err != nil {
		return nil, err
	}
	if resDB.RowsAffected == 0 {
		return nil, errs.ErrNotFound
	}

	return mapping.UserModelToDBO(&userModel), nil
}

// GetProductByUUID возвращает сущность Product из БД--
func (ths *SQLStore) GetProductByUUID(productDBO *dbo.ProductDBO) (*dbo.ProductDBO, error) {
	if err := ths.verifyPointer(productDBO); err != nil {
		return nil, err
	}

	productModel := mapping.ProductDBOtoModel(productDBO)
	resDB := ths.db.Find(&productModel, "uuid", productDBO.UUID)
	if err := resDB.Error; err != nil {
		return nil, err
	}
	if resDB.RowsAffected == 0 {
		return nil, errs.ErrNotFound
	}

	return mapping.ProductModelToDBO(productModel), nil
}

// GetCartItemsByUUIDs возвращает агрегированный список CardItem включая общую сумму TotalPrice по CardItemUUIDs --
func (ths *SQLStore) GetCartItemsByUUIDs(cartItemUUIDs []string) ([]dbo.CartItemDBO, error) {
	if err := ths.verifyPointer(cartItemUUIDs); err != nil {
		return nil, err
	}

	var cartItemsModel []models.CartItemModel

	resDB := ths.db.Table("cart_items").
		Where("cart_items.uuid IN (?)", cartItemUUIDs).
		Scan(&cartItemsModel)
	if err := resDB.Error; err != nil {
		return nil, err
	}
	if resDB.RowsAffected == 0 {
		return nil, errs.ErrNotFound
	}

	return mapping.CartItemsModelToDBO(cartItemsModel), nil
}
