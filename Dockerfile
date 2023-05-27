FROM golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./myapp ./
RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]