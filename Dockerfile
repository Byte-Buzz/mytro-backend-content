FROM golang:1.26-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/backend-tile ./cmd/backend-tile && \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/migrate ./cmd/migrate

FROM gcr.io/distroless/static-debian13:nonroot
WORKDIR /app

COPY --from=build /out/ /app/
COPY migrations ./migrations

ENV GIN_MODE=release

EXPOSE 8080
ENTRYPOINT ["/app/backend-tile"]
