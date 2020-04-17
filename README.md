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
awaitrmq [flags]
```

### Options

```
  -l, --dnslookup      When true, awaits a succesful dnslookup before proceeding. (default true)
  -h, --help           help for awaitrmq
  -i, --interval int   Interval between attempts to check in milliseconds (default 2000)
  -t, --timeout int    Timeout to stop waiting in milliseconds. Pass 0 to timeout in ~ 290 years.
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)