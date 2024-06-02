FROM golang:1.22.3 as builder

WORKDIR /
COPY go.mod /
COPY *.go ./
RUN go build -o /app

FROM migrate/migrate

COPY --from=builder ./app ./

ENTRYPOINT ["/app"]