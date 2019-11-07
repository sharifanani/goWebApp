package Pages

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = fmt.Fprint(w, "Home page!!")
}
