#!/bin/sh

go get ./... && go build && export $(cat .env | xargs) && MODE=cmd ./adventuretime