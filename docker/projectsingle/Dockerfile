FROM golang
VOLUME /usr/src/myapp
ENV GOPROXY=https://goproxy.cn,direct
ENV LANG=C.UTF-8
WORKDIR /usr/src/myapp
COPY .. .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app ./...
CMD ["app"]