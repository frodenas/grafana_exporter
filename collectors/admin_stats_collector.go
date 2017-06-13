package collectors

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/frodenas/grafana_exporter/grafana"
)

type AdminStatsCollector struct {
	grafanaClient                   grafana.Client
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
}

func NewAdminStatsCollector(grafanaClient grafana.Client) *AdminStatsCollector {
	totalAlertsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "alerts_total",
			Help:      "Total number of Grafana Alerts.",
		},
	)

	totalDashboardsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "dashboards_total",
			Help:      "Total number of Grafana Dashboards.",
		},
	)

	totalDatasourcesMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "datasources_total",
			Help:      "Total number of Grafana Datasources.",
		},
	)

	totalOrgsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "orgs_total",
			Help:      "Total number of Grafana Orgs.",
		},
	)

	totalPlaylistsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "playlists_total",
			Help:      "Total number of Grafana Playlists.",
		},
	)

	totalSnapshotsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "snapshots_total",
			Help:      "Total number of Grafana Snapshots.",
		},
	)

	totalStarredMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "starred_total",
			Help:      "Total number of Grafana Dashboards Starred.",
		},
	)

	totalTagsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "tags_total",
			Help:      "Total number of Grafana Tags.",
		},
	)

	totalUsersMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "users_total",
			Help:      "Total number of Grafana Users.",
		},
	)

	scrapesTotalMetric := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "scrapes_total",
			Help:      "Total number of Grafana Admin Stats scrapes.",
		},
	)

	scrapeErrorsTotalMetric := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "scrape_errors_total",
			Help:      "Total number of Grafana Admin Stats scrape errors.",
		},
	)

	lastScrapeErrorMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "last_scrape_error",
			Help:      "Whether the last metrics scrape from Grafana Admin Stats resulted in an error (1 for error, 0 for success).",
		},
	)

	lastScrapeTimestampMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "last_scrape_timestamp",
			Help:      "Number of seconds since 1970 since last metrics scrape from Grafana Admin Stats.",
		},
	)

	lastScrapeDurationSecondsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "last_scrape_duration_seconds",
			Help:      "Duration of the last metrics scrape from Grafana Admin Stats.",
		},
	)

	adminStatsCollector := &AdminStatsCollector{
		grafanaClient:                   grafanaClient,
		totalAlertsMetric:               totalAlertsMetric,
		totalDashboardsMetric:           totalDashboardsMetric,
		totalDatasourcesMetric:          totalDatasourcesMetric,
		totalOrgsMetric:                 totalOrgsMetric,
		totalPlaylistsMetric:            totalPlaylistsMetric,
		totalSnapshotsMetric:            totalSnapshotsMetric,
		totalStarredMetric:              totalStarredMetric,
		totalTagsMetric:                 totalTagsMetric,
		totalUsersMetric:                totalUsersMetric,
		scrapesTotalMetric:              scrapesTotalMetric,
		scrapeErrorsTotalMetric:         scrapeErrorsTotalMetric,
		lastScrapeErrorMetric:           lastScrapeErrorMetric,
		lastScrapeTimestampMetric:       lastScrapeTimestampMetric,
		lastScrapeDurationSecondsMetric: lastScrapeDurationSecondsMetric,
	}

	return adminStatsCollector
}

func (c *AdminStatsCollector) Describe(ch chan<- *prometheus.Desc) {
	c.totalAlertsMetric.Describe(ch)
	c.totalDashboardsMetric.Describe(ch)
	c.totalDatasourcesMetric.Describe(ch)
	c.totalOrgsMetric.Describe(ch)
	c.totalPlaylistsMetric.Describe(ch)
	c.totalSnapshotsMetric.Describe(ch)
	c.totalStarredMetric.Describe(ch)
	c.totalTagsMetric.Describe(ch)
	c.totalUsersMetric.Describe(ch)
	c.scrapesTotalMetric.Describe(ch)
	c.scrapeErrorsTotalMetric.Describe(ch)
	c.lastScrapeErrorMetric.Describe(ch)
	c.lastScrapeTimestampMetric.Describe(ch)
	c.lastScrapeDurationSecondsMetric.Describe(ch)
}

func (c *AdminStatsCollector) Collect(ch chan<- prometheus.Metric) {
	var begun = time.Now()

	errorMetric := float64(0)
	if err := c.reportAdminStatsMetrics(ch); err != nil {
		errorMetric = float64(1)
		c.scrapeErrorsTotalMetric.Inc()
		log.Errorf("Error while getting Grafana Admin Stats metrics: %s", err)
	}
	c.scrapeErrorsTotalMetric.Collect(ch)

	c.scrapesTotalMetric.Inc()
	c.scrapesTotalMetric.Collect(ch)

	c.lastScrapeErrorMetric.Set(errorMetric)
	c.lastScrapeErrorMetric.Collect(ch)

	c.lastScrapeTimestampMetric.Set(float64(time.Now().Unix()))
	c.lastScrapeTimestampMetric.Collect(ch)

	c.lastScrapeDurationSecondsMetric.Set(time.Since(begun).Seconds())
	c.lastScrapeDurationSecondsMetric.Collect(ch)
}

func (c *AdminStatsCollector) reportAdminStatsMetrics(ch chan<- prometheus.Metric) error {
	adminStats, err := c.grafanaClient.GetAdminStats()
	if err != nil {
		return err
	}

	c.totalAlertsMetric.Set(float64(adminStats.AlertCount))
	c.totalAlertsMetric.Collect(ch)

	c.totalDashboardsMetric.Set(float64(adminStats.DashboardCount))
	c.totalDashboardsMetric.Collect(ch)

	c.totalDatasourcesMetric.Set(float64(adminStats.DatasourceCount))
	c.totalDatasourcesMetric.Collect(ch)

	c.totalOrgsMetric.Set(float64(adminStats.OrgCount))
	c.totalOrgsMetric.Collect(ch)

	c.totalPlaylistsMetric.Set(float64(adminStats.PlaylistCount))
	c.totalPlaylistsMetric.Collect(ch)

	c.totalSnapshotsMetric.Set(float64(adminStats.SnapshotCount))
	c.totalSnapshotsMetric.Collect(ch)

	c.totalStarredMetric.Set(float64(adminStats.StarredCount))
	c.totalStarredMetric.Collect(ch)

	c.totalTagsMetric.Set(float64(adminStats.TagCount))
	c.totalTagsMetric.Collect(ch)

	c.totalUsersMetric.Set(float64(adminStats.UserCount))
	c.totalUsersMetric.Collect(ch)

	return nil
}
