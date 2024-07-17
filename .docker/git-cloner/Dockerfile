# --- start build stage #1
FROM registry.access.redhat.com/ubi9/go-toolset:1.20 AS builder

COPY . .

RUN go version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod=vendor -buildvcs=false -ldflags="-s -w" -o /tmp/openshift-builds-git-cloner ./cmd/git

# --- start build stage #2
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

LABEL \
    com.redhat.component="openshift-builds-git-cloner" \
    name="openshift-builds/git-cloner" \
    version=${CI_CONTAINER_VERSION} \
    summary="Red Hat OpenShift Builds Git Cloner" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Git Cloner" \
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

COPY --from=builder /tmp/openshift-builds-git-cloner .

ENTRYPOINT ["./openshift-builds-git-cloner"]

