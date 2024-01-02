package finschia

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/a10adotapp/a10a.app/ent"
	entfinschiaitemtoken "github.com/a10adotapp/a10a.app/ent/finschiaitemtoken"
	entfinschiaitemtokenactivity "github.com/a10adotapp/a10a.app/ent/finschiaitemtokenactivity"
	"github.com/a10adotapp/a10a.app/lib/slices"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

var (
	MillionArthurs = struct {
		BlockchainAddress string
		ContractID        string
	}{
		BlockchainAddress: "link10pgvx8pn5qwgwv066g93jlqux6mnsp0kajg9hp",
		ContractID:        "db8c702e",
	}
)

type Transaction struct {
	FirstMessageType string `json:"first_message_type"`
	Info             struct {
		Hash        string `json:"hash"`
		BlockHeight uint32 `json:"block_height"`
		BlockTime   uint64 `json:"block_time"`
		Position    uint32 `json:"position"`
	} `json:"transaction"`
}

type Message struct {
	Type     string `json:"type"`
	RawValue string `json:"value"`
	Value    MessageValue
}

type MessageValue struct {
	Params []struct {
		TokenType string `json:"tokenType"`
	} `json:"params"`
}

type ItemToken struct {
	Profile struct {
		Name string `json:"item_token_name"`
	} `json:"profile"`
	RawProperties string `json:"meta"`
	Properties    struct {
		Series       *string
		GearCategory *string
		GearRarity   *string
	}
}

func (s FinschiaService) GetMillionArthurs(ctx context.Context) error {
	lastActivity, err := s.entClient.FinschiaItemTokenActivity.Query().
		Order(ent.Desc(entfinschiaitemtokenactivity.FieldActivatedAt)).
		First(ctx)
	if ent.MaskNotFound(err) != nil {
		return err
	}

	transactions, err := fetchTransactions()
	if err != nil {
		return err
	}

	for _, transaction := range transactions {
		time.Sleep(1 * time.Second)

		activatedAt := time.Unix(int64(transaction.Info.BlockTime/1000), int64(transaction.Info.BlockTime%1000))

		if lastActivity != nil && activatedAt.Before(lastActivity.ActivatedAt) {
			return nil
		}

		messages, err := fetchMessages(transaction)
		if err != nil {
			slog.Error("GetMillionArthurs", slog.Any("error", err))

			continue
		}

		for _, message := range messages {
			if message.Type == "/lbm.collection.v1.MsgMintNFT" {
				for _, param := range message.Value.Params {
					itemToken, err := s.entClient.FinschiaItemToken.Query().
						Where(
							entfinschiaitemtoken.ContractID(MillionArthurs.ContractID),
							entfinschiaitemtoken.TokenType(param.TokenType),
						).
						First(ctx)
					if ent.MaskNotFound(err) != nil {
						slog.Error("GetMillionArthurs", slog.Any("error", err))

						continue
					}

					if itemToken == nil {
						itemTokenData, err := fetchItemToken(MillionArthurs.ContractID, param.TokenType)
						if err != nil {
							slog.Error("GetMillionArthurs", slog.Any("error", err))

							continue
						}

						itemToken, err = s.entClient.FinschiaItemToken.Create().
							SetContractID(MillionArthurs.ContractID).
							SetTokenType(param.TokenType).
							SetName(itemTokenData.Profile.Name).
							Save(ctx)
						if err != nil {
							slog.Error("GetMillionArthurs", slog.Any("error", err))

							continue
						}

						_, err = s.entClient.FinschiaItemTokenMillionArthursProperty.Create().
							SetFinschiaItemToken(itemToken).
							SetNillableSeries(itemTokenData.Properties.Series).
							SetNillableGearCategory(itemTokenData.Properties.GearCategory).
							SetNillableGearRarity(itemTokenData.Properties.GearRarity).
							Save(ctx)
						if err != nil {
							slog.Error("GetMillionArthurs", slog.Any("error", err))

							continue
						}
					}

					activity, err := s.entClient.FinschiaItemTokenActivity.Query().
						Where(
							entfinschiaitemtokenactivity.TransactionHash(transaction.Info.Hash),
						).
						First(ctx)
					if ent.MaskNotFound(err) != nil {
						slog.Error("GetMillionArthurs", slog.Any("error", err))

						continue
					}

					if activity == nil {
						_, err = s.entClient.FinschiaItemTokenActivity.Create().
							SetFinschiaItemToken(itemToken).
							SetActivityType(message.Type).
							SetActivatedAt(activatedAt).
							Save(ctx)
						if err != nil {
							slog.Error("GetMillionArthurs", slog.Any("error", err))

							continue
						}
					}
				}
			}
		}
	}

	return nil
}

