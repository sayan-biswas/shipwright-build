FROM registry.access.redhat.com/ubi9/go-toolset@sha256:b23212b6a296e95cbed5dd415d4cd5567db9e8057aa08e291f1fddc087cea944 AS builder

COPY . .

RUN CGO_ENABLED=0 GO111MODULE=on go build -a -mod=vendor -ldflags="-s -w" -o openshift-builds-image-bundler ./cmd/bundle

FROM registry.access.redhat.com/ubi9-minimal@sha256:9ffc5b7c447ba1918778c60e028216c8a98e3593aec0d3eca330817bc2e31e2b

RUN \
  microdnf clean all && \
  rm -rf /var/cache/yum && \
  echo 'nonroot:x:1000:1000:nonroot:/:/sbin/nologin' > /etc/passwd && \
  echo 'nonroot:x:1000:' > /etc/group && \
  mkdir /.docker && \
  chown 1000:1000 /.docker

COPY --from=builder /opt/app-root/src/openshift-builds-image-bundler /ko-app/bundle
COPY LICENSE /licenses/

ENTRYPOINT ["/ko-app/bundle"]

LABEL \
    com.redhat.component="openshift-builds-image-bundler" \
    name="openshift-builds/image-bundler" \
    version="v1.1.0" \
    summary="Red Hat OpenShift Builds Image Bundler" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Image Bundler" \
    io.k8s.description="Red Hat OpenShift Builds Image Bundler" \
    io.k8s.display-name="Red Hat OpenShift Builds Image Bundler" \
    io.openshift.tags="builds,image-bundler"
