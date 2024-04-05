FROM golang:latest
WORKDIR /src
COPY go.* ./ 
RUN go mod download 
COPY . /src
RUN go build -o /main
EXPOSE 50052
ENTRYPOINT ["/main"]
