#!/usr/bin/env bash

clear

echo -e "\e[34mPreparing...\e[0m"

OUTPUT_FOLDER="./bin/"

# Make sure the output folder exists
if [[ ! -d "$OUTPUT_FOLDER" ]]; then
    mkdir "$OUTPUT_FOLDER"
fi

echo -e "\e[93mBuilding...\e[0m"

go build -o "$OUTPUT_FOLDER/listener" 2> /tmp/rts.error

if [ $? -eq 0 ]; then
    echo -e "\e[1;32mDone\e[0m"
else
    ERROR_MSG=$(< /tmp/rts.error)
    echo
    echo -e "\e[4;31mSorry! Build process failed:\e[0m"
    echo -e "\e[31m$ERROR_MSG\e[0m"
fi