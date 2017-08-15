# FROM golang:alpine3.6
# RUN mkdir -p /app
# WORKDIR /app
# ADD . /app
# RUN go build ./main.go
# CMD ["./main"]

FROM scratch
ADD main /
CMD ["/main"]