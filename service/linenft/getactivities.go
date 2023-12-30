package linenft

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/a10adotapp/a10a.app/ent"
	entlinenft "github.com/a10adotapp/a10a.app/ent/linenft"
	entlinenftactivity "github.com/a10adotapp/a10a.app/ent/linenftactivity"
)

type Activity struct {
	ActivityType string  `json:"activityType"`
	SaleID       uint32  `json:"saleId"`
	SaleType     string  `json:"saleType"`
	CurrencyType string  `json:"currencyType"`
	Price        float32 `json:"price"`
	NFTID        uint32  `json:"nftId"`
	CreatedAt    uint64  `json:"createdAt"`
}

type SaleData struct {
	Token struct {
		ContractID  string `json:"contractId"`
		TokenType   string `json:"tokenType"`
		Description string `json:"nftDescription"`
	} `json:"goods"`
	Item struct {
		TokenIndex    string             `json:"tokenIndex"`
		Name          string             `json:"name"`
		ContentURL    string             `json:"improContentsUrl"`
		RawProperties map[string]*string `json:"properties"`
		Properties    struct {
			Series       *string
			OMJ          *string
			GearCategory *string
			GearRarity   *string
		} `json:"-"`
	} `json:"nft"`
	CurrencyType string  `json:"currencyType"`
	Price        float32 `json:"price"`
}

func (s LINENFTService) GetActivities(ctx context.Context) error {
	activities, err := fetchActivities()
	if err != nil {
		return err
	}

	for _, activity := range activities {
		lineNFT, err := s.entClient.LINENFT.Query().
			Where(
				entlinenft.LineNftID(activity.NFTID),
			).
			First(ctx)
		if ent.MaskNotFound(err) != nil {
			slog.Error("GetActivities", slog.Any("error", err))

			continue
		}

		if lineNFT == nil {
			saleData, err := fetchSaleData(activity.SaleID)
			if err != nil {
				slog.Error("GetActivities", slog.Any("error", err))

				continue
			}

			lineNFT, err = s.entClient.LINENFT.Create().
				SetLineNftID(activity.NFTID).
				SetContractID(saleData.Token.ContractID).
				SetTokenType(saleData.Token.TokenType).
				SetTokenIndex(saleData.Item.TokenIndex).
				SetTokenName(saleData.Item.Name).
				SetTokenDescription(saleData.Token.Description).
				SetTokenContentURL(saleData.Item.ContentURL).
				Save(ctx)
			if err != nil {
				slog.Error("GetActivities", slog.Any("error", err))

				continue
			}

			_, err = s.entClient.LINENFTMillionArthursProperty.Create().
				SetLineNft(lineNFT).
				SetNillableSeries(saleData.Item.Properties.Series).
				SetNillableGearCategory(saleData.Item.Properties.GearCategory).
				SetNillableGearRarity(saleData.Item.Properties.GearRarity).
				SetNillableOmj(saleData.Item.Properties.OMJ).
				Save(ctx)
			if err != nil {
				slog.Error("GetActivities", slog.Any("error", err))

				continue
			}
		}

		lineNFTActivity, err := s.entClient.LINENFTActivity.Query().
			Where(
				entlinenftactivity.ActivityType(activity.ActivityType),
				entlinenftactivity.SaleID(activity.SaleID),
			).
			First(ctx)
		if ent.MaskNotFound(err) != nil {
			slog.Error("GetActivities", slog.Any("error", err))

			continue
		}

		if lineNFTActivity != nil {
			continue
		}

		activatedAt := time.Unix(int64(activity.CreatedAt/1000), int64(activity.CreatedAt%1000))

		_, err = s.entClient.LINENFTActivity.Create().
			SetLineNft(lineNFT).
			SetActivityType(activity.ActivityType).
			SetSaleID(activity.SaleID).
			SetSaleType(activity.SaleType).
			SetCurrencyType(activity.CurrencyType).
			SetPrice(activity.Price).
			SetActivatedAt(activatedAt).
			Save(ctx)
		if err != nil {
			slog.Error("GetActivities", slog.Any("error", err))

			continue
		}
	}

	return nil
}

func fetchActivities() ([]Activity, error) {
	res, err := http.Get("https://nft.line.me/api/v1/notice/activity?limit=100")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resData := []Activity{}

	if err = json.Unmarshal(resBody, &resData); err != nil {
		return nil, err
	}

	return resData, nil
}

func fetchSaleData(saleID uint32) (*SaleData, error) {
	res, err := http.Get(fmt.Sprintf("https://nft.line.me/api/v1/marketplace/%d", saleID))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resData := SaleData{}

	if err = json.Unmarshal(resBody, &resData); err != nil {
		return nil, err
	}

	for key, value := range resData.Item.RawProperties {
		switch key {
		case "🍭シリーズ":
			resData.Item.Properties.Series = value
		case "💰OMJ":
			resData.Item.Properties.OMJ = value
		case "⚙ギアカテゴリ":
			resData.Item.Properties.GearCategory = value
		case "⚙ギアレアリティ":
			resData.Item.Properties.GearRarity = value
		}
	}

	return &resData, nil
}
