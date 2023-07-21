#!/bin/bash

# Check if a directory exists with the same name as the provided string
check_directory() {
  local dir_name="$1"
  if [ -d "$dir_name" ]; then
    echo "Directory '$dir_name' exists."
    exit 0
  else
    echo "Directory '$dir_name' does not exist."
    exit 1
  fi
}

# Validate the input argument
if [ $# -ne 1 ]; then
  echo "Usage: $0 <directory_name>"
  exit 1
fi

directory_name="$1"

# Call the function to check if the directory exists
check_directory "$directory_name"
