FROM golang:1.20.4

ARG IS_DEV=false

ENV GIN_MODE=release

WORKDIR /app

COPY . .

RUN if [ "$IS_DEV" = "true" ] ; then go install github.com/cosmtrek/air@v1.43.0 ; fi

RUN if [ "$IS_DEV" = "false" ] ; then go build main.go ; fi

EXPOSE $PORT

ENTRYPOINT [ "./main" ]
