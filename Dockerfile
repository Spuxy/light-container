FROM golang:latest
WORKDIR /service
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o container .
CMD [ "sleep", "5000" ]
