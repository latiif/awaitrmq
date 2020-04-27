FROM golang:alpine AS build-env
COPY . /awaitrmq
WORKDIR /awaitrmq
RUN apk add --no-cache make~=4.2
RUN apk add --no-cache git~=2.24
RUN make clean linux
FROM scratch
COPY --from=build-env /awaitrmq/bin/awaitrmq-linux-amd64 /go/bin/awaitrmq
ENTRYPOINT ["/go/bin/awaitrmq"]
