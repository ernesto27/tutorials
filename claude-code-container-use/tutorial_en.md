# Tutorial: Claude Code and Container Use

This tutorial walks you through an example of how to use Claude Code together with container-use to run background tasks inside a container on a separate Git branch.

- Setup Claude Code
- Create a Go API project
- Using Claude Code
- Installing Container Use
- Example: background tasks with container-use

## Setup Claude Code

Claude Code is a CLI tool that lets you interact with the Anthropic Claude models (Sonnet 4, Sonnet 3.7, etc.) directly from your terminal. You don’t need an editor (VS Code, Cursor, Windsurf), although you can integrate Claude Code with them if desired.

### Prerequisites

Make sure you have Node.js installed, then run:

```bash
npm install -g @anthropic-ai/claude-code
```

### Initial Configuration

Run Claude Code inside your project directory. For example:

```bash
mkdir claude-code-example && cd claude-code-example
claude
```

You will be prompted for permissions—press **Yes**.

### Authentication Options

Claude Code offers two options:

1. **Claude Subscription**
2. **Anthropic Account Console**

If you already have a Claude account, choose the first option. Otherwise, select the second option to set up an API key automatically for you.

### Manual API Key Setup (Optional)

If you have your own API key, you can configure it manually:

```bash
cd ~/.claude
```

Create or update `settings.json`:

```json
{
  "apiKeyHelper": "~/.claude/anthropic_key.sh"
}
```

Create the script:

```bash
nano anthropic_key.sh
```

Add your key:

```bash
echo "sk-..."
```

Once configured, you can use Claude Code in your project directory.

![Claude Code Setup](https://raw.githubusercontent.com/ernesto27/tutorials/refs/heads/master/claude-code-container-use/image1.png)

## Create a Go API Project

To test Claude Code with container-use, we’ll build a simple Go API.

[Go Documentation](https://go.dev/)

In your `claude-code-example` folder, add two files.

### go.mod

```go
module goapi

go 1.23
```

### main.go

```go
package main

import (
  "encoding/json"
  "net/http"
)

type Response struct {
  Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  resp := Response{Message: "Hello World"}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(resp)
}

func main() {
  http.HandleFunc("/hello", helloHandler)
  http.ListenAndServe(":8080", nil)
}
```

Run the service:

```bash
go run main.go
```

Test with curl:

```bash
curl http://localhost:8080/hello
```

Initialize Git (required for container-use):

```bash
git init
git add .
git commit -m "Initial commit"
```

## Using Claude Code

Claude Code uses a `CLAUDE.md` file to guide its behavior. Inside your project, start a session:

```bash
claude /init
```

You’ll be asked to authorize various tools.

![](https://raw.githubusercontent.com/ernesto27/tutorials/refs/heads/master/claude-code-container-use/image2.png)

[Generated CLAUDE.md file](https://github.com/ernesto27/tutorials/commit/975140d21d0f116a68dfea28c0e55feae95bdc57)

Now, ask Claude to add an endpoint that returns the API version (defined in `version.json`):

```bash
claude "Create an endpoint that returns the API version as defined in version.json"
```

Claude displays changes incrementally. Accept or modify as needed.

[Claude-generated changes](https://github.com/ernesto27/tutorials/commit/5f80e12ea9aad4a51fece7a34de5db72c921ba4f)

This shows a basic interactive workflow. Next, let’s run tasks in the background without affecting our main branch.

## Claude Code with Container Use

To isolate work on its own branch, we’ll leverage containers. We’ll use the [container-use](https://github.com/dagger/container-use) project by the Dagger team.

### Install Container Use

```bash
curl -fsSL https://raw.githubusercontent.com/dagger/container-use/main/install.sh | bash
```

Once installed, configure Claude Code to use container-use’s MCP plugin:

```bash
claude mcp add container-use -- cu stdio
```

Add the container-use rules to your `CLAUDE.md`:

```bash
curl https://raw.githubusercontent.com/dagger/container-use/main/rules/agent.md >> CLAUDE.md
```

### Running Background Tasks

Now request tests and refactoring as a background task:

```bash
claude "Create tests for the API endpoints and add a folder named controllers for the handlers, using container-use MCP"
```

Claude runs inside a container and will prompt for permission to run tools.

![](https://raw.githubusercontent.com/ernesto27/tutorials/refs/heads/master/claude-code-container-use/image3.png)

Select **Yes, and don’t ask again for container-use:environment_create** to skip future prompts.

List your background tasks:

```bash
cu list
```

```
ID                TITLE                           CREATED         UPDATED

devoted-squirrel  Go API Testing and Refactoring  5 minutes ago   11 seconds ago
```

View logs:

```bash
cu logs
```

Check out the new branch to see changes:

![](https://raw.githubusercontent.com/ernesto27/tutorials/refs/heads/master/claude-code-container-use/image4.png)

After reviewing and approving the changes, merge them into your main branch:

```bash
cu merge devoted-squirrel main
```

[Changes merged](https://github.com/ernesto27/tutorials/commit/75bb0b1c7323a5ef236d668a4dfb44093728ec6e)

## Summary

In this tutorial, you learned how to integrate Claude Code with container-use to:

- Configure and authenticate Claude Code in your terminal
- Create a Go API project and add endpoints interactively
- Install and set up container-use for isolated tasks
- Run background tasks in a container, review and merge changes

This workflow provides reproducibility, isolation, and a parallel development process without impacting your main branch.
