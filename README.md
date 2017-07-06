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
| `grafana_admin_stats_alerts_total` | Total number of Grafana Alerts | |
| `grafana_admin_stats_dashboards_total` | Total number of Grafana Dashboards | |
| `grafana_admin_stats_datasources_total` | Total number of Grafana Datasources | |
| `grafana_admin_stats_orgs_total` | Total number of Grafana Orgs | |
| `grafana_admin_stats_playlists_total` | Total number of Grafana Playlists | |
| `grafana_admin_stats_snapshots_total` | Total number of Grafana Snapshots | |
| `grafana_admin_stats_starred_total` | Total number of Grafana Dashboards Starred | |
| `grafana_admin_stats_tags_total` | Total number of Grafana Tags | |
| `grafana_admin_stats_users_total` | Total number of Grafana Users | |
| `grafana_admin_stats_scrapes_total` | Total number of Grafana Admin Stats scrapes | |
| `grafana_admin_stats_scrape_errors_total` | Total number of Grafana Admin Stats scrape errors | |
| `grafana_admin_stats_last_scrape_error` | Whether the last metrics scrape from Grafana Admin Stats resulted in an error (`1` for error, `0` for success) | |
| `grafana_admin_stats_last_scrape_timestamp` | Number of seconds since 1970 since last metrics scrape from Grafana Admin Stats | |
| `grafana_admin_stats_last_scrape_duration_seconds` | Duration of the last metrics scrape from Grafana Admin Stats | |

The exporter returns the following Grafana Metrics:

| Metric | Description | Labels |
| ------ | ----------- | ------ |
| `grafana_metrics_alerting_active_alerts` | Number of active alerts | |
| `grafana_metrics_alerting_notifications_sent` | Total number of alert notifications sent | `type` |
| `grafana_metrics_alerting_results` | Total number of alerting results | `state` |
| `grafana_metrics_api_responses` | Total number of API responses | `status_code` |
| `grafana_metrics_api_user_signup` | Total number of API user signups | `state` |
| `grafana_metrics_page_responses` | Total number of Page responses | `status_code` |
| `grafana_metrics_proxy_responses` | Total number of Proxy responses | `status_code` |
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
