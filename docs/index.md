# Canary

[![Go Report Card][2]][3]
[![Docker Repository on Quay][4]][5]

A canary http server to run as a docker container. It can be run inside a
kubernetes cluster. More status tests and routes may be added as the
kubernetes cluster project develops.

- route `/status` returns HTTP 200 and a JSON body `{"status", "OK"}`
- remaining routes return HTTP 200 and `OK`

## Usage

<!-- markdownlint-disable MD029 MD032 -->
1. `docker pull quay.io/philoserf/canary`
2. create a container or a pod, etc.
3. `curl -si http://address-to-created-item/status`
<!-- markdownlint-enable -->

---

Copyright 2018 by Mark Ayers. License: [Apache 2.0][1]

[1]: LICENSE.md "Apache 2.0 software license"
[2]: https://goreportcard.com/badge/github.com/philoserf/canary "Go Report Card"
[3]: https://goreportcard.com/report/github.com/philoserf/canary "Go Report Card"
[4]: https://quay.io/repository/philoserf/canary/status "Docker imaage status"
[5]: https://quay.io/repository/philoserf/canary "Docker image on Quay"
