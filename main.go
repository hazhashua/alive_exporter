package main

import (
	"net/http"

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

	//Create a new instance of the foocollector and
	//register it with the prometheus client.
	foo := collector.newFooCollector()
	prometheus.MustRegister(foo)

	go collector.generateValue(foo.channel)
	go collector.getValueLoop(foo.channel)

	// ch := chan <- prometheus.Metric
	// foo.Collectx(make(chan<- prometheus.Metric), 100)

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
