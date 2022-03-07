package models

var ProductTname string = "product"
var SizeTname string = ProductTname + "_size"
var GenderTname string = ProductTname + "_gender"
var CatagoryTname string = ProductTname + "_catagory"

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Catagory  int     `json:"catagory"`
	Size      int     `json:"size"`
	Gender    int     `json:"gender"`
	Price     float32 `json:"price"`
	Quantiry  float32 `json:"quantity"`
	Ispadding bool    `json:"ispadding"`
}

type ProductForm struct {
	Name     string `json:"name"`
	Catagory string `json:"catagory"`
	Size     string `json:"size"`
	Gender   string `json:"gender"`
	Price    string `json:"price"`
	Quantiry string `json:"quantity"`
}

type Product_size struct {
	Product_size_id   int    `json:"product_size_id"`
	Product_size_name string `json:"product_size_name"`
}
type Product_catagory struct {
	Product_catagory_id   int    `json:"product_catagory_id"`
	Product_catagory_name string `json:"product_catagory_name"`
}
type Product_gender struct {
	Product_gender_id   int    `json:"product_gender_id"`
	Product_gender_name string `json:"product_gender_name"`
}

func (product *Product) TableName() string {
	return "product"
}
