#!/bin/bash

for image in $(ls /home/bob/images)
do      
        if [[ $image = *.jpeg ]]
                then
                new_name=$(echo $image | sed 's/jpeg/jpg/g')
                mv images/$image images/$new_name
        fi
done
