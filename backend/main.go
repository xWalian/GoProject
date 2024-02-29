package main

import (
	"log"
	"main/src/application/storage"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/xWalian/GoProject/backend/docs"
	"github.com/xWalian/GoProject/backend/src/application/models"
	"github.com/xWalian/GoProject/backend/src/application/storage"
	"gorm.io/gorm"
)

type Order struct {
	Product_id int `json: "product_id"`
	Quantity   int `json: "quantity"`
	Owner      int `json: "owner"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateOrder(context *fiber.Ctx) error {
	order := Order{}
	err := context.BodyParser(&order)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&order).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create order"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book has been added"})
	return nil
}

func (r *Repository) GetOrders(context *fiber.Ctx) error {
	orderModels := &[]models.Orders{}
	err := r.DB.Find(orderModels).Error

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "could not get orders"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "orders fetched successfully",
		"data":    orderModels,
	})
	return nil
}
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-order", r.CreateOrder)
	api.Delete("delete-order", r.DeleteOrder)
	api.Get("/get-orders/:id", r.GetOrderById)
	api.Get("/orders", r.GetOrders)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	r := Repository{
		DB: db,
	}
	config := &storage.Config{
		Host:     os.Gatenv("DB_HOST"),
		Port:     os.Gatenv("DB_PORT"),
		Password: os.Gatenv("DB_PASS"),
		User:     os.Gatenv("DB_USER"),
		SSLMode:  os.Gatenv("DB_SSLMODE"),
		DBName:   os.Gatenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load the database")
	}

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	r.SetupRoutes(app)
	app.Listen(":8080")
}
