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

		alertingActiveAlertsValue               = 1
		alertingNotificationsSentLineCount      = 2
		alertingNotificationsSentEmailCount     = 3
		alertingNotificationsSentOpsgenieCount  = 4
		alertingNotificationsSentPagerdutyCount = 5
		alertingNotificationsSentPushoverCount  = 6
		alertingNotificationsSentSensuCount     = 7
		alertingNotificationsSentSlackCount     = 8
		alertingNotificationsSentTelegramCount  = 9
		alertingNotificationsSentThreemaCount   = 10
		alertingNotificationsSentVictoropsCount = 11
		alertingNotificationsSentWebhookCount   = 12
		alertingResultStateAlertingCount        = 13
		alertingResultStateNoDataCount          = 14
		alertingResultStateOkCount              = 15
		alertingResultStatePausedCount          = 16
		alertingResultStatePendingCount         = 17
		apiAdminUserCreateCount                 = 18
		apiDashboardSnapshotCreateCount         = 19
		apiDashboardSnapshotExternalCount       = 20
		apiDashboardSnapshotGetCount            = 21
		apiLoginOauthCount                      = 22
		apiLoginPostCount                       = 23
		apiOrgCreateCount                       = 24
		apiRespStatusCode200Count               = 25
		apiRespStatusCode404Count               = 26
		apiRespStatusCode500Count               = 27
		apiRespStatusCodeUnknownCount           = 28
		apiUserSignupCompletedCount             = 29
		apiUserSignupInviteCount                = 30
		apiUserSignupStartedCount               = 31
		awsCloudwatchGetMetricStatisticsCount   = 32
		awsCloudwatchListMetricsCount           = 33
		instanceStartCount                      = 34
		modelsDashboardInsertCount              = 35
		pageRespStatusCode200Count              = 36
		pageRespStatusCode404Count              = 37
		pageRespStatusCode500Count              = 38
		pageRespStatusCodeUnknownCount          = 39
		proxyRespStatusCode200Count             = 40
		proxyRespStatusCode404Count             = 41
		proxyRespStatusCode500Count             = 42
		proxyRespStatusCodeUnknownCount         = 43
		statsTotalsStatDashboardsValue          = 44
		statsTotalsStatOrgsValue                = 45
		statsTotalsStatPlaylistsValue           = 46
		statsTotalsStatUsersValue               = 47

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

		alertingNotificationsSentMetric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "alerting_notifications_sent",
				Help:      "Total number of alert notifications sent.",
			},
			[]string{"type"},
		)
		alertingNotificationsSentMetric.WithLabelValues("line").Add(float64(alertingNotificationsSentLineCount))
		alertingNotificationsSentMetric.WithLabelValues("email").Add(float64(alertingNotificationsSentEmailCount))
		alertingNotificationsSentMetric.WithLabelValues("opsgenie").Add(float64(alertingNotificationsSentOpsgenieCount))
		alertingNotificationsSentMetric.WithLabelValues("pagerduty").Add(float64(alertingNotificationsSentPagerdutyCount))
		alertingNotificationsSentMetric.WithLabelValues("pushover").Add(float64(alertingNotificationsSentPushoverCount))
		alertingNotificationsSentMetric.WithLabelValues("sensu").Add(float64(alertingNotificationsSentSensuCount))
		alertingNotificationsSentMetric.WithLabelValues("slack").Add(float64(alertingNotificationsSentSlackCount))
		alertingNotificationsSentMetric.WithLabelValues("telegram").Add(float64(alertingNotificationsSentTelegramCount))
		alertingNotificationsSentMetric.WithLabelValues("threema").Add(float64(alertingNotificationsSentThreemaCount))
		alertingNotificationsSentMetric.WithLabelValues("victorops").Add(float64(alertingNotificationsSentVictoropsCount))
		alertingNotificationsSentMetric.WithLabelValues("webhook").Add(float64(alertingNotificationsSentWebhookCount))

		alertingResultsMetric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "alerting_results",
				Help:      "Total number of alerting results.",
			},
			[]string{"state"},
		)
		alertingResultsMetric.WithLabelValues("alerting").Add(float64(alertingResultStateAlertingCount))
		alertingResultsMetric.WithLabelValues("no_data").Add(float64(alertingResultStateNoDataCount))
		alertingResultsMetric.WithLabelValues("ok").Add(float64(alertingResultStateOkCount))
		alertingResultsMetric.WithLabelValues("paused").Add(float64(alertingResultStatePausedCount))
		alertingResultsMetric.WithLabelValues("pending").Add(float64(alertingResultStatePendingCount))

		apiResponsesMetric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_responses",
				Help:      "Total number of API responses.",
			},
			[]string{"status_code"},
		)
		apiResponsesMetric.WithLabelValues("200").Add(float64(apiRespStatusCode200Count))
		apiResponsesMetric.WithLabelValues("404").Add(float64(apiRespStatusCode404Count))
		apiResponsesMetric.WithLabelValues("500").Add(float64(apiRespStatusCode500Count))
		apiResponsesMetric.WithLabelValues("unknown").Add(float64(apiRespStatusCodeUnknownCount))

		apiUserSignupsMetric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "api_user_signups",
				Help:      "Total number of API user signups.",
			},
			[]string{"state"},
		)
		apiUserSignupsMetric.WithLabelValues("completed").Add(float64(apiUserSignupCompletedCount))
		apiUserSignupsMetric.WithLabelValues("invite").Add(float64(apiUserSignupInviteCount))
		apiUserSignupsMetric.WithLabelValues("started").Add(float64(apiUserSignupStartedCount))

		pageResponsesMetric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "page_responses",
				Help:      "Total number of Page responses.",
			},
			[]string{"status_code"},
		)
		pageResponsesMetric.WithLabelValues("200").Add(float64(pageRespStatusCode200Count))
		pageResponsesMetric.WithLabelValues("404").Add(float64(pageRespStatusCode404Count))
		pageResponsesMetric.WithLabelValues("500").Add(float64(pageRespStatusCode500Count))
		pageResponsesMetric.WithLabelValues("unknown").Add(float64(pageRespStatusCodeUnknownCount))

		proxyResponsesMetric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "grafana",
				Subsystem: "metrics",
				Name:      "proxy_responses",
				Help:      "Total number of Proxy responses.",
			},
			[]string{"status_code"},
		)
		proxyResponsesMetric.WithLabelValues("200").Add(float64(proxyRespStatusCode200Count))
		proxyResponsesMetric.WithLabelValues("404").Add(float64(proxyRespStatusCode404Count))
		proxyResponsesMetric.WithLabelValues("500").Add(float64(proxyRespStatusCode500Count))
		proxyResponsesMetric.WithLabelValues("unknown").Add(float64(proxyRespStatusCodeUnknownCount))

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
			Eventually(descriptions).Should(Receive(Equal(alertingResultsMetric.WithLabelValues("alerting").Desc())))
		})

		It("returns a grafana_metrics_api_responses metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiResponsesMetric.WithLabelValues("200").Desc())))
		})

		It("returns a grafana_metrics_api_user_signups metric description", func() {
			Eventually(descriptions).Should(Receive(Equal(apiUserSignupsMetric.WithLabelValues("completed").Desc())))
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
				APIUserSignupComplete: grafana.Counter{
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

		It("returns a grafana_metrics_api_user_signups metric with a state completed", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiUserSignupsMetric.WithLabelValues("completed"))))
		})

		It("returns a grafana_metrics_api_user_signups metric with a state inivte", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiUserSignupsMetric.WithLabelValues("invite"))))
		})

		It("returns a grafana_metrics_api_user_signups metric with a state started", func() {
			Eventually(metrics).Should(Receive(PrometheusMetric(apiUserSignupsMetric.WithLabelValues("started"))))
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
