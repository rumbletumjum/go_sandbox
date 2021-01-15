package thing

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "I love you %s", request.URL.Path[1:])
}

func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
