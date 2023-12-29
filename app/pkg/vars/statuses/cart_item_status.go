package status

// CartItemStatus --
type CartItemStatus string

const (
	// CartItemNew - новая позиция в корзине покупок
	CartItemNew CartItemStatus = "new"
	// CartItemProcessing - позиция находится в процессе обработки
	CartItemProcessing CartItemStatus = "processing"
	// CartItemCompleted - позиция успешно обработана
	CartItemCompleted CartItemStatus = "completed"
	// CartItemCanceled - позиция отменена
	CartItemCanceled CartItemStatus = "canceled"
)

func (ths CartItemStatus) String() string {
	return string(ths)
}
