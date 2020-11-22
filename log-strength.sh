#!/bin/bash

IP="192.168.8.1"
SID=""

while true; do
  ST=$(./get-wifi-strength -ip $IP -sid $SID)
  NOW=$(date +"%F %T")

  echo "${NOW}, ${ST}" >>strength.csv
  sleep 300
done
