# Used in production and staging

# Please read Docker & K8s best practices before updating this Dockerfile
# https://zendesk.atlassian.net/wiki/spaces/PAAS/pages/133398613/How+to+deploy+a+project+on+Kubernetes#HowtodeployaprojectonKubernetes-Best-practices

# FROM golang:1.13 (Please make sure this is up-to-date with the sha below)
FROM golang@sha256:c94c082fbfd00975dfef89d439ff9e0059e1816175093f5a2e80541acb8f9352 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080
