FROM ubuntu:jammy

ENV GOPATH /go

RUN apt update && \
    apt install -y golang && \
    apt-get purge -y --auto-remove

COPY go.mod .
COPY go.sum .

RUN GO11MODULE=on go mod download

COPY . .

RUN go build -tags=release -buildvcs=false .

FROM ubuntu:jammy

RUN apt update && \
    apt install -y curl && \
    apt-get purge -y --auto-remove

ENTRYPOINT [ "/fly-exporter" ]
