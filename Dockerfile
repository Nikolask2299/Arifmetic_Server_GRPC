FROM golang:1.21

WORKDIR /arifm

COPY . .

RUN make

CMD [ "./arifm-server", "--config=./config_file/configuration.json"]