# --- start build stage #1
FROM registry.access.redhat.com/ubi9/go-toolset:1.21.11-2.1720624888 AS builder

COPY . .

RUN go version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod=vendor -buildvcs=false -ldflags="-s -w" -o /tmp/openshift-builds-webhook ./cmd/shipwright-build-webhook

# --- start build stage #2
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

LABEL \
    com.redhat.component="openshift-builds-webhook" \
    name="openshift-builds/webhook" \
    version=${CI_CONTAINER_VERSION} \
    summary="Red Hat OpenShift Builds webhook" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds webhook" \
    io.k8s.display-name="Red Hat OpenShift Builds webhook" \
    io.openshift.tags="builds,conversion-webhook"

COPY --from=builder /tmp/openshift-builds-webhook .

ENTRYPOINT ["./openshift-builds-webhook"]
