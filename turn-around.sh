#!/bin/bash

go build .
sudo rm /usr/local/bin/terminal-launcher
sudo mv terminal-launcher /usr/local/bin/terminal-launcher
