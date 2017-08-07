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

var _ = Describe("MetricsCollectors", func() {
	var (
		grafanaClient *grafanafakes.FakeClient

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

		alertingActiveAlertsValue               = 1
		alertingExecutionTimeCount              = 2
		alertingExecutionTimeMax                = 3
		alertingExecutionTimeMean               = 4.1
		alertingExecutionTimeMin                = 5
		alertingExecutionTimeP25                = 6.1
		alertingExecutionTimeP75                = 7.1
		alertingExecutionTimeP90                = 8.1
		alertingExecutionTimeP99                = 9.1
		alertingExecutionTimeStd                = 10.1
		alertingNotificationsSentLineCount      = 11
		alertingNotificationsSentDingDingCount  = 12
		alertingNotificationsSentEmailCount     = 13
		alertingNotificationsSentOpsgenieCount  = 14
		alertingNotificationsSentPagerdutyCount = 15
		alertingNotificationsSentPushoverCount  = 16
		alertingNotificationsSentSensuCount     = 17
		alertingNotificationsSentSlackCount     = 18
		alertingNotificationsSentTelegramCount  = 19
		alertingNotificationsSentThreemaCount   = 20
		alertingNotificationsSentVictoropsCount = 21
		alertingNotificationsSentWebhookCount   = 22
		alertingResultStateAlertingCount        = 23
		alertingResultStateNoDataCount          = 24
		alertingResultStateOkCount              = 25
		alertingResultStatePausedCount          = 26
		alertingResultStatePendingCount         = 27
		apiAdminUserCreateCount                 = 28
		apiDashboardGetCount                    = 29
		apiDashboardGetMax                      = 30
		apiDashboardGetMean                     = 31.1
		apiDashboardGetMin                      = 32
		apiDashboardGetP25                      = 33.1
		apiDashboardGetP75                      = 34.1
		apiDashboardGetP90                      = 35.1
		apiDashboardGetP99                      = 36.1
		apiDashboardGetStd                      = 37.1
		apiDashboardSaveCount                   = 38
		apiDashboardSaveMax                     = 39
		apiDashboardSaveMean                    = 40.1
		apiDashboardSaveMin                     = 41
		apiDashboardSaveP25                     = 42.1
		apiDashboardSaveP75                     = 43.1
		apiDashboardSaveP90                     = 44.1
		apiDashboardSaveP99                     = 45.1
		apiDashboardSaveStd                     = 46.1
		apiDashboardSearchCount                 = 47
		apiDashboardSearchMax                   = 48
		apiDashboardSearchMean                  = 49.1
		apiDashboardSearchMin                   = 50
		apiDashboardSearchP25                   = 51.1
		apiDashboardSearchP75                   = 52.1
		apiDashboardSearchP90                   = 53.1
		apiDashboardSearchP99                   = 54.1
		apiDashboardSearchStd                   = 55.1
		apiDashboardSnapshotCreateCount         = 56
		apiDashboardSnapshotExternalCount       = 57
		apiDashboardSnapshotGetCount            = 58
		apiDataproxyRequestAllCount             = 59
		apiDataproxyRequestAllMax               = 60
		apiDataproxyRequestAllMean              = 61.1
		apiDataproxyRequestAllMin               = 62
		apiDataproxyRequestAllP25               = 63.1
		apiDataproxyRequestAllP75               = 64.1
		apiDataproxyRequestAllP90               = 65.1
		apiDataproxyRequestAllP99               = 66.1
		apiDataproxyRequestAllStd               = 67.1
		apiLoginOauthCount                      = 68
		apiLoginPostCount                       = 69
		apiOrgCreateCount                       = 70
		apiRespStatusCode200Count               = 71
		apiRespStatusCode404Count               = 72
		apiRespStatusCode500Count               = 73
		apiRespStatusCodeUnknownCount           = 74
		apiUserSignupCompletedCount             = 75
		apiUserSignupInviteCount                = 76
		apiUserSignupStartedCount               = 77
		awsCloudwatchGetMetricStatisticsCount   = 78
		awsCloudwatchListMetricsCount           = 79
		instanceStartCount                      = 80
		modelsDashboardInsertCount              = 81
		pageRespStatusCode200Count              = 82
		pageRespStatusCode404Count              = 83
		pageRespStatusCode500Count              = 84
		pageRespStatusCodeUnknownCount          = 85
		proxyRespStatusCode200Count             = 86
		proxyRespStatusCode404Count             = 87
		proxyRespStatusCode500Count             = 88
		proxyRespStatusCodeUnknownCount         = 89
		statsTotalsStatDashboardsValue          = 90
		statsTotalsStatOrgsValue                = 91
		statsTotalsStatPlaylistsValue           = 92
		statsTotalsStatUsersValue               = 93

		metricsCollector *MetricsCollector
	)

	BeforeEach(func() {
		grafanaClient = &grafanafakes.FakeClient{}

		alertingActiveAlertsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "alerting_active_alerts",
				Help:      "Number of active alerts.",
			},
		)
		alertingActiveAlertsMetric.Set(float64(alertingActiveAlertsValue))

		alertingExecutionTimeMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "alerting_execution_time",
				Help:      "Alerting execution time.",
			},
			[]string{"metric"},
		)
		alertingExecutionTimeMetric.WithLabelValues("count").Set(float64(alertingExecutionTimeCount))
		alertingExecutionTimeMetric.WithLabelValues("max").Set(float64(alertingExecutionTimeMax))
		alertingExecutionTimeMetric.WithLabelValues("mean").Set(alertingExecutionTimeMean)
		alertingExecutionTimeMetric.WithLabelValues("min").Set(float64(alertingExecutionTimeMin))
		alertingExecutionTimeMetric.WithLabelValues("p25").Set(alertingExecutionTimeP25)
		alertingExecutionTimeMetric.WithLabelValues("p75").Set(alertingExecutionTimeP75)
		alertingExecutionTimeMetric.WithLabelValues("p90").Set(alertingExecutionTimeP90)
		alertingExecutionTimeMetric.WithLabelValues("p99").Set(alertingExecutionTimeP99)
		alertingExecutionTimeMetric.WithLabelValues("std").Set(alertingExecutionTimeStd)

		alertingNotificationsSentMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "alerting_notifications_sent",
				Help:      "Number of alert notifications sent.",
			},
			[]string{"type"},
		)
		alertingNotificationsSentMetric.WithLabelValues("line").Set(float64(alertingNotificationsSentLineCount))
		alertingNotificationsSentMetric.WithLabelValues("dingding").Set(float64(alertingNotificationsSentDingDingCount))
		alertingNotificationsSentMetric.WithLabelValues("email").Set(float64(alertingNotificationsSentEmailCount))
		alertingNotificationsSentMetric.WithLabelValues("opsgenie").Set(float64(alertingNotificationsSentOpsgenieCount))
		alertingNotificationsSentMetric.WithLabelValues("pagerduty").Set(float64(alertingNotificationsSentPagerdutyCount))
		alertingNotificationsSentMetric.WithLabelValues("pushover").Set(float64(alertingNotificationsSentPushoverCount))
		alertingNotificationsSentMetric.WithLabelValues("sensu").Set(float64(alertingNotificationsSentSensuCount))
		alertingNotificationsSentMetric.WithLabelValues("slack").Set(float64(alertingNotificationsSentSlackCount))
		alertingNotificationsSentMetric.WithLabelValues("telegram").Set(float64(alertingNotificationsSentTelegramCount))
		alertingNotificationsSentMetric.WithLabelValues("threema").Set(float64(alertingNotificationsSentThreemaCount))
		alertingNotificationsSentMetric.WithLabelValues("victorops").Set(float64(alertingNotificationsSentVictoropsCount))
		alertingNotificationsSentMetric.WithLabelValues("webhook").Set(float64(alertingNotificationsSentWebhookCount))

		alertingResultsMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "alerting_results",
				Help:      "Number of alerting results.",
			},
			[]string{"state"},
		)
		alertingResultsMetric.WithLabelValues("alerting").Set(float64(alertingResultStateAlertingCount))
		alertingResultsMetric.WithLabelValues("no_data").Set(float64(alertingResultStateNoDataCount))
		alertingResultsMetric.WithLabelValues("ok").Set(float64(alertingResultStateOkCount))
		alertingResultsMetric.WithLabelValues("paused").Set(float64(alertingResultStatePausedCount))
		alertingResultsMetric.WithLabelValues("pending").Set(float64(alertingResultStatePendingCount))

		apiAdminUserCreateMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_admin_user_create",
				Help:      "Number of calls to Admin User Create API.",
			},
		)
		apiAdminUserCreateMetric.Set(float64(apiAdminUserCreateCount))

		apiDashboardGetMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_dashboard_get",
				Help:      "Dashboard Get API times.",
			},
			[]string{"metric"},
		)
		apiDashboardGetMetric.WithLabelValues("count").Set(float64(apiDashboardGetCount))
		apiDashboardGetMetric.WithLabelValues("max").Set(float64(apiDashboardGetMax))
		apiDashboardGetMetric.WithLabelValues("mean").Set(apiDashboardGetMean)
		apiDashboardGetMetric.WithLabelValues("min").Set(float64(apiDashboardGetMin))
		apiDashboardGetMetric.WithLabelValues("p25").Set(apiDashboardGetP25)
		apiDashboardGetMetric.WithLabelValues("p75").Set(apiDashboardGetP75)
		apiDashboardGetMetric.WithLabelValues("p90").Set(apiDashboardGetP90)
		apiDashboardGetMetric.WithLabelValues("p99").Set(apiDashboardGetP99)
		apiDashboardGetMetric.WithLabelValues("std").Set(apiDashboardGetStd)

		apiDashboardSaveMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_dashboard_save",
				Help:      "Dashboard Save API times.",
			},
			[]string{"metric"},
		)
		apiDashboardSaveMetric.WithLabelValues("count").Set(float64(apiDashboardSaveCount))
		apiDashboardSaveMetric.WithLabelValues("max").Set(float64(apiDashboardSaveMax))
		apiDashboardSaveMetric.WithLabelValues("mean").Set(apiDashboardSaveMean)
		apiDashboardSaveMetric.WithLabelValues("min").Set(float64(apiDashboardSaveMin))
		apiDashboardSaveMetric.WithLabelValues("p25").Set(apiDashboardSaveP25)
		apiDashboardSaveMetric.WithLabelValues("p75").Set(apiDashboardSaveP75)
		apiDashboardSaveMetric.WithLabelValues("p90").Set(apiDashboardSaveP90)
		apiDashboardSaveMetric.WithLabelValues("p99").Set(apiDashboardSaveP99)
		apiDashboardSaveMetric.WithLabelValues("std").Set(apiDashboardSaveStd)

		apiDashboardSearchMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_dashboard_search",
				Help:      "Dashboard Search API times.",
			},
			[]string{"metric"},
		)
		apiDashboardSearchMetric.WithLabelValues("count").Set(float64(apiDashboardSearchCount))
		apiDashboardSearchMetric.WithLabelValues("max").Set(float64(apiDashboardSearchMax))
		apiDashboardSearchMetric.WithLabelValues("mean").Set(apiDashboardSearchMean)
		apiDashboardSearchMetric.WithLabelValues("min").Set(float64(apiDashboardSearchMin))
		apiDashboardSearchMetric.WithLabelValues("p25").Set(apiDashboardSearchP25)
		apiDashboardSearchMetric.WithLabelValues("p75").Set(apiDashboardSearchP75)
		apiDashboardSearchMetric.WithLabelValues("p90").Set(apiDashboardSearchP90)
		apiDashboardSearchMetric.WithLabelValues("p99").Set(apiDashboardSearchP99)
		apiDashboardSearchMetric.WithLabelValues("std").Set(apiDashboardSearchStd)

		apiDashboardSnapshotCreateMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_dashboard_snapshot_create",
				Help:      "Number of calls to Dashboard Snapshot Create API.",
			},
		)
		apiDashboardSnapshotCreateMetric.Set(float64(apiDashboardSnapshotCreateCount))

		apiDashboardSnapshotExternalMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_dashboard_snapshot_external",
				Help:      "Number of calls to Dashboard Snapshot External API.",
			},
		)
		apiDashboardSnapshotExternalMetric.Set(float64(apiDashboardSnapshotExternalCount))

		apiDashboardSnapshotGetMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_dashboard_snapshot_get",
				Help:      "Number of calls to Dashboard Snapshot Get API.",
			},
		)
		apiDashboardSnapshotGetMetric.Set(float64(apiDashboardSnapshotGetCount))

		apiDataproxyRequestAllMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_dataproxy_request_all",
				Help:      "Dataproxy request API times.",
			},
			[]string{"metric"},
		)
		apiDataproxyRequestAllMetric.WithLabelValues("count").Set(float64(apiDataproxyRequestAllCount))
		apiDataproxyRequestAllMetric.WithLabelValues("max").Set(float64(apiDataproxyRequestAllMax))
		apiDataproxyRequestAllMetric.WithLabelValues("mean").Set(apiDataproxyRequestAllMean)
		apiDataproxyRequestAllMetric.WithLabelValues("min").Set(float64(apiDataproxyRequestAllMin))
		apiDataproxyRequestAllMetric.WithLabelValues("p25").Set(apiDataproxyRequestAllP25)
		apiDataproxyRequestAllMetric.WithLabelValues("p75").Set(apiDataproxyRequestAllP75)
		apiDataproxyRequestAllMetric.WithLabelValues("p90").Set(apiDataproxyRequestAllP90)
		apiDataproxyRequestAllMetric.WithLabelValues("p99").Set(apiDataproxyRequestAllP99)
		apiDataproxyRequestAllMetric.WithLabelValues("std").Set(apiDataproxyRequestAllStd)

		apiLoginOauthMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_login_oauth",
				Help:      "Number of calls to Login OAuth API.",
			},
		)
		apiLoginOauthMetric.Set(float64(apiLoginOauthCount))

		apiLoginPostMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_login_post",
				Help:      "Number of calls to Login Post API.",
			},
		)
		apiLoginPostMetric.Set(float64(apiLoginPostCount))

		apiOrgCreateMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_org_create",
				Help:      "Number of calls to Org Create API.",
			},
		)
		apiOrgCreateMetric.Set(float64(apiOrgCreateCount))

		apiResponsesMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_responses",
				Help:      "Number of API responses.",
			},
			[]string{"code"},
		)
		apiResponsesMetric.WithLabelValues("200").Set(float64(apiRespStatusCode200Count))
		apiResponsesMetric.WithLabelValues("404").Set(float64(apiRespStatusCode404Count))
		apiResponsesMetric.WithLabelValues("500").Set(float64(apiRespStatusCode500Count))
		apiResponsesMetric.WithLabelValues("unknown").Set(float64(apiRespStatusCodeUnknownCount))

		apiUserSignupsCompletedMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_user_signups_completed",
				Help:      "Number of API User Signups completed.",
			},
		)
		apiUserSignupsCompletedMetric.Set(float64(apiUserSignupCompletedCount))

		apiUserSignupsInviteMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_user_signups_invite",
				Help:      "Number of API User Signups invite.",
			},
		)
		apiUserSignupsInviteMetric.Set(float64(apiUserSignupInviteCount))

		apiUserSignupsStartedMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_user_signups_started",
				Help:      "Number of API User Signups started.",
			},
		)
		apiUserSignupsStartedMetric.Set(float64(apiUserSignupStartedCount))

		awsCloudwatchGetMetricStatisticsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "aws_cloudwatch_get_metric_statistics",
				Help:      "Number of calls to AWS CloudWatch Get Metric Statistics API.",
			},
		)
		awsCloudwatchGetMetricStatisticsMetric.Set(float64(awsCloudwatchGetMetricStatisticsCount))

		awsCloudwatchListMetricsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "aws_cloudwatch_list_metrics",
				Help:      "Number of calls to AWS CloudWatch List Metrics API.",
			},
		)
		awsCloudwatchListMetricsMetric.Set(float64(awsCloudwatchListMetricsCount))

		instanceStartMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "instance_start",
				Help:      "Number of Instance Starts.",
			},
		)
		instanceStartMetric.Set(float64(instanceStartCount))

		modelsDashboardInsertMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "models_dashboard_insert",
				Help:      "Number of Dashboard inserts.",
			},
		)
		modelsDashboardInsertMetric.Set(float64(modelsDashboardInsertCount))

		pageResponsesMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "page_responses",
				Help:      "Number of Page responses.",
			},
			[]string{"code"},
		)
		pageResponsesMetric.WithLabelValues("200").Set(float64(pageRespStatusCode200Count))
		pageResponsesMetric.WithLabelValues("404").Set(float64(pageRespStatusCode404Count))
		pageResponsesMetric.WithLabelValues("500").Set(float64(pageRespStatusCode500Count))
		pageResponsesMetric.WithLabelValues("unknown").Set(float64(pageRespStatusCodeUnknownCount))

		proxyResponsesMetric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "proxy_responses",
				Help:      "Number of Proxy responses.",
			},
			[]string{"code"},
		)
		proxyResponsesMetric.WithLabelValues("200").Set(float64(proxyRespStatusCode200Count))
		proxyResponsesMetric.WithLabelValues("404").Set(float64(proxyRespStatusCode404Count))
		proxyResponsesMetric.WithLabelValues("500").Set(float64(proxyRespStatusCode500Count))
		proxyResponsesMetric.WithLabelValues("unknown").Set(float64(proxyRespStatusCodeUnknownCount))

		dashboardsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "dashboards",
				Help:      "Number of dashboards.",
			},
		)
		dashboardsMetric.Set(float64(statsTotalsStatDashboardsValue))

		orgsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "orgs",
				Help:      "Number of orgs.",
			},
		)
		orgsMetric.Set(float64(statsTotalsStatOrgsValue))

		playlistsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "playlists",
				Help:      "Number of playlists.",
			},
		)
		playlistsMetric.Set(float64(statsTotalsStatPlaylistsValue))

		usersMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "users",
				Help:      "Number of users.",
			},
		)
		usersMetric.Set(float64(statsTotalsStatUsersValue))

		scrapesTotalMetric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "scrapes_total",
				Help:      "Total number of Grafana metrics scrapes.",
			},
		)
		scrapesTotalMetric.Inc()

		scrapeErrorsTotalMetric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "scrape_errors_total",
				Help:      "Total number of Grafana metrics scrape errors.",
			},
		)

		lastScrapeErrorMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "last_scrape_error",
				Help:      "Whether the last metrics scrape from Grafana resulted in an error (1 for error, 0 for success).",
			},
		)
		lastScrapeErrorMetric.Set(0)

		lastScrapeTimestampMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "last_scrape_timestamp",
				Help:      "Number of seconds since 1970 since last metrics scrape from Grafana.",
			},
		)

		lastScrapeDurationSecondsMetric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "last_scrape_duration_seconds",
				Help:      "Duration of the last metrics scrape from Grafana.",
			},
		)
	})

	JustBeforeEach(func() {
		metricsCollector = NewMetricsCollector(grafanaClient)
	})

	Describe("Describe", func() {
		var (
			descriptions chan *prometheus.Desc
		)

		BeforeEach(func() {
			descriptions = make(chan *prometheus.Desc)
		})

		JustBeforeEach(func() {
			go metricsCollector.Describe(descriptions)
		})

		It("returns a grafana_metrics_alerting_active_alerts metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(alertingActiveAlertsMetric.Desc())))
		})

		It("returns a grafana_metrics_alerting_execution_time metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(alertingExecutionTimeMetric.WithLabelValues("count").Desc())))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(alertingNotificationsSentMetric.WithLabelValues("line").Desc())))
		})

		It("returns a grafana_metrics_alerting_results metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(alertingResultsMetric.WithLabelValues("200").Desc())))
		})

		It("returns a grafana_metrics_api_admin_user_create metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiAdminUserCreateMetric.Desc())))
		})

		It("returns a grafana_metrics_api_dashboard_get metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiDashboardGetMetric.WithLabelValues("count").Desc())))
		})

		It("returns a grafana_metrics_api_dashboard_save metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiDashboardSaveMetric.WithLabelValues("count").Desc())))
		})

		It("returns a grafana_metrics_api_dashboard_search metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiDashboardSearchMetric.WithLabelValues("count").Desc())))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_create metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiDashboardSnapshotCreateMetric.Desc())))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_external metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiDashboardSnapshotExternalMetric.Desc())))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_get metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiDashboardSnapshotGetMetric.Desc())))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiDataproxyRequestAllMetric.WithLabelValues("count").Desc())))
		})

		It("returns a grafana_metrics_api_login_oauth metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiLoginOauthMetric.Desc())))
		})

		It("returns a grafana_metrics_api_login_post metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiLoginPostMetric.Desc())))
		})

		It("returns a grafana_metrics_api_org_create metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiOrgCreateMetric.Desc())))
		})

		It("returns a grafana_metrics_api_responses metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiResponsesMetric.WithLabelValues("200").Desc())))
		})

		It("returns a grafana_metrics_api_user_signups_completed metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiUserSignupsCompletedMetric.Desc())))
		})

		It("returns a grafana_metrics_api_user_signups_invite metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiUserSignupsInviteMetric.Desc())))
		})

		It("returns a grafana_metrics_api_user_signups_started metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiUserSignupsStartedMetric.Desc())))
		})

		It("returns a grafana_metrics_aws_cloudwatch_get_metric_statistics metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(awsCloudwatchGetMetricStatisticsMetric.Desc())))
		})

		It("returns a grafana_metrics_aws_cloudwatch_list_metrics metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(awsCloudwatchListMetricsMetric.Desc())))
		})

		It("returns a grafana_metrics_instance_start metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(instanceStartMetric.Desc())))
		})

		It("returns a grafana_metrics_models_dashboard_insert metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(modelsDashboardInsertMetric.Desc())))
		})

		It("returns a grafana_metrics_page_responses metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(pageResponsesMetric.WithLabelValues("200").Desc())))
		})

		It("returns a grafana_metrics_proxy_responses metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(proxyResponsesMetric.WithLabelValues("200").Desc())))
		})

		It("returns a grafana_metrics_dashboards metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(dashboardsMetric.Desc())))
		})

		It("returns a grafana_metrics_orgs metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(orgsMetric.Desc())))
		})

		It("returns a grafana_metrics_playlists metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(playlistsMetric.Desc())))
		})

		It("returns a grafana_metrics_users metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(usersMetric.Desc())))
		})

		It("returns a grafana_metrics_scrapes_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(scrapesTotalMetric.Desc())))
		})

		It("returns a grafana_metrics_scrape_errors_total metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(scrapeErrorsTotalMetric.Desc())))
		})

		It("returns a grafana_metrics_last_scrape_error metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(lastScrapeErrorMetric.Desc())))
		})

		It("returns a grafana_metrics_last_scrape_timestamp metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(lastScrapeTimestampMetric.Desc())))
		})

		It("returns a grafana_metrics_last_scrape_duration_seconds metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(lastScrapeDurationSecondsMetric.Desc())))
		})
	})

	Describe("Collect", func() {
		var (
			metricsResponse grafana.Metrics

			metrics chan prometheus.Metric
		)

		BeforeEach(func() {
			metricsResponse = grafana.Metrics{
				AlertingActiveAlerts: grafana.Gauge{
					Value: int64(alertingActiveAlertsValue),
				},
				AlertingExecutionTime: grafana.Timer{
					Count: int64(alertingExecutionTimeCount),
					Max:   int64(alertingExecutionTimeMax),
					Mean:  float64(alertingExecutionTimeMean),
					Min:   int64(alertingExecutionTimeMin),
					P25:   float64(alertingExecutionTimeP25),
					P75:   float64(alertingExecutionTimeP75),
					P90:   float64(alertingExecutionTimeP90),
					P99:   float64(alertingExecutionTimeP99),
					Std:   float64(alertingExecutionTimeStd),
				},
				AlertingNotificationsSentLine: grafana.Counter{
					Count: int64(alertingNotificationsSentLineCount),
				},
				AlertingNotificationsSentDingDing: grafana.Counter{
					Count: int64(alertingNotificationsSentDingDingCount),
				},
				AlertingNotificationsSentEmail: grafana.Counter{
					Count: int64(alertingNotificationsSentEmailCount),
				},
				AlertingNotificationsSentOpsgenie: grafana.Counter{
					Count: int64(alertingNotificationsSentOpsgenieCount),
				},
				AlertingNotificationsSentPagerduty: grafana.Counter{
					Count: int64(alertingNotificationsSentPagerdutyCount),
				},
				AlertingNotificationsSentPushover: grafana.Counter{
					Count: int64(alertingNotificationsSentPushoverCount),
				},
				AlertingNotificationsSentSensu: grafana.Counter{
					Count: int64(alertingNotificationsSentSensuCount),
				},
				AlertingNotificationsSentSlack: grafana.Counter{
					Count: int64(alertingNotificationsSentSlackCount),
				},
				AlertingNotificationsSentTelegram: grafana.Counter{
					Count: int64(alertingNotificationsSentTelegramCount),
				},
				AlertingNotificationsSentThreema: grafana.Counter{
					Count: int64(alertingNotificationsSentThreemaCount),
				},
				AlertingNotificationsSentVictorops: grafana.Counter{
					Count: int64(alertingNotificationsSentVictoropsCount),
				},
				AlertingNotificationsSentWebhook: grafana.Counter{
					Count: int64(alertingNotificationsSentWebhookCount),
				},
				AlertingResultStateAlerting: grafana.Counter{
					Count: int64(alertingResultStateAlertingCount),
				},
				AlertingResultStateNoData: grafana.Counter{
					Count: int64(alertingResultStateNoDataCount),
				},
				AlertingResultStateOk: grafana.Counter{
					Count: int64(alertingResultStateOkCount),
				},
				AlertingResultStatePaused: grafana.Counter{
					Count: int64(alertingResultStatePausedCount),
				},
				AlertingResultStatePending: grafana.Counter{
					Count: int64(alertingResultStatePendingCount),
				},
				APIAdminUserCreate: grafana.Counter{
					Count: int64(apiAdminUserCreateCount),
				},
				APIDashboardGet: grafana.Timer{
					Count: int64(apiDashboardGetCount),
					Max:   int64(apiDashboardGetMax),
					Mean:  float64(apiDashboardGetMean),
					Min:   int64(apiDashboardGetMin),
					P25:   float64(apiDashboardGetP25),
					P75:   float64(apiDashboardGetP75),
					P90:   float64(apiDashboardGetP90),
					P99:   float64(apiDashboardGetP99),
					Std:   float64(apiDashboardGetStd),
				},
				APIDashboardSave: grafana.Timer{
					Count: int64(apiDashboardSaveCount),
					Max:   int64(apiDashboardSaveMax),
					Mean:  float64(apiDashboardSaveMean),
					Min:   int64(apiDashboardSaveMin),
					P25:   float64(apiDashboardSaveP25),
					P75:   float64(apiDashboardSaveP75),
					P90:   float64(apiDashboardSaveP90),
					P99:   float64(apiDashboardSaveP99),
					Std:   float64(apiDashboardSaveStd),
				},
				APIDashboardSearch: grafana.Timer{
					Count: int64(apiDashboardSearchCount),
					Max:   int64(apiDashboardSearchMax),
					Mean:  float64(apiDashboardSearchMean),
					Min:   int64(apiDashboardSearchMin),
					P25:   float64(apiDashboardSearchP25),
					P75:   float64(apiDashboardSearchP75),
					P90:   float64(apiDashboardSearchP90),
					P99:   float64(apiDashboardSearchP99),
					Std:   float64(apiDashboardSearchStd),
				},
				APIDashboardSnapshotCreate: grafana.Counter{
					Count: int64(apiDashboardSnapshotCreateCount),
				},
				APIDashboardSnapshotExternal: grafana.Counter{
					Count: int64(apiDashboardSnapshotExternalCount),
				},
				APIDashboardSnapshotGet: grafana.Counter{
					Count: int64(apiDashboardSnapshotGetCount),
				},
				APIDataproxyRequestAll: grafana.Timer{
					Count: int64(apiDataproxyRequestAllCount),
					Max:   int64(apiDataproxyRequestAllMax),
					Mean:  float64(apiDataproxyRequestAllMean),
					Min:   int64(apiDataproxyRequestAllMin),
					P25:   float64(apiDataproxyRequestAllP25),
					P75:   float64(apiDataproxyRequestAllP75),
					P90:   float64(apiDataproxyRequestAllP90),
					P99:   float64(apiDataproxyRequestAllP99),
					Std:   float64(apiDataproxyRequestAllStd),
				},
				APILoginOauth: grafana.Counter{
					Count: int64(apiLoginOauthCount),
				},
				APILoginPost: grafana.Counter{
					Count: int64(apiLoginPostCount),
				},
				APIOrgCreate: grafana.Counter{
					Count: int64(apiOrgCreateCount),
				},
				APIRespStatusCode200: grafana.Counter{
					Count: int64(apiRespStatusCode200Count),
				},
				APIRespStatusCode404: grafana.Counter{
					Count: int64(apiRespStatusCode404Count),
				},
				APIRespStatusCode500: grafana.Counter{
					Count: int64(apiRespStatusCode500Count),
				},
				APIRespStatusCodeUnknown: grafana.Counter{
					Count: int64(apiRespStatusCodeUnknownCount),
				},
				APIUserSignupCompleted: grafana.Counter{
					Count: int64(apiUserSignupCompletedCount),
				},
				APIUserSignupInvite: grafana.Counter{
					Count: int64(apiUserSignupInviteCount),
				},
				APIUserSignupStarted: grafana.Counter{
					Count: int64(apiUserSignupStartedCount),
				},
				AWSCloudwatchGetMetricStatistics: grafana.Counter{
					Count: int64(awsCloudwatchGetMetricStatisticsCount),
				},
				AWSCloudwatchListMetrics: grafana.Counter{
					Count: int64(awsCloudwatchListMetricsCount),
				},
				InstanceStart: grafana.Counter{
					Count: int64(instanceStartCount),
				},
				ModelsDashboardInsert: grafana.Counter{
					Count: int64(modelsDashboardInsertCount),
				},
				PageRespStatusCode200: grafana.Counter{
					Count: int64(pageRespStatusCode200Count),
				},
				PageRespStatusCode404: grafana.Counter{
					Count: int64(pageRespStatusCode404Count),
				},
				PageRespStatusCode500: grafana.Counter{
					Count: int64(pageRespStatusCode500Count),
				},
				PageRespStatusCodeUnknown: grafana.Counter{
					Count: int64(pageRespStatusCodeUnknownCount),
				},
				ProxyRespStatusCode200: grafana.Counter{
					Count: int64(proxyRespStatusCode200Count),
				},
				ProxyRespStatusCode404: grafana.Counter{
					Count: int64(proxyRespStatusCode404Count),
				},
				ProxyRespStatusCode500: grafana.Counter{
					Count: int64(proxyRespStatusCode500Count),
				},
				ProxyRespStatusCodeUnknown: grafana.Counter{
					Count: int64(proxyRespStatusCodeUnknownCount),
				},
				StatsTotalsStatDashboards: grafana.Gauge{
					Value: int64(statsTotalsStatDashboardsValue),
				},
				StatsTotalsStatOrgs: grafana.Gauge{
					Value: int64(statsTotalsStatOrgsValue),
				},
				StatsTotalsStatPlaylists: grafana.Gauge{
					Value: int64(statsTotalsStatPlaylistsValue),
				},
				StatsTotalsStatUsers: grafana.Gauge{
					Value: int64(statsTotalsStatUsersValue),
				},
			}
			grafanaClient.GetMetricsReturns(metricsResponse, nil)

			metrics = make(chan prometheus.Metric)
		})

		JustBeforeEach(func() {
			go metricsCollector.Collect(metrics)
		})

		It("returns a grafana_metrics_alerting_active_alerts metric", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingActiveAlertsMetric)))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric count", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("count"))))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric max", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("max"))))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric mean", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("mean"))))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric min", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("min"))))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric p25", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("p25"))))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric p75", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("p75"))))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric p90", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("p90"))))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric p99", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("p99"))))
		})

		It("returns a grafana_metrics_alerting_execution_time metric with a metric std", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingExecutionTimeMetric.WithLabelValues("std"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type line", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("line"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type dingding", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("dingding"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type email", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("email"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type opsgenie", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("opsgenie"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type pagerduty", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("pagerduty"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type pushover", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("pushover"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type sensu", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("sensu"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type slack", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("slack"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type telegram", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("telegram"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type threema", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("threema"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type victorops", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("victorops"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type webhook", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("webhook"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state alerting", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("alerting"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state no_data", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("no_data"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state ok", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("ok"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state paused", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("paused"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state pending", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("pending"))))
		})

		It("returns a grafana_metrics_api_admin_user_create metric with a state completed", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiAdminUserCreateMetric)))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric count", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("count"))))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric max", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("max"))))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric mean", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("mean"))))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric min", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("min"))))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric p25", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("p25"))))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric p75", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("p75"))))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric p90", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("p90"))))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric p99", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("p99"))))
		})

		It("returns a grafana_metrics_api_dashboard_get metric with a metric std", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardGetMetric.WithLabelValues("std"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric count", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("count"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric max", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("max"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric mean", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("mean"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric min", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("min"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric p25", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("p25"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric p75", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("p75"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric p90", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("p90"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric p99", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("p99"))))
		})

		It("returns a grafana_metrics_api_dashboard_save metric with a metric std", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSaveMetric.WithLabelValues("std"))))
		})

		It("returns a grafana_metrics_api_dashboard_search metric with a metric count", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("count"))))
		})

		It("returns a grafana_metrics_api_api_dashboard_search with a metric max", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("max"))))
		})

		It("returns a grafana_metrics_api_dashboard_search metric with a metric mean", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("mean"))))
		})

		It("returns a grafana_metrics_api_dashboard_search metric with a metric min", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("min"))))
		})

		It("returns a grafana_metrics_api_dashboard_search metric with a metric p25", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("p25"))))
		})

		It("returns a grafana_metrics_api_dashboard_search metric with a metric p75", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("p75"))))
		})

		It("returns a grafana_metrics_api_dashboard_search metric with a metric p90", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("p90"))))
		})

		It("returns a grafana_metrics_api_dashboard_search metric with a metric p99", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("p99"))))
		})

		It("returns a grafana_metrics_api_dashboard_search metric with a metric std", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSearchMetric.WithLabelValues("std"))))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_create metric with a state completed", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSnapshotCreateMetric)))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_external metric with a state completed", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSnapshotExternalMetric)))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_get metric with a state completed", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDashboardSnapshotGetMetric)))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric count", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("count"))))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric max", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("max"))))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric mean", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("mean"))))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric min", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("min"))))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric p25", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("p25"))))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric p75", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("p75"))))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric p90", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("p90"))))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric p99", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("p99"))))
		})

		It("returns a grafana_metrics_api_dataproxy_request_all metric with a metric std", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiDataproxyRequestAllMetric.WithLabelValues("std"))))
		})

		It("returns a grafana_metrics_api_login_post metric with a state completed", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiLoginPostMetric)))
		})

		It("returns a grafana_metrics_api_org_create metric with a state completed", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiOrgCreateMetric)))
		})

		It("returns a grafana_metrics_api_responses metric with a status_code 200", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiResponsesMetric.WithLabelValues("unknown"))))
		})

		It("returns a grafana_metrics_api_responses metric with a status_code 404", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiResponsesMetric.WithLabelValues("404"))))
		})

		It("returns a grafana_metrics_api_responses metric with a status_code 500", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiResponsesMetric.WithLabelValues("500"))))
		})

		It("returns a grafana_metrics_api_responses metric with a status_code unknown", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiResponsesMetric.WithLabelValues("unknown"))))
		})

		It("returns a grafana_metrics_api_user_signups_completed metric with a state completed", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiUserSignupsCompletedMetric)))
		})

		It("returns a grafana_metrics_api_user_signups_invite metric with a state inivte", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiUserSignupsInviteMetric)))
		})

		It("returns a grafana_metrics_api_user_signups_started metric with a state started", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(apiUserSignupsStartedMetric)))
		})

		It("returns a grafana_metrics_aws_cloudwatch_get_metric_statistics metric with a state started", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(awsCloudwatchGetMetricStatisticsMetric)))
		})

		It("returns a grafana_metrics_aws_cloudwatch_list_metrics metric with a state started", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(awsCloudwatchListMetricsMetric)))
		})

		It("returns a grafana_metrics_instance_start metric with a state started", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(instanceStartMetric)))
		})

		It("returns a grafana_metrics_models_dashboard_insert metric with a state started", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(modelsDashboardInsertMetric)))
		})

		It("returns a grafana_metrics_page_responses metric with a status_code 200", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(pageResponsesMetric.WithLabelValues("200"))))
		})

		It("returns a grafana_metrics_page_responses metric with a status_code 404", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(pageResponsesMetric.WithLabelValues("404"))))
		})

		It("returns a grafana_metrics_page_responses metric with a status_code 500", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(pageResponsesMetric.WithLabelValues("500"))))
		})

		It("returns a grafana_metrics_page_responses metric with a status_code unknown", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(pageResponsesMetric.WithLabelValues("unknown"))))
		})

		It("returns a grafana_metrics_proxy_responses metric with a status_code 200", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(proxyResponsesMetric.WithLabelValues("200"))))
		})

		It("returns a grafana_metrics_proxy_responses metric with a status_code 404", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(proxyResponsesMetric.WithLabelValues("404"))))
		})

		It("returns a grafana_metrics_proxy_responses metric with a status_code 500", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(proxyResponsesMetric.WithLabelValues("500"))))
		})

		It("returns a grafana_metrics_proxy_responses metric with a status_code unknown", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(proxyResponsesMetric.WithLabelValues("unknown"))))
		})

		It("returns a grafana_metrics_dashboards metric", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(dashboardsMetric)))
		})

		It("returns a grafana_metrics_orgs metric", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(orgsMetric)))
		})

		It("returns a grafana_metrics_playlists metric", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(playlistsMetric)))
		})

		It("returns a grafana_metrics_users metric", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(usersMetric)))
		})

		It("returns a grafana_metrics_scrapes_total metric", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(scrapesTotalMetric)))
		})

		It("returns a grafana_metrics_scrape_errors_total", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(scrapeErrorsTotalMetric)))
		})

		It("returns a grafana_metrics_last_scrape_error metric", func() {
			Eventually(metrics, "2s").Should(Receive(PrometheusMetric(lastScrapeErrorMetric)))
		})

		Context("when it fails to list the security groups", func() {
			BeforeEach(func() {
				grafanaClient.GetMetricsReturns(metricsResponse, errors.New("error"))

				scrapeErrorsTotalMetric.Inc()
				lastScrapeErrorMetric.Set(1)
			})

			It("returns a grafana_metrics_scrape_errors_total metric", func() {
				Eventually(metrics).Should(Receive(PrometheusMetric(scrapeErrorsTotalMetric)))
			})

			It("returns a grafana_metrics_last_scrape_error metric", func() {
				Eventually(metrics).Should(Receive(PrometheusMetric(lastScrapeErrorMetric)))
			})
		})
	})
})
