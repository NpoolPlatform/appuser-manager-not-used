FROM jrei/systemd-centos:7

COPY AppUserManager.viper.yaml /etc/AppUserManager/AppUserManager.viper.yaml
COPY appuser-manager /usr/local/bin/appuser-manager

CMD ["appuser-manager", "run"]
