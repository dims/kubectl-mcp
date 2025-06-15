# kubectl MCP (Model Context Protocol) Plugin

## Overview

The `kubectl-mcp` plugin implements a [Model Context Protocol (MCP)](https://github.com/modelcontextprotocol) server for kubectl. This server enables AI assistants like Amazon Q to interact with kubectl functionality directly through a standardized protocol, allowing for seamless integration of kubectl commands into AI-powered workflows.

## What is MCP?

Model Context Protocol (MCP) is an open protocol that standardizes how applications provide context and tools to Large Language Models (LLMs). It enables AI assistants to:

1. Discover available tools and their capabilities
2. Execute commands and receive structured responses
3. Provide contextual information to enhance AI interactions

## How the kubectl MCP Plugin Works

The kubectl MCP plugin exposes all kubectl commands as tools that can be invoked by AI assistants. The implementation consists of several key components:

### Core Components

1. **`main.go`**: Entry point that starts the MCP server
2. **`pkg/mcp/server.go`**: Creates and configures the MCP server with all kubectl commands
3. **`pkg/mcp/tools.go`**: Handles the registration of kubectl commands as MCP tools

### Implementation Details

- The server uses the `mcp-go` library to implement the Model Context Protocol
- It dynamically registers all kubectl commands (get, apply, delete, etc.) as MCP tools
- Each command's help text, flags, and parameters are exposed through the protocol
- The server runs as a stdio server, communicating through standard input/output

### Command Registration Process

1. The server recursively traverses the kubectl command tree
2. For each command, it extracts:
   - Command description and usage information
   - Available flags and their descriptions
   - Required and optional parameters

## Using the kubectl MCP Plugin with Amazon Q Chat

To use the kubectl MCP plugin with Amazon Q Chat, you need to register it in the Amazon Q configuration file.

Create or edit the file at `$HOME/.aws/amazonq/mcp.json` with the following content:

```json
{
  "mcpServers": {
    "kubectl": {
      "command": "kubectl-mcp",
      "args": []
    }
  }
}
```

This configuration tells Amazon Q Chat to:
1. Register a server named "kubectl"
2. Use the `kubectl-mcp` command to start the MCP server
3. Make all kubectl commands available as tools with the prefix `kubectl___`

## Benefits

- **Seamless Integration**: AI assistants can execute kubectl commands directly
- **Structured Responses**: Commands return structured data that can be parsed by AI models
- **Discoverability**: AI assistants can discover available commands and their parameters
- **Context-Aware**: Provides rich context about Kubernetes resources

## Installation

To install the kubectl MCP plugin:

1. Build the plugin:
   ```bash
   make kubectl-mcp
   ```

2. Move the binary to a location in your PATH:
   ```bash
   sudo mv kubectl-mcp /usr/local/bin/
   ```

3. Verify the installation:
   ```bash
   kubectl plugin list | grep mcp
   ```

## Quick test for MCP server

To quickly test the MCP server, you can run the following command in your terminal:

```bash
echo '{"jsonrpc":"2.0","method":"tools/list","id":1}' | kubectl-mcp | jq
```

## Recommendations

- Monitor background operations for long-running commands. Commands such as resource creation have a 45-second timeout for command responses but the processes continue in background.
- For complex kubectl operations, consider using the plugin in combination with kubectl apply -f for declarative resource management.
