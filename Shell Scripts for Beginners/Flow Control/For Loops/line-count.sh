#!/bin/bash

for file in $(ls)
do
    echo Line count of $file is $(cat $file | wc -l)
done