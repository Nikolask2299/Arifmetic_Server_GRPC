FROM golang:1.21

WORKDIR /arifm

COPY . .

RUN  go build -o arifm-server services/cmd/server/main.go 

CMD [ "./arifm-server", "--config=./config_file/configuration.json"]