FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
  go build -trimpath -ldflags="-s -w" -o /out/example-ci-go-backend ./cmd/server

FROM scratch

USER 65532:65532

COPY --from=build /out/example-ci-go-backend /example-ci-go-backend

EXPOSE 8080

ENTRYPOINT ["/example-ci-go-backend"]
