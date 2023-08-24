//go:build embed

package public

import "embed"

//go:embed dist/*
var ebox embed.FS

func init() {
	IsEmbedded = true
}
