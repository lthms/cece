#!/bin/bash
set -euo pipefail

# =============================================================================
# DEPRECATED: This script is deprecated.
#
# CeCe is now a Claude Code plugin. To uninstall:
#
#   1. Remove the plugin from your Claude Code settings
#   2. Delete .claude/rules/cece.md and .claude/cece.local.md from your projects
#
# This script is kept for backwards compatibility during the transition.
# It removes files created by the old symlink-based installation.
# =============================================================================

echo "WARNING: uninstall.sh is deprecated. CeCe is now a Claude Code plugin."
echo "This script removes files from the old symlink-based installation."
echo ""

CLAUDE_DIR="$HOME/.claude"

removed=0

echo "Uninstalling CeCe from $CLAUDE_DIR..."

# Remove CLAUDE.md if it's a symlink
if [[ -L "$CLAUDE_DIR/CLAUDE.md" ]]; then
    rm "$CLAUDE_DIR/CLAUDE.md"
    echo "  Removed CLAUDE.md"
    removed=$((removed + 1))
fi

# Remove all cece-*.md files in subdirectories
for dir in rules commands agents; do
    if [[ -d "$CLAUDE_DIR/$dir" ]]; then
        for file in "$CLAUDE_DIR/$dir"/cece-*.md; do
            if [[ -e "$file" || -L "$file" ]]; then
                rm "$file"
                echo "  Removed $dir/$(basename "$file")"
                removed=$((removed + 1))
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
