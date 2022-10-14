package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/klauspost/compress/zstd"
)

// Compress compress data and encode
func Compress(data []byte) string {
	var str string
	if strings.Contains(http.DetectContentType(data), "text/plain") {
		var buf bytes.Buffer
		w, err := zstd.NewWriter(&buf)
		if err != nil {
			goto failed
		}
		if _, err := w.Write(data); err == nil {
			w.Close()
			str = "$2$" + base64.StdEncoding.EncodeToString(buf.Bytes())
		}
	}
failed:
	if len(str) == 0 {
		str = "$0$" + base64.StdEncoding.EncodeToString(data)
	}
	return str
}

// Decompress decode data and decompress
func Decompress(str string) ([]byte, error) {
	switch {
	case strings.HasPrefix(str, "$0$"):
		return base64.StdEncoding.DecodeString(str[3:])
	case strings.HasPrefix(str, "$1$"):
		b64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str[3:]))
		r, err := gzip.NewReader(b64)
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		_, err = io.Copy(&buf, r)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	case strings.HasPrefix(str, "$2$"):
		b64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str[3:]))
		r, err := zstd.NewReader(b64)
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		_, err = io.Copy(&buf, r)
		if err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	default:
		return nil, fmt.Errorf("invalid data: %s", str)
	}
}
