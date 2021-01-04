FROM golang:1.15.6-alpine3.12 AS build

ADD . /var/www/lear_go_with_tests

WORKDIR /var/www/lear_go_with_tests/
