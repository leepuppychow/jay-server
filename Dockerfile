FROM golang

ADD . /go/src/github.com/leepuppychow/jay_medtronic

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/lib/pq
RUN go get github.com/raja/argon2pw

RUN go install github.com/leepuppychow/jay_medtronic

ENTRYPOINT /go/bin/jay_medtronic

EXPOSE 3000