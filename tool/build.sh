#!/bin/bash

RUN_NAME="go_fixtool"

mkdir output

export GO111MODULE=on

go build -a -o output/${RUN_NAME}