FROM docker.io/library/golang:1 as build

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./server.go
USER 65534:65534

FROM gcr.io/distroless/static-debian12:latest
WORKDIR /
COPY --from=build /app/server /usr/local/bin/server
COPY --from=build /app/migrations /migrations

ENTRYPOINT ["/usr/local/bin/server"]
