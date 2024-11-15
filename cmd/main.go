package main

import (
	"log"
	"sync_btim/config"
	"sync_btim/entity"
	"sync_btim/handlers"
	"sync_btim/routes"
	"sync_btim/service"
	"sync_btim/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	// Memuat variabel environment dari file .env
	utils.LoadEnv()

	// Inisialisasi koneksi ke dua PostgreSQL database
	dbPostgres1, dbPostgres2, err := config.InitDBs()
	if err != nil {
		log.Fatal("Failed to connect to databases: ", err)
	}

	// var DB1_NAME1 = os.Getenv("DB1_NAME")
	controller := initialApp(dbPostgres1, dbPostgres2)

	// Inisialisasi Fiber
	app := fiber.New()
	// app.Use(cors.New())

	routes.APIRoutes(app, controller.TablenameHandler)

	log.Fatal(app.Listen(":3000"))

}

type APP struct {
	TablenameHandler *handlers.TablenameHandler
}

func initialApp(db1 *gorm.DB, db2 *gorm.DB) APP {

	var handler = APP{}

	TNRepo := entity.NewTablesNameRepository(db1, db2)
	TNservice := service.NewUserService(TNRepo)
	TNHandler := handlers.NewTablenameHandler(TNservice)

	handler.TablenameHandler = TNHandler

	return handler
}
