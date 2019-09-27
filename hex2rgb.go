package main

import (
	"fmt"
	"net/http"

	/* "log" */
	"github.com/gorilla/mux"
	"github.com/lucasb-eyer/go-colorful"

)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/convert/hex-rgb/{hex}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		hex := params["hex"]
		c, err := colorful.Hex("#" + hex)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(w, "RGB(%v, %v, %v)", c.R*255, c.G*255, c.B*255)

	})

	http.ListenAndServe(":3000", r)
}
