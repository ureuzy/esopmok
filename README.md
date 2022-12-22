# esopmok

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

esopmok does the opposite of [Kompose](https://kompose.io/). That is, it converts the kubernetes manifest to docker-compose.yml


# 🚀 Quick Start

## Install Krew

https://krew.sigs.k8s.io/docs/user-guide/setup/install/

## Install kubectl-esopmok




## Usage

```text
$ kubectl esopmok --help
Convert kubernetes manifest to docker-compose.yml

Usage:
  kubectl-esopmok [flags]

Examples:
kubectl esopmok deploy [deployment name]
```

examples

```text
$ kubectl esopmok deploy grafana
name: esopmok
services:
    grafana-download-dashboards:
        command:
            - /bin/sh
        deploy:
            replicas: 1
        image: curlimages/curl:7.85.0
        pull_policy: if_not_present
        volumes:
            - type: volume
              source: config
              target: /etc/grafana/download_dashboards.sh
            - type: volume
              source: storage
              target: /var/lib/grafana
            - type: volume
              source: auth-generic-oauth-secret-mount
              target: /etc/secrets/auth_generic_oauth
    grafana-grafana:
        deploy:
            replicas: 1
        environment:
            GF_PATHS_DATA: /var/lib/grafana/
            GF_PATHS_LOGS: /var/log/grafana
            GF_PATHS_PLUGINS: /var/lib/grafana/plugins
            GF_PATHS_PROVISIONING: /etc/grafana/provisioning
            GF_SECURITY_ADMIN_PASSWORD: ""
            GF_SECURITY_ADMIN_USER: ""
        image: grafana/grafana:9.1.5
        ports:
            - target: 3000
              protocol: TCP
        pull_policy: if_not_present
        volumes:
            - type: volume
              source: config
              target: /etc/grafana/grafana.ini
            - type: volume
              source: storage
              target: /var/lib/grafana
            - type: volume
              source: config
              target: /etc/grafana/provisioning/datasources/datasources.yaml
            - type: volume
              source: config
              target: /etc/grafana/provisioning/dashboards/dashboardproviders.yaml
            - type: volume
              source: sc-dashboard-volume
              target: /tmp/dashboards
            - type: volume
              source: sc-dashboard-provider
              target: /etc/grafana/provisioning/dashboards/sc-dashboardproviders.yaml
            - type: volume
              source: auth-generic-oauth-secret-mount
              target: /etc/secrets/auth_generic_oauth
    grafana-grafana-sc-dashboard:
        deploy:
            replicas: 1
        environment:
            FOLDER: /tmp/dashboards
            LABEL: grafana_dashboard
            METHOD: WATCH
            RESOURCE: both
        image: quay.io/kiwigrid/k8s-sidecar:1.19.2
        pull_policy: if_not_present
        volumes:
            - type: volume
              source: sc-dashboard-volume
              target: /tmp/dashboards
volumes:
    auth-generic-oauth-secret-mount:
        name: auth-generic-oauth-secret-mount
    config:
        name: config
    dashboards-default:
        name: dashboards-default
    sc-dashboard-provider:
        name: sc-dashboard-provider
    sc-dashboard-volume:
        name: sc-dashboard-volume
    storage:
        name: storage
```
