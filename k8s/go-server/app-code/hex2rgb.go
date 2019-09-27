package main

import (
	"fmt"
	"net/http"

	/* "log" */
	"github.com/gorilla/mux"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
			Name: "go_app_hex2rgb_handler_requests_total",
			Help: "The total number of processed events",
		})
	)

	r := mux.NewRouter()
	r.HandleFunc("/convert/hex-rgb/{hex}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		hex := params["hex"]
		c, err := colorful.Hex("#" + hex)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(w, "RGB(%v, %v, %v)", c.R*255, c.G*255, c.B*255)
		opsProcessed.Inc()
	})

	r.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":3000", r)
}
