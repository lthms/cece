#!/bin/bash
set -euo pipefail

CLAUDE_DIR="$HOME/.claude"
REPO_DIR="$(cd "$(dirname "$0")" && pwd)"

# Validate we're in the repo root
if [[ ! -f "$REPO_DIR/CLAUDE.md" ]]; then
    echo "Error: CLAUDE.md not found. Run this script from the cece repo."
    exit 1
fi

# Collect all files to symlink
files_to_link=()
files_to_link+=("CLAUDE.md")

for dir in rules commands agents; do
    if [[ -d "$REPO_DIR/$dir" ]]; then
        for file in "$REPO_DIR/$dir"/*; do
            if [[ -f "$file" ]]; then
                files_to_link+=("$dir/$(basename "$file")")
            fi
        done
    fi
done

# Check for existing files that would conflict
conflicts=()
for file in "${files_to_link[@]}"; do
    target="$CLAUDE_DIR/$file"
    if [[ -e "$target" ]]; then
        conflicts+=("$target")
    fi
done

if [[ ${#conflicts[@]} -gt 0 ]]; then
    echo "Error: The following files already exist:"
    for conflict in "${conflicts[@]}"; do
        echo "  $conflict"
    done
    echo ""
    echo "Remove them manually before running install."
    exit 1
fi

# Create directories
mkdir -p "$CLAUDE_DIR"
for dir in rules commands agents; do
    if [[ -d "$REPO_DIR/$dir" ]]; then
        mkdir -p "$CLAUDE_DIR/$dir"
    fi
done

# Create symlinks
echo "Installing CeCe to $CLAUDE_DIR..."
for file in "${files_to_link[@]}"; do
    ln -s "$REPO_DIR/$file" "$CLAUDE_DIR/$file"
    echo "  Linked $file"
done

echo ""
echo "Done. Linked ${#files_to_link[@]} files."
