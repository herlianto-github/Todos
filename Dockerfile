FROM golang:1.17-alpine
WORKDIR /app
COPY . ./
RUN go mod download
COPY *.go ./
RUN go build -o /todos
CMD ["/todos"]
