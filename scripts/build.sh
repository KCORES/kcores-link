#!/bin/sh

cd src/
env GO111MODULE=on go build -ldflags "-H=windowsgui"

cp ./kcores-link.exe ../release/