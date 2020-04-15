FROM golang:alpine

COPY . $GOPATH/src/github.com/ardifirmansyah/duren
WORKDIR $GOPATH/src/github.com/ardifirmansyah/duren
ADD . .

# RUN go get -u github.com/golang/dep/cmd/dep
# RUN $GOPATH/bin/dep ensure -v

RUN go install -v $GOPATH/src/github.com/ardifirmansyah/duren

COPY ./files/docker/wait-for.sh wait-for.sh
RUN chmod +x wait-for.sh

CMD sh ./wait-for.sh postgres:5432 -- duren