#!/bin/bash

for mission in $(cat mission-names.txt)
do
    create-and-launch-rocket $mission
done