#!/usr/bin/env sh

for var in "$@"
do
    out=$(basename -s .lvl $var)
    ./con $var > "level/$out.map"
done
