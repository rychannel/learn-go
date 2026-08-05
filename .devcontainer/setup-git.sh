#!/bin/bash

# Source the local environment file if it exists
if [ -f /workspaces/learn-go/.env.local ]; then
    source /workspaces/learn-go/.env.local
    
    # Configure git if variables are set
    if [ -n "$GIT_USER_NAME" ] && [ -n "$GIT_USER_EMAIL" ]; then
        git config --global user.name "$GIT_USER_NAME"
        git config --global user.email "$GIT_USER_EMAIL"
        echo "✓ Git configured: $GIT_USER_NAME <$GIT_USER_EMAIL>"
    else
        echo "⚠ Git credentials not found in .env.local"
    fi
else
    echo "⚠ .env.local not found. Copy .env.local.example to .env.local and configure your git credentials."
fi

# Show Go version
go version
