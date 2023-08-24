//go:build !embed

package public

import "embed"

var ebox embed.FS

func init() {
	IsEmbedded = false
}
