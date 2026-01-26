package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type signalInteractionArgs struct {
	Name    string `json:"name" jsonschema:"Your configured name from project_setup Identity"`
	Mode    string `json:"mode" jsonschema:"Your current mode indicator (e.g. 🐱, ✨, 🔥, 📋, 🧠, 🔬)"`
	Type    string `json:"type" jsonschema:"The interaction type: clarification, approval, or blocker"`
	Message string `json:"message" jsonschema:"The message to display in the notification"`
}

func runMCPServer() error {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "cece",
		Version: "1.0.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "signal_interaction",
		Description: "Signal an interaction to the user via desktop notification. Call this when you need user input for a clarification, approval, or blocker.",
	}, handleSignalInteraction)

	slog.Debug("starting MCP server")
	return server.Run(context.Background(), &mcp.StdioTransport{})
}

func handleSignalInteraction(ctx context.Context, req *mcp.CallToolRequest, args signalInteractionArgs) (*mcp.CallToolResult, any, error) {
	// Map interaction type to icon
	icon := "dialog-question"
	if args.Type == "blocker" {
		icon = "dialog-warning"
	}

	// Send notification asynchronously to avoid blocking the conversation
	name := args.Name
	if name == "" {
		name = "CeCe"
	}
	title := fmt.Sprintf("%s %s", args.Mode, name)
	cmd := exec.Command("notify-send", "-i", icon, title, args.Message)
	if err := cmd.Start(); err != nil {
		slog.Warn("notify-send failed to start", "error", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "notification sent"},
		},
	}, nil, nil
}

func ensureMCPServer() error {
	// Check if already configured
	cmd := exec.Command("claude", "mcp", "get", "cece")
	if err := cmd.Run(); err == nil {
		slog.Debug("MCP server already configured")
		return nil
	}

	// Add it
	fmt.Printf("Configuring CeCe MCP server...\n")
	addCmd := exec.Command("claude", "mcp", "add", "cece", "cece", "mcp")
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("failed to configure MCP server: %w", err)
	}

	return nil
}
