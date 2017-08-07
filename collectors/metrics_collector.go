package collectors

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/frodenas/grafana_exporter/grafana"
)

type MetricsCollector struct {
	grafanaClient                          grafana.Client
	alertingActiveAlertsMetric             prometheus.Gauge
	alertingExecutionTimeMetric            *prometheus.GaugeVec
	alertingNotificationsSentMetric        *prometheus.GaugeVec
	alertingResultsMetric                  *prometheus.GaugeVec
	apiAdminUserCreateMetric               prometheus.Gauge
	apiDashboardGetMetric                  *prometheus.GaugeVec
	apiDashboardSaveMetric                 *prometheus.GaugeVec
	apiDashboardSearchMetric               *prometheus.GaugeVec
	apiDashboardSnapshotCreateMetric       prometheus.Gauge
	apiDashboardSnapshotExternalMetric     prometheus.Gauge
	apiDashboardSnapshotGetMetric          prometheus.Gauge
	apiDataproxyRequestAllMetric           *prometheus.GaugeVec
	apiLoginOauthMetric                    prometheus.Gauge
	apiLoginPostMetric                     prometheus.Gauge
	apiOrgCreateMetric                     prometheus.Gauge
	apiResponsesMetric                     *prometheus.GaugeVec
	apiUserSignupsCompletedMetric          prometheus.Gauge
	apiUserSignupsInviteMetric             prometheus.Gauge
	apiUserSignupsStartedMetric            prometheus.Gauge
	awsCloudwatchGetMetricStatisticsMetric prometheus.Gauge
	awsCloudwatchListMetricsMetric         prometheus.Gauge
	instanceStartMetric                    prometheus.Gauge
	modelsDashboardInsertMetric            prometheus.Gauge
	pageResponsesMetric                    *prometheus.GaugeVec
	proxyResponsesMetric                   *prometheus.GaugeVec
	dashboardsMetric                       prometheus.Gauge
	orgsMetric                             prometheus.Gauge
	playlistsMetric                        prometheus.Gauge
	usersMetric                            prometheus.Gauge
	scrapesTotalMetric                     prometheus.Counter
	scrapeErrorsTotalMetric                prometheus.Counter
	lastScrapeErrorMetric                  prometheus.Gauge
	lastScrapeTimestampMetric              prometheus.Gauge
	lastScrapeDurationSecondsMetric        prometheus.Gauge
}

