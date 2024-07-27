FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM builder as build-server

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/server

FROM scratch as server

COPY --from=build-server /app/server /app/server

CMD ["/app/server"]