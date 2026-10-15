FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopducteo.com ./cmd/server

FROM scratch

COPY ./config/docker.yaml /config/local.yaml

COPY --from=builder /build/crm.shopducteo.com /

ENTRYPOINT [ "/crm.shopducteo.com", "config/local.yaml" ]
