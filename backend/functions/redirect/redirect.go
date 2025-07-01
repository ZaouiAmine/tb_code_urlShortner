package lib

import (
	"github.com/taubyte/go-sdk/event"
	"net/http"
)

//export redirect
func redirect(e event.Event) uint32 {
	http.Redirect(w, r, "https://google.com", http.StatusFound)
	return 0
}
