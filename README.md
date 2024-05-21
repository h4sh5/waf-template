# WAFfles - a WAF template

A golang HTTP proxy that forwards requests to an origin URL after checking a blocking function.

## usage

`go build`

`./waffles <url of origin>`

example (if origin URL is at `http://localhost:8080`)

`./waffles http://localhost:8080`

