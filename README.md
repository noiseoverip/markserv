# markserve

go build main.go
echo DIR=YOURDIR >> /etc/markser.env
sudo cp markserv.service /etc/systemd/system/markserv.service
sudo systemctl enable markserv
sudo systemctl start markserv

