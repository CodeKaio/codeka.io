#!/bin/bash

regex='\[\/\/\]: # \(INCLUDECODE (.*)\)'

# Process input line by line
while IFS= read -r line; do
  # Extract the filename using regex
  if [[ "$line" =~ $regex ]]; then
    filename="${BASH_REMATCH[1]}"

    # Find the file
    if [[ $? -eq 0 ]]; then
      # Output the file content wrapped in code blocks
      echo '```'
      cat "$filename"
      echo ''
      echo '```'
    fi
  else
    # Output the line as is
    echo "$line"
  fi
done

exit 0