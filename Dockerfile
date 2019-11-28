# Used in production and staging

# Please read Docker & K8s best practices before updating this Dockerfile
# https://zendesk.atlassian.net/wiki/spaces/PAAS/pages/133398613/How+to+deploy+a+project+on+Kubernetes#HowtodeployaprojectonKubernetes-Best-practices

# FROM golang:1.13 (Please make sure this is up-to-date with the sha below)
FROM golang@sha256:c94c082fbfd00975dfef89d439ff9e0059e1816175093f5a2e80541acb8f9352 as builder

WORKDIR /app

COPY . .
COPY REVISION /REVISION

RUN go mod tidy

# FROM gcr.io/docker-images-180022/base/alpine:3.10

# WORKDIR /app

# COPY --from=builder /app/bin bin
# COPY --from=builder /app/scripts scripts
# COPY --from=builder /app/migrations migrations

EXPOSE 8080

# User nobody
USER 65534

ENTRYPOINT ["dumb-init"]
