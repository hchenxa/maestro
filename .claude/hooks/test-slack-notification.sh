#!/bin/bash
#
# Test Slack Notification
# This script tests the Slack webhook configuration
#

set -e

HOOK_DIR="$(cd "$(dirname "$0")" && pwd)"

# Load configuration
CONFIG_FILE="$HOOK_DIR/config.sh"
if [ ! -f "$CONFIG_FILE" ]; then
    echo "ERROR: Config file not found: $CONFIG_FILE"
    echo ""
    echo "To create the config file:"
    echo "  cd $HOOK_DIR"
    echo "  cp config.sh.example config.sh"
    echo "  # Edit config.sh and set your SLACK_WEBHOOK_URL"
    exit 1
fi

source "$CONFIG_FILE"

if [ -z "$SLACK_WEBHOOK_URL" ]; then
    echo "ERROR: SLACK_WEBHOOK_URL is not set in $CONFIG_FILE"
    echo ""
    echo "Please edit $CONFIG_FILE and set your Slack webhook URL"
    exit 1
fi

echo "Testing Slack notification..."
echo "Webhook URL: ${SLACK_WEBHOOK_URL:0:30}..."
echo ""

# Send test notification
"$HOOK_DIR/deployment-monitor.sh" notify "TEST" "This is a test notification from Maestro deployment monitor"

echo ""
echo "Test notification sent!"
echo "Check your Slack channel to verify it was received."
