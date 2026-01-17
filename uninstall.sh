#!/bin/bash
set -euo pipefail

CLAUDE_DIR="$HOME/.claude"
REPO_DIR="$(cd "$(dirname "$0")" && pwd)"

# Validate we're in the repo root
if [[ ! -f "$REPO_DIR/CLAUDE.md" ]]; then
    echo "Error: CLAUDE.md not found. Run this script from the cece repo."
    exit 1
fi

# Find and remove symlinks pointing to this repo
removed=0

remove_if_symlink_to_repo() {
    local target="$1"
    local relative="$2"

    if [[ -L "$target" ]]; then
        local link_target
        link_target="$(readlink "$target")"
        if [[ "$link_target" == "$REPO_DIR/$relative" ]]; then
            rm "$target"
            echo "  Removed $relative"
            ((removed++))
        fi
    fi
}

echo "Uninstalling CeCe from $CLAUDE_DIR..."

# Check CLAUDE.md
remove_if_symlink_to_repo "$CLAUDE_DIR/CLAUDE.md" "CLAUDE.md"

# Check files in subdirectories
for dir in rules commands agents; do
    if [[ -d "$CLAUDE_DIR/$dir" ]]; then
        for file in "$CLAUDE_DIR/$dir"/*; do
            if [[ -L "$file" ]]; then
                basename="$(basename "$file")"
                remove_if_symlink_to_repo "$file" "$dir/$basename"
            fi
        done
    fi
done

echo ""
if [[ $removed -eq 0 ]]; then
    echo "No symlinks found pointing to this repo."
else
    echo "Done. Removed $removed symlinks."
fi
