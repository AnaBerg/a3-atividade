package client

import (
	"fmt"
	"go-atividade-a3/address"
	"go-atividade-a3/database"
	"go-atividade-a3/validator"

	"github.com/gofiber/fiber"
)

func CreateClient(c *fiber.Ctx) {
	db := database.DBConn

	var client Client
	c.BodyParser(&client)

	if err := validator.ValidateStruct(client); err != nil {
		c.Status(500).JSON(err)
		return
	}

	db.Create(&client)

	c.JSON(clientToClientDTO(client, true))
}

func GetClientsByCity(c *fiber.Ctx) {
	cidade := c.Params("cidade")
	db := database.DBConn

	var addresses []address.Address
	db.Where("Cidade LIKE ?", "%"+cidade+"%").Find(&addresses)

	if len(addresses) == 0 {
		c.JSON(nil)
		return
	}

	var clientsDTO []ClientDTO
	for _, address := range addresses {
		var client Client
		db.Find(&client, address.ClientID)

		clientsDTO = append(clientsDTO, clientToClientDTO(client, false))
	}

	address := addresses[0]
	c.JSON(ClientsByCityDTO{Cidade: address.Cidade, Uf: address.Uf, Clientes: clientsDTO})
}

func GetClients(c *fiber.Ctx) {
	db := database.DBConn

	var clients []Client
	db.Preload("Endereco").Find(&clients)

	var clientsDTO []ClientDTO

	for _, client := range clients {

		clientsDTO = append(clientsDTO, clientToClientDTO(client, true))
	}

	c.JSON(clientsDTO)
}

func DeleteClient(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var client Client
	db.Find(&client, id)

	fmt.Println(client.Nome)

	if client.ID == 0 {
		c.Status(500).Send("No client found with given ID")
		return
	}

	db.Delete(&client, id)

	c.Send("Client successfully deleted")

}
