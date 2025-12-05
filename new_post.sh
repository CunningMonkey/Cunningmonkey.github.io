#!/bin/bash

# Check if a title was provided
if [ -z "$1" ]; then
  echo "Usage: $0 \"Post Title\""
  echo "Please provide a title for the new post."
  exit 1
fi

TITLE="$1"
DATE=$(date +%Y-%m-%d)
# Convert title to lowercase, replace spaces with hyphens, and remove non-alphanumeric characters for filename
FILENAME=$(echo "$TITLE" | tr '[:upper:]' '[:lower:]' | sed 's/ /-/g' | sed 's/[^a-z0-9-]//g').md
FILEPATH="content/$FILENAME"

# Check if file already exists
if [ -f "$FILEPATH" ]; then
  echo "Error: File $FILEPATH already exists."
  exit 1
fi

# Create the file with frontmatter
cat <<EOF > "$FILEPATH"
---
title: "$TITLE"
date: $DATE
summary: "Summary of your new post."
---

## $TITLE

Write your content here...
EOF

echo "Created new post: $FILEPATH"
