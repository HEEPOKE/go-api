package controller

import (
	"Backend/go-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/Shirtgo?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	db.Create(&model.User{FirstName: "D42", LastName: "a"})

	var user model.User
	db.First(&user, 1)
	db.First(&user, "code = ?", "D42")

	db.Model(&user).Update("Price", 200)

	db.Model(&user).Updates(model.User{FirstName: "D42", LastName: "a"})
	db.Model(&user).Updates(map[string]interface{}{"Firstname": 200, "Lastname": "F42"})

	db.Delete(&user, 1)
}