FROM golang:1.18.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build -o /rest-api-product-mangagement

EXPOSE 8080

CMD [ "/rest-api-product-mangagement" ]
