package linenft

import (
	"github.com/a10adotapp/a10a.app/ent"
)

type LINENFTService struct {
	entClient *ent.Client
}

func NewLINENFTService(entClient *ent.Client) LINENFTService {
	return LINENFTService{
		entClient: entClient,
	}
}
