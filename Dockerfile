FROM golang:1.13-alpine

WORKDIR /opt/code/
ADD ./ /opt/code/

RUN apk --update add git gcc musl-dev

RUN go mod download
RUN CGO_ENABLED=0 go build -o bin/blog cmd/blog/blog.go

FROM alpine

WORKDIR /app

EXPOSE 8080

COPY --from=0 /opt/code/bin/blog /app/
COPY --from=0 /opt/code/configs/config.yml /app/configs/config.yml
COPY --from=0 /opt/code/static/ /app/static/

ENTRYPOINT ["./blog"]