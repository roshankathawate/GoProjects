FROM golang:latest 
ADD . /go/src 
WORKDIR /go/src/math 
RUN go install math
CMD ["/bin/bash"]
