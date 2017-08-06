# Grafana Prometheus Exporter [![Build Status](https://travis-ci.org/frodenas/grafana_exporter.png)](https://travis-ci.org/frodenas/grafana_exporter)

A [Prometheus][prometheus] exporter for [Grafana][grafana] metrics.

## Installation

### Binaries

Download the already existing [binaries][binaries] for your platform:

```bash
$ ./grafana_exporter <flags>
```

### From source

Using the standard `go install` (you must have [Go][golang] already installed in your local machine):

```bash
$ go install github.com/frodenas/grafana_exporter
$ grafana_exporter <flags>
```

### Docker

To run the grafana exporter as a Docker container, run:

```bash
$ docker run -p 9261:9261 frodenas/grafana-exporter <flags>
```

### Cloud Foundry

The exporter can be deployed to an already existing [Cloud Foundry][cloudfoundry] environment:

```bash
$ git clone https://github.com/frodenas/grafana_exporter.git
$ cd grafana_exporter
```

Modify the included [application manifest file][manifest] to include the desired properties. Then you can push the exporter to your Cloud Foundry environment:

```bash
$ cf push
```

### BOSH

This exporter can be deployed using the [Prometheus BOSH Release][prometheus-boshrelease].

## Usage

### Flags

| Flag / Environment Variable | Required | Default | Description |
| --------------------------- | -------- | ------- | ----------- |
| `grafana.uri`<br />`GRAFANA_EXPORTER_GRAFANA_URI` | Yes | | Grafana URI |
| `grafana.username`<br />`GRAFANA_EXPORTER_GRAFANA_USERNAME` | No | | Grafana Username |
| `grafana.password`<br />`GRAFANA_EXPORTER_GRAFANA_PASSWORD` | No | | Grafana Password |
| `grafana.skip-ssl-verify`<br />`GRAFANA_EXPORTER_GRAFANA_SKIP_SSL_VERIFY` | No | `false` | Disable Grafana SSL Verify |
| `web.listen-address`<br />`GRAFANA_EXPORTER_WEB_LISTEN_ADDRESS` | No | `:9261` | Address to listen on for web interface and telemetry |
| `web.telemetry-path`<br />`GRAFANA_EXPORTER_WEB_TELEMETRY_PATH` | No | `/metrics` | Path under which to expose Prometheus metrics |

### Metrics

The exporter returns the following [Admin Stats][admin-stats] metrics:

| Metric | Description | Labels |
| ------ | ----------- | ------ |
| `grafana_admin_stats_alerts` | Number of Grafana Alerts | |
| `grafana_admin_stats_dashboards` | Number of Grafana Dashboards | |
| `grafana_admin_stats_datasources` | Number of Grafana Datasources | |
| `grafana_admin_stats_orgs` | Number of Grafana Orgs | |
| `grafana_admin_stats_playlists` | Number of Grafana Playlists | |
| `grafana_admin_stats_db_snapshots` | Number of Grafana Snapshots | |
| `grafana_admin_stats_starred_db` | Number of Grafana Dashboards Starred | |
| `grafana_admin_stats_db_tags` | Number of Grafana Tags | |
| `grafana_admin_stats_users` | Number of Grafana Admin Stats scrapes | |
| `grafana_admin_stats_scrape_errors_total` | Total number of Grafana Admin Stats scrape errors | |
| `grafana_admin_stats_last_scrape_error` | Whether the last metrics scrape from Grafana Admin Stats resulted in an error (`1` for error, `0` for success) | |
| `grafana_admin_stats_last_scrape_timestamp` | Number of seconds since 1970 since last metrics scrape from Grafana Admin Stats | |
| `grafana_admin_stats_last_scrape_duration_seconds` | Duration of the last metrics scrape from Grafana Admin Stats | |

The exporter returns the following Grafana Metrics:

| Metric | Description | Labels |
| ------ | ----------- | ------ |
| `grafana_metrics_alerting_active_alerts` | Number of active alerts | |
| `grafana_metrics_alerting_notifications_sent` | Number of alert notifications sent | `type` |
| `grafana_metrics_alerting_results` | Number of alerting results | `state` |
| `grafana_metrics_api_admin_user_create` | Number of calls to Admin User Create API | |
| `grafana_metrics_api_dashboard_snapshot_create` | Number of calls to Dashboard Snapshot Create API | |
| `grafana_metrics_api_dashboard_snapshot_external` | Number of calls to Dashboard Snapshot External API | |
| `grafana_metrics_api_dashboard_snapshot_get` | Number of calls to Dashboard Snapshot Get API | |
| `grafana_metrics_api_login_oauth` | Number of calls to Login OAuth API | |
| `grafana_metrics_api_login_post` | Number of calls to Login Post API | |
| `grafana_metrics_api_org_create` | Number of calls to Org Create API | |
| `grafana_metrics_api_responses` | Number of API responses | `code` |
| `grafana_metrics_api_user_signups_completed` | Number of API User Signups completed | |
| `grafana_metrics_api_user_signups_invite` | Number of API User Signups invite | |
| `grafana_metrics_api_user_signups_started` | Number of API User Signups started | |
| `grafana_metrics_aws_cloudwatch_get_metric_statistics` | Number of calls to AWS CloudWatch Get Metric Statistics API | |
| `grafana_metrics_aws_cloudwatch_list_metric` | Number of calls to AWS CloudWatch List Metrics API | |
| `grafana_metrics_instance_start` | Number of Instance Starts | |
| `grafana_metrics_models_dashboard_insert` | Number of Dashboard inserts | |
| `grafana_metrics_page_responses` | Number of Page responses | `code` |
| `grafana_metrics_proxy_responses` | Number of Proxy responses | `code` |
| `grafana_metrics_dashboards` | Number of dashboards | |
| `grafana_metrics_orgs` | Number of orgs | |
| `grafana_metrics_playlists` | Number of playlists | |
| `grafana_metrics_users` | Number of users | |
| `grafana_metrics_scrapes_total` | Total number of Grafana metrics scrapes | |
| `grafana_metrics_scrape_errors_total` | Total number of Grafana metrics scrape errors | |
| `grafana_metrics_last_scrape_error` | Whether the last metrics scrape from Grafana resulted in an error (`1` for error, `0` for success) | |
| `grafana_metrics_last_scrape_timestamp` | Number of seconds since 1970 since last metrics scrape from Grafana | |
| `grafana_metrics_last_scrape_duration_seconds` | Duration of the last metrics scrape from Grafana | |

## Contributing

Refer to the [contributing guidelines][contributing].

## License

Apache License 2.0, see [LICENSE][license].

[admin-stats]: http://docs.grafana.org/http_api/admin/#grafana-stats
[binaries]: https://github.com/frodenas/grafana_exporter/releases
[cloudfoundry]: https://www.cloudfoundry.org/
[contributing]: https://github.com/frodenas/grafana_exporter/blob/master/CONTRIBUTING.md
[golang]: https://golang.org/
[grafana]: https://grafana.com/
[license]: https://github.com/frodenas/grafana_exporter/blob/master/LICENSE
[manifest]: https://github.com/frodenas/grafana_exporter/blob/master/manifest.yml
[prometheus]: https://prometheus.io/
[prometheus-boshrelease]: https://github.com/cloudfoundry-community/prometheus-boshrelease
