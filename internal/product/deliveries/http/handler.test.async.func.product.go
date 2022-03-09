package http

import (
	"fmt"
	"shirt-shop/internal/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func aaa1(add int, age chan int) {
	time.Sleep(4 * time.Second)
	tmepage := 35
	age <- tmepage + add // this is return main func.
	fmt.Println("async function message finish,.,.1")
}
func aaa2(add int, age chan int) {
	time.Sleep(2 * time.Second)
	tmepage := 30
	age <- tmepage + add // this is return main func.
	fmt.Println("async function message finish,.,.2")
}
func testThread(i int) {
	time.Sleep(3 * time.Second)
	fmt.Println("async function testThread :", i)
}

func FindProduct(phh *productHandler, id string, c chan models.Product, i int) {

	fmt.Println("query product", i)
	productS, err := phh.productUC.CheckProductID(id)
	if err != nil {
		fmt.Println(err)
	}
	c <- *productS
}

func (ph *productHandler) AsyncFunc(c *fiber.Ctx) error {
	payload := struct {
		Ids []int `json:"ids"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	productChannel := make(chan models.Product)

	var resultproductChannel []models.Product

	for i, val := range payload.Ids {
		// go func(i int, val int) {
		// 	fmt.Println(val)
		// 	productS, err := ph.productUC.CheckProductID(strconv.FormatInt(int64(val), 10))
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	resultproductChannel = append(resultproductChannel, *productS)
		// }(i, val)
		// fmt.Println("loop print ", i)
		go FindProduct(ph, strconv.FormatInt(int64(val), 10), productChannel, i)
		// // fmt.Println(<-prodmap)
		resultproductChannel = append(resultproductChannel, <-productChannel)
	}

	// fmt.Println("-----------------------")
	// fmt.Println(resultproductChannel)
	// fmt.Println("-----------------------")

	// age1 := make(chan int)
	// age2 := make(chan int)

	// fmt.Println("start async function 1 2")

	// go aaa1(5, age1)
	// go aaa2(6, age2)

	// for i := 1; i < 10; i++ {
	// 	// multithreading
	// 	go testThread(i)
	// }

	// // wait until return
	// proson_age2 := <-age2
	// proson_age1 := <-age1
	// fmt.Println("wait until aaa1 and aaa2 return")
	// fmt.Println("proson age : ", proson_age1)
	// fmt.Println("proson age : ", proson_age2)

	// product := make(chan map[int]int)
	// // product

	// go func() {
	// 	mappy := make(map[int]int)
	// 	for i := 0; i < 10; i++ {
	// 		mappy[i] = i + 10
	// 	}
	// 	product <- mappy
	// 	close(product)
	// }()

	// products := <-product
	// fmt.Println("Product channel ", products)
	// for pd := range products {
	// 	fmt.Println("Product : ", pd)
	// }
	return c.Status(200).JSON(fiber.Map{"result": resultproductChannel})
}
