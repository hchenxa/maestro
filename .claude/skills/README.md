# Maestro Claude Skills

This directory contains custom Claude Code skills for Maestro development and operations.

## Available Skills

### setup-maestro-cluster

Sets up a long-running Maestro cluster environment using Azure ARO-HCP infrastructure.

**Usage:**
```bash
/setup-maestro-cluster
```

**What it does:**
1. Verifies Azure CLI installation and login status
2. Checks that you're logged into the "ARO Hosted Control Planes" Azure account
3. Clones the ARO-HCP repository to a temporary location
4. Sets required environment variables (USER, PERSIST, GITHUB_ACTIONS, GOTOOLCHAIN)
5. Runs `make personal-dev-env` to deploy the environment
6. Monitors and reports deployment status

**Prerequisites:**
- Azure CLI installed (`brew install azure-cli` on macOS)
- Logged into correct Azure account: `az login`
- Valid Azure permissions for resource creation

**Environment Variables Set:**
- `USER=oasis`
- `PERSIST=true`
- `GITHUB_ACTIONS=true`
- `GOTOOLCHAIN=go1.24.4`

## Hooks

### deployment-monitor.sh

A hook that monitors long-running deployment processes and sends notifications.

**Features:**
- Desktop notifications (macOS/Linux)
- Slack notifications via webhook
- Customizable status messages

**Configuration:**

To enable Slack notifications:

1. Create a Slack webhook:
   - Go to https://api.slack.com/messaging/webhooks
   - Create an Incoming Webhook for your channel
   - Copy the webhook URL

2. Configure the hook:
   ```bash
   cd .claude/hooks
   cp config.sh.example config.sh
   # Edit config.sh and set your SLACK_WEBHOOK_URL
   ```

3. Or set it as an environment variable:
   ```bash
   export SLACK_WEBHOOK_URL="https://hooks.slack.com/services/YOUR/WEBHOOK/URL"
   ```

**Usage:**
```bash
# Monitor a deployment
.claude/hooks/deployment-monitor.sh monitor

# Send a completion notification
.claude/hooks/deployment-monitor.sh notify "COMPLETE" "Deployment finished successfully"

# Send a failure notification
.claude/hooks/deployment-monitor.sh notify "FAILED" "Deployment failed with errors"
```

The hook will:
1. Send a Slack notification (if configured) with color-coded messages
2. Send desktop notifications on macOS (via osascript) or Linux (via notify-send)

## How Skills Work

Skills are invoked in Claude Code using the `/` prefix followed by the skill name. When you run a skill:

1. Claude Code reads the skill markdown file
2. Executes the bash script in the Implementation section
3. Returns the output to you in the chat

Skills are a powerful way to automate complex, multi-step workflows that you perform frequently.

## Creating New Skills

To create a new skill:

1. Create a new `.md` file in `.claude/skills/`
2. Include these sections:
   - Title and description
   - Prerequisites
   - Usage example
   - Implementation (bash script in a code block)
3. Make sure the bash script is well-commented and handles errors

See `setup-maestro-cluster.md` as an example.
