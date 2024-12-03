FROM golang:1.23@sha256:2b01164887cecfa796f6bff6787c8c597d0e0e09e0694428fdde4a343303eb60

WORKDIR /go/src/dyplomat
COPY . /go/src/dyplomat

RUN go install -v ./...
CMD ["dyplomat"]
