package grafana_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/frodenas/grafana_exporter/grafana"
)

var _ = Describe("HTTPClient", func() {
	var (
		server *ghttp.Server
		client Client
		err    error

		username      = "fake-username"
		password      = "fake-password"
		skipSSLVerify = true
	)

	BeforeEach(func() {
		server = ghttp.NewServer()

		client, err = NewHTTPClient(server.URL(), username, password, skipSSLVerify)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GetAdminStats", func() {
		var (
			statusCode         int
			adminStats         AdminStats
			adminStatsResponse AdminStats
		)

		BeforeEach(func() {
			statusCode = http.StatusOK
			adminStatsResponse = AdminStats{
				AlertCount:      1,
				DashboardCount:  2,
				DatasourceCount: 3,
				OrgCount:        4,
				PlaylistCount:   5,
				DBSnapshotCount: 6,
				StarredDBCount:  7,
				DBTagCount:      8,
				UserCount:       9,
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/api/admin/stats"),
					ghttp.VerifyBasicAuth(username, password),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &adminStatsResponse),
				),
			)
		})

		JustBeforeEach(func() {
			adminStats, err = client.GetAdminStats()
		})

		It("returns the admin stats", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(adminStats).To(Equal(adminStatsResponse))
		})

		Context("when it fails to get the admin stats", func() {
			BeforeEach(func() {
				statusCode = http.StatusInternalServerError
			})

			It("returns an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Error getting admin stats, http status code: 500"))
			})
		})
	})

	Describe("GetMetrics", func() {
		var (
			statusCode      int
			metrics         Metrics
			metricsResponse Metrics
		)

		BeforeEach(func() {
			statusCode = http.StatusOK
			metricsResponse = Metrics{
				AlertingActiveAlerts: Gauge{
					Value: int64(1),
				},
				AlertingNotificationsSentLine: Counter{
					Count: int64(2),
				},
				AlertingExecutionTime: Timer{
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
				AlertingNotificationsSentDingDing: Counter{
					Count: int64(3),
				},
				AlertingNotificationsSentEmail: Counter{
					Count: int64(4),
				},
				AlertingNotificationsSentOpsgenie: Counter{
					Count: int64(5),
				},
				AlertingNotificationsSentPagerduty: Counter{
					Count: int64(6),
				},
				AlertingNotificationsSentPushover: Counter{
					Count: int64(7),
				},
				AlertingNotificationsSentSensu: Counter{
					Count: int64(8),
				},
				AlertingNotificationsSentSlack: Counter{
					Count: int64(9),
				},
				AlertingNotificationsSentTelegram: Counter{
					Count: int64(10),
				},
				AlertingNotificationsSentThreema: Counter{
					Count: int64(11),
				},
				AlertingNotificationsSentVictorops: Counter{
					Count: int64(12),
				},
				AlertingNotificationsSentWebhook: Counter{
					Count: int64(13),
				},
				AlertingResultStateAlerting: Counter{
					Count: int64(14),
				},
				AlertingResultStateNoData: Counter{
					Count: int64(15),
				},
				AlertingResultStateOk: Counter{
					Count: int64(16),
				},
				AlertingResultStatePaused: Counter{
					Count: int64(17),
				},
				AlertingResultStatePending: Counter{
					Count: int64(18),
				},
				APIAdminUserCreate: Counter{
					Count: int64(19),
				},
				APIDashboardGet: Timer{
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
				APIDashboardSave: Timer{
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
				APIDashboardSearch: Timer{
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
				APIDashboardSnapshotCreate: Counter{
					Count: int64(20),
				},
				APIDashboardSnapshotExternal: Counter{
					Count: int64(21),
				},
				APIDashboardSnapshotGet: Counter{
					Count: int64(22),
				},
				APIDataproxyRequestAll: Timer{
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
				APILoginOauth: Counter{
					Count: int64(23),
				},
				APILoginPost: Counter{
					Count: int64(24),
				},
				APIOrgCreate: Counter{
					Count: int64(25),
				},
				APIRespStatusCode200: Counter{
					Count: int64(26),
				},
				APIRespStatusCode404: Counter{
					Count: int64(27),
				},
				APIRespStatusCode500: Counter{
					Count: int64(28),
				},
				APIRespStatusCodeUnknown: Counter{
					Count: int64(29),
				},
				APIUserSignupCompleted: Counter{
					Count: int64(30),
				},
				APIUserSignupInvite: Counter{
					Count: int64(31),
				},
				APIUserSignupStarted: Counter{
					Count: int64(32),
				},
				AWSCloudwatchGetMetricStatistics: Counter{
					Count: int64(33),
				},
				AWSCloudwatchListMetrics: Counter{
					Count: int64(34),
				},
				InstanceStart: Counter{
					Count: int64(35),
				},
				ModelsDashboardInsert: Counter{
					Count: int64(36),
				},
				PageRespStatusCode200: Counter{
					Count: int64(37),
				},
				PageRespStatusCode404: Counter{
					Count: int64(38),
				},
				PageRespStatusCode500: Counter{
					Count: int64(39),
				},
				PageRespStatusCodeUnknown: Counter{
					Count: int64(40),
				},
				ProxyRespStatusCode200: Counter{
					Count: int64(41),
				},
				ProxyRespStatusCode404: Counter{
					Count: int64(42),
				},
				ProxyRespStatusCode500: Counter{
					Count: int64(43),
				},
				ProxyRespStatusCodeUnknown: Counter{
					Count: int64(44),
				},
				StatsTotalsStatDashboards: Gauge{
					Value: int64(45),
				},
				StatsTotalsStatOrgs: Gauge{
					Value: int64(46),
				},
				StatsTotalsStatPlaylists: Gauge{
					Value: int64(47),
				},
				StatsTotalsStatUsers: Gauge{
					Value: int64(48),
				},
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/api/metrics"),
					ghttp.VerifyBasicAuth(username, password),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &metricsResponse),
				),
			)
		})

		JustBeforeEach(func() {
			metrics, err = client.GetMetrics()
		})

		It("returns the metrics", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(metrics).To(Equal(metricsResponse))
		})

		Context("when it fails to get the metrics", func() {
			BeforeEach(func() {
				statusCode = http.StatusInternalServerError
			})

			It("returns an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Error getting metrics, http status code: 500"))
			})
		})
	})
})
