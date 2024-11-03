FROM registry.access.redhat.com/ubi9/go-toolset@sha256:97e30a01caeece72ee967013e7c7af777ea4ee93840681ddcfe38a87eb4c084a AS builder

COPY . .

RUN CGO_ENABLED=0 GO111MODULE=on go build -a -mod=vendor -ldflags="-s -w" -o openshift-builds-image-bundler ./cmd/bundle

FROM registry.access.redhat.com/ubi9-minimal@sha256:c0e70387664f30cd9cf2795b547e4a9a51002c44a4a86aa9335ab030134bf392

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
