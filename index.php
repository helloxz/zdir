<?php
	error_reporting(E_ALL^E_NOTICE^E_WARNING^E_DEPRECATED);
	//获取控制器
	$c = @$_GET['c'];
	//进行过滤
	$c = strip_tags($c);
	//读取版本号
	$version = @file_get_contents("./functions/version.txt");
	//载入配置文件
	include_once("./config.php");
	//载入zdir类
	include_once("./functions/zdir.class.php");
	//根据不同的请求载入不同的方法
	//如果没有请求控制器
	if((!isset($c)) || ($c == '')){
		//载入主页
		include_once("./functions/home.php");
	}
	else if($c == 'indexes'){
		echo '非法请求！';
		exit;
	}
	else{
		include_once("./functions/".$c.'.php');
	}
	
?>