package client

import "github.com/gofiber/fiber"

func Routes(app *fiber.App) {
	app.Get("/clients", GetClients)
	app.Get("/clients/:cidade", GetClientsByCity)
	app.Post("/client", CreateClient)
	app.Delete("/client/:id", DeleteClient)
}
