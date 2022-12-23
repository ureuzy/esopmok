# esopmok

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

esopmok does the opposite of [Kompose](https://kompose.io/). That is, it converts the kubernetes manifest to docker-compose.yml

![demo](https://raw.github.com/wiki/ureuzy/esopmok/images/demo.gif)

# 🚀 Quick Start

## Install Krew

https://krew.sigs.k8s.io/docs/user-guide/setup/install/

## Install kubectl-esopmok

Binary Download [here](https://github.com/ureuzy/esopmok/releases)

```text
$ mv kubectl-esopmok ~/.krew/bin
```

## Usage

```text
$ kubectl esopmok --help
Convert kubernetes manifest to docker-compose.yml

Usage:
  kubectl-esopmok [flags]

Examples:
kubectl esopmok deploy [deployment name]
```

As an example, convert the Grafana Deployment to docker-compose.yml.

```text
$ kubectl get deploy grafana
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
grafana   1/1     1            1           85d
```

To convert, enter the following command.

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

~~~
```

Additionally, it can be run as follows

```text
$ kubectl esopmok deploy grafana | docker compose -f - up
```
