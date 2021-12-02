package client

import "go-atividade-a3/address"

func clientToClientDTO(client Client, includeAddress bool) ClientDTO {
	clientDTO := ClientDTO{
		ID: client.ID, Nome: client.Nome}

	if includeAddress {
		address := address.AddressDTO{
			Cidade: client.Endereco.Cidade,
			Uf:     client.Endereco.Uf}

		clientDTO.Endereco = &address
	}

	return clientDTO

}
