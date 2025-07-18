#!/bin/bash

regex='\[\/\/\]: # \(DIAGRAM (.*)\)'

# Process input line by line
while IFS= read -r line; do
  # Extract the filename using regex
  if [[ "$line" =~ $regex ]]; then
    filename="${BASH_REMATCH[1]}"

    # Find the file
    if [[ $? -eq 0 ]]; then
      drawio -x -f png --width 1024 -t "$filename" >/dev/null 2>&1
      exported_file=$(basename "$filename" .drawio)
      # Output the file content wrapped in code blocks
      echo "![]($exported_file.png)"
    fi
  else
    # Output the line as is
    echo "$line"
  fi
done

exit 0