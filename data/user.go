package data
import (
		"gorm.io/gorm"
  	//"gorm.io/driver/sqlite"
		"github.com/Rogstaa/gobux/database"

)


 type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func GetProduct() (Product, error) {
  db := database.Database.Db
  var product Product
  if err := db.First(&product, 1).Error; err != nil {
    return Product{}, err
  }
  return product, nil
}

func InsertProduct(code string, price uint) error {
  db := database.Database.Db
  newProduct := Product{
    Code:  code,
    Price: price,
  }
  return db.Create(&newProduct).Error
}

/*
func GetProduct() {
  	db := database.Database.Db
	var product Product
  	db.First(&product, 1)
	println(product.Code)
}

*/
