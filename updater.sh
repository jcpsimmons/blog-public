#!/bin/bash

# check if there are any updates on github, pull and restart the server
git fetch origin
if [ $(git rev-parse HEAD) != $(git rev-parse @{u}) ]; then
    git pull
    sudo systemctl restart jcsblogd.service
fi