func NewMetricsCollector(grafanaClient grafana.Client) *MetricsCollector {
	alertingActiveAlertsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "alerting_active_alerts",
			Help:      "Number of active alerts.",
		},
	)

	alertingExecutionTimeMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "alerting_execution_time",
			Help:      "Alerting execution time.",
		},
		[]string{"metric"},
	)

	alertingNotificationsSentMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "alerting_notifications_sent",
			Help:      "Number of alert notifications sent.",
		},
		[]string{"type"},
	)

	alertingResultsMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "alerting_results",
			Help:      "Number of alerting results.",
		},
		[]string{"state"},
	)

	apiAdminUserCreateMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_admin_user_create",
			Help:      "Number of calls to Admin User Create API.",
		},
	)

	apiDashboardGetMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_dashboard_get",
			Help:      "Dashboard Get API times.",
		},
		[]string{"metric"},
	)

	apiDashboardSaveMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_dashboard_save",
			Help:      "Dashboard Save API times.",
		},
		[]string{"metric"},
	)

	apiDashboardSearchMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_dashboard_search",
			Help:      "Dashboard Search API times.",
		},
		[]string{"metric"},
	)

	apiDashboardSnapshotCreateMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_dashboard_snapshot_create",
			Help:      "Number of calls to Dashboard Snapshot Create API.",
		},
	)

	apiDashboardSnapshotExternalMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_dashboard_snapshot_external",
			Help:      "Number of calls to Dashboard Snapshot External API.",
		},
	)

	apiDashboardSnapshotGetMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_dashboard_snapshot_get",
			Help:      "Number of calls to Dashboard Snapshot Get API.",
		},
	)

	apiDataproxyRequestAllMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_dataproxy_request_all",
			Help:      "Dataproxy request API times.",
		},
		[]string{"metric"},
	)

	apiLoginOauthMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_login_oauth",
			Help:      "Number of calls to Login OAuth API.",
		},
	)

	apiLoginPostMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_login_post",
			Help:      "Number of calls to Login Post API.",
		},
	)

	apiOrgCreateMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_org_create",
			Help:      "Number of calls to Org Create API.",
		},
	)

	apiResponsesMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_responses",
			Help:      "Number of API responses.",
		},
		[]string{"code"},
	)

	apiUserSignupsCompletedMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_user_signups_completed",
			Help:      "Number of API User Signups completed.",
		},
	)

	apiUserSignupsInviteMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_user_signups_invite",
			Help:      "Number of API User Signups invite.",
		},
	)

	apiUserSignupsStartedMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "api_user_signups_started",
			Help:      "Number of API User Signups started.",
		},
	)

	awsCloudwatchGetMetricStatisticsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "aws_cloudwatch_get_metric_statistics",
			Help:      "Number of calls to AWS CloudWatch Get Metric Statistics API.",
		},
	)

	awsCloudwatchListMetricsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "aws_cloudwatch_list_metrics",
			Help:      "Number of calls to AWS CloudWatch List Metrics API.",
		},
	)

	instanceStartMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "instance_start",
			Help:      "Number of Instance Starts.",
		},
	)

	modelsDashboardInsertMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "models_dashboard_insert",
			Help:      "Number of Dashboard inserts.",
		},
	)

	pageResponsesMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "page_responses",
			Help:      "Number of Page responses.",
		},
		[]string{"code"},
	)

	proxyResponsesMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grafana",
			Subsystem: "metrics",
			Name:      "proxy_responses",
			Help:      "Number of Proxy responses.",
		},
		[]string{"code"},
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
		grafanaClient:                          grafanaClient,
		alertingActiveAlertsMetric:             alertingActiveAlertsMetric,
		alertingExecutionTimeMetric:            alertingExecutionTimeMetric,
		alertingNotificationsSentMetric:        alertingNotificationsSentMetric,
		alertingResultsMetric:                  alertingResultsMetric,
		apiAdminUserCreateMetric:               apiAdminUserCreateMetric,
		apiDashboardGetMetric:                  apiDashboardGetMetric,
		apiDashboardSaveMetric:                 apiDashboardSaveMetric,
		apiDashboardSearchMetric:               apiDashboardSearchMetric,
		apiDashboardSnapshotCreateMetric:       apiDashboardSnapshotCreateMetric,
		apiDashboardSnapshotExternalMetric:     apiDashboardSnapshotExternalMetric,
		apiDashboardSnapshotGetMetric:          apiDashboardSnapshotGetMetric,
		apiDataproxyRequestAllMetric:           apiDataproxyRequestAllMetric,
		apiLoginOauthMetric:                    apiLoginOauthMetric,
		apiLoginPostMetric:                     apiLoginPostMetric,
		apiOrgCreateMetric:                     apiOrgCreateMetric,
		apiResponsesMetric:                     apiResponsesMetric,
		apiUserSignupsCompletedMetric:          apiUserSignupsCompletedMetric,
		apiUserSignupsInviteMetric:             apiUserSignupsInviteMetric,
		apiUserSignupsStartedMetric:            apiUserSignupsStartedMetric,
		awsCloudwatchGetMetricStatisticsMetric: awsCloudwatchGetMetricStatisticsMetric,
		awsCloudwatchListMetricsMetric:         awsCloudwatchListMetricsMetric,
		instanceStartMetric:                    instanceStartMetric,
		modelsDashboardInsertMetric:            modelsDashboardInsertMetric,
		pageResponsesMetric:                    pageResponsesMetric,
		proxyResponsesMetric:                   proxyResponsesMetric,
		dashboardsMetric:                       dashboardsMetric,
		orgsMetric:                             orgsMetric,
		playlistsMetric:                        playlistsMetric,
		usersMetric:                            usersMetric,
		scrapesTotalMetric:                     scrapesTotalMetric,
		scrapeErrorsTotalMetric:                scrapeErrorsTotalMetric,
		lastScrapeErrorMetric:                  lastScrapeErrorMetric,
		lastScrapeTimestampMetric:              lastScrapeTimestampMetric,
		lastScrapeDurationSecondsMetric:        lastScrapeDurationSecondsMetric,
	}

	return metricsCollector
}

