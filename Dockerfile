ARG GO_VERSION=1.18.0

# Stage 1: dev environment
FROM golang:${GO_VERSION}-alpine AS dev
RUN go env -w GOPROXY=direct
RUN apk add --no-cache git tzdata
RUN apk --no-cache add ca-certificates && update-ca-certificates
WORKDIR /src
RUN go install github.com/cespare/reflex@latest
COPY . .

# Stage 1: build
FROM dev as build
WORKDIR /src
COPY --from=dev /src .
RUN CGO_ENABLED=0 GOOS=linux go build \ 
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo -o app .

# Stage 1: production app
FROM gcr.io/distroless/static AS prod
WORKDIR /app
COPY --from=build /src .
CMD ["./app"]