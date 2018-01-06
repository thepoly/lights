FROM golang:latest

ENV GOPATH /go/app
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN git clone https://github.com/thepoly/lights.git $GOPATH
RUN chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/server/
RUN cd $GOPATH/src/server/

RUN go get

EXPOSE 8080

RUN go build -o lights

CMD ["./lights"]
