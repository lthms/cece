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

# Check existing files and categorize them
errors=()
to_remove=()
to_create=()

for file in "${files_to_link[@]}"; do
    target="$CLAUDE_DIR/$file"
    expected="$REPO_DIR/$file"

    if [[ -L "$target" ]]; then
        actual="$(readlink "$target")"
        if [[ "$actual" == "$expected" ]]; then
            # Correct symlink, nothing to do
            :
        elif [[ ! -e "$target" ]]; then
            # Dangling symlink, remove and recreate
            to_remove+=("$file")
            to_create+=("$file")
        else
            # Symlink to wrong location
            errors+=("$target is a symlink to $actual (expected $expected)")
        fi
    elif [[ -e "$target" ]]; then
        # Regular file or directory exists
        errors+=("$target exists and is not a symlink")
    else
        # Doesn't exist, create it
        to_create+=("$file")
    fi
done

# Also check for dangling cece-* symlinks not in repo
for dir in rules commands agents; do
    if [[ -d "$CLAUDE_DIR/$dir" ]]; then
        for file in "$CLAUDE_DIR/$dir"/cece-*; do
            if [[ -L "$file" && ! -e "$file" ]]; then
                relative="$dir/$(basename "$file")"
                # Only remove if not already in our to_remove list
                if [[ ! " ${to_remove[*]:-} " =~ " $relative " ]]; then
                    to_remove+=("$relative")
                fi
            fi
        done
    fi
done

if [[ ${#errors[@]} -gt 0 ]]; then
    echo "Error: The following files have conflicts:"
    for error in "${errors[@]}"; do
        echo "  $error"
    done
    exit 1
fi

# Create directories
mkdir -p "$CLAUDE_DIR"
for dir in rules commands agents; do
    if [[ -d "$REPO_DIR/$dir" ]]; then
        mkdir -p "$CLAUDE_DIR/$dir"
    fi
done

# Remove dangling symlinks
if [[ ${#to_remove[@]} -gt 0 ]]; then
    echo "Removing dangling symlinks..."
    for file in "${to_remove[@]}"; do
        rm "$CLAUDE_DIR/$file"
        echo "  Removed $file"
    done
fi

# Create symlinks
if [[ ${#to_create[@]} -gt 0 ]]; then
    echo "Creating symlinks..."
    for file in "${to_create[@]}"; do
        ln -s "$REPO_DIR/$file" "$CLAUDE_DIR/$file"
        echo "  Linked $file"
    done
fi

echo ""
if [[ ${#to_remove[@]} -eq 0 && ${#to_create[@]} -eq 0 ]]; then
    echo "Already up to date."
else
    echo "Done."
fi
