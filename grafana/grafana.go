package grafana

type Client interface {
	GetAdminStats() (AdminStats, error)
	GetMetrics() (Metrics, error)
}

type AdminStats struct {
	AlertCount      int `json:"alert_count"`
	DashboardCount  int `json:"dashboard_count"`
	DatasourceCount int `json:"data_source_count"`
	OrgCount        int `json:"org_count"`
	PlaylistCount   int `json:"playlist_count"`
	SnapshotCount   int `json:"db_snapshot_count"`
	StarredCount    int `json:"starred_db_count"`
	TagCount        int `json:"db_tag_count"`
	UserCount       int `json:"user_count"`
}

type Metrics struct {
	AlertingActiveAlerts               Gauge   `json:"alerting.active_alerts"`
	AlertingExecutionTime              Timer   `json:"alerting.execution_time"`
	AlertingNotificationsSentLine      Counter `json:"alerting.notifications_sent.type_LINE"`
	AlertingNotificationsSentEmail     Counter `json:"alerting.notifications_sent.type_email"`
	AlertingNotificationsSentOpsgenie  Counter `json:"alerting.notifications_sent.type_opsgenie"`
	AlertingNotificationsSentPagerduty Counter `json:"alerting.notifications_sent.type_pagerduty"`
	AlertingNotificationsSentPushover  Counter `json:"alerting.notifications_sent.type_pushover"`
	AlertingNotificationsSentSensu     Counter `json:"alerting.notifications_sent.type_sensu"`
	AlertingNotificationsSentSlack     Counter `json:"alerting.notifications_sent.type_slack"`
	AlertingNotificationsSentTelegram  Counter `json:"alerting.notifications_sent.type_telegram"`
	AlertingNotificationsSentThreema   Counter `json:"alerting.notifications_sent.type_threema"`
	AlertingNotificationsSentVictorops Counter `json:"alerting.notifications_sent.type_victorops"`
	AlertingNotificationsSentWebhook   Counter `json:"alerting.notifications_sent.type_webhook"`
	AlertingResultStateAlerting        Counter `json:"alerting.result.state_alerting"`
	AlertingResultStateNoData          Counter `json:"alerting.result.state_no_data"`
	AlertingResultStateOk              Counter `json:"alerting.result.state_ok"`
	AlertingResultStatePaused          Counter `json:"alerting.result.state_paused"`
	AlertingResultStatePending         Counter `json:"alerting.result.state_pending"`
	APIAdminUserCreate                 Counter `json:"api.admin.user_create"`
	APIDashboardGet                    Timer   `json:"api.dashboard.get"`
	APIDashboardSave                   Timer   `json:"api.dashboard.save"`
	APIDashboardSearch                 Timer   `json:"api.dashboard.search"`
	APIDashboardSnapshotCreate         Counter `json:"api.dashboard_snapshot.create"`
	APIDashboardSnapshotExternal       Counter `json:"api.dashboard_snapshot.external"`
	APIDashboardSnapshotGet            Counter `json:"api.dashboard_snapshot.get"`
	APIDataproxyRequestAll             Timer   `json:"api.dataproxy.request.all"`
	APILoginOauth                      Counter `json:"api.login.oauth"`
	APILoginPost                       Counter `json:"api.login.post"`
	APIOrgCreate                       Counter `json:"api.org.create"`
	APIRespStatusCode200               Counter `json:"api.resp_status.code_200"`
	APIRespStatusCode404               Counter `json:"api.resp_status.code_404"`
	APIRespStatusCode500               Counter `json:"api.resp_status.code_500"`
	APIRespStatusCodeUnknown           Counter `json:"api.resp_status.code_unknown"`
	APIUserSignupCompleted             Counter `json:"api.user.signup_completed"`
	APIUserSignupInvite                Counter `json:"api.user.signup_invite"`
	APIUserSignupStarted               Counter `json:"api.user.signup_started"`
	AWSCloudwatchGetMetricStatistics   Counter `json:"aws.cloudwatch.get_metric_statistics"`
	AWSCloudwatchListMetrics           Counter `json:"aws.cloudwatch.list_metrics"`
	InstanceStart                      Counter `json:"instance_start"`
	ModelsDashboardInsert              Counter `json:"models.dashboard.insert"`
	PageRespStatusCode200              Counter `json:"page.resp_status.code_200"`
	PageRespStatusCode404              Counter `json:"page.resp_status.code_404"`
	PageRespStatusCode500              Counter `json:"page.resp_status.code_500"`
	PageRespStatusCodeUnknown          Counter `json:"page.resp_status.code_unknown"`
	ProxyRespStatusCode200             Counter `json:"proxy.resp_status.code_200"`
	ProxyRespStatusCode404             Counter `json:"proxy.resp_status.code_404"`
	ProxyRespStatusCode500             Counter `json:"proxy.resp_status.code_500"`
	ProxyRespStatusCodeUnknown         Counter `json:"proxy.resp_status.code_unknown"`
	StatsTotalsStatDashboards          Gauge   `json:"stat_totals.stat_dashboards"`
	StatsTotalsStatOrgs                Gauge   `json:"stat_totals.stat_orgs"`
	StatsTotalsStatPlaylists           Gauge   `json:"stat_totals.stat_playlists"`
	StatsTotalsStatUsers               Gauge   `json:"stat_totals.stat_users"`
}

type Counter struct {
	Count int64 `json:"count,omitempty"`
}

type Gauge struct {
	Value int64 `json:"value,omitempty"`
}

type Timer struct {
	Count int64   `json:"count,omitempty"`
	Max   int64   `json:"max,omitempty"`
	Mean  float64 `json:"mean,omitempty"`
	Min   int64   `json:"min,omitempty"`
	P25   float64 `json:"p25,omitempty"`
	P75   float64 `json:"p75,omitempty"`
	P90   float64 `json:"p90,omitempty"`
	P99   float64 `json:"p99,omitempty"`
	Std   float64 `json:"std,omitempty"`
}
