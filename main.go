package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name`
	Email     string `json:"email"`
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@/golang_auth?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database!")
	}

	database.AutoMigrate(&User{})

	DB = database
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	DB.Create(&user)

	return c.JSON(user)
}

func main() {
	//Connect to Database!
	ConnectDatabase()

	app := fiber.New()
	//Route to Register user
	app.Get("/register", Register)

	app.Listen(":7000")

}
