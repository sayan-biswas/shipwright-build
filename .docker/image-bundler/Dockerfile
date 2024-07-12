# --- start build stage #1
FROM registry.access.redhat.com/ubi9/go-toolset:1.20 AS builder

COPY . .

RUN go version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod=vendor -buildvcs=false -ldflags="-s -w" -o /tmp/openshift-builds-image-bundler ./cmd/bundle

# --- start build stage #2
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest

LABEL \
    com.redhat.component="openshift-builds-image-bundler" \
    name="openshift-builds/image-bundler" \ 
    version=${CI_CONTAINER_VERSION} \
    summary="Red Hat OpenShift Builds Image Bundler" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Image Bundler" \
    io.k8s.display-name="Red Hat OpenShift Image Bundler" \
    io.openshift.tags="builds,image-bundler"

RUN \
  microdnf clean all && \
  rm -rf /var/cache/yum && \
  echo 'nonroot:x:1000:1000:nonroot:/:/sbin/nologin' > /etc/passwd && \
  echo 'nonroot:x:1000:' > /etc/group && \
  mkdir /.docker && \
  chown 1000:1000 /.docker

COPY --from=builder /tmp/openshift-builds-image-bundler .

ENTRYPOINT ["./openshift-builds-image-bundler"]
