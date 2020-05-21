# what-is-this-secret

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
