FROM registry.access.redhat.com/ubi9/go-toolset@sha256:b23212b6a296e95cbed5dd415d4cd5567db9e8057aa08e291f1fddc087cea944 AS builder

COPY . .

RUN CGO_ENABLED=0 GO111MODULE=on go build -a -mod=vendor -ldflags="-s -w" -o openshift-builds-controller ./cmd/shipwright-build-controller

FROM registry.access.redhat.com/ubi9-minimal@sha256:9ffc5b7c447ba1918778c60e028216c8a98e3593aec0d3eca330817bc2e31e2b

COPY --from=builder /opt/app-root/src/openshift-builds-controller .
COPY LICENSE /licenses/

ENTRYPOINT ["./openshift-builds-controller"]

LABEL \
    com.redhat.component="openshift-builds-controller" \
    name="openshift-builds/controller" \
    version="v1.1.0" \
    summary="Red Hat OpenShift Builds Controller" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Controller" \
    io.k8s.description="Red Hat OpenShift Builds Controller" \
    io.k8s.display-name="Red Hat OpenShift Builds Controller" \
    io.openshift.tags="builds,controller"