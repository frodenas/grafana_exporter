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

		totalAlertsMetric               prometheus.Gauge
		totalDashboardsMetric           prometheus.Gauge
		totalDatasourcesMetric          prometheus.Gauge
		totalOrgsMetric                 prometheus.Gauge
		totalPlaylistsMetric            prometheus.Gauge
		totalSnapshotsMetric            prometheus.Gauge
		totalStarredMetric              prometheus.Gauge
		totalTagsMetric                 prometheus.Gauge
		totalUsersMetric                prometheus.Gauge
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

		totalAlertsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "alerts_total",
				Help:      "Total number of Grafana Alerts.",
			},
		)
		totalAlertsMetric.Set(float64(alertCount))

		totalDashboardsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "dashboards_total",
				Help:      "Total number of Grafana Dashboards.",
			},
		)
		totalDashboardsMetric.Set(float64(dashboardCount))

		totalDatasourcesMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "datasources_total",
				Help:      "Total number of Grafana Datasources.",
			},
		)
		totalDatasourcesMetric.Set(float64(datasourceCount))

		totalOrgsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "orgs_total",
				Help:      "Total number of Grafana Orgs.",
			},
		)
		totalOrgsMetric.Set(float64(orgCount))

		totalPlaylistsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "playlists_total",
				Help:      "Total number of Grafana Playlists.",
			},
		)
		totalPlaylistsMetric.Set(float64(playlistCount))

		totalSnapshotsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "snapshots_total",
				Help:      "Total number of Grafana Snapshots.",
			},
		)
		totalSnapshotsMetric.Set(float64(snapshotCount))

		totalStarredMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "starred_total",
				Help:      "Total number of Grafana Dashboards Starred.",
			},
		)
		totalStarredMetric.Set(float64(starredCount))

		totalTagsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "tags_total",
				Help:      "Total number of Grafana Tags.",
			},
		)
		totalTagsMetric.Set(float64(tagCount))

		totalUsersMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "admin_stats",
				Name:      "users_total",
				Help:      "Total number of Grafana Users.",
			},
		)
		totalUsersMetric.Set(float64(userCount))

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

		It("returns a grafana_admin_stats_alerts_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalAlertsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_dashboards_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalDashboardsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_datasources_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalDatasourcesMetric.Desc())))
		})

		It("returns a grafana_admin_stats_orgs_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalOrgsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_playlists_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalPlaylistsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_snapshots_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalSnapshotsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_starred_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalStarredMetric.Desc())))
		})

		It("returns a grafana_admin_stats_tags_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalTagsMetric.Desc())))
		})

		It("returns a grafana_admin_stats_users_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(totalUsersMetric.Desc())))
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
				SnapshotCount:   snapshotCount,
				StarredCount:    starredCount,
				TagCount:        tagCount,
				UserCount:       userCount,
			}
			grafanaClient.GetAdminStatsReturns(adminStatsResponse, nil)

			metrics = make(chan prometheus.Metric)
		})

		JustBeforeEach(func() {
			go adminStatsCollector.Collect(metrics)
		})

		It("returns a grafana_admin_stats_alerts_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalAlertsMetric)))
		})

		It("returns a grafana_admin_stats_dashboards_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalDashboardsMetric)))
		})

		It("returns a grafana_admin_stats_datasources_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalDatasourcesMetric)))
		})

		It("returns a grafana_admin_stats_orgs_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalOrgsMetric)))
		})

		It("returns a grafana_admin_stats_playlists_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalPlaylistsMetric)))
		})

		It("returns a grafana_admin_stats_snapshots_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalSnapshotsMetric)))
		})

		It("returns a grafana_admin_stats_starred_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalStarredMetric)))
		})

		It("returns a grafana_admin_stats_tags_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalTagsMetric)))
		})

		It("returns a grafana_admin_stats_users_total metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(totalUsersMetric)))
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
