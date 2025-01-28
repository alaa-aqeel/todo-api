FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache go 

ENV GOPATH /go

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go install github.com/cosmtrek/air@v1.49.0

COPY go.mod go.sum ./
RUN go mod download

COPY .env /app
COPY .air.toml /app
COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
