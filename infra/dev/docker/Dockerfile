FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -o /out/api ./cmd

FROM alpine:3.20
COPY --from=build /out/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
