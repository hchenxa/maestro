#!/bin/bash
#
# Deployment Monitor Hook
# This hook monitors long-running deployment processes and notifies when complete
#
# Usage: Can be called after triggering a deployment to monitor its progress
#
# Configuration:
# Set SLACK_WEBHOOK_URL environment variable or create .claude/hooks/config.sh

set -e

HOOK_NAME="deployment-monitor"

# Load configuration if exists
CONFIG_FILE="$(dirname "$0")/config.sh"
if [ -f "$CONFIG_FILE" ]; then
    source "$CONFIG_FILE"
fi

# Function to check if a process is still running
check_process() {
    local pid=$1
    if ps -p "$pid" > /dev/null 2>&1; then
        return 0  # Process is running
    else
        return 1  # Process is not running
    fi
}

# Function to monitor deployment
monitor_deployment() {
    local start_time=$(date +%s)
    local max_wait_time=3600  # 1 hour max

    echo "[$HOOK_NAME] Monitoring deployment process..."

    while [ $(($(date +%s) - start_time)) -lt $max_wait_time ]; do
        # Check if deployment is complete by looking for success indicators
        # This is a placeholder - adjust based on actual deployment markers

        # Sleep for 30 seconds before next check
        sleep 30

        echo "[$HOOK_NAME] Deployment still in progress... (elapsed: $(($(date +%s) - start_time))s)"
    done

    echo "[$HOOK_NAME] Maximum wait time reached"
}

# Function to send Slack notification
send_slack_notification() {
    local status=$1
    local message=$2
    local webhook_url=$3

    if [ -z "$webhook_url" ]; then
        return 1
    fi

    # Determine color based on status
    local color="good"
    local emoji=":white_check_mark:"
    if [[ "$status" == "FAILED" ]]; then
        color="danger"
        emoji=":x:"
    elif [[ "$status" == "COMPLETE" ]]; then
        color="good"
        emoji=":white_check_mark:"
    fi

    # Create JSON payload
    local payload=$(cat <<EOF
{
  "attachments": [
    {
      "color": "$color",
      "title": "$emoji Maestro Deployment $status",
      "text": "$message",
      "footer": "Maestro Deployment Monitor",
      "ts": $(date +%s)
    }
  ]
}
EOF
)

    # Send to Slack
    curl -X POST -H 'Content-type: application/json' \
        --data "$payload" \
        "$webhook_url" \
        --silent --show-error
}

# Function to send notification
notify_completion() {
    local status=$1
    local message=$2

    echo ""
    echo "=========================================="
    echo "[$HOOK_NAME] DEPLOYMENT $status"
    echo "Message: $message"
    echo "Time: $(date)"
    echo "=========================================="
    echo ""

    # Send Slack notification if webhook is configured
    if [ -n "$SLACK_WEBHOOK_URL" ]; then
        echo "[$HOOK_NAME] Sending Slack notification..."
        if send_slack_notification "$status" "$message" "$SLACK_WEBHOOK_URL"; then
            echo "[$HOOK_NAME] Slack notification sent successfully"
        else
            echo "[$HOOK_NAME] Failed to send Slack notification"
        fi
    fi

    # Also send system notification if available
    if command -v osascript &> /dev/null; then
        # macOS notification
        osascript -e "display notification \"$message\" with title \"Maestro Deployment $status\""
    elif command -v notify-send &> /dev/null; then
        # Linux notification
        notify-send "Maestro Deployment $status" "$message"
    fi
}

# Main execution
case "${1:-monitor}" in
    monitor)
        monitor_deployment
        ;;
    notify)
        notify_completion "${2:-COMPLETE}" "${3:-Deployment finished}"
        ;;
    *)
        echo "Usage: $0 {monitor|notify}"
        exit 1
        ;;
esac
