package finschia

import "github.com/a10adotapp/a10a.app/ent"

type FinschiaService struct {
	entClient *ent.Client
}

func NewFinschiaService(entClient *ent.Client) FinschiaService {
	return FinschiaService{
		entClient: entClient,
	}
}
