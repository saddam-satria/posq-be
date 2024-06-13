FROM golang:1-alpine3.19 as builder

WORKDIR /app/posq-be 

COPY go.mod go.sum ./

RUN go mod download

COPY commons ./commons 

COPY domains ./domains 

COPY middlewares ./middlewares 

COPY models ./models  

COPY repositories ./repositories  

COPY services ./services  

COPY utils ./utils  

COPY main.go migrate.go seed.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o dist/main .

FROM alpine:latest

WORKDIR /app/posq-be 

COPY --from=builder /app/posq-be/dist ./dist

CMD ["./dist/main"]