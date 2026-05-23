#!/bin/bash

go build .
sudo rm /usr/local/bin/alacritty-launcher
sudo mv alacritty-launcher /usr/local/bin/alacritty-launcher
