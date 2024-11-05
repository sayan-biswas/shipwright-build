FROM registry.access.redhat.com/ubi9/go-toolset@sha256:97e30a01caeece72ee967013e7c7af777ea4ee93840681ddcfe38a87eb4c084a AS builder

COPY . .

RUN CGO_ENABLED=0 GO111MODULE=on go build -a -mod=vendor -ldflags="-s -w" -o openshift-builds-webhook ./cmd/shipwright-build-webhook

FROM registry.access.redhat.com/ubi9-minimal@sha256:c0e70387664f30cd9cf2795b547e4a9a51002c44a4a86aa9335ab030134bf392

COPY --from=builder /opt/app-root/src/openshift-builds-webhook .
COPY LICENSE /licenses/

ENTRYPOINT ["./openshift-builds-webhook"]

LABEL \
    com.redhat.component="openshift-builds-webhook" \
    name="openshift-builds/webhook" \
    version="v1.1.0" \
    summary="Red Hat OpenShift Builds webhook" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds webhook" \
    io.k8s.description="Red Hat OpenShift Builds Webhook" \
    io.k8s.display-name="Red Hat OpenShift Builds Webhook" \
    io.openshift.tags="builds,conversion-webhook"
