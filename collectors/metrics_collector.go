package collectors

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/frodenas/grafana_exporter/grafana"
)

type MetricsCollector struct {
	grafanaClient                   grafana.Client
	alertingActiveAlertsMetric      prometheus.Gauge
	alertingNotificationsSentMetric *prometheus.CounterVec
	alertingResultsMetric           *prometheus.CounterVec
	apiResponsesMetric              *prometheus.CounterVec
	apiUserSignupsMetric            *prometheus.CounterVec
	pageResponsesMetric             *prometheus.CounterVec
	proxyResponsesMetric            *prometheus.CounterVec
	dashboardsMetric                prometheus.Gauge
	orgsMetric                      prometheus.Gauge
	playlistsMetric                 prometheus.Gauge
	usersMetric                     prometheus.Gauge
	scrapesTotalMetric              prometheus.Counter
	scrapeErrorsTotalMetric         prometheus.Counter
	lastScrapeErrorMetric           prometheus.Gauge
	lastScrapeTimestampMetric       prometheus.Gauge
	lastScrapeDurationSecondsMetric prometheus.Gauge
}

func NewMetricsCollector(grafanaClient grafana.Client) (*MetricsCollector, error) {
	alertingActiveAlertsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "alerting_active_alerts",
			Help:      "Number of active alerts.",
		},
	)

	alertingNotificationsSentMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "alerting_notifications_sent",
			Help:      "Total number of alert notifications sent.",
		},
		[]string{"type"},
	)

	alertingResultsMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "alerting_results",
			Help:      "Total number of alerting results.",
		},
		[]string{"state"},
	)

	apiResponsesMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_responses",
			Help:      "Total number of API responses.",
		},
		[]string{"status_code"},
	)

	apiUserSignupsMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_user_signups",
			Help:      "Total number of API user signups.",
		},
		[]string{"state"},
	)

	pageResponsesMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "page_responses",
			Help:      "Total number of Page responses.",
		},
		[]string{"status_code"},
	)

	proxyResponsesMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "proxy_responses",
			Help:      "Total number of Proxy responses.",
		},
		[]string{"status_code"},
	)

	dashboardsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "dashboards",
			Help:      "Number of dashboards.",
		},
	)

	orgsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "orgs",
			Help:      "Number of orgs.",
		},
	)

	playlistsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "playlists",
			Help:      "Number of playlists.",
		},
	)

	usersMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "users",
			Help:      "Number of users.",
		},
	)

	scrapesTotalMetric := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "scrapes_total",
			Help:      "Total number of Grafana metrics scrapes.",
		},
	)

	scrapeErrorsTotalMetric := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "scrape_errors_total",
			Help:      "Total number of Grafana metrics scrape errors.",
		},
	)

	lastScrapeErrorMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "last_scrape_error",
			Help:      "Whether the last metrics scrape from Grafana resulted in an error (1 for error, 0 for success).",
		},
	)

	lastScrapeTimestampMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "last_scrape_timestamp",
			Help:      "Number of seconds since 1970 since last metrics scrape from Grafana.",
		},
	)

	lastScrapeDurationSecondsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "last_scrape_duration_seconds",
			Help:      "Duration of the last metrics scrape from Grafana.",
		},
	)

	metricsCollector := &MetricsCollector{
		grafanaClient:                   grafanaClient,
		alertingActiveAlertsMetric:      alertingActiveAlertsMetric,
		alertingNotificationsSentMetric: alertingNotificationsSentMetric,
		alertingResultsMetric:           alertingResultsMetric,
		apiResponsesMetric:              apiResponsesMetric,
		apiUserSignupsMetric:            apiUserSignupsMetric,
		pageResponsesMetric:             pageResponsesMetric,
		proxyResponsesMetric:            proxyResponsesMetric,
		dashboardsMetric:                dashboardsMetric,
		orgsMetric:                      orgsMetric,
		playlistsMetric:                 playlistsMetric,
		usersMetric:                     usersMetric,
		scrapesTotalMetric:              scrapesTotalMetric,
		scrapeErrorsTotalMetric:         scrapeErrorsTotalMetric,
		lastScrapeErrorMetric:           lastScrapeErrorMetric,
		lastScrapeTimestampMetric:       lastScrapeTimestampMetric,
		lastScrapeDurationSecondsMetric: lastScrapeDurationSecondsMetric,
	}

	return metricsCollector, nil
}

func (c *MetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	c.alertingActiveAlertsMetric.Describe(ch)
	c.alertingNotificationsSentMetric.Describe(ch)
	c.alertingResultsMetric.Describe(ch)
	c.apiResponsesMetric.Describe(ch)
	c.apiUserSignupsMetric.Describe(ch)
	c.pageResponsesMetric.Describe(ch)
	c.proxyResponsesMetric.Describe(ch)
	c.dashboardsMetric.Describe(ch)
	c.orgsMetric.Describe(ch)
	c.playlistsMetric.Describe(ch)
	c.usersMetric.Describe(ch)
	c.scrapesTotalMetric.Describe(ch)
	c.scrapeErrorsTotalMetric.Describe(ch)
	c.lastScrapeErrorMetric.Describe(ch)
	c.lastScrapeTimestampMetric.Describe(ch)
	c.lastScrapeDurationSecondsMetric.Describe(ch)
}

