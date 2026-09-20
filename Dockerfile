FROM node:24-alpine AS node-build-env
LABEL authors="Vidar André Bø"

WORKDIR /data/
ENV CI="TRUE"

COPY ./client/ /data/

RUN corepack enable

RUN pnpm install --frozen-lockfile

RUN pnpm build

FROM golang:1.26-alpine AS go-build-env

WORKDIR /data/

COPY ./api/ /data/

RUN go build cmd/api/main.go

FROM alpine:3

WORKDIR /data/

COPY --from=node-build-env /data/dist/ ./static
COPY --from=go-build-env /data/main .

ENTRYPOINT ["/data/main"]