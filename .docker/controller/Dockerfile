# --- start build stage #1
FROM registry.access.redhat.com/ubi9/go-toolset:1.21@sha256:5c948cdfd0132e982426bc9d3a81eeae66871080ef274abdde1a4a8303509188 AS builder

COPY . .

RUN go version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -mod=vendor -buildvcs=false -ldflags="-s -w" -o /tmp/openshift-builds-controller ./cmd/shipwright-build-controller

# --- start build stage #2
FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4@sha256:a7d837b00520a32502ada85ae339e33510cdfdbc8d2ddf460cc838e12ec5fa5a

# add labels
LABEL \
    com.redhat.component="openshift-builds-controller" \
    name="openshift-builds/controller" \
    version=${CI_CONTAINER_VERSION} \
    summary="Red Hat OpenShift Builds Controller" \
    maintainer="openshift-builds@redhat.com" \
    description="Red Hat OpenShift Builds Controller" \
    io.k8s.display-name="Red Hat OpenShift Builds Controller" \
    io.openshift.tags="builds,controller"

COPY --from=builder /tmp/openshift-builds-controller .

ENTRYPOINT ["./openshift-builds-controller"]
