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
		alertingNotificationsSentMetric        *prometheus.GaugeVec
		alertingResultsMetric                  *prometheus.GaugeVec
		apiAdminUserCreateMetric               prometheus.Gauge
		apiDashboardSnapshotCreateMetric       prometheus.Gauge
		apiDashboardSnapshotExternalMetric     prometheus.Gauge
		apiDashboardSnapshotGetMetric          prometheus.Gauge
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
		alertingNotificationsSentLineCount      = 2
		alertingNotificationsSentDingDingCount  = 3
		alertingNotificationsSentEmailCount     = 4
		alertingNotificationsSentOpsgenieCount  = 5
		alertingNotificationsSentPagerdutyCount = 6
		alertingNotificationsSentPushoverCount  = 7
		alertingNotificationsSentSensuCount     = 8
		alertingNotificationsSentSlackCount     = 9
		alertingNotificationsSentTelegramCount  = 10
		alertingNotificationsSentThreemaCount   = 11
		alertingNotificationsSentVictoropsCount = 12
		alertingNotificationsSentWebhookCount   = 13
		alertingResultStateAlertingCount        = 14
		alertingResultStateNoDataCount          = 15
		alertingResultStateOkCount              = 16
		alertingResultStatePausedCount          = 17
		alertingResultStatePendingCount         = 18
		apiAdminUserCreateCount                 = 19
		apiDashboardSnapshotCreateCount         = 20
		apiDashboardSnapshotExternalCount       = 21
		apiDashboardSnapshotGetCount            = 22
		apiLoginOauthCount                      = 23
		apiLoginPostCount                       = 24
		apiOrgCreateCount                       = 25
		apiRespStatusCode200Count               = 26
		apiRespStatusCode404Count               = 27
		apiRespStatusCode500Count               = 28
		apiRespStatusCodeUnknownCount           = 29
		apiUserSignupCompletedCount             = 30
		apiUserSignupInviteCount                = 31
		apiUserSignupStartedCount               = 32
		awsCloudwatchGetMetricStatisticsCount   = 33
		awsCloudwatchListMetricsCount           = 34
		instanceStartCount                      = 35
		modelsDashboardInsertCount              = 36
		pageRespStatusCode200Count              = 37
		pageRespStatusCode404Count              = 38
		pageRespStatusCode500Count              = 39
		pageRespStatusCodeUnknownCount          = 40
		proxyRespStatusCode200Count             = 41
		proxyRespStatusCode404Count             = 42
		proxyRespStatusCode500Count             = 43
		proxyRespStatusCodeUnknownCount         = 44
		statsTotalsStatDashboardsValue          = 45
		statsTotalsStatOrgsValue                = 46
		statsTotalsStatPlaylistsValue           = 47
		statsTotalsStatUsersValue               = 48

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

		It("returns a grafana_metrics_alerting_notifications_sent metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(alertingNotificationsSentMetric.WithLabelValues("line").Desc())))
		})

		It("returns a grafana_metrics_alerting_results metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(alertingResultsMetric.WithLabelValues("200").Desc())))
		})

		It("returns a grafana_metrics_api_admin_user_create metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiAdminUserCreateMetric.Desc())))
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
					Count: int64(1),
					Max:   int64(2),
					Mean:  float64(1.1),
					Min:   int64(3),
					P25:   float64(2.2),
					P75:   float64(3.3),
					P90:   float64(3.4),
					P99:   float64(3.5),
					Std:   float64(3.6),
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
					Count: int64(1),
					Max:   int64(2),
					Mean:  float64(1.1),
					Min:   int64(3),
					P25:   float64(2.2),
					P75:   float64(3.3),
					P90:   float64(3.4),
					P99:   float64(3.5),
					Std:   float64(3.6),
				},
				APIDashboardSave: grafana.Timer{
					Count: int64(1),
					Max:   int64(2),
					Mean:  float64(1.1),
					Min:   int64(3),
					P25:   float64(2.2),
					P75:   float64(3.3),
					P90:   float64(3.4),
					P99:   float64(3.5),
					Std:   float64(3.6),
				},
				APIDashboardSearch: grafana.Timer{
					Count: int64(1),
					Max:   int64(2),
					Mean:  float64(1.1),
					Min:   int64(3),
					P25:   float64(2.2),
					P75:   float64(3.3),
					P90:   float64(3.4),
					P99:   float64(3.5),
					Std:   float64(3.6),
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
					Count: int64(1),
					Max:   int64(2),
					Mean:  float64(1.1),
					Min:   int64(3),
					P25:   float64(2.2),
					P75:   float64(3.3),
					P90:   float64(3.4),
					P99:   float64(3.5),
					Std:   float64(3.6),
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

		It("returns a grafana_admin_alerting_active_alerts metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingActiveAlertsMetric)))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type line", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("line"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type dingding", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("dingding"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type email", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("email"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type opsgenie", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("opsgenie"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type pagerduty", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("pagerduty"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type pushover", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("pushover"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type sensu", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("sensu"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type slack", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("slack"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type telegram", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("telegram"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type threema", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("threema"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type victorops", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("victorops"))))
		})

		It("returns a grafana_metrics_alerting_notifications_sent metric with a type webhook", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingNotificationsSentMetric.WithLabelValues("webhook"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state alerting", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("alerting"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state no_data", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("no_data"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state ok", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("ok"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state paused", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("paused"))))
		})

		It("returns a grafana_metrics_alerting_results metric with a state pending", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(alertingResultsMetric.WithLabelValues("pending"))))
		})

		It("returns a grafana_metrics_api_admin_user_create metric with a state completed", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiAdminUserCreateMetric)))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_create metric with a state completed", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiDashboardSnapshotCreateMetric)))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_external metric with a state completed", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiDashboardSnapshotExternalMetric)))
		})

		It("returns a grafana_metrics_api_dashboard_snapshot_get metric with a state completed", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiDashboardSnapshotGetMetric)))
		})

		It("returns a grafana_metrics_api_login_post metric with a state completed", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiLoginPostMetric)))
		})

		It("returns a grafana_metrics_api_org_create metric with a state completed", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiOrgCreateMetric)))
		})

		It("returns a grafana_metrics_api_responses metric with a status_code 200", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiResponsesMetric.WithLabelValues("unknown"))))
		})

		It("returns a grafana_metrics_api_responses metric with a status_code 404", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiResponsesMetric.WithLabelValues("404"))))
		})

		It("returns a grafana_metrics_api_responses metric with a status_code 500", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiResponsesMetric.WithLabelValues("500"))))
		})

		It("returns a grafana_metrics_api_responses metric with a status_code unknown", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiResponsesMetric.WithLabelValues("unknown"))))
		})

		It("returns a grafana_metrics_api_user_signups_completed metric with a state completed", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiUserSignupsCompletedMetric)))
		})

		It("returns a grafana_metrics_api_user_signups_invite metric with a state inivte", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiUserSignupsInviteMetric)))
		})

		It("returns a grafana_metrics_api_user_signups_started metric with a state started", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiUserSignupsStartedMetric)))
		})

		It("returns a grafana_metrics_aws_cloudwatch_get_metric_statistics metric with a state started", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(awsCloudwatchGetMetricStatisticsMetric)))
		})

		It("returns a grafana_metrics_aws_cloudwatch_list_metrics metric with a state started", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(awsCloudwatchListMetricsMetric)))
		})

		It("returns a grafana_metrics_instance_start metric with a state started", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(instanceStartMetric)))
		})

		It("returns a grafana_metrics_imodels_dashboard_insert metric with a state started", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(modelsDashboardInsertMetric)))
		})

		It("returns a grafana_metrics_page_responses metric with a status_code 200", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(pageResponsesMetric.WithLabelValues("200"))))
		})

		It("returns a grafana_metrics_page_responses metric with a status_code 404", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(pageResponsesMetric.WithLabelValues("404"))))
		})

		It("returns a grafana_metrics_page_responses metric with a status_code 500", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(pageResponsesMetric.WithLabelValues("500"))))
		})

		It("returns a grafana_metrics_page_responses metric with a status_code unknown", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(pageResponsesMetric.WithLabelValues("unknown"))))
		})

		It("returns a grafana_metrics_proxy_responses metric with a status_code 200", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(proxyResponsesMetric.WithLabelValues("200"))))
		})

		It("returns a grafana_metrics_proxy_responses metric with a status_code 404", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(proxyResponsesMetric.WithLabelValues("404"))))
		})

		It("returns a grafana_metrics_proxy_responses metric with a status_code 500", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(proxyResponsesMetric.WithLabelValues("500"))))
		})

		It("returns a grafana_metrics_proxy_responses metric with a status_code unknown", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(proxyResponsesMetric.WithLabelValues("unknown"))))
		})

		It("returns a grafana_admin_dashboards metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(dashboardsMetric)))
		})

		It("returns a grafana_admin_orgs metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(orgsMetric)))
		})

		It("returns a grafana_admin_playlists metric", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(playlistsMetric)))
		})

		It("returns a grafana_admin_users metric", func() {
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
				grafanaClient.GetMetricsReturns(metricsResponse, errors.New("error"))

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
