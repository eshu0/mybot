systemctl stop mybot

cp /home/pi/go/src/github.com/eshu0/mybot/mybot.service /etc/systemd/system/mybot.service
 
systemctl daemon-reload

systemctl enable mybot.service

systemctl start mybot.service

systemctl status mybot.service




