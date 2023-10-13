FROM golang:alpine
COPY go.mod ./
COPY go.sum ./
COPY ./ ./ 
RUN go mod download
RUN go build -o goApi .
CMD ["./goApi"]