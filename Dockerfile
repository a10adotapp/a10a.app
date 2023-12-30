FROM golang:1.21 AS build

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go build -o /go/bin/app /go/src/app

FROM scratch AS runner
# FROM debian:bookworm AS runner

COPY --from=build /go/bin/app /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/app"]

EXPOSE 3000
