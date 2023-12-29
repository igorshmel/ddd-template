package status

// OrderStatus --
type OrderStatus string

const (
	// OrderNew - новый заказ
	OrderNew OrderStatus = "new"
	// OrderProcessing - заказ находится в процессе обработки
	OrderProcessing OrderStatus = "processing"
	// OrderCompleted - заказ успешно обработан
	OrderCompleted OrderStatus = "completed"
	// OrderCanceled - ордер отменен
	OrderCanceled OrderStatus = "canceled"
)

func (ths OrderStatus) String() string {
	return string(ths)
}
