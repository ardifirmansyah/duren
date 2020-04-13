FROM golang:1.12

COPY . $GOPATH/src/github.com/ardifirmansyah/duren
WORKDIR $GOPATH/src/github.com/ardifirmansyah/duren
ADD . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN $GOPATH/bin/dep ensure -v

RUN go build -o duren

COPY ./files/docker/wait-for.sh wait-for.sh
RUN chmod +x wait-for.sh

CMD sh ./wait-for.sh postgres:5432 -- duren