package collectors_test

import (
	"errors"
	"flag"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/frodenas/grafana_exporter/grafana"
	"github.com/frodenas/grafana_exporter/grafana/grafanafakes"
	"github.com/prometheus/client_golang/prometheus"

	. "github.com/frodenas/grafana_exporter/collectors"
	. "github.com/frodenas/grafana_exporter/utils/test_matchers"
)

func init() {
	flag.Set("log.level", "fatal")
}

var _ = Describe("AdminStatsCollectors", func() {
	var (
		grafanaClient *grafanafakes.FakeClient

		alertsMetric                    prometheus.Gauge
		dashboardsMetric                prometheus.Gauge
		datasourcesMetric               prometheus.Gauge
		orgsMetric                      prometheus.Gauge
		playlistsMetric                 prometheus.Gauge
		dbSnapshotsMetric               prometheus.Gauge
		starredDBMetric                 prometheus.Gauge
		dbTagsMetric                    prometheus.Gauge
		usersMetric                     prometheus.Gauge
		scrapesTotalMetric              prometheus.Counter
		scrapeErrorsTotalMetric         prometheus.Counter
		lastScrapeErrorMetric           prometheus.Gauge
		lastScrapeTimestampMetric       prometheus.Gauge
		lastScrapeDurationSecondsMetric prometheus.Gauge

		alertCount      = 1
		dashboardCount  = 2
		datasourceCount = 3
		orgCount        = 4
		playlistCount   = 5
		snapshotCount   = 6
		starredCount    = 7
		tagCount        = 8
		userCount       = 9

		adminStatsCollector *AdminStatsCollector
	)

	BeforeEach(func() {
		grafanaClient = &grafanafakes.FakeClient{}

		alertsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "alerts",
				Help:      "Number of Grafana Alerts.",
			},
		)

		dashboardsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "dashboards",
				Help:      "Number of Grafana Dashboards.",
			},
		)

		datasourcesMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "datasources",
				Help:      "Number of Grafana Datasources.",
			},
		)

		orgsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "orgs",
				Help:      "Number of Grafana Orgs.",
			},
		)

		playlistsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "playlists",
				Help:      "Number of Grafana Playlists.",
			},
		)

		dbSnapshotsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "db_snapshots",
				Help:      "Number of Grafana Snapshots.",
			},
		)

		starredDBMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "starred_db",
				Help:      "Number of Grafana Dashboards Starred.",
			},
		)

		dbTagsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "db_tags",
				Help:      "Number of Grafana Tags.",
			},
		)

		usersMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "users",
				Help:      "Number of Grafana Users.",
			},
		)

		scrapesTotalMetric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "scrapes_total",
				Help:      "Total number of Grafana Admin Stats scrapes.",
			},
		)
		scrapesTotalMetric.Inc()

		scrapeErrorsTotalMetric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "scrape_errors_total",
				Help:      "Total number of Grafana Admin Stats scrape errors.",
			},
		)

		lastScrapeErrorMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "last_scrape_error",
				Help:      "Whether the last metrics scrape from Grafana Admin Stats resulted in an error (1 for error, 0 for success).",
			},
		)
		lastScrapeErrorMetric.Set(0)

		lastScrapeTimestampMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "last_scrape_timestamp",
				Help:      "Number of seconds since 1970 since last metrics scrape from Grafana Admin Stats.",
			},
		)

		lastScrapeDurationSecondsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "last_scrape_duration_seconds",
				Help:      "Duration of the last metrics scrape from Grafana Admin Stats.",
			},
		)
	})

	JustBeforeEach(func() {
		adminStatsCollector = NewAdminStatsCollector(grafanaClient)
	})

	Describe("Describe", func() {
		var (
			descriptions chan *prometheus.Desc
		)

		BeforeEach(func() {
			descriptions = make(chan *prometheus.Desc)
		})

		JustBeforeEach(func() {
			go adminStatsCollector.Describe(descriptions)
		})

		It("returns a grafana_admin_stats_alerts metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(alertsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_dashboards metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(dashboardsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_datasources metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(datasourcesMetric.Desc())))
		})

		It("returns a grafana_admin_stats_orgs metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(orgsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_playlists metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(playlistsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_db_snapshots metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(dbSnapshotsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_starred_db metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(starredDBMetric.Desc())))
		})

		It("returns a grafana_admin_stats_db_tags metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(dbTagsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_users metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(usersMetric.Desc())))
		})

		It("returns a grafana_admin_stats_scrapes_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(scrapesTotalMetric.Desc())))
		})

		It("returns a grafana_admin_stats_scrape_errors_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(scrapeErrorsTotalMetric.Desc())))
		})

		It("returns a grafana_admin_stats_last_scrape_error metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(lastScrapeErrorMetric.Desc())))
		})

		It("returns a grafana_admin_stats_last_scrape_timestamp metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(lastScrapeTimestampMetric.Desc())))
		})

		It("returns a grafana_admin_stats_last_scrape_duration_seconds metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(lastScrapeDurationSecondsMetric.Desc())))
		})
	})

	Describe("Collect", func() {
		var (
			adminStatsResponse grafana.AdminStats

			metrics chan prometheus.Metric
		)

		BeforeEach(func() {
			adminStatsResponse = grafana.AdminStats{
				AlertCount:      alertCount,
				DashboardCount:  dashboardCount,
				DatasourceCount: datasourceCount,
				OrgCount:        orgCount,
				PlaylistCount:   playlistCount,
				DBSnapshotCount: snapshotCount,
				StarredDBCount:  starredCount,
				DBTagCount:      tagCount,
				UserCount:       userCount,
			}
			grafanaClient.GetAdminStatsReturns(adminStatsResponse, nil)

			metrics = make(chan prometheus.Metric)
		})

		JustBeforeEach(func() {
			go adminStatsCollector.Collect(metrics)
		})

		It("returns a grafana_admin_stats_alerts metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertsMetric)))
		})

		It("returns a grafana_admin_stats_dashboards metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(dashboardsMetric)))
		})

		It("returns a grafana_admin_stats_datasources metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(datasourcesMetric)))
		})

		It("returns a grafana_admin_stats_orgs metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(orgsMetric)))
		})

		It("returns a grafana_admin_stats_playlists metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(playlistsMetric)))
		})

		It("returns a grafana_admin_stats_db_snapshots metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(dbSnapshotsMetric)))
		})

		It("returns a grafana_admin_stats_starred_db metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(starredDBMetric)))
		})

		It("returns a grafana_admin_stats_db_tags metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(dbTagsMetric)))
		})

		It("returns a grafana_admin_stats_users metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(usersMetric)))
		})

		It("returns a grafana_admin_stats_scrapes_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(scrapesTotalMetric)))
		})

		It("returns a grafana_admin_stats_scrape_errors_total", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(scrapeErrorsTotalMetric)))
		})

		It("returns a grafana_admin_stats_last_scrape_error metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(lastScrapeErrorMetric)))
		})

		Context("when it fails to list the security groups", func() {
			BeforeEach(func() {
				grafanaClient.GetAdminStatsReturns(adminStatsResponse, errors.New("error"))

				scrapeErrorsTotalMetric.Inc()
				lastScrapeErrorMetric.Set(1)
			})

			It("returns a grafana_admin_stats_scrape_errors_total metric", func() {
				Eventually(metrics).Should(Receive(PrometheusMetric(scrapeErrorsTotalMetric)))
			})

			It("returns a grafana_admin_stats_last_scrape_error metric", func() {
				Eventually(metrics).Should(Receive(PrometheusMetric(lastScrapeErrorMetric)))
			})
		})
	})
})
