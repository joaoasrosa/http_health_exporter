package http_health_exporter

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :9545")
	log.Fatal(http.ListenAndServe(":9545", nil))
}
