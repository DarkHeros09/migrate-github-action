#!/bin/sh

sh -c '
    migrate -version
    migrate -path migrations -database postgresql://postgres:secret@localhost:6666/database?sslmode=disable -verbose up
'
