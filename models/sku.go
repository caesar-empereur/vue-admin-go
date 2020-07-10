package models

type Sku struct {
	Id string `orm:"column(id);pk" json:"id"`

	Name string `orm:"column(name)" json:"name"`

	Price string `orm:"column(price)" json:"price"`

	Spec string `orm:"column(spec)" json:"spec"`

	ShipFee int `orm:"column(ship_fee)" json:"ship_fee"`

	Count int `orm:"column(count)" json:"count"`

	ImageUrl string `orm:"column(image_url)" json:"image_url"`
}

func (sku *Sku) TableName() string {
	return "t_sku"
}
