package requests

type OrdersItemScheduleLine struct {
	OrderID                               int      `json:"OrderID"`
	OrderItem                             int      `json:"OrderItem"`
	ScheduleLine                          int      `json:"ScheduleLine"`
	RequestedDeliveryDate                 string   `json:"RequestedDeliveryDate"`
	ConfirmedDeliveryDate                 string   `json:"ConfirmedDeliveryDate"`
	OriginalOrderQuantityInBaseUnit       float32  `json:"OriginalOrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheck float32  `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	OpenConfirmedQuantityInBaseUnit       *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                 *bool    `json:"StockIsFullyConfirmed"`
	ItemScheduleLineDeliveryBlockStatus   *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
}
