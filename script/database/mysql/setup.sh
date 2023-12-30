#!/bin/bash


# 确保脚本以 root 权限运行
if [ "$(id -u)" != "0" ]; then
   echo "该脚本需要以 root 权限运行。" 1>&2
   exit 1
fi

apt update -y --fix-missing && apt upgrade -y && apt autoremove -y

#!/bin/bash

echo -e "\e[92m----------------------------------S1.系统工具---------------------------------------\e[39m"
apt install -y git curl wget sudo nano lsof htop makeself dmidecode shellcheck bsdmainutils
# -------------------------------------------------------------------------------------------------
echo -e "\e[92m----------------------------------S2.网络工具---------------------------------------\e[39m"
apt install -y netstat ntpdate net-tools traceroute netscript-2.4 inetutils-ping dnsutils

# 设置 MySQL 密码
MYSQL_ROOT_PASSWORD="123456"

# 预先设置 MySQL 和 phpMyAdmin 的安装选项
echo "mysql-server mysql-server/root_password password $MYSQL_ROOT_PASSWORD" | debconf-set-selections
echo "mysql-server mysql-server/root_password_again password $MYSQL_ROOT_PASSWORD" | debconf-set-selections
echo "phpmyadmin phpmyadmin/reconfigure-webserver multiselect apache2" | debconf-set-selections
echo "phpmyadmin phpmyadmin/dbconfig-install boolean true" | debconf-set-selections
echo "phpmyadmin phpmyadmin/mysql/admin-user string root" | debconf-set-selections
echo "phpmyadmin phpmyadmin/mysql/admin-pass password $MYSQL_ROOT_PASSWORD" | debconf-set-selections
echo "phpmyadmin phpmyadmin/mysql/app-pass password $MYSQL_ROOT_PASSWORD" | debconf-set-selections
echo "phpmyadmin phpmyadmin/app-password-confirm password $MYSQL_ROOT_PASSWORD" | debconf-set-selections

# 更新软件包列表
apt update

# 安装 MySQL
echo "正在安装 MySQL..."
apt install -y default-mysql-server

# 安装 phpMyAdmin
echo "正在安装 phpMyAdmin..."
apt install -y phpmyadmin

# 安装必要的工具
echo "正在安装必要的工具..."
apt install -y sed

echo "安装完成。"