func (c *MetricsCollector) Collect(ch chan<- prometheus.Metric) {
	var begun = time.Now()

	errorMetric := float64(0)
	if err := c.reportMetrics(ch); err != nil {
		errorMetric = float64(1)
		c.scrapeErrorsTotalMetric.Inc()
		log.Errorf("Error while getting Grafana metrics: %s", err)
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

func (c *MetricsCollector) reportMetrics(ch chan<- prometheus.Metric) error {
	metrics, err := c.grafanaClient.GetMetrics()
	if err != nil {
		return err
	}

	c.alertingActiveAlertsMetric.Set(float64(metrics.AlertingActiveAlerts.Value))
	c.alertingActiveAlertsMetric.Collect(ch)

	c.alertingNotificationsSentMetric.Reset()
	c.alertingNotificationsSentMetric.WithLabelValues("line").Add(float64(metrics.AlertingNotificationsSentLine.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("email").Add(float64(metrics.AlertingNotificationsSentEmail.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("opsgenie").Add(float64(metrics.AlertingNotificationsSentOpsgenie.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("pagerduty").Add(float64(metrics.AlertingNotificationsSentPagerduty.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("pushover").Add(float64(metrics.AlertingNotificationsSentPushover.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("sensu").Add(float64(metrics.AlertingNotificationsSentSensu.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("slack").Add(float64(metrics.AlertingNotificationsSentSlack.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("telegram").Add(float64(metrics.AlertingNotificationsSentTelegram.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("threema").Add(float64(metrics.AlertingNotificationsSentThreema.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("victorops").Add(float64(metrics.AlertingNotificationsSentVictorops.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("webhook").Add(float64(metrics.AlertingNotificationsSentWebhook.Count))
	c.alertingNotificationsSentMetric.Collect(ch)

	c.alertingResultsMetric.Reset()
	c.alertingResultsMetric.WithLabelValues("alerting").Add(float64(metrics.AlertingResultStateAlerting.Count))
	c.alertingResultsMetric.WithLabelValues("no_data").Add(float64(metrics.AlertingResultStateNoData.Count))
	c.alertingResultsMetric.WithLabelValues("ok").Add(float64(metrics.AlertingResultStateOk.Count))
	c.alertingResultsMetric.WithLabelValues("paused").Add(float64(metrics.AlertingResultStatePaused.Count))
	c.alertingResultsMetric.WithLabelValues("pending").Add(float64(metrics.AlertingResultStatePending.Count))
	c.alertingResultsMetric.Collect(ch)

	c.apiResponsesMetric.Reset()
	c.apiResponsesMetric.WithLabelValues("200").Add(float64(metrics.APIRespStatusCode200.Count))
	c.apiResponsesMetric.WithLabelValues("404").Add(float64(metrics.APIRespStatusCode404.Count))
	c.apiResponsesMetric.WithLabelValues("500").Add(float64(metrics.APIRespStatusCode500.Count))
	c.apiResponsesMetric.WithLabelValues("unknown").Add(float64(metrics.APIRespStatusCodeUnknown.Count))
	c.apiResponsesMetric.Collect(ch)

	c.apiUserSignupsMetric.Reset()
	c.apiUserSignupsMetric.WithLabelValues("completed").Add(float64(metrics.APIUserSignupComplete.Count))
	c.apiUserSignupsMetric.WithLabelValues("invite").Add(float64(metrics.APIUserSignupInvite.Count))
	c.apiUserSignupsMetric.WithLabelValues("started").Add(float64(metrics.APIUserSignupStarted.Count))
	c.apiUserSignupsMetric.Collect(ch)

	c.pageResponsesMetric.Reset()
	c.pageResponsesMetric.WithLabelValues("200").Add(float64(metrics.PageRespStatusCode200.Count))
	c.pageResponsesMetric.WithLabelValues("404").Add(float64(metrics.PageRespStatusCode404.Count))
	c.pageResponsesMetric.WithLabelValues("500").Add(float64(metrics.PageRespStatusCode500.Count))
	c.pageResponsesMetric.WithLabelValues("unknown").Add(float64(metrics.PageRespStatusCodeUnknown.Count))
	c.pageResponsesMetric.Collect(ch)

	c.proxyResponsesMetric.Reset()
	c.proxyResponsesMetric.WithLabelValues("200").Add(float64(metrics.ProxyRespStatusCode200.Count))
	c.proxyResponsesMetric.WithLabelValues("404").Add(float64(metrics.ProxyRespStatusCode404.Count))
	c.proxyResponsesMetric.WithLabelValues("500").Add(float64(metrics.ProxyRespStatusCode500.Count))
	c.proxyResponsesMetric.WithLabelValues("unknown").Add(float64(metrics.ProxyRespStatusCodeUnknown.Count))
	c.proxyResponsesMetric.Collect(ch)

	c.dashboardsMetric.Set(float64(metrics.StatsTotalsStatDashboards.Value))
	c.dashboardsMetric.Collect(ch)

	c.orgsMetric.Set(float64(metrics.StatsTotalsStatOrgs.Value))
	c.orgsMetric.Collect(ch)

	c.playlistsMetric.Set(float64(metrics.StatsTotalsStatPlaylists.Value))
	c.playlistsMetric.Collect(ch)

	c.usersMetric.Set(float64(metrics.StatsTotalsStatUsers.Value))
	c.usersMetric.Collect(ch)

	return nil
}
