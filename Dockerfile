FROM golang:1.9-alpine3.7
COPY CDmon.go /home/dfalc/go/src/CDmon/CDmon.go
RUN go build /home/dfalc/go/src/CDmon/CDmon.go
CMD ./CDmon

