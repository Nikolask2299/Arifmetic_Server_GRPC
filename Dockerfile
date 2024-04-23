FROM golang:1.22

WORKDIR /arifm

COPY . .

RUN  go build -o arifm-server services/cmd/server/main.go 

CMD [ "./arifm-server", "--config=./config_file/configuration.json"]