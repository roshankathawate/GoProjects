FROM golang:latest 

ADD . /go/src 
WORKDIR /go/src/math 
RUN go build
CMD ["/bin/bash"]
