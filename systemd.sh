#!/usr/bin/bash
useradd webhook
groupadd webhook
usermod -a -G  webhook webhook
# mv /root/webhook /usr/local/bin/
chown webhook:webhook /usr/local/bin/prometheus
chmod a+x /usr/local/bin/webhook


cat <<EOF | sudo tee /etc/systemd/system/webhook.service

[Unit]
Description=webhook
After=network-online.target

[Service]
User=webhook
Group=webhook
Type=simple
ExecStart=/bin/sh -c "/usr/local/bin/wechatrobot-webhook -RobotKey='99cf40db-e0c7-4731-904b-d809cfb1570d'"
Restart=always

[Install]
WantedBy=multi-user.target

EOF

systemctl daemon-reload
systemctl enable webhook
systemctl restart webhook
systemctl status webhook

