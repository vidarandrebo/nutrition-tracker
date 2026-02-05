FROM node:25-alpine AS node-build-env
LABEL authors="Vidar André Bø"

WORKDIR /data/

COPY ./client/ /data/

RUN npm ci

RUN npm run build

FROM golang:1.25-alpine AS go-build-env

WORKDIR /data/

COPY ./api/ /data/

RUN go build cmd/api/main.go

FROM alpine:3

WORKDIR /data/

COPY --from=node-build-env /data/dist/ ./static
COPY --from=go-build-env /data/main .

ENTRYPOINT ["/data/main"]