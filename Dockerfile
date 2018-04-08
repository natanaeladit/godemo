FROM golang:latest
ADD . /go/src/github.com/natanaeladit/godemo
RUN go get -d -v ./...
RUN go install github.com/natanaeladit/godemo
ENTRYPOINT /go/bin/godemo
EXPOSE 8080