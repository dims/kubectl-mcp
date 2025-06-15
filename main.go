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

// Package main implements the kubectl-mcp plugin for Model Context Protocol support
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"k8s.io/kubernetes/cmd/kubectl-mcp/pkg/mcp"
)

func main() {
	// Create the MCP server
	mcpServer, err := mcp.NewKubectlMCPServer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating MCP server: %v\n", err)
		os.Exit(1)
	}

	// Create a stdio server
	stdioServer := server.NewStdioServer(mcpServer)

	// Start the server
	if err := stdioServer.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting MCP server: %v\n", err)
		os.Exit(1)
	}
}
