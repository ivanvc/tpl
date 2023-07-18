package template

import (
	"github.com/dustin/go-humanize"
	"github.com/spf13/cast"
)

// Defines the function map defining the file_size function.
var FileSizeFunc = map[string]any{"file_size": fileSize}

// Formats as a human-readable file size from a bytes size.
func fileSize(v any) string {
	s := cast.ToUint64(v)
	return humanize.Bytes(s)
}
