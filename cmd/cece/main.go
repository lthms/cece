package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

//go:embed system_prompt.md
var systemPrompt string

const (
	pluginID        = "cece@lthms-cece"
	marketplaceName = "lthms-cece"
	marketplaceRepo = "lthms/cece"
)

type Marketplace struct {
	Name string `json:"name"`
}

type Plugin struct {
	ID          string `json:"id"`
	Version     string `json:"version"`
	Enabled     bool   `json:"enabled"`
	InstallPath string `json:"installPath"`
}

func main() {
	setupLogger()

	// Handle mcp subcommand
	if len(os.Args) > 1 && os.Args[1] == "mcp" {
		if err := runMCPServer(); err != nil {
			fmt.Fprintf(os.Stderr, "cece mcp: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "cece: %v\n", err)
		os.Exit(1)
	}
}

func setupLogger() {
	level := slog.LevelInfo
	if os.Getenv("CECE_DEBUG") != "" {
		level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(logger)
}

func run() error {
	// Find claude binary
	claudePath, err := exec.LookPath("claude")
	if err != nil {
		return fmt.Errorf("claude not found in PATH: %w", err)
	}

	// Ensure plugin is installed
	plugin, err := ensurePlugin()
	if err != nil {
		return fmt.Errorf("plugin check failed: %w", err)
	}

	// Ensure MCP server is configured
	if err := ensureMCPServer(); err != nil {
		return fmt.Errorf("MCP server check failed: %w", err)
	}

	// Read project config
	projectConfig, hasConfig, err := readProjectConfig()
	if err != nil {
		return fmt.Errorf("failed to read project config: %w", err)
	}

	// Compose full system prompt (embedded prompt + project config)
	fullPrompt := composeSystemPrompt(systemPrompt, projectConfig)
	slog.Debug("composed system prompt", "content", fullPrompt)

	// Build arguments (plugin.InstallPath no longer needed since plugin is installed)
	_ = plugin // Plugin verification done, but path not needed for args
	args := buildArgs(os.Args[1:], fullPrompt, hasConfig)
	slog.Debug("built args", "hasConfig", hasConfig, "argCount", len(args))

	// Exec into claude
	return syscall.Exec(claudePath, append([]string{"claude"}, args...), os.Environ())
}

func ensureMarketplace() error {
	cmd := exec.Command("claude", "plugin", "marketplace", "list", "--json")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list marketplaces: %w", err)
	}

	var marketplaces []Marketplace
	if err := json.Unmarshal(output, &marketplaces); err != nil {
		return fmt.Errorf("failed to parse marketplace list: %w", err)
	}

	for _, m := range marketplaces {
		if m.Name == marketplaceName {
			return nil
		}
	}

	fmt.Fprintf(os.Stderr, "Adding CeCe marketplace...\n")
	addCmd := exec.Command("claude", "plugin", "marketplace", "add", marketplaceRepo)
	addCmd.Stdout = os.Stdout
	addCmd.Stderr = os.Stderr
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("failed to add marketplace: %w", err)
	}

	return nil
}

func ensurePlugin() (*Plugin, error) {
	// Ensure marketplace is configured first
	if err := ensureMarketplace(); err != nil {
		return nil, err
	}

	// Run claude plugin list --json
	cmd := exec.Command("claude", "plugin", "list", "--json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list plugins: %w", err)
	}

	var plugins []Plugin
	if err := json.Unmarshal(output, &plugins); err != nil {
		return nil, fmt.Errorf("failed to parse plugin list: %w", err)
	}

	// Check if plugin is installed and enabled
	for _, p := range plugins {
		if p.ID == pluginID {
			if !p.Enabled {
				return nil, fmt.Errorf("plugin %s is installed but disabled", pluginID)
			}
			return &p, nil
		}
	}

	// Plugin not installed, install it
	fmt.Fprintf(os.Stderr, "Installing CeCe plugin...\n")
	installCmd := exec.Command("claude", "plugin", "install", pluginID)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to install plugin: %w", err)
	}

	// Re-check plugin list
	cmd = exec.Command("claude", "plugin", "list", "--json")
	output, err = cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list plugins after install: %w", err)
	}

	if err := json.Unmarshal(output, &plugins); err != nil {
		return nil, fmt.Errorf("failed to parse plugin list: %w", err)
	}

	for _, p := range plugins {
		if p.ID == pluginID {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("plugin %s not found after installation", pluginID)
}

func readProjectConfig() (string, bool, error) {
	// Look for .cece/config.md in current directory
	configPath := ".cece/config.md"

	content, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// No config file
			return "", false, nil
		}
		return "", false, fmt.Errorf("failed to read %s: %w", configPath, err)
	}

	return string(content), true, nil
}

func composeSystemPrompt(systemPrompt, projectConfig string) string {
	var sb strings.Builder
	sb.WriteString(systemPrompt)
	sb.WriteString("\n\n<project_setup>\n")
	sb.WriteString(projectConfig)
	sb.WriteString("\n</project_setup>\n")
	return sb.String()
}

const setupPrompt = "Hello! Let's set you up to contribute to this project."

func buildArgs(originalArgs []string, systemPromptContent string, hasConfig bool) []string {
	var args []string
	var userAppendPrompt string
	skipNext := false

	// Parse original args to find --append-system-prompt (we merge it with ours)
	// All other args pass through unchanged, including --plugin-dir
	for i, arg := range originalArgs {
		if skipNext {
			skipNext = false
			continue
		}

		if arg == "--append-system-prompt" && i+1 < len(originalArgs) {
			userAppendPrompt = originalArgs[i+1]
			skipNext = true
			continue
		}

		if strings.HasPrefix(arg, "--append-system-prompt=") {
			userAppendPrompt = strings.TrimPrefix(arg, "--append-system-prompt=")
			continue
		}

		args = append(args, arg)
	}

	// Compose final system prompt (ours + user's)
	finalPrompt := systemPromptContent
	if userAppendPrompt != "" {
		finalPrompt = finalPrompt + "\n\n" + userAppendPrompt
	}
	args = append(args, "--append-system-prompt", finalPrompt)

	// Add setup prompt if no config exists
	if !hasConfig {
		args = append(args, setupPrompt)
	}

	return args
}
