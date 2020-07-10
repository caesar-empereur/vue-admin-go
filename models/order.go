package models

type Order struct {
	Id string `orm:"column(id);pk" json:"id"`

	OrderNo string `orm:"column(order_no)" json:"order_no"`

	BuyerName string `orm:"column(buyer_name)" json:"buyer_name"`

	SellerName string `orm:"column(seller_name)" json:"seller_name"`

	SkuId string `orm:"column(sku_id)" json:"sku_id"`

	SkuName string `orm:"column(sku_name)" json:"sku_name"`

	Amount int `orm:"column(amount)" json:"amount"`

	Status string `orm:"column(status)" json:"status"`
}

func (order *Order) TableName() string {
	return "t_order"
}
