cp /home/pi/go/src/github.com/eshu0/mybot/mybot.service /etc/systemd/system/mybot.service

systemctl stop mybot
 
systemctl daemon-reload

systemctl enable example.service

systemctl start mybot.service

systemctl status mybot.service




