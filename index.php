<?php
/**
* 
* Zdir首页入口文件
* @author	xiaoz<xiaoz93@outlook.com>
* @version	null
*/
error_reporting(E_ALL^E_NOTICE^E_WARNING^E_DEPRECATED);
//获取控制器
$c = @$_GET['c'];
//进行过滤
$c = strip_tags($c);
//读取版本号
$version = @file_get_contents("./functions/version.txt");
//载入配置文件
if( !file_exists('./config.php') ) {
	exit('<h3>配置文件不存在，请将config.simple.php复制一份并命名为config.php</h3>');
}

require("./config.php");
//载入zdir类
include_once("./functions/zdir.class.php");
//获取密码
$password = $config['password'];
//获取用户名
$username = $config['username'];
//如果开启了验证
if(@$config['auth'] === TRUE){
	//基本验证
	if ( ($_SERVER['PHP_AUTH_PW'] !== $password) || ($_SERVER['PHP_AUTH_USER'] !== $username) ){
		header('WWW-Authenticate: Basic realm="Please verify."');
		header('HTTP/1.0 401 Unauthorized');
		exit('权限不足！');
	}
}
//根据不同的请求载入不同的方法
//如果没有请求控制器
if((!isset($c)) || ($c == '')){
	//载入主页
	include_once("./functions/home.php");
}
//不允许访问的控制器
// else if($c == 'indexes'){
// 	echo '非法请求！';
// 	exit;
// }
//如果是文件管理器
else if($c == 'admin') {
	//如果当前目录存在
	if( file_exists('./functions/admin.php') ) {
		header("Location: ./functions/admin.php");
		exit;
	}
	else{
		header("Location: ./zdir/functions/admin.php");
		exit;
	}
}
else{
	include_once("./functions/".$c.'.php');
}
