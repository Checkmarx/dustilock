#!/bin/sh -l

export SCAN_PATH="/github/workspace/"
export ARGS=""

if [ -n "$INPUT_RECURSIVE" ]; then
    export ARGS="$ARGS -r"
fi

if [ -n "$INPUT_AUDIT" ]; then
    export ARGS="$ARGS -a"
fi

if [ -n "$INPUT_PATH" ]; then
    export SCAN_PATH="$SCAN_PATH$INPUT_PATH"
fi

/bin/dusti-lock -p $SCAN_PATH$ARGS