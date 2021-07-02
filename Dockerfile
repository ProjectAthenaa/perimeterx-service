#build stage
FROM golang:1.16.0-alpine3.13 AS build-env

ARG GH_TOKEN
RUN mkdir /app
ADD ./src /app
WORKDIR /app
RUN --mount=type=cache,target=/root/.cache/go-build
RUN go mod download
RUN go build -o goapp


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /app/goapp /app/

EXPOSE 3000 3000

ENTRYPOINT ./goapp