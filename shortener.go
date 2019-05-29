package url_shortener

import "encoding/base64"

//Shortener url shortener interface
type Shortener interface {
	Shorten(string) string
	Resolve(string) string
}

//URLShortener wrapper for url shortener service
type URLShortener struct {
	Mapper map[string]string
}

//NewURLShortener create new url shortener service
func NewURLShortener() *URLShortener {
	mapper := make(map[string]string)
	return &URLShortener{
		Mapper: mapper,
	}
}

//Shorten shorten input url and store for resolving
func (u *URLShortener) Shorten(url string) string {
	shortened := shorten(url)
	u.Mapper[shortened] = url
	return shortened
}

//Shorten shorten input url and store for resolving
func (u *URLShortener) Resolve(short string) string {
	if original, ok := u.Mapper[short]; ok {
		return original
	}
	return ""
}

func shorten(url string) string {
	return base64.StdEncoding.EncodeToString([]byte(url))
}
