package changokushi

import "github.com/a10adotapp/a10a.app/ent"

type ChangokushiService struct {
	entClient *ent.Client
}

func NewChangokushiService(entClient *ent.Client) *ChangokushiService {
	return &ChangokushiService{
		entClient: entClient,
	}
}
