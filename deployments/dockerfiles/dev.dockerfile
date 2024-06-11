FROM golang:1-alpine3.19 as builder

WORKDIR /app/posq-be 

COPY go.mod go.sum ./

RUN go mod download && \ 
    go get -u github.com/cosmtrek/air && \
    go install github.com/cosmtrek/air

COPY commons ./commons 

COPY domains ./domains 

COPY middlewares ./middlewares 

COPY models ./models  

COPY repositories ./repositories  

COPY services ./services  

COPY utils ./utils  

COPY main.go .air.toml migrate.go seed.go./

CMD [ "air", "-c", ".air.toml" ]