FROM golang:latest as build

WORKDIR /app

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download

CMD go run github.com/cosmtrek/air@v1.43.0 \
    --build.cmd "make build" --build.bin "/tmp/bin/api" --build.delay "100" \
    --build.exclude_dir "" \
    --build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
    --misc.clean_on_exit "true"

