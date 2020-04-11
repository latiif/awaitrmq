FROM golang:alpine AS build-env
COPY . /awaitrmq
WORKDIR /awaitrmq
RUN apk add --update make
RUN apk add --update git
RUN make clean linux
FROM scratch
COPY --from=build-env /awaitrmq/bin/awaitrmq-linux-amd64 /go/bin/awaitrmq
ENTRYPOINT ["/go/bin/awaitrmq"]
