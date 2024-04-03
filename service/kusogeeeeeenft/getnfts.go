package kusogeeeeeenft

import (
	"context"
	"log/slog"

	"github.com/a10adotapp/a10a.app/ent"
	entkusogeeeeeenft "github.com/a10adotapp/a10a.app/ent/kusogeeeeeenft"
	"github.com/a10adotapp/a10a.app/service/kusogeeeeeenft/kusogeeeeeeapi"
)

func (s *KusogeeeeeeNFTService) GetNFTs(ctx context.Context) error {
	apiClient := kusogeeeeeeapi.NewClient()

	page := 1

	for {
		nfts, pagination, err := apiClient.NFTs(ctx, page)
		if err != nil {
			return err
		}

		slog.InfoContext(
			ctx, "[KusogeeeeeeNFTService.GetNFTs] data fetched",
			slog.Any("page", pagination),
		)

		if len(nfts) == 0 {
			break
		}

		page = page + 1

		for _, nft := range nfts {
			kusogeeeeeeNFT, err := s.entClient.KusogeeeeeeNFT.Query().
				Where(
					entkusogeeeeeenft.URI(nft.ID),
				).
				First(ctx)
			if ent.MaskNotFound(err) != nil {
				return err
			}

			if kusogeeeeeeNFT != nil {
				isUpdated := false
				isUpdated = isUpdated || (nft.Status != kusogeeeeeeNFT.Status)
				isUpdated = isUpdated || (nft.Amount != kusogeeeeeeNFT.Price)

				if !isUpdated {
					slog.InfoContext(
						ctx, "[KusogeeeeeeNFTService.GetNFTs] may be duplicated",
						slog.Any("saved_nft", kusogeeeeeeNFT),
						slog.Any("nft_data", nft),
					)

					return nil
				}
			}

			if kusogeeeeeeNFT == nil {
				kusogeeeeeeNFT, err = s.entClient.KusogeeeeeeNFT.Create().
					SetURI(nft.ID).
					SetType(nft.MetaAttributeString("type")).
					SetName(nft.Name).
					SetStatus(nft.Status).
					SetPrice(nft.Amount).
					SetPublishedAt(nft.PublishedAt).
					Save(ctx)
				if err != nil {
					return err
				}
			}

			kusogeeeeeeNFT, err = s.entClient.KusogeeeeeeNFT.UpdateOne(kusogeeeeeeNFT).
				SetStatus(nft.Status).
				SetPrice(nft.Amount).
				SetPublishedAt(nft.PublishedAt).
				SetNillableWeaponRank(nft.MetaAttributeNillableInt("等級")).
				SetNillableWeaponType(nft.MetaAttributeNillableString("攻撃種")).
				SetNillableWeaponVitality(nft.MetaAttributeNillableInt("体力")).
				SetNillableWeaponStrength(nft.MetaAttributeNillableInt("攻撃")).
				SetNillableWeaponPhysicalDefense(nft.MetaAttributeNillableInt("武防")).
				SetNillableWeaponMagicalDefense(nft.MetaAttributeNillableInt("知防")).
				SetNillableWeaponAgility(nft.MetaAttributeNillableInt("俊敏")).
				SetNillableCharacterRank(nft.MetaAttributeNillableString("rank")).
				SetNillableCharacterTotalSupply(nft.MetaAttributeNillableInt("total_supply")).
				Save(ctx)
			if err != nil {
				return nil
			}

			_, err = s.entClient.KusogeeeeeeNFTChangeLog.Create().
				SetKusogeeeeeeNftID(kusogeeeeeeNFT.ID).
				SetStatus(kusogeeeeeeNFT.Status).
				SetPrice(kusogeeeeeeNFT.Price).
				Save(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
