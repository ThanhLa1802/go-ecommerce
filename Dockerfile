# ---------- BUILD STAGE ----------
FROM golang:1.25-alpine AS builder

WORKDIR /build

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o crm.shopdev.com ./cmd/server


# ---------- RUNTIME STAGE ----------
FROM alpine:latest

WORKDIR /app

# copy binary
COPY --from=builder /build/crm.shopdev.com .

# copy config
COPY ./configs ./configs

# create log folder
RUN mkdir -p /app/storages/logs

EXPOSE 8082

ENTRYPOINT [ "./crm.shopdev.com", "config/local.yaml" ]