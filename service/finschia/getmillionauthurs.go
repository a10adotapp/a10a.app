package finschia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func (FinschiaService) GetMillionAuthurs(ctx context.Context) error {
	reqData := []struct {
		Key    string   `json:"key"`
		Values []string `json:"values"`
	}{
		{
			Key: "💰OMJ",
		},
	}

	for i := 0; i <= 100; i++ {
		reqData[0].Values = append(reqData[0].Values, fmt.Sprintf("%d", i))
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://nft.line.me/api/v1/marketplace/brand/30/filter?size=1&page=1&sort=PRICE_LOW", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resData := struct {
		Content []struct {
			ID    uint32  `json:"id"`
			Price float32 `json:"price"`
			NFT   struct {
				Name string `json:"name"`
			} `json:"nft"`
		} `json:"content"`
	}{}

	if err = json.Unmarshal(resBody, &resData); err != nil {
		return err
	}

	if len(resData.Content) == 0 {
		return nil
	}

	slog.Info(
		"GetMillionAuthurs",
		slog.Uint64("id", uint64(resData.Content[0].ID)),
		slog.Float64("price", float64(resData.Content[0].Price)),
		slog.String("name", resData.Content[0].NFT.Name),
	)

	if resData.Content[0].Price > 0.005 {
		return nil
	}

	bot, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		return err
	}

	bot.PushMessage(&messaging_api.PushMessageRequest{
		To: os.Getenv("LINE_MESSAGE_RECEIVER"),
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: strings.Join([]string{
					fmt.Sprintf("Name: %s", resData.Content[0].NFT.Name),
					fmt.Sprintf("Price: %f", resData.Content[0].Price),
					fmt.Sprintf("https://nft.line.me/marketplace/sale/%d", resData.Content[0].ID),
				}, "\n"),
			},
		},
	}, "")

	return nil
}
