package collectors

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/frodenas/grafana_exporter/grafana"
)

type AdminStatsCollector struct {
	grafanaClient                   grafana.Client
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
}

func NewAdminStatsCollector(grafanaClient grafana.Client) *AdminStatsCollector {
	alertsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "alerts",
			Help:      "Number of Grafana Alerts.",
		},
	)

	dashboardsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "dashboards",
			Help:      "Number of Grafana Dashboards.",
		},
	)

	datasourcesMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "datasources",
			Help:      "Number of Grafana Datasources.",
		},
	)

	orgsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "orgs",
			Help:      "Number of Grafana Orgs.",
		},
	)

	playlistsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "playlists",
			Help:      "Number of Grafana Playlists.",
		},
	)

	dbSnapshotsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "db_snapshots",
			Help:      "Number of Grafana Snapshots.",
		},
	)

	starredDBMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "starred_db",
			Help:      "Number of Grafana Dashboards Starred.",
		},
	)

	dbTagsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "db_tags",
			Help:      "Number of Grafana Tags.",
		},
	)

	usersMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "admin_stats",
			Name:      "users",
			Help:      "Number of Grafana Users.",
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
		alertsMetric:                    alertsMetric,
		dashboardsMetric:                dashboardsMetric,
		datasourcesMetric:               datasourcesMetric,
		orgsMetric:                      orgsMetric,
		playlistsMetric:                 playlistsMetric,
		dbSnapshotsMetric:               dbSnapshotsMetric,
		starredDBMetric:                 starredDBMetric,
		dbTagsMetric:                    dbTagsMetric,
		usersMetric:                     usersMetric,
		scrapesTotalMetric:              scrapesTotalMetric,
		scrapeErrorsTotalMetric:         scrapeErrorsTotalMetric,
		lastScrapeErrorMetric:           lastScrapeErrorMetric,
		lastScrapeTimestampMetric:       lastScrapeTimestampMetric,
		lastScrapeDurationSecondsMetric: lastScrapeDurationSecondsMetric,
	}

	return adminStatsCollector
}

func (c *AdminStatsCollector) Describe(ch chan<- *prometheus.Desc) {
	c.alertsMetric.Describe(ch)
	c.dashboardsMetric.Describe(ch)
	c.datasourcesMetric.Describe(ch)
	c.orgsMetric.Describe(ch)
	c.playlistsMetric.Describe(ch)
	c.dbSnapshotsMetric.Describe(ch)
	c.starredDBMetric.Describe(ch)
	c.dbTagsMetric.Describe(ch)
	c.usersMetric.Describe(ch)
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

	c.alertsMetric.Set(float64(adminStats.AlertCount))
	c.alertsMetric.Collect(ch)

	c.dashboardsMetric.Set(float64(adminStats.DashboardCount))
	c.dashboardsMetric.Collect(ch)

	c.datasourcesMetric.Set(float64(adminStats.DatasourceCount))
	c.datasourcesMetric.Collect(ch)

	c.orgsMetric.Set(float64(adminStats.OrgCount))
	c.orgsMetric.Collect(ch)

	c.playlistsMetric.Set(float64(adminStats.PlaylistCount))
	c.playlistsMetric.Collect(ch)

	c.dbSnapshotsMetric.Set(float64(adminStats.DBSnapshotCount))
	c.dbSnapshotsMetric.Collect(ch)

	c.starredDBMetric.Set(float64(adminStats.StarredDBCount))
	c.starredDBMetric.Collect(ch)

	c.dbTagsMetric.Set(float64(adminStats.DBTagCount))
	c.dbTagsMetric.Collect(ch)

	c.usersMetric.Set(float64(adminStats.UserCount))
	c.usersMetric.Collect(ch)

	return nil
}
