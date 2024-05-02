package kusogeeeeeeapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/a10adotapp/a10a.app/lib/xslices"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) NFTs(ctx context.Context, page int) ([]*NFT, *Pagination, error) {
	url := "https://api.kusogeeeeee.com/v1/player/collections/01HP6H0VE1XSB9NPHHGRTTQPCN/nfts/search"

	reqData := &RequestData{
		Page:    page,
		PerPage: 50,
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		slog.ErrorContext(ctx, "failed serialize request body", slog.Any("err", err))

		return nil, nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate http request", slog.Any("err", err), slog.String("url", url))

		return nil, nil, err
	}

	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		slog.ErrorContext(ctx, "failed to complete http request", slog.Any("err", err), slog.String("url", url))

		return nil, nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		slog.ErrorContext(ctx, "failed to read response body", slog.Any("err", err))

		return nil, nil, err
	}

	resData := &ResponseData{}

	if err := json.Unmarshal(resBody, resData); err != nil {
		slog.ErrorContext(
			ctx,
			"failed to parse response body",
			slog.Any("err", err),
			slog.String("url", url),
			slog.Int("response_status", res.StatusCode),
			slog.String("response_body", string(resBody)),
		)

		return nil, nil, err
	}

	return xslices.Map(resData.NFTItems, func(nftItem *NFTItem) *NFT {
		return nftItem.NFT
	}), resData.Pagination, nil
}

type NFTItem struct {
	NFT *NFT `json:"nft"`
}

type ResponseData struct {
	NFTItems   []*NFTItem  `json:"nfts"`
	Pagination *Pagination `json:"pagination"`
}

type RequestData struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalCount  int `json:"total_count"`
	TotalPages  int `json:"total_pages"`
}
