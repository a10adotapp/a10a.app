package kusogeeeeeeapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) NFTs(ctx context.Context, page int) ([]*NFT, *Pagination, error) {
	endpointURL, err := url.Parse("https://api.kusogeeeeee.com/v1/player/nfts")
	if err != nil {
		return nil, nil, err
	}

	queryParams := endpointURL.Query()
	queryParams.Set("page", strconv.Itoa(page))
	queryParams.Set("per_page", "20")
	queryParams.Set("sort_order", "updated_at_desc")
	queryParams.Add("filter_statuses", "minted")
	queryParams.Add("filter_statuses", "locked")
	queryParams.Add("filter_statuses", "hold")
	queryParams.Add("filter_statuses", "on_sale")
	queryParams.Add("filter_statuses", "sold")

	endpointURL.RawQuery = queryParams.Encode()

	res, err := http.Get(endpointURL.String())
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	resData := &ResponseData{}

	if err := json.Unmarshal(resBody, resData); err != nil {
		return nil, nil, err
	}

	return resData.NFTs, resData.Pagination, nil
}

type ResponseData struct {
	NFTs       []*NFT      `json:"nfts"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalCount  int `json:"total_count"`
	TotalPages  int `json:"total_pages"`
}
