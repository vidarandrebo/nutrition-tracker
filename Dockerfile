FROM node:24.11.1-alpine3.23 AS node-build-env
LABEL authors="Vidar André Bø"

WORKDIR /data/
ENV CI="TRUE"

COPY ./client/ /data/

RUN corepack enable

RUN pnpm install --frozen-lockfile

RUN pnpm build

FROM golang:1.27.1-alpine3.24 AS go-build-env

WORKDIR /data/

COPY ./api/ /data/

RUN go build cmd/api/main.go

FROM alpine:3.24.2

WORKDIR /data/

COPY --from=node-build-env /data/dist/ ./static
COPY --from=go-build-env /data/main .

ENTRYPOINT ["/data/main"]