FROM golang
ADD . /go/src/github.com/williamzion/chatter
RUN go install github.com/williamzion/chatter
ENTRYPOINT /go/bin/chatter
EXPOSE 8080