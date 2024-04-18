#!/bin/bash

#  The script is pretty straightforward. It checks if there are any updates on the remote repository, pulls the updates, builds the binary, and restarts the server.
#  The script is scheduled to run every 5 minutes using a cron job.
#  # Path: /etc/cron.d/jcsblog
# */5 * * * * root /usr/local/bin/updater.sh

#  The script also logs the update process using the  logger  command.
#  # Path: /etc/rsyslog.d/jcsblog.conf

git fetch origin

if [ $(git rev-parse HEAD) != $(git rev-parse @{u}) ]; then
    # update syslog
    logger -t jcsblogupdater "Updating jcsblog"
    git pull
    go build -o /usr/local/bin/jcsblog .

    # check if tmux session persistent-server exists
    if tmux has-session -t persistent-server 2>/dev/null; then
        logger -t jcsblogupdater "Killing existing persistent-server tmux session"
        tmux send-keys -t persistent-server:0.0 C-c
        tmux send-keys -t persistent-server:0.0 "exit" Enter
    fi

    logger -t jcsblogupdater "Starting new persistent-server tmux session"
    tmux new-session -d -s persistent-server
    tmux send-keys -t persistent-server:0.0 "/usr/local/bin/jcsblog" Enter
    logger -t jcsblogupdater "jcsblog updated and restarted"
else
    logger -t jcsblogupdater "No updates found"
fi
