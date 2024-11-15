package routes

import (
	"sync_btim/handlers"

	"github.com/gofiber/fiber/v2"
)

func APIRoutes(app *fiber.App, handler *handlers.TablenameHandler) {

	// app.Get(":source/data/:table_name", handler.MigrationExec)
	// app.Post("source/field/:tablename", tablenameHandler.GetDataInTable)
	app.Post("sync/client_id", handler.SyncClientIDBTIM)
	// app.Post("migration/client", tablenameHandler.MigrationExec)

}
