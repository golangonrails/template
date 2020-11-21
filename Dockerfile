FROM golang:1.15.5-alpine3.12 AS builder
ARG ARG_GOPROXY="https://goproxy.io"
ENV APP_NAME="app" APP_DIR="/app" GOPROXY=$ARG_GOPROXY GO111MODULE=on
WORKDIR ${APP_DIR}
COPY src ${APP_DIR}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

FROM alpine:3.12.1 AS runner
ENV APP_NAME="app" APP_DIR="/app"
COPY --from=builder ${APP_DIR}/${APP_NAME}  ${APP_DIR}/${APP_NAME}
WORKDIR ${APP_DIR}

ENTRYPOINT ["/app/app"]
