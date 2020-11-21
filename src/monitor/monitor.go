/*
  Server Monitor Port Module
*/
package monitor

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Serve(addr string) error {
	indexPage := []byte(
		`<html>
  <head>
    <title>Monitor</title>
  </head>
  <body>
    <h1>Monitor</h1>
    <ul>
      <li>
        <a href="/metrics">Prometheus Metrics</a>
      </li>
      <li>
        <a href="/debug/pprof">Pprof</a>
      </li>
    </ul>
  </body>
</html>`)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write(indexPage)
	})
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(addr, nil)
}
