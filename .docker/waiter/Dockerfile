# --- start build stage #1
FROM registry.access.redhat.com/ubi9/go-toolset:1.21.11-7 AS builder

COPY . .

RUN go version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod=vendor -buildvcs=false -ldflags="-s -w" -o /tmp/openshift-builds-waiter ./cmd/waiter

# --- start build stage #2
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

LABEL \
    com.redhat.component="openshift-builds-waiter" \
    name="openshift-builds/waiter" \
    version=${CI_CONTAINER_VERSION} \
    summary="Red Hat OpenShift Builds Waiter" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Waiter" \
    io.k8s.display-name="Red Hat OpenShift Builds Waiter" \
    io.openshift.tags="builds,waiter"

RUN \
  microdnf --assumeyes --nodocs install tar && \
  microdnf clean all && \
  rm -rf /var/cache/yum && \
  echo 'nonroot:x:1000:1000:nonroot:/:/sbin/nologin' > /etc/passwd && \
  echo 'nonroot:x:1000:' > /etc/group && \
  mkdir /.docker && \
  chown 1000:1000 /.docker

COPY --from=builder /tmp/openshift-builds-waiter .

ENTRYPOINT ["./openshift-builds-waiter"]
