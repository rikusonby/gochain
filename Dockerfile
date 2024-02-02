FROM alpine:latest

ENV TZ=Asia/Shanghai
RUN apk update && apk add tzdata go && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime

WORKDIR /www

RUN go install github.com/cosmtrek/air@latest

COPY .air.toml /www/.air.toml

COPY . /www

RUN go mod tidy

EXPOSE 2313

CMD ["/root/go/bin/air"]
