# compile sass

FROM ghcr.io/dart-musl/dart:latest AS dart

ARG BUF_VERSION=1.61.0
ARG SASS_VERSION=1.96.0

RUN \
# get buf
curl -Lo /usr/local/bin/buf \
"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(uname -s)-$(uname -m)" && \
chmod +x /usr/local/bin/buf && \
# get dart-sass and compile
git clone https://github.com/sass/dart-sass.git /dart-sass && \
cd /dart-sass && \
git checkout ${SASS_VERSION} && \
dart pub get && \
dart run grinder protobuf && \
dart compile exe bin/sass.dart -o /sass && \
# cleanup so we dont save all this
cd / && \
rm -f /usr/local/bin/buf && \
rm -rf /root/.cache /root/.dart-tool /root/.pub-cache /dart-sass

# compile site

FROM docker.io/golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN \
GOEXPERIMENT=greenteagc \
CGO_ENABLED=0 GOOS=linux \
go build -ldflags="-s -w" -o blahaj.quest

# create image

FROM scratch

WORKDIR /
ENV PATH=/:$PATH

# should only be one file
COPY --from=dart /lib/ld-musl-*.so.1 /lib/
COPY --from=dart /sass /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /app/blahaj.quest /

CMD ["/blahaj.quest"]
