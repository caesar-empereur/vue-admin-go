package models

type Order struct {
	Id string `orm:"column(id);pk" json:"id"`

	OrderNo string `orm:"column(order_no)" json:"orderNo"`

	BuyerName string `orm:"column(buyer_name)" json:"buyerName"`

	SellerName string `orm:"column(seller_name)" json:"sellerName"`

	SkuId string `orm:"column(sku_id)" json:"skuId"`

	SkuName string `orm:"column(sku_name)" json:"skuName"`

	Amount int `orm:"column(amount)" json:"amount"`

	Status string `orm:"column(status)" json:"status"`
}

func (order *Order) TableName() string {
	return "t_order"
}
