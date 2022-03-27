package a

import (
	"net/http"
	"util"
)

func issue39() {
	res, _ := http.Get("http://example.com/") // OK
	defer util.Close(res.Body)
}
