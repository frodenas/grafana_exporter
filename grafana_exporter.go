package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"

	"github.com/frodenas/grafana_exporter/collectors"
	"github.com/frodenas/grafana_exporter/grafana"
)

var (
	grafanaURI = flag.String(
		"grafana.uri", "",
		"Grafana URI ($GRAFANA_EXPORTER_GRAFANA_URI).",
	)

	grafanaUsername = flag.String(
		"grafana.username", "",
		"Grafana Username ($GRAFANA_EXPORTER_GRAFANA_USERNAME).",
	)

	grafanaPassword = flag.String(
		"grafana.password", "",
		"Grafana Password ($GRAFANA_EXPORTER_GRAFANA_PASSWORD).",
	)

	grafanaSkipSSLValidation = flag.Bool(
		"grafana.skip-ssl-verify", false,
		"Disable Grafana SSL Verify ($GRAFANA_EXPORTER_GRAFANA_SKIP_SSL_VERIFY).",
	)

	listenAddress = flag.String(
		"web.listen-address", ":9261",
		"Address to listen on for web interface and telemetry ($GRAFANA_EXPORTER_WEB_LISTEN_ADDRESS).",
	)

	metricsPath = flag.String(
		"web.telemetry-path", "/metrics",
		"Path under which to expose Prometheus metrics ($GRAFANA_EXPORTER_WEB_TELEMETRY_PATH).",
	)

	showVersion = flag.Bool(
		"version", false,
		"Print version information.",
	)
)

func init() {
	prometheus.MustRegister(version.NewCollector("grafana_exporter"))
}

func overrideFlagsWithEnvVars() {
	overrideWithEnvVar("GRAFANA_EXPORTER_GRAFANA_URI", grafanaURI)
	overrideWithEnvVar("GRAFANA_EXPORTER_GRAFANA_USERNAME", grafanaUsername)
	overrideWithEnvVar("GRAFANA_EXPORTER_GRAFANA_PASSWORD", grafanaPassword)
	overrideWithEnvBool("GRAFANA_EXPORTER_GRAFANA_SKIP_SSL_VERIFY", grafanaSkipSSLValidation)
	overrideWithEnvVar("GRAFANA_EXPORTER_WEB_LISTEN_ADDRESS", listenAddress)
	overrideWithEnvVar("GRAFANA_EXPORTER_WEB_TELEMETRY_PATH", metricsPath)
}

func overrideWithEnvVar(name string, value *string) {
	envValue := os.Getenv(name)
	if envValue != "" {
		*value = envValue
	}
}

func overrideWithEnvBool(name string, value *bool) {
	envValue := os.Getenv(name)
	if envValue != "" {
		var err error
		*value, err = strconv.ParseBool(envValue)
		if err != nil {
			log.Fatalf("Invalid `%s`: %s", name, err)
		}
	}
}

func main() {
	flag.Parse()
	overrideFlagsWithEnvVars()

	if *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("grafana_exporter"))
		os.Exit(0)
	}

	if *grafanaURI == "" {
		log.Error("Flag `grafana.uri` is required")
		os.Exit(1)
	}

	log.Infoln("Starting grafana_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	grafanaClient, err := grafana.NewHTTPClient(*grafanaURI, *grafanaUsername, *grafanaPassword, *grafanaSkipSSLValidation)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	adminStatsCollector := collectors.NewAdminStatsCollector(grafanaClient)
	prometheus.MustRegister(adminStatsCollector)

	metricsCollector := collectors.NewMetricsCollector(grafanaClient)
	prometheus.MustRegister(metricsCollector)

	http.Handle(*metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Grafana Exporter</title></head>
             <body>
             <h1>Grafana Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
