package kusogeeeeeenft

import "github.com/a10adotapp/a10a.app/ent"

type KusogeeeeeeNFTService struct {
	entClient *ent.Client
}

func NewKusogeeeeeeNFTService(entClient *ent.Client) *KusogeeeeeeNFTService {
	return &KusogeeeeeeNFTService{
		entClient: entClient,
	}
}
