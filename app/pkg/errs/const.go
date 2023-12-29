package errs

import (
	"errors"
)

var (
	ErrSyntax                  = errors.New("ERR_SYNTAX")           // ErrSyntax - syntax errors
	ErrParseRequest            = errors.New("ERR_PARSE_REQUEST")    // ErrParseRequest - parse request errors
	ErrEmptyPointer            = errors.New("ERR_EMPTY_POINTER_DB") // ErrEmptyPointer - empty point errors
	ErrEmptyData               = errors.New("ERR_EMPTY_DATA")
	ErrCreateUserFailed        = errors.New("ERR_CREATE_USER_FAILED")
	ErrCreateCartItemFailed    = errors.New("ERR_CREATE_CART_ITEM_FAILED")
	ErrCreateOrderFailed       = errors.New("ERR_CREATE_ORDER_FAILED")
	ErrCreateProductFailed     = errors.New("ERR_CREATE_PRODUCT_FAILED")
	ErrGetProductFailed        = errors.New("ERR_GET_PRODUCT_FAILED")
	ErrReduceUserBalanceFailed = errors.New("ERR_REDUCE_USER_BALANCE_FAILED")
	ErrUpdateUser              = errors.New("ERR_UPDATE_USER")
	ErrUpdateCartItems         = errors.New("ERR_UPDATE_CART_ITEMS")
	ErrUpdateProduct           = errors.New("ERR_UPDATE_PRODUCT")
	ErrNotFound                = errors.New("ERR_NOT_FOUND")
	ErrUserNotFound            = errors.New("ERR_USER_NOT_FOUND")
	ErrCartItemsNotFound       = errors.New("ERR_CART_ITEMS_NOT_FOUND")
	ErrQuantityIsNotEnough     = errors.New("ERR_QUANTITY_IS_NOT_ENOUGH")
	ErrInvalidStatus           = errors.New("ERR_INVALID_STATUS")
)
