# Setup Maestro Long-Running Cluster

Sets up a long-running Maestro cluster environment using Azure ARO-HCP infrastructure.

## Description

This skill automates the deployment of a Maestro long-running cluster by:
1. Verifying Azure CLI is installed and configured
2. Checking Azure account login status (must be "ARO Hosted Control Planes")
3. Cloning the ARO-HCP repository
4. Setting required environment variables
5. Running the personal development environment setup
6. Monitoring the deployment process

## Prerequisites

- Azure CLI must be installed
- Must be logged into the correct Azure account ("ARO Hosted Control Planes")
- Internet connectivity for cloning repository
- Required permissions for Azure resource creation

## Usage

```bash
/setup-maestro-cluster
```

## Steps

1. **Verify Azure CLI Installation**
   - Check if `az` command is available
   - If not found, fail with instructions to install Azure CLI

2. **Verify Azure Account Login**
   - Run `az account show` to check current account
   - Verify the account name contains "ARO Hosted Control Planes"
   - If not logged in or wrong account, fail and ask user to login first

3. **Clone ARO-HCP Repository**
   - Create a temporary directory
   - Clone `github.com/Azure/ARO-HCP` repository
   - Change to the cloned directory

4. **Configure Environment**
   - Export `USER=oasis`
   - Export `PERSIST=true`
   - Export `GITHUB_ACTIONS=true`
   - Export `GOTOOLCHAIN=go1.24.4`

5. **Deploy Personal Dev Environment**
   - Run `make personal-dev-env`
   - Monitor the process for completion
   - Report success or failure

6. **Cleanup**
   - Return to original directory
   - Optionally clean up temporary clone (or keep for debugging)

## Implementation

```bash
#!/bin/bash
set -e

echo "Starting Maestro long-running cluster setup..."

# Step 1: Check if Azure CLI is installed
if ! command -v az &> /dev/null; then
    echo "ERROR: Azure CLI is not installed."
    echo "Please install it from: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli"
    exit 1
fi

echo "✓ Azure CLI is installed"

# Step 2: Verify Azure account login
if ! az account show &> /dev/null; then
    echo "ERROR: Not logged into Azure."
    echo "Please run: az login"
    exit 1
fi

ACCOUNT_NAME=$(az account show --query "name" -o tsv)
echo "Current Azure account: $ACCOUNT_NAME"

if [[ ! "$ACCOUNT_NAME" =~ "ARO Hosted Control Planes" ]]; then
    echo "ERROR: Not logged into the correct Azure account."
    echo "Expected account containing 'ARO Hosted Control Planes', but got: $ACCOUNT_NAME"
    echo "Please login to the correct account using: az login"
    exit 1
fi

echo "✓ Logged into correct Azure account"

# Step 3: Clone ARO-HCP repository
TEMP_DIR=$(mktemp -d)
echo "Cloning ARO-HCP repository to: $TEMP_DIR"

if ! git clone https://github.com/Azure/ARO-HCP "$TEMP_DIR/ARO-HCP"; then
    echo "ERROR: Failed to clone ARO-HCP repository"
    exit 1
fi

echo "✓ Repository cloned successfully"

# Step 4 & 5: Configure environment and deploy
pushd "$TEMP_DIR/ARO-HCP" > /dev/null

echo "Setting environment variables..."
export USER=oasis
export PERSIST=true
export GITHUB_ACTIONS=true
export GOTOOLCHAIN=go1.24.4

echo "USER=$USER"
echo "PERSIST=$PERSIST"
echo "GITHUB_ACTIONS=$GITHUB_ACTIONS"
echo "GOTOOLCHAIN=$GOTOOLCHAIN"

echo ""
echo "Starting personal-dev-env deployment..."
echo "This may take several minutes..."
echo ""

if make personal-dev-env; then
    echo ""
    echo "✓ Deployment completed successfully!"
    echo "ARO-HCP repository location: $TEMP_DIR/ARO-HCP"
else
    echo ""
    echo "ERROR: Deployment failed!"
    popd > /dev/null
    exit 1
fi

popd > /dev/null

echo ""
echo "Setup complete!"
echo "Note: The ARO-HCP repository has been cloned to: $TEMP_DIR/ARO-HCP"
echo "You can navigate there to manage the environment."
```

## Notes

- The deployment process can take several minutes to complete
- The cloned repository is placed in a temporary directory and retained for debugging
- All required environment variables are set automatically
- The script will fail fast if prerequisites are not met
