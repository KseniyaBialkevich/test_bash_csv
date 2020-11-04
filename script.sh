#!/bin/bash
for file in `find ./files -type f -name "*.csv"`
do
awk -F ',' '{if (NR!=1) {print $2}}' $file;
done > ./interim.csv
awk '!a[$1]++' ./interim.csv | awk 'NF' > ./result.csv
rm ./interim.csv
