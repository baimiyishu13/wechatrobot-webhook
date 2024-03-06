#!/usr/bin/bash

function main() {
    /usr/local/bin/wechat-webhook -RobotKey $1 -addr $2 &
    for (( ; ; )); do
       sleep 60
    done
}

main "$1" "$2"
