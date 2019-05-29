package url_shortener

import (
	"encoding/base64"
	"testing"
)

func TestShortener(t *testing.T) {
	svc := NewURLShortener()
	longURLs := []string{
		`https://google.com`,
		``,
		`sdsdsdsdsd

			eripweopoeir`,
	}
	for _, url := range longURLs {
		svc.Shorten(url)
		if svc.Resolve(base64.StdEncoding.EncodeToString([]byte(url))) != url {
			t.Errorf("Resolve() = %q, want %q", svc.Resolve(base64.StdEncoding.EncodeToString([]byte(url))), url)
		}
	}
}
