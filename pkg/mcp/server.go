/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package mcp implements the Model Context Protocol (MCP) server functionality for kubectl
package mcp

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"k8s.io/kubectl/pkg/cmd"
)

// NewKubectlMCPServer creates and configures an MCP server for kubectl
// It sets up the command structure and registers all kubectl commands as MCP tools
func NewKubectlMCPServer() (*server.MCPServer, error) {
	// Get kubectl version
	version := "v0.0.0" // This will be replaced with actual version in production

	// Create a new MCP server with the specified name and version
	s := server.NewMCPServer("kubectl", version, server.WithInstructions("MCP server for kubectl"))
	
	// Create kubectl command
	kubectlCmd := cmd.NewDefaultKubectlCommand()

	// Register all kubectl commands as MCP tools
	if err := registerTools(s, kubectlCmd); err != nil {
		return nil, fmt.Errorf("failed to register kubectl commands as MCP tools: %w", err)
	}

	return s, nil
}