func (c *MetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	c.alertingActiveAlertsMetric.Describe(ch)
	c.alertingExecutionTimeMetric.Describe(ch)
	c.alertingNotificationsSentMetric.Describe(ch)
	c.alertingResultsMetric.Describe(ch)
	c.apiAdminUserCreateMetric.Describe(ch)
	c.apiDashboardGetMetric.Describe(ch)
	c.apiDashboardSaveMetric.Describe(ch)
	c.apiDashboardSearchMetric.Describe(ch)
	c.apiDashboardSnapshotCreateMetric.Describe(ch)
	c.apiDashboardSnapshotExternalMetric.Describe(ch)
	c.apiDashboardSnapshotGetMetric.Describe(ch)
	c.apiDataproxyRequestAllMetric.Describe(ch)
	c.apiLoginOauthMetric.Describe(ch)
	c.apiLoginPostMetric.Describe(ch)
	c.apiOrgCreateMetric.Describe(ch)
	c.apiResponsesMetric.Describe(ch)
	c.apiUserSignupsCompletedMetric.Describe(ch)
	c.apiUserSignupsInviteMetric.Describe(ch)
	c.apiUserSignupsStartedMetric.Describe(ch)
	c.awsCloudwatchGetMetricStatisticsMetric.Describe(ch)
	c.awsCloudwatchListMetricsMetric.Describe(ch)
	c.instanceStartMetric.Describe(ch)
	c.modelsDashboardInsertMetric.Describe(ch)
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

	c.alertingExecutionTimeMetric.WithLabelValues("count").Set(float64(metrics.AlertingExecutionTime.Count))
	c.alertingExecutionTimeMetric.WithLabelValues("max").Set(float64(metrics.AlertingExecutionTime.Max))
	c.alertingExecutionTimeMetric.WithLabelValues("mean").Set(metrics.AlertingExecutionTime.Mean)
	c.alertingExecutionTimeMetric.WithLabelValues("min").Set(float64(metrics.AlertingExecutionTime.Min))
	c.alertingExecutionTimeMetric.WithLabelValues("p25").Set(metrics.AlertingExecutionTime.P25)
	c.alertingExecutionTimeMetric.WithLabelValues("p75").Set(metrics.AlertingExecutionTime.P75)
	c.alertingExecutionTimeMetric.WithLabelValues("p90").Set(metrics.AlertingExecutionTime.P90)
	c.alertingExecutionTimeMetric.WithLabelValues("p99").Set(metrics.AlertingExecutionTime.P99)
	c.alertingExecutionTimeMetric.WithLabelValues("std").Set(metrics.AlertingExecutionTime.Std)
	c.alertingExecutionTimeMetric.Collect(ch)

	c.alertingNotificationsSentMetric.WithLabelValues("line").Set(float64(metrics.AlertingNotificationsSentLine.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("dingding").Set(float64(metrics.AlertingNotificationsSentDingDing.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("email").Set(float64(metrics.AlertingNotificationsSentEmail.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("opsgenie").Set(float64(metrics.AlertingNotificationsSentOpsgenie.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("pagerduty").Set(float64(metrics.AlertingNotificationsSentPagerduty.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("pushover").Set(float64(metrics.AlertingNotificationsSentPushover.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("sensu").Set(float64(metrics.AlertingNotificationsSentSensu.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("slack").Set(float64(metrics.AlertingNotificationsSentSlack.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("telegram").Set(float64(metrics.AlertingNotificationsSentTelegram.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("threema").Set(float64(metrics.AlertingNotificationsSentThreema.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("victorops").Set(float64(metrics.AlertingNotificationsSentVictorops.Count))
	c.alertingNotificationsSentMetric.WithLabelValues("webhook").Set(float64(metrics.AlertingNotificationsSentWebhook.Count))
	c.alertingNotificationsSentMetric.Collect(ch)

	c.alertingResultsMetric.WithLabelValues("alerting").Set(float64(metrics.AlertingResultStateAlerting.Count))
	c.alertingResultsMetric.WithLabelValues("no_data").Set(float64(metrics.AlertingResultStateNoData.Count))
	c.alertingResultsMetric.WithLabelValues("ok").Set(float64(metrics.AlertingResultStateOk.Count))
	c.alertingResultsMetric.WithLabelValues("paused").Set(float64(metrics.AlertingResultStatePaused.Count))
	c.alertingResultsMetric.WithLabelValues("pending").Set(float64(metrics.AlertingResultStatePending.Count))
	c.alertingResultsMetric.Collect(ch)

	c.apiAdminUserCreateMetric.Set(float64(metrics.APIAdminUserCreate.Count))
	c.apiAdminUserCreateMetric.Collect(ch)

	c.apiDashboardGetMetric.WithLabelValues("count").Set(float64(metrics.APIDashboardGet.Count))
	c.apiDashboardGetMetric.WithLabelValues("max").Set(float64(metrics.APIDashboardGet.Max))
	c.apiDashboardGetMetric.WithLabelValues("mean").Set(metrics.APIDashboardGet.Mean)
	c.apiDashboardGetMetric.WithLabelValues("min").Set(float64(metrics.APIDashboardGet.Min))
	c.apiDashboardGetMetric.WithLabelValues("p25").Set(metrics.APIDashboardGet.P25)
	c.apiDashboardGetMetric.WithLabelValues("p75").Set(metrics.APIDashboardGet.P75)
	c.apiDashboardGetMetric.WithLabelValues("p90").Set(metrics.APIDashboardGet.P90)
	c.apiDashboardGetMetric.WithLabelValues("p99").Set(metrics.APIDashboardGet.P99)
	c.apiDashboardGetMetric.WithLabelValues("std").Set(metrics.APIDashboardGet.Std)
	c.apiDashboardGetMetric.Collect(ch)

	c.apiDashboardSaveMetric.WithLabelValues("count").Set(float64(metrics.APIDashboardSave.Count))
	c.apiDashboardSaveMetric.WithLabelValues("max").Set(float64(metrics.APIDashboardSave.Max))
	c.apiDashboardSaveMetric.WithLabelValues("mean").Set(metrics.APIDashboardSave.Mean)
	c.apiDashboardSaveMetric.WithLabelValues("min").Set(float64(metrics.APIDashboardSave.Min))
	c.apiDashboardSaveMetric.WithLabelValues("p25").Set(metrics.APIDashboardSave.P25)
	c.apiDashboardSaveMetric.WithLabelValues("p75").Set(metrics.APIDashboardSave.P75)
	c.apiDashboardSaveMetric.WithLabelValues("p90").Set(metrics.APIDashboardSave.P90)
	c.apiDashboardSaveMetric.WithLabelValues("p99").Set(metrics.APIDashboardSave.P99)
	c.apiDashboardSaveMetric.WithLabelValues("std").Set(metrics.APIDashboardSave.Std)
	c.apiDashboardSaveMetric.Collect(ch)

	c.apiDashboardSearchMetric.WithLabelValues("count").Set(float64(metrics.APIDashboardSearch.Count))
	c.apiDashboardSearchMetric.WithLabelValues("max").Set(float64(metrics.APIDashboardSearch.Max))
	c.apiDashboardSearchMetric.WithLabelValues("mean").Set(metrics.APIDashboardSearch.Mean)
	c.apiDashboardSearchMetric.WithLabelValues("min").Set(float64(metrics.APIDashboardSearch.Min))
	c.apiDashboardSearchMetric.WithLabelValues("p25").Set(metrics.APIDashboardSearch.P25)
	c.apiDashboardSearchMetric.WithLabelValues("p75").Set(metrics.APIDashboardSearch.P75)
	c.apiDashboardSearchMetric.WithLabelValues("p90").Set(metrics.APIDashboardSearch.P90)
	c.apiDashboardSearchMetric.WithLabelValues("p99").Set(metrics.APIDashboardSearch.P99)
	c.apiDashboardSearchMetric.WithLabelValues("std").Set(metrics.APIDashboardSearch.Std)
	c.apiDashboardSearchMetric.Collect(ch)

	c.apiDashboardSnapshotCreateMetric.Set(float64(metrics.APIDashboardSnapshotCreate.Count))
	c.apiDashboardSnapshotCreateMetric.Collect(ch)

	c.apiDashboardSnapshotExternalMetric.Set(float64(metrics.APIDashboardSnapshotExternal.Count))
	c.apiDashboardSnapshotExternalMetric.Collect(ch)

	c.apiDashboardSnapshotGetMetric.Set(float64(metrics.APIDashboardSnapshotGet.Count))
	c.apiDashboardSnapshotGetMetric.Collect(ch)

	c.apiDataproxyRequestAllMetric.WithLabelValues("count").Set(float64(metrics.APIDataproxyRequestAll.Count))
	c.apiDataproxyRequestAllMetric.WithLabelValues("max").Set(float64(metrics.APIDataproxyRequestAll.Max))
	c.apiDataproxyRequestAllMetric.WithLabelValues("mean").Set(metrics.APIDataproxyRequestAll.Mean)
	c.apiDataproxyRequestAllMetric.WithLabelValues("min").Set(float64(metrics.APIDataproxyRequestAll.Min))
	c.apiDataproxyRequestAllMetric.WithLabelValues("p25").Set(metrics.APIDataproxyRequestAll.P25)
	c.apiDataproxyRequestAllMetric.WithLabelValues("p75").Set(metrics.APIDataproxyRequestAll.P75)
	c.apiDataproxyRequestAllMetric.WithLabelValues("p90").Set(metrics.APIDataproxyRequestAll.P90)
	c.apiDataproxyRequestAllMetric.WithLabelValues("p99").Set(metrics.APIDataproxyRequestAll.P99)
	c.apiDataproxyRequestAllMetric.WithLabelValues("std").Set(metrics.APIDataproxyRequestAll.Std)
	c.apiDataproxyRequestAllMetric.Collect(ch)

	c.apiLoginOauthMetric.Set(float64(metrics.APILoginOauth.Count))
	c.apiLoginOauthMetric.Collect(ch)

	c.apiLoginPostMetric.Set(float64(metrics.APILoginPost.Count))
	c.apiLoginPostMetric.Collect(ch)

	c.apiOrgCreateMetric.Set(float64(metrics.APIOrgCreate.Count))
	c.apiOrgCreateMetric.Collect(ch)

	c.apiResponsesMetric.WithLabelValues("200").Set(float64(metrics.APIRespStatusCode200.Count))
	c.apiResponsesMetric.WithLabelValues("404").Set(float64(metrics.APIRespStatusCode404.Count))
	c.apiResponsesMetric.WithLabelValues("500").Set(float64(metrics.APIRespStatusCode500.Count))
	c.apiResponsesMetric.WithLabelValues("unknown").Set(float64(metrics.APIRespStatusCodeUnknown.Count))
	c.apiResponsesMetric.Collect(ch)

	c.apiUserSignupsCompletedMetric.Set(float64(metrics.APIUserSignupCompleted.Count))
	c.apiUserSignupsCompletedMetric.Collect(ch)

	c.apiUserSignupsInviteMetric.Set(float64(metrics.APIUserSignupInvite.Count))
	c.apiUserSignupsInviteMetric.Collect(ch)

	c.apiUserSignupsStartedMetric.Set(float64(metrics.APIUserSignupStarted.Count))
	c.apiUserSignupsStartedMetric.Collect(ch)

	c.awsCloudwatchGetMetricStatisticsMetric.Set(float64(metrics.AWSCloudwatchGetMetricStatistics.Count))
	c.awsCloudwatchGetMetricStatisticsMetric.Collect(ch)

	c.awsCloudwatchListMetricsMetric.Set(float64(metrics.AWSCloudwatchListMetrics.Count))
	c.awsCloudwatchListMetricsMetric.Collect(ch)

	c.instanceStartMetric.Set(float64(metrics.InstanceStart.Count))
	c.instanceStartMetric.Collect(ch)

	c.modelsDashboardInsertMetric.Set(float64(metrics.ModelsDashboardInsert.Count))
	c.modelsDashboardInsertMetric.Collect(ch)

	c.pageResponsesMetric.WithLabelValues("200").Set(float64(metrics.PageRespStatusCode200.Count))
	c.pageResponsesMetric.WithLabelValues("404").Set(float64(metrics.PageRespStatusCode404.Count))
	c.pageResponsesMetric.WithLabelValues("500").Set(float64(metrics.PageRespStatusCode500.Count))
	c.pageResponsesMetric.WithLabelValues("unknown").Set(float64(metrics.PageRespStatusCodeUnknown.Count))
	c.pageResponsesMetric.Collect(ch)

	c.proxyResponsesMetric.WithLabelValues("200").Set(float64(metrics.ProxyRespStatusCode200.Count))
	c.proxyResponsesMetric.WithLabelValues("404").Set(float64(metrics.ProxyRespStatusCode404.Count))
	c.proxyResponsesMetric.WithLabelValues("500").Set(float64(metrics.ProxyRespStatusCode500.Count))
	c.proxyResponsesMetric.WithLabelValues("unknown").Set(float64(metrics.ProxyRespStatusCodeUnknown.Count))
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
