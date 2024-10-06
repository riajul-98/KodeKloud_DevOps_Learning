#!/bin/bash

NUM1=$1
NUM2=$2

while true
do
    echo "1. Add"
    echo "2. Subtract"
    echo "3. Multiply"
    echo "4. Divide"
    echo "5. Quit"
    read -p "Please choose an option: " operation
        
    if [ $operation -eq 1 ]
    do
        answer=$(( $NUM1 + $NUM2 ))
        echo "Answer=$answer"
    elif [ $operation -eq 2 ]
    do
        answer=$(( $NUM1 - $NUM2 ))
        echo "Answer=$answer"
    elif [ $operation -eq 3 ]
    do
        answer=$(( $NUM1 * $NUM2 ))
        echo "Answer=$answer"
    elif [ $operation -eq 4 ]
    do
        answer=$(( $NUM1 / $NUM2 ))
        echo "Answer=$answer"
    elif [ $operation -eq 5 ]
    do
        break
    fi
done