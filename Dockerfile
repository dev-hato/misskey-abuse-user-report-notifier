# checkov:skip=CKV_DOCKER_2
FROM golang:1.22.3 AS develop

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
ARG TARGETPLATFORM

WORKDIR /go/app
COPY go.mod go.sum ./
COPY tools.go ./
RUN go install github.com/cosmtrek/air
COPY . .
RUN mapfile -t PLATFORM < <(echo "${TARGETPLATFORM}" | tr '/' ' ') \
    && CGO_ENABLED=0 GOOS=linux GOARCH=${PLATFORM[2]} go build -o ./app \
    && rm -rf /go/pkg/mod/dario.cat/mergo@*/.vscode \
              /go/pkg/mod/github.com/*/*/Dockerfile \
              /go/pkg/mod/github.com/pelletier/go-toml/*/Dockerfile \
              /go/pkg/mod/golang.org/x/sys@*/unix/linux/Dockerfile \
              /usr/local/go/src/crypto/internal/boring/Dockerfile \
              /usr/local/go/src/crypto/internal/nistec/fiat/Dockerfile

RUN useradd -l -m -s /bin/bash -N -u "1000" "nonroot" \
    && chown -R nonroot /go/app/

RUN find / -type f -perm /u+s -ignore_readdir_race -exec chmod u-s {} \; \
    && find / -type f -perm /g+s -ignore_readdir_race -exec chmod g-s {} \;

USER nonroot

CMD ["air", "-c", ".air.toml"]

FROM scratch

COPY --from=develop /etc/group /etc/group
COPY --from=develop /etc/passwd /etc/passwd
COPY --from=develop /etc/shadow /etc/shadow
USER nonroot

WORKDIR /go/app
COPY --from=develop /go/app/app ./
COPY --from=develop /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./app"]
