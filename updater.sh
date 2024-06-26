#!/bin/bash

# Determine the directory where the script is located
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")

# Change to the directory of the script
cd "$SCRIPT_DIR" || exit

# Fetch latest changes from the origin
git fetch origin

# Check if the HEAD commit differs from the upstream
if [ $(git rev-parse HEAD) != $(git rev-parse @{u}) ]; then
    # Log update
    logger -t jcsblogupdater "Updating jcsblog"
    git pull
    go build -o /usr/local/bin/jcsblog .

    # Check and terminate existing tmux session
    if tmux has-session -t persistent-server 2>/dev/null; then
        logger -t jcsblogupdater "Killing existing persistent-server tmux session"
        tmux send-keys -t persistent-server:0.0 C-c
        tmux send-keys -t persistent-server:0.0 "exit" Enter
    fi

    # Start a new tmux session
    logger -t jcsblogupdater "Starting new persistent-server tmux session"
    tmux new-session -d -s persistent-server
    tmux send-keys -t persistent-server:0.0 "/usr/local/bin/jcsblog" Enter
    logger -t jcsblogupdater "jcsblog updated and restarted"
else
    logger -t jcsblogupdater "No updates found"
fi
