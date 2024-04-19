FROM golang:alpine AS build
COPY . .
RUN GO11MODULE=on go mod download
RUN go build -tags=release -buildvcs=false .

FROM golang:alpine
RUN apk add --no-cache curl
COPY --from=build /go/fly-exporter /fly-exporter
ENTRYPOINT [ "/fly-exporter" ]
