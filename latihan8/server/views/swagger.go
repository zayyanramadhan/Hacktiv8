package views

import "sesi8-latihan/server/models"

type GetAllPeopleSwagger struct {
	Response
	Payload []models.Person
}

type CreatePeopleSwagger struct {
	ResponseCreate
}

type UpdatePeopleSwagger struct {
	ResponseUpdate
}

type DeletePeopleSwagger struct {
	ResponseDelete
}
