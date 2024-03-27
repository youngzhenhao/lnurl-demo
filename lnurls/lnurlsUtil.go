package lnurls

import (
	"github.com/fiatjaf/go-lnurl"
)

func Encode(url string) string {
	en, _ := lnurl.LNURLEncode(url)
	return en
}
func Decode(ln string) string {
	de, _ := lnurl.LNURLDecode(ln)
	return de
}
