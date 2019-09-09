FROM golang:1.12-alpine

WORKDIR /opt/code/
ADD ./ /opt/code/

RUN apk --update add git

RUN go get
RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 \
    go build  -o bin/blog cmd/blog/blog.go

FROM alpine

WORKDIR /app

EXPOSE 8080

COPY --from=0 /opt/code/bin/blog /app/
COPY --from=0 /opt/code/configs/config.yml /app/configs/config.yml

ENTRYPOINT ["./blog"]