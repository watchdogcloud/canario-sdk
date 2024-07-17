package examples

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zakhaev26/canario/pkg/canario"
)

// example as mock server
func MainMock() {

	cio := canario.NewCanario()
	go cio.RunPeriodicMetrics()

	r := mux.NewRouter()
	fmt.Println("hmm")
	r.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"echo": "Hello Canario!",
		})
	})

	http.ListenAndServe(":3400", r)
}
