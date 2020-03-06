package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, s string) {
	fmt.Fprintf(w, "Hello, %v", s)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
