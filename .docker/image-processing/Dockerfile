# --- start build stage #1
FROM registry.access.redhat.com/ubi9/go-toolset:1.21@sha256:5c948cdfd0132e982426bc9d3a81eeae66871080ef274abdde1a4a8303509188 AS builder

COPY . .

RUN go version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod=vendor -buildvcs=false -ldflags="-s -w" -o /tmp/openshift-builds-image-processing ./cmd/image-processing

# --- start build stage #2
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4@sha256:a7d837b00520a32502ada85ae339e33510cdfdbc8d2ddf460cc838e12ec5fa5a

LABEL \
    com.redhat.component="openshift-builds-image-processing" \
    name="openshift-builds/image-processing" \
    version=${CI_CONTAINER_VERSION} \
    summary="Red Hat OpenShift Builds Image Processing" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Image Processing" \
    io.k8s.display-name="Red Hat OpenShift Builds Image Processing" \
    io.openshift.tags="builds,image-processing"

RUN \
  microdnf clean all && \
  rm -rf /var/cache/yum && \
  echo 'nonroot:x:1000:1000:nonroot:/:/sbin/nologin' > /etc/passwd && \
  echo 'nonroot:x:1000:' > /etc/group && \
  mkdir /.docker && \
  chown 1000:1000 /.docker

COPY --from=builder /tmp/openshift-builds-image-processing .

ENTRYPOINT ["./openshift-builds-image-processing"]
