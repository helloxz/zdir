<?php
	$siteinfo = array(
		"title"		=>	"Zdir 实用的目录列表程序",
		"keywords"	=>	"zdir,h5ai,Directory Lister,Fdscript,目录列表,目录索引",
		"description"	=>	"Zdir是一款使用PHP开发的目录列表程序，简单实用，免费开源。"
	);

	//需要忽略的目录
	$ignore	= array(
		".",
		".git",
		".user.ini",
		"favicon.ico",
		"functions",
		"config.php",
		"index.php",
		"static",
		"LICENSE",
		"template",
		"cache.php",
		"indexes.php"
	);
	//设置IP与密码
	$config = array(
		"allowip"	=>	array(
			//"0.0.0.0",			//注意设置为0.0.0.0则不限制IP，更多说明请参考帮助文档：https://doc.xiaoz.me/#/zdir/
			"::1",
			"127.0.0.1",
			"192.168.1."
		),
		"password"	=>	"xiaoz.me",
		"auth"		=>	FALSE
	);
	//设置上传目录,以网站根目录开始
	$uploadDir = "/upload";
?>
