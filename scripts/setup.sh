#!/bin/bash

# Number of folders to create
num_folders=100

# Number of files in each folder
num_files=100
DIR_NAME="folders"

# Create a directory to store the folders
mkdir -p "${DIR_NAME}"


# Loop to create folders
for ((i=1; i<=num_folders; i++))
do
    folder_name="${DIR_NAME}/folder$i"
    mkdir -p "$folder_name"

    # Loop to create files in each folder
    for ((j=1; j<=num_files; j++))
    do
        file_name="file$j.txt"
        touch "$folder_name/$file_name"
    done
done