func fetchTransactions() ([]Transaction, error) {
	requestURL, err := url.Parse("https://explorer.blockchain.line.me")
	if err != nil {
		return nil, err
	}

	requestURL.Path = path.Join(
		path.Join(requestURL.Path, "v2", "finschia-2"),
		path.Join("addresses", MillionArthurs.BlockchainAddress),
		"transactions",
	)

	requestURL.RawQuery = func() string {
		query := url.Values{}

		query.Set("size", "500")
		query.Set("search_from", "top")

		return query.Encode()
	}()

	res, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resData := struct {
		Transactions []Transaction `json:"transactions"`
	}{}

	if err = json.Unmarshal(resBody, &resData); err != nil {
		return nil, err
	}

	return resData.Transactions, nil
}

func fetchMessages(transaction Transaction) ([]Message, error) {
	requestURL, err := url.Parse("https://explorer.blockchain.line.me")
	if err != nil {
		return nil, err
	}

	requestURL.Path = path.Join(
		path.Join(requestURL.Path, "v2", "finschia-2"),
		path.Join("blocks", fmt.Sprintf("%d", transaction.Info.BlockHeight)),
		path.Join("transactions", fmt.Sprintf("%d", transaction.Info.Position)),
		"messages",
	)

	requestURL.RawQuery = func() string {
		query := url.Values{}

		query.Set("size", "100")
		query.Set("search_from", "top")

		return query.Encode()
	}()

	res, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resData := struct {
		Messages []Message `json:"messages"`
	}{}

	if err = json.Unmarshal(resBody, &resData); err != nil {
		return nil, err
	}

	resData.Messages = slices.Map(resData.Messages, make([]Message, 0), func(message Message) Message {
		message.Value = MessageValue{}
		if err = json.Unmarshal([]byte(message.RawValue), &message.Value); err != nil {
			slog.Error("GetMillionArthurs", slog.String("step", "fetchMessages"), slog.Any("error", err))
		}

		return message
	})

	return resData.Messages, nil
}

func fetchItemToken(contractID string, tokenType string) (*ItemToken, error) {
	requestURL, err := url.Parse("https://explorer.blockchain.line.me")
	if err != nil {
		return nil, err
	}

	requestURL.Path = path.Join(
		path.Join(requestURL.Path, "v1", "finschia-2"),
		path.Join("item-tokens", contractID),
		fmt.Sprintf("%s00000001", tokenType),
	)

	res, err := http.Get(requestURL.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resData := ItemToken{}

	if err = json.Unmarshal(resBody, &resData); err != nil {
		return nil, err
	}

	properties := map[string]any{}

	if err = json.Unmarshal([]byte(resData.RawProperties), &properties); err != nil {
		return nil, err
	}

	for key, value := range properties {
		switch key {
		case "🍭シリーズ":
			if value, ok := value.(string); ok {
				resData.Properties.Series = &value
			}
		case "⚙ギアカテゴリ":
			if value, ok := value.(string); ok {
				resData.Properties.GearCategory = &value
			}
		case "⚙ギアレアリティ":
			if value, ok := value.(string); ok {
				resData.Properties.GearRarity = &value
			}
		}
	}

	return &resData, nil
}

func sendLINEMessage() error {
	bot, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		return err
	}

	bot.PushMessage(&messaging_api.PushMessageRequest{
		To: os.Getenv("LINE_MESSAGE_RECEIVER"),
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: "Hello, World!",
			},
		},
	}, "")

	return nil
}
