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