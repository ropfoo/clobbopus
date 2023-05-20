#!/bin/sh

curl -L https://github.com/ropfoo/clobbopus/releases/download/Latest/clobbopus_data --output clobbopus_data
curl -L https://github.com/ropfoo/clobbopus/releases/download/Latest/clobbopus_server --output clobbopus_server

chmod +x clobbopus_server
chmod +x clobbopus_data