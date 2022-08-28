#!/usr/bin/env sh

for var in "$@"
do
    out=$(basename -s .lvl $var)
    ./con $var > "maps/$out.map"
done
