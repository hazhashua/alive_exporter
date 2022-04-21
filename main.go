package main

import (
	"net/http"
	// "github.com/hazhashua/alive_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {

	// c := make(chan int)
	// go generateValue(c)
	// for {
	// 	fmt.Println(getValue(c))
	// }

	// aPIV1Services, err := UnmarshalAPIV1Services("bytesxxx")
	// bytes, err = aPIV1Services.Marshal()

	getUrl("http://124.65.131.14:38080/api/v1/services")

	//Create a new instance of the foocollector and
	//register it with the prometheus client.
	foo := newFooCollector()
	prometheus.MustRegister(foo)

	go generateValue(foo.channel)
	go getValueLoop(foo.channel)

	// ch := chan <- prometheus.Metric
	// foo.Collectx(make(chan<- prometheus.Metric), 100)

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
