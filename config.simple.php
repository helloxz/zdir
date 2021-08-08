<?php
	#默认读取同级目录
	$thedir = str_replace('/zdir','',__DIR__);
	$thedir = str_replace('\zdir','',__DIR__);
	$siteinfo = array(
		"site_name"	=>	"<i class='fa fa-sitemap' aria-hidden='true'></i> Zdir",
		"title"		=>	"Zdir 实用的目录列表程序",
		"keywords"	=>	"zdir,h5ai,Directory Lister,Fdscript,目录列表,目录索引",
		"description"	=>	"Zdir是一款使用PHP开发的目录列表程序，简单实用，免费开源。"
	);

	//需要忽略的目录
	$ignore	= array(
		".",
		".git",
		".user.ini",
		".htaccess",
		"favicon.ico",
		"functions",
		"config.php",
		"index.php",
		"static",
		"LICENSE",
		"template",
		"cache.php",
		"indexes.php",
		"zdir"
	);
	//设置参数
	$config = array(
		//thedir为需要读取的目录，如:/data/wwwroot/soft.xiaoz.org，仅将zdir放在子目录的情况下需要配置此项，其它请勿配置此选项
		"thedir"	=>	$thedir,
		"allowip"	=>	array(
			//"0.0.0.0",			//注意设置为0.0.0.0则不限制IP，更多说明请参考帮助文档：https://doc.xiaoz.me/#/zdir/
			"::1",
			"127.0.0.1",
			"192.168.1."
		),
		"username"	=>	"zdir",			//用户名
		"password"	=>	"xiaoz.me",		//密码
		"auth"		=>	FALSE			//是否开启访问验证
	);	
?>