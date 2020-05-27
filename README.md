# what-is-this-secret

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/madhuakula/what-is-this-secret/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/madhuakula/what-is-this-secret.svg)](https://github.com/madhuakula/what-is-this-secret/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/madhuakula/what-is-this-secret)](https://goreportcard.com/report/github.com/madhuakula/what-is-this-secret)


[https://wits.madhuakula.com](https://wits.madhuakula.com)

This program is a simple secrets checker utility, give any found secret key and it identifies which service that key belongs to in the real-world

## Usage

* Run using CLI

```bash
what-is-this-secret search -k "AKIAIOSFODNN7EXAMPLE"
```

* Run in server mode, then navigate to [http://localhost:8000](http://localhost:8000)

```bash
what-is-this-secret server
```

![](what-is-this-secret.png)

* Using the curl with API in server mode

```bash
curl -H "Content-Type: application/json" -X POST -d '{"key":"AKIAIOSFODNN7EXAMPLE"}' http://localhost:8000
```

## References

TBD
