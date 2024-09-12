#!/bin/bash

for package in $(cat install-packages.txt)
do
    sudo apt install $package
done