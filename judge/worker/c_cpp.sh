#!/bin/bash

# default values
TIME_LIMIT=2s

# parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --src)
            shift
            SRC_FILE=$1
            shift
            ;;
        --cases)
            shift
            CASES=($@)
            shift $#
            ;;
        --time)
            shift
            TIME_LIMIT=$1
            shift
            ;;
        *)
            echo "Unknown argument: $1"
            exit 1
            ;;
    esac
done

# get file extension
EXTENSION=${SRC_FILE##*.}

# set compile and run command based on file extension
case $EXTENSION in
    c)
        COMPILE_CMD="gcc -O2 -std=c99 -o program $SRC_FILE"
        RUN_CMD="./program"
        ;;
    cpp)
        COMPILE_CMD="g++ -O2 -std=c++11 -o program $SRC_FILE"
        RUN_CMD="./program"
        ;;
    *)
        echo "Unsupported file extension: $EXTENSION"
        exit 1
        ;;
esac

# compile the source code
if $COMPILE_CMD; then
    echo "compile: success"
else
    echo "compile: failed"
    exit 1
fi

# run the test cases
total_cases=${#CASES[@]}
pass_cases=0

echo "---"

for ((i=0; i<$total_cases; i+=2))
do
    in_file=${CASES[$i]}
    out_file=${CASES[$i+1]}
    if timeout $TIME_LIMIT $RUN_CMD < $in_file | diff - $out_file; then
        echo "case $((i/2+1)): passed"
        ((pass_cases++))
    else
        echo "case $((i/2+1)): failed"
    fi
done

echo "---"
echo "time: $SECONDS ms"
echo "memory: $(echo "$(awk '/VmPeak/{print $2}' /proc/self/status) * 1024" | bc -l) bytes"

# exit with the number of passed test cases as the return code
exit $pass_cases
