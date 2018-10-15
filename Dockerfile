FROM golang

RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/pipeline-service
WORKDIR /go/src/pipeline-service

RUN dep ensure
RUN go build

EXPOSE 8000

CMD ./pipeline-service
