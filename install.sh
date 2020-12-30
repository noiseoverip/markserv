#!/bin/bash
go build

sudo systemctl stop markserv
sudo cp markserv /usr/local/bin/markserv
sudo cp markserv.service /etc/systemd/system/markserv.service
sudo systemctl daemon-reload
sudo systemctl restart markserv
sudo systemctl enable markserv

