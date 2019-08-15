FROM golang:1.12 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY cmd/scan/go.mod .
COPY cmd/scan/go.sum .

RUN go mod download

COPY cmd/scan .
RUN ls -la
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o main

FROM scratch
COPY --from=builder /app/main /app/

CMD ["/app/main"]