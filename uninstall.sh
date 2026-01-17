#!/bin/bash
set -euo pipefail

CLAUDE_DIR="$HOME/.claude"

removed=0

echo "Uninstalling CeCe from $CLAUDE_DIR..."

# Remove CLAUDE.md if it's a symlink
if [[ -L "$CLAUDE_DIR/CLAUDE.md" ]]; then
    rm "$CLAUDE_DIR/CLAUDE.md"
    echo "  Removed CLAUDE.md"
    ((removed++))
fi

# Remove all cece-* files in subdirectories
for dir in rules commands agents; do
    if [[ -d "$CLAUDE_DIR/$dir" ]]; then
        for file in "$CLAUDE_DIR/$dir"/cece-*; do
            if [[ -e "$file" || -L "$file" ]]; then
                rm "$file"
                echo "  Removed $dir/$(basename "$file")"
                ((removed++))
            fi
        done
    fi
done

echo ""
if [[ $removed -eq 0 ]]; then
    echo "No CeCe files found."
else
    echo "Done. Removed $removed files."
fi
