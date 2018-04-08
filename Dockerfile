FROM golang:latest
ADD . /go/src/github.com/natanaeladit/godemo
RUN go get -d -v ./...
# Build the godemo command inside the container.
RUN go install github.com/natanaeladit/godemo
# Run the godemo command when the container starts.
ENTRYPOINT /go/bin/godemo
# http server listens on port 8080.
EXPOSE 8080