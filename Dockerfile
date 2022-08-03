FROM golang:1.17 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIT_TERMINAL_PROMPT=1

COPY ./go.mod ./go.sum ${GOPATH}/src/github.com/kvendingoldo/aws-cognito-lambdas/
COPY ./internal ${GOPATH}/src/github.com/kvendingoldo/aws-cognito-lambdas/internal
COPY ./cmd  ${GOPATH}/src/github.com/kvendingoldo/aws-cognito-lambdas/cmd
WORKDIR ${GOPATH}/src/github.com/kvendingoldo/aws-cognito-lambdas
RUN go get ./
RUN go build -ldflags="-s -w" -o lambda .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder go/src/github.com/kvendingoldo/aws-cognito-lambdas/lambda /app/
WORKDIR /app
ENTRYPOINT ["/app/lambda"]
