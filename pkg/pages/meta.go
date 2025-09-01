package pages

import (
	"maps"

	"github.com/e-kucheriavyi/gossrng/configs"
)

func NewMetaMap(params map[string]string) map[string]string {
	meta := map[string]string{
		"title":       configs.FallbackTitle,
		"description": configs.FallbackDescription,
		"keywords":    configs.FallbackKeywords,
		"date":        "",
	}

	maps.Copy(meta, params)

	return meta
}
