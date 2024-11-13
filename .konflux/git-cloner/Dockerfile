FROM registry.access.redhat.com/ubi9/go-toolset@sha256:b23212b6a296e95cbed5dd415d4cd5567db9e8057aa08e291f1fddc087cea944 AS builder

COPY . .

RUN CGO_ENABLED=0 GO111MODULE=on go build -a -mod=vendor -ldflags="-s -w" -o openshift-builds-git-cloner ./cmd/git

FROM registry.access.redhat.com/ubi9-minimal@sha256:9ffc5b7c447ba1918778c60e028216c8a98e3593aec0d3eca330817bc2e31e2b

RUN \
  microdnf --assumeyes --nodocs install git git-lfs && \
  microdnf clean all && \
  rm -rf /var/cache/yum && \
  echo 'nonroot:x:1000:1000:nonroot:/:/sbin/nologin' > /etc/passwd && \
  echo 'nonroot:x:1000:' > /etc/group && \
  mkdir /.docker && chown 1000:1000 /.docker && \
  mkdir /.ssh && chown 1000:1000 /.ssh

COPY --from=builder /opt/app-root/src/openshift-builds-git-cloner /ko-app/git
COPY LICENSE /licenses/

ENTRYPOINT ["/ko-app/git"]

LABEL \
    com.redhat.component="openshift-builds-git-cloner" \
    name="openshift-builds/git-cloner" \
    version="v1.1.0" \
    summary="Red Hat OpenShift Builds Git Cloner" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Git Cloner" \
    io.k8s.description="Red Hat OpenShift Builds Git Cloner" \
    io.k8s.display-name="Red Hat OpenShift Builds Git Cloner" \
    io.openshift.tags="builds,git-cloner"
