# awaitrmq
[![Go Report Card](https://goreportcard.com/badge/github.com/latiif/awaitrmq)](https://goreportcard.com/report/github.com/latiif/awaitrmq) ![Docker Pulls](https://img.shields.io/docker/pulls/latiif/awaitrmq) ![Docker Latest](https://img.shields.io/docker/v/latiif/awaitrmq)

awaitrmq is a smart [Kubernetes init container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers) for RabbitMQ.

## Background
One of the multiple uses of init containers is running them before the main container to ensure that a service is up using `dnslookup`. The problem with this approach in Kubernetes is that the service can be up but the underlying pod(s) might still not be  ready. awaitrmq tries to solve this issue for RabbitMQ.

## Usage

## awaitrmq

Awaits for a RabbitMQ Service

### Synopsis

Smartly awaits for a RabbitMQ Service to actually be running

```
awaitrmq RabbitMQ-Address[:port number] [flags]
```

### Options

RabbitMQ-Address MUST be at position 1.

Port number is optional. (default `5672`).
```
  -h, --help              help for awaitrmq
  -i, --interval string   Interval between attempts to check. (default 2s)
  -t, --timeout string    Timeout to stop waiting in milliseconds. Pass 0 to timeout in ~ 290 years. (default 0)
  -v, --verbose           Sets output to verbose. (default false)
```

### Example
In both examples, we use `awaitrmq` to await a RabbitMQ instance running on `localhost` port `5672`. Interval is set to `1` second and timeout limit is `5` minutes. No credentials are needed.

#### As a binary

```bash
awaitrmq localhost -v -i=1s -t=5m
```

#### As an init container
When defining containers that depend on the message bus, add this snippet under `spec`.
```YAML
      initContainers:
        - name:   messagebus-init-container
          image:  latiif/awaitrmq
          args:   ["rabbitmq-svc","-v","-i=1s","-t=5m"]
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
