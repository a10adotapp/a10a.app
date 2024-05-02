package changokushi

import (
	"context"
	"log/slog"

	"github.com/a10adotapp/a10a.app/ent"
	entchangokushiweapon "github.com/a10adotapp/a10a.app/ent/changokushiweapon"
	entchangokushiweaponchangelog "github.com/a10adotapp/a10a.app/ent/changokushiweaponchangelog"
	"github.com/a10adotapp/a10a.app/service/changokushi/kusogeeeeeeapi"
)

var page = 1

func (s *ChangokushiService) GetWeapons(ctx context.Context) error {
	apiClient := kusogeeeeeeapi.NewClient()

	nfts, pagination, err := apiClient.NFTs(ctx, page)
	if err != nil {
		slog.ErrorContext(ctx, "failed to fetch nfts", slog.Any("err", err))

		return err
	}

	if len(nfts) == 0 {
		slog.InfoContext(ctx, "no items on the page", slog.Any("page", pagination))

		page = 1

		return nil
	}

	for _, nft := range nfts {
		changokushiWeapon, err := s.entClient.ChangokushiWeapon.Query().
			Where(
				entchangokushiweapon.URI(nft.ID),
			).
			First(ctx)
		if ent.MaskNotFound(err) != nil {
			slog.ErrorContext(ctx, "failed to find changokushi weapon", slog.Any("err", err))

			continue
		}

		if changokushiWeapon == nil {
			changokushiWeapon, err = s.entClient.ChangokushiWeapon.Create().
				SetURI(nft.ID).
				SetType(nft.MetaAttributeString("type")).
				SetName(nft.Name).
				SetRank(nft.MetaAttributeInt("等級")).
				SetType(nft.MetaAttributeString("攻撃種")).
				SetSkill(nft.MetaAttributeString("スキル")).
				SetVitality(nft.MetaAttributeInt("体力")).
				SetStrength(nft.MetaAttributeInt("攻撃")).
				SetPhysicalDefense(nft.MetaAttributeInt("武防")).
				SetMagicalDefense(nft.MetaAttributeInt("知防")).
				SetAgility(nft.MetaAttributeInt("俊敏")).
				Save(ctx)
			if err != nil {
				slog.ErrorContext(ctx, "failed to save changokushi weapon", slog.Any("err", err))

				continue
			}
		}

		changokushiWeaponChangeLog, err := s.entClient.ChangokushiWeaponChangeLog.Query().
			Where(
				entchangokushiweaponchangelog.ChangokushiWeaponID(changokushiWeapon.ID),
			).
			Order(
				ent.Desc(entchangokushiweaponchangelog.FieldPublishedAt),
			).
			First(ctx)
		if ent.MaskNotFound(err) != nil {
			slog.ErrorContext(ctx, "failed to find changokushi weapon change log", slog.Any("err", err))

			continue
		}

		needsCreate := changokushiWeaponChangeLog == nil
		needsCreate = needsCreate || (changokushiWeaponChangeLog.Status != nft.Status)
		needsCreate = needsCreate || (changokushiWeaponChangeLog.Price != nft.Amount)

		if needsCreate {
			if _, err := s.entClient.ChangokushiWeaponChangeLog.Create().
				SetChangokushiWeaponID(changokushiWeapon.ID).
				SetStatus(nft.Status).
				SetPrice(nft.Amount).
				SetPublishedAt(nft.PublishedAt).
				Save(ctx); err != nil {
				slog.ErrorContext(
					ctx,
					"failed to save changokushi weapon change log",
					slog.Any("err", err),
					slog.Any("nft", nft),
				)

				continue
			}
		}
	}

	page = page + 1

	return nil
}
