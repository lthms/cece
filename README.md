# CeCe

A coding assistant configuration for Claude Code.

## Installation

Clone this repository anywhere on your system:

```bash
git clone https://github.com/lthms/cece.git
cd cece
./install.sh
```

The install script creates symlinks in `~/.claude/` for:
- `CLAUDE.md`
- Files in `rules/`, `commands/`, and `agents/`

The script will fail if any target files already exist. Remove conflicting
files manually before running install.

## Uninstallation

```bash
./uninstall.sh
```

This removes only symlinks pointing to this repo. Other files in `~/.claude/`
are left untouched.

## Post-install Setup

After installation, configure your git identity for CeCe:

```bash
git config --global cece.name "CeCe"
git config --global cece.email "your-cece-email@example.com"
git config --global cece.defaultMode "peer"
```

Then run `/setup` in Claude Code to create `.claude/cece.local.md` for each
project you work on.
