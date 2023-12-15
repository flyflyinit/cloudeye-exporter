package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteCatalogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCatalogResponse struct{}"
	}

	return strings.Join([]string{"DeleteCatalogResponse", string(data)}, " ")
}