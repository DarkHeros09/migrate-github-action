#!/bin/sh

sh -c '
    migrate -version
    migrate -path $INPUT_PATH -database postgresql://postgres:secret@localhost:6666/cshop?sslmode=disable -verbose up
'
