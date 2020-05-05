#!/bin/bash
#####		Docker for Zdir一键安装脚本		#####
#####		Update:2020-05-05				#####
#####		Author:xiaoz					#####

#导入环境变量
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/bin:/sbin
export PATH

#检查并安装Docker
function check_docker(){
	echo '-------------------------------------------'
	docker_path=$(which docker)
	if [ -e "${docker_path}" ]
	then
		echo 'Docker已安装，继续执行'
	else
		read -p "Docker未安装，是否安装Docker?(y/n):" is_docker
		if [ $is_docker == 'y' ]
			then
				curl -fsSL https://get.docker.com -o get-docker.sh
				sh get-docker.sh
			else
				echo '放弃安装！'
				echo '-------------------------------------------'
				exit
		fi
	fi
	#启动docker
	systemctl start docker
	echo '-------------------------------------------'
}

#安装前的准备
function ready(){
	#创建用户和用户组
	groupadd www
	useradd -M -g www www -s /sbin/nologin
	#CentOS
	if [ -e "/usr/bin/yum" ]
	then
		yum -y update
		yum -y install unzip wget
	else
		#更新软件，否则可能make命令无法安装
		apt-get update
		apt-get install -y unzip wget
	fi
}

#安装Zdir
function install_zdir(){
	echo '-------------------------------------------'
	read -p "请输入Zdir安装目录（如果留空，则默认为/data/wwwroot/zdir）:" zdir_path
	#如果路径为空
	if [ -z "${zdir_path}" ]
	then
		zdir_path='/data/wwwroot/zdir'
	fi
	#创建目录
	mkdir -p $zdir_path
	#下载源码
	wget -O ${zdir_path}/zdir.zip https://github.com/helloxz/zdir/archive/master.zip
	#进入目录
	cd $zdir_path
	unzip -o zdir.zip
	mv zdir-master/* ./
	rm -rf zdir-master
	#重命名配置文件
	mv config.simple.php config.php
	echo '-------------------------------------------'
	#设置文件管理器密码
	read -p "请设置文件管理器密码:" zdir_pass
	#如果密码为空，循环让用户输入
	while [ -z "${zdir_pass}" ]
	do
		read -p "请设置文件管理器密码:" zdir_pass
	done
	#设置密码
	sed -i "s/\"xiaoz.me\"/\"${zdir_pass}\"/g" ${zdir_path}/config.php
	#设置用户组权限
	chown -R www:www $zdir_path
}
#自动放行端口
function chk_firewall(){
	if [ -e "/etc/sysconfig/iptables" ]
	then
		iptables -I INPUT -p tcp --dport 1080 -j ACCEPT
		service iptables save
		service iptables restart
	elif [ -e "/etc/firewalld/zones/public.xml" ]
	then
		firewall-cmd --zone=public --add-port=1080/tcp --permanent
		firewall-cmd --reload
	elif [ -e "/etc/ufw/before.rules" ]
	then
		sudo ufw allow 1080/tcp
	fi
}

#运行容器
function zdir_run(){
	docker run --name="zdir"  \
    -d -p 1080:80 --restart=always \
    -v ${zdir_path}:/data/wwwroot/default \
    helloz/caddy-php:v1.3 \
    /usr/sbin/start.sh
    #获取ip
	osip=$(curl -4s https://api.ip.sb/ip)
    echo '-------------------------------------------'
    echo '安装完毕，请访问http://'${osip}:1080
    echo 'Zdir安装路径为:'${zdir_path}
    echo 'Zdir用户名为:zdir,密码为:'${zdir_pass}
    echo '如需帮助，请访问:https://dwz.ovh/zdir'
    echo '-------------------------------------------'
}

check_docker
ready
install_zdir
chk_firewall
zdir_run