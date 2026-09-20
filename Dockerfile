FROM node:26.8.2-alpine3.23 AS node-build-env
LABEL authors="Vidar André Bø"

WORKDIR /data/
ENV CI="TRUE"

COPY ./client/ /data/

RUN apk add --no-cache pnpm

RUN pnpm install --frozen-lockfile

RUN pnpm build

FROM golang:1.27.1-alpine3.24 AS go-build-env

WORKDIR /data/

COPY ./api/ /data/

RUN go build -o nutrition-tracker cmd/api/main.go

FROM alpine:3.24.2

WORKDIR /data/

COPY --from=node-build-env /data/dist/ ./static
COPY --from=go-build-env /data/nutrition-tracker .
COPY --from=go-build-env /data/appsettings.json .

ENTRYPOINT ["/data/nutrition-tracker"]