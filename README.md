# awaitrmq

awaitrmq is a smart [Kubernetes init container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers) for RabbitMQ.

## Background
One of the multiple uses of init containers is running them before the main container to ensure that a service is up using `dnslookup`. The problem with this approach in Kubernetes is that the service can be up but the underlying pod(s) might still not be  ready. awaitrmq tries to solve this issue for RabbitMQ.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)