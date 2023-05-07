FROM golang:latest AS build
WORKDIR /app
COPY *.go ./
RUN go build counter.go

FROM debian:bullseye-slim
COPY --from=build /app/counter /app
EXPOSE 8080
ENTRYPOINT ["/app"]
