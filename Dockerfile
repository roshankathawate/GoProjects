FROM golang:latest 
#RUN mkdir /app
#RUN mkdir -p /goproject/math
ADD . /go/src 
#ENV GOPATH=/go/src/
WORKDIR /go/src/math 
RUN go install math
#til/numberutil/numberutil.go 
#RUN go build  . 
CMD ["/bin/bash"]
