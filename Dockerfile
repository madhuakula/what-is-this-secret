FROM golang@sha256:244a736db4a1d2611d257e7403c729663ce2eb08d4628868f9d9ef2735496659 as builder
LABEL MAINTAINER="Madhu Akula"

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/src/madhuakula/what-is-this-secret/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/what-is-this-secret .

# STEP 2 build a small image
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/bin/what-is-this-secret /go/bin/what-is-this-secret

# Use an unprivileged user.
USER appuser:appuser

# Run the what-is-this-secret binary.
ENTRYPOINT ["/go/bin/what-is-this-secret"]