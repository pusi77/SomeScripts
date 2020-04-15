#!/bin/bash

username="CHANGEME"
ip_address="CHANGEME"

usage="Usage: ./phdis.sh interval\ne.g.: To stop Pi-Hole for 10 min -> \"./phdis.sh 10\"\n"

function check_changeme(){
    if [ "$username" == "CHANGEME" ] || [ "$ip_address" == "CHANGEME" ]; then
        echo "Please change username and ip_address values hardcoded in the script."
        exit 1
    fi
}

function check_args() {
    if ! [ "$#" -eq 1 ]; then
        echo "Wrong number of parameters!"
        printf "$usage"
        exit 1
    elif [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
        printf "$usage"
        exit 0
    elif ! [[ $1 =~ ^[0-9]+$ ]]; then 
        echo "The argument must be an integer."
        printf "$usage"
        exit 1
    fi
}

check_changeme
check_args "$@"
echo "Connecting to $username@$ip_address"...
ssh "$username"@"$ip_address" "
    echo Connected
    pihole disable $1m
    echo Connection closed"
