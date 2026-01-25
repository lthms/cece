# Notifications

CeCe can send desktop notifications when it needs your attention during command
modes. This helps when you're multitasking—CeCe works autonomously, and you get
a ping when input is needed.

## How It Works

When CeCe encounters an interaction pattern (clarification, approval, or
blocker) with a policy of `ask`, it calls the `signal_interaction` MCP tool
before displaying the prompt. This triggers a desktop notification.

**Interaction types:**

| Type | Icon | Meaning |
|------|------|---------|
| `clarification` | dialog-question | CeCe needs more information |
| `approval` | dialog-question | CeCe needs your sign-off |
| `blocker` | dialog-warning | CeCe cannot proceed as planned |

## Requirements

The notification system uses `notify-send`, which works with any
freedesktop-compliant notification daemon.

Make sure `notify-send` is installed (usually part of `libnotify`):

```bash
# Arch Linux
pacman -S libnotify

# Debian/Ubuntu
apt install libnotify-bin
```

## MCP Server

The notification feature is provided by an MCP server built into the `cece`
binary. When you run `cece`, it automatically configures Claude Code to use this
server.

**Manual configuration** (if needed):

```bash
claude mcp add cece cece mcp
```

This tells Claude Code to spawn `cece mcp` as an MCP server.

## The `signal_interaction` Tool

The MCP server exposes a single tool:

**Name:** `signal_interaction`

**Parameters:**
- `name` (string) — The configured name from project_setup Identity (falls back to "CeCe")
- `mode` (string) — The current mode indicator (e.g., 🐱, ✨, 🔥)
- `type` (string) — The interaction type: `clarification`, `approval`, or `blocker`
- `message` (string) — The message to display in the notification

**Example call:**

```json
{
  "name": "CeCe",
  "mode": "✨",
  "type": "clarification",
  "message": "Which authentication method should we use?"
}
```

This sends a desktop notification with title "✨ CeCe" and the message as the
body. The icon reflects the interaction type (question or warning).

## Customization

The notification behavior depends on your notification daemon's configuration.
Refer to your notification daemon's documentation for customization options
(positioning, styling, timeouts, etc.).
