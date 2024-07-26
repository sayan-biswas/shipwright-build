# --- start build stage #1
FROM registry.access.redhat.com/ubi9/go-toolset:1.21@sha256:5c948cdfd0132e982426bc9d3a81eeae66871080ef274abdde1a4a8303509188 AS builder

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# This are vendors dependencies. Copy this layer first to take advantage of cache hits.
COPY vendor/ vendor/

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY version/ version/

RUN go version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod=vendor -ldflags="-s -w" -o openshift-builds-git-cloner ./cmd/git

# --- start build stage #2
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4@sha256:a7d837b00520a32502ada85ae339e33510cdfdbc8d2ddf460cc838e12ec5fa5a

LABEL \
    com.redhat.component="openshift-builds-git-cloner" \
    name="openshift-builds/git-cloner" \
    version=${CI_CONTAINER_VERSION} \
    summary="Red Hat OpenShift Builds Git Cloner" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Git Cloner" \
    io.k8s.description="Red Hat OpenShift Builds Git Cloner" \
    io.k8s.display-name="Red Hat OpenShift Builds Git Cloner" \
    io.openshift.tags="builds,git-cloner"

RUN \
  microdnf --assumeyes --nodocs install git git-lfs && \
  microdnf clean all && \
  rm -rf /var/cache/yum && \
  echo 'nonroot:x:1000:1000:nonroot:/:/sbin/nologin' > /etc/passwd && \
  echo 'nonroot:x:1000:' > /etc/group && \
  mkdir /.docker && chown 1000:1000 /.docker && \
  mkdir /.ssh && chown 1000:1000 /.ssh

COPY --from=builder /opt/app-root/src/openshift-builds-git-cloner .

ENTRYPOINT ["./openshift-builds-git-cloner"]

