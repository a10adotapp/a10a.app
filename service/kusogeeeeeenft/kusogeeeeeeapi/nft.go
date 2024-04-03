package kusogeeeeeeapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type NFT struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	Status         string            `json:"status"`
	Collection     map[string]string `json:"collection"`
	MetaAttributes []*MetaAttribute  `json:"meta_attributes"`
	Amount         int               `json:"amount"`
	PublishedAt    time.Time         `json:"published_at"`
}

type MetaAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (nft *NFT) String() string {
	str := fmt.Sprintf("%s (%s) %s (%d)", nft.Name, nft.ID, nft.Status, nft.Amount)

	options := nft.Collection

	for _, metaAttribute := range nft.MetaAttributes {
		options[metaAttribute.Key] = metaAttribute.Value
	}

	if jsonData, err := json.Marshal(options); err == nil {
		str = fmt.Sprintf("%s %s", str, jsonData)
	}

	str = fmt.Sprintf("%s %s", str, nft.PublishedAt.Format(time.RFC3339))

	return str
}

func (nft *NFT) MetaAttributeString(key string) string {
	for _, metaAttribute := range nft.MetaAttributes {
		if metaAttribute.Key == key {
			return metaAttribute.Value
		}
	}

	return ""
}

func (nft *NFT) MetaAttributeNillableString(key string) *string {
	for _, metaAttribute := range nft.MetaAttributes {
		if metaAttribute.Key == key {
			return &metaAttribute.Value
		}
	}

	return nil
}

func (nft *NFT) MetaAttributeInt(key string) int {
	for _, metaAttribute := range nft.MetaAttributes {
		if metaAttribute.Key == key {
			if value, err := strconv.Atoi(metaAttribute.Value); err == nil {
				return value
			}
		}
	}

	return 0
}

func (nft *NFT) MetaAttributeNillableInt(key string) *int {
	for _, metaAttribute := range nft.MetaAttributes {
		if metaAttribute.Key == key {
			if value, err := strconv.Atoi(metaAttribute.Value); err == nil {
				return &value
			}
		}
	}

	return nil
}
