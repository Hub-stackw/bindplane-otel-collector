# Prometheus Metrics with Google Cloud

The Prometheus Receiver can be used to send prometheus style metrics scraped from prometheus exporters to Google Cloud Monitoring. See the [documentation](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusexecreceiver) for more details.

## Prerequisites

See the [prerequisites](../prerequisites.md) doc for Google Cloud prerequisites.

## Configuration

An example configuration is located [here](./config.yaml).

1. Copy [config.yaml](./config.yaml) to `/opt/observiq-otel-collector/config.yaml`
2. Restart the collector: `sudo systemctl restart observiq-otel-collector`

## Setup

This example assumes that [node exporter](https://github.com/prometheus/node_exporter) is running on the host system, on port 9100. Generally, all prometheus style exporters are supported, such as the [aerospike exporter](https://github.com/aerospike/aerospike-prometheus-exporter).

If you wish to scrape metrics from external systems, update `config.yaml`'s prometheus scrape config with a list of IP addresses:
```yaml
- targets:
    - '127.0.0.1:9100'
    - '10.128.1.30:9100'
    - '10.128.1.31:9100'
    - '10.128.1.32:9100'
    - '10.128.1.33:9100'
```

## Metrics

Metrics can be found with the `custom.googleapis.com` prefix.

Example MQL query for [node exporter's](https://github.com/prometheus/node_exporter) `node_cpu_seconds_total`:
```
fetch global
| metric 'custom.googleapis.com/node_cpu_seconds_total'
| align rate(1m)
| every 1m
```