#!/bin/sh

sh -c '
    migrate -version
    migrate $INPUT_COMMAND
'
