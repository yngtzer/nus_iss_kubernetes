ARG VERSION=alpine3.16

FROM golang:${VERSION} AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY *.go .

RUN go mod download 

RUN go build -o dov-bear main.go

# Compling language does not need source files after compiled
FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/dov-bear .
COPY public public
COPY templates templates

ENV PORT=8080
EXPOSE ${PORT}

CMD [ "./dov-bear" ]

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD curl -s http://localhost:${PORT}/healthz || exit 1