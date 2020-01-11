#!/usr/bin/env bash

clear

function push() {
    printf "\nExecuting git push command\n"
    git push
}

printf "\e[34mPreparing...\e[0m\n"

OUTPUT_FOLDER="./bin/"

# Make sure the output folder exists
if [[ ! -d "$OUTPUT_FOLDER" ]]; then
    mkdir "$OUTPUT_FOLDER"
fi

printf "\e[34mBuilding...\e[0m\n"

go build -o "$OUTPUT_FOLDER/listener" 2> /tmp/rts.error

if [ $? -eq 0 ]; then
    printf "\e[1;32mDone\e[0m\n"

    # If the main go build process exits successfully
    # we go ahead to execute the sub command
    SUB_COMMAND=$1
    if [[ -n "$1" ]]; then
        case $1 in
            push)
            push
            ;;
        esac
    fi
else
    ERROR_MSG=$(< /tmp/rts.error)
    printf "\n\e[1;4;31mFAILED:\e[0m\n"
    printf "\e[31m$ERROR_MSG\e[0m\n"
fi
