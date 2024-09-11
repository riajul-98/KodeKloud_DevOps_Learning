#!/bin/bash

NUM1=$1
NUM2=$2
NUM3=$3

SUM=$(( NUM1 + NUM2 + NUM3 ))

Average=$( echo "$SUM / 3" | bc -l)

echo $Average