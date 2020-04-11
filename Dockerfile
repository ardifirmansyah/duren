FROM golang:alpine

COPY . $GOPATH/src/github.com/ardifirmansyah/duren
WORKDIR $GOPATH/src/github.com/ardifirmansyah/duren
ADD . .

# RUN dep ensure -v && go build -o
RUN go install -v github.com/ardifirmansyah/duren

COPY ./files/docker/wait-for.sh wait-for.sh
RUN chmod +x wait-for.sh

CMD sh ./wait-for.sh postgres:5432 -- duren