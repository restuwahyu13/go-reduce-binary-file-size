FROM golang:1.20.3-alpine
USER ${USER}

COPY ./go.mod ./
COPY ./ ./

ENV GO111MODULE="on" \
  CGO_ENABLED="0" \
  GOPATH="/"

RUN apk update \
  && apk upgrade \
  && apk add upx

RUN go build --ldflags "-r" -o main . \
  && upx --no-progress -9 ./main \
  && upx -t ./main

EXPOSE 3000
ENTRYPOINT ["./main"]