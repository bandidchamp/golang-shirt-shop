package models

var ProductTname string = "product"
var SizeTname string = ProductTname + "_size"
var GenderTname string = ProductTname + "_gender"
var CatagoryTname string = ProductTname + "_catagory"

type Product struct {
	Id        int     `json:"id" gorm:"size:64"`
	Name      string  `json:"name" gorm:"size:4096" validate:"required"`
	Catagory  int     `json:"catagory" gorm:"size:64" validate:"required"`
	Size      int     `json:"size" gorm:"size:64" validate:"required"`
	Gender    int     `json:"gender" gorm:"size:64" validate:"required"`
	Price     float32 `json:"price" gorm:"size:64" validate:"required"`
	Quantiry  float32 `json:"quantity" gorm:"size:64" validate:"required"`
	Ispadding bool    `json:"ispadding" gorm:"size:1"`
}

type ProductForm struct {
	Name      string `json:"name"`
	Catagory  string `json:"catagory"`
	Size      string `json:"size"`
	Gender    string `json:"gender"`
	Price     string `json:"price"`
	Quantiry  string `json:"quantity"`
	Ispadding string `json:"ispadding"`
}

type Product_size struct {
	Product_size_id   int    `json:"product_size_id" gorm:"size:64"`
	Product_size_name string `json:"product_size_name" gorm:"size:2048"`
}
type Product_catagory struct {
	Product_catagory_id   int    `json:"product_catagory_id" gorm:"size:64"`
	Product_catagory_name string `json:"product_catagory_name" gorm:"size:2048"`
}
type Product_gender struct {
	Product_gender_id   int    `json:"product_gender_id" gorm:"size:64"`
	Product_gender_name string `json:"product_gender_name" gorm:"size:2048"`
}

type ProductFilter struct {
	Name     string `json:"name"`
	Catagory string `json:"catagory"`
	Size     string `json:"size"`
	Gender   string `json:"gender"`
	Limit    string `json:"limit"`
	Offset   string `json:"offset"`
	OrderBy  string `json:"orderby"`
}

func (product *Product) TableName() string {
	return "product"
}
