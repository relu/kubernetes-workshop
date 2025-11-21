FROM docker.io/golang:1.25-alpine AS build

COPY . /build
WORKDIR /build

EXPOSE 3000

RUN CGO_ENABLED=0 go build -o app ./main.go

FROM gcr.io/distroless/static

COPY --from=build /build/app /app

CMD ["/app"]
