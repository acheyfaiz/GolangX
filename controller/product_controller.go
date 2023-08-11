package Controller

import (
	db "golang-exercise/config"
	"golang-exercise/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetListAllProduct(c *fiber.Ctx) error {

	// to set limit/pagination
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	// way to assign an array
	var list []model.ProductInMinishops

	// var first model.ProductInMinishops

	// to get first record only
	//result := db.DB.Find(&list)

	// to get all record
	// db.DB.Select("*").Find(&list).Count(&count)

	// get record by limit (pagination)
	db.DB.Select("*").Limit(limit).Offset(skip).Find(&list).Count(&count)

	// query start
	// data := db.DB.Select("*").Limit(10).Offset(skip).Find(&list).Count(&count)
	// if data != nil {
	// 	fmt.Println(data)
	// }

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"message": "List All Product in Minishop",
		"data": list,
	})

}

func GetListProductByPid(c *fiber.Ctx) error {

	// to get the request
	pid := c.Params("pid")

	// to declare model slice to store multiple products
	var pidModel []model.ProductInMinishops

	db.DB.Select("pid, productid, itemsold, dateapply, status").
		Where("pid=?", pid).
		Order("dateapply DESC").
		Find(&pidModel)

	// idk what is len :v
	if len(pidModel) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": false,
			"message": "No products found for the given pid",
			"error": map[string]interface{}{},
		})
	}

	// to store the product data
	productDataList := make([]map[string]interface{}, 0)

	for _, product := range pidModel {
		productData := map[string]interface{}{
			"pid": product.Pid,
			"productid": product.Productid,
			"itemsold": product.Itemsold,
			"dateapply": product.Dateapply,
			"statusproduct": product.Status,
		}
		productDataList = append(productDataList, productData)
	}

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"message": "success",
		"data": productDataList,
	})

}



func GetListProduct(c *fiber.Ctx) error {
	return nil
}