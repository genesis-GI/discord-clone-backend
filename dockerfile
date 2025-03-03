FROM golang:latest

WORKDIR /app

COPY . . 

RUN go build -o main .

EXPOSE 8050

CMD [ "./main"]