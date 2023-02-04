# syntax = docker/dockerfile:experimental
# Build Container
FROM golang:1.19 as builder

ENV GO111MODULE on
ENV GOPRIVATE=github.com/latonaio
WORKDIR /go/src/github.com/latonaio

COPY . .
RUN go mod download
RUN go build -o convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier ./

# Runtime Container
FROM alpine:3.14
RUN apk add --no-cache libc6-compat
ENV SERVICE=convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier \
    APP_DIR="${AION_HOME}/${POSITION}/${SERVICE}"

WORKDIR ${AION_HOME}

COPY --from=builder /go/src/github.com/latonaio/convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier .

CMD ["./convert-to-dpfm-supply-chain-relationship-from-sap-bp-supplier"]
