<?php
	//载入文件
	if ( $config['thedir'] == '' ){
		include_once("./zdir.class.php");
	}
	else{
		include_once($config['thedir']."/zdir/functions/zdir.class.php");
	}
	//获取当前路径
	$thedir = __DIR__;
	$thedir = str_replace("\\","/",$thedir);
	$thedir = str_replace("/functions","",$thedir);
	$thedir = str_replace("/zdir","",$thedir);
	#$thedir = str_replace("");
	//获取文件路径
	$filepath = $_POST['file'];
	$filepath = con_coding($filepath,FALSE);
	//var_dump($filepath);

	//如果文件不存在
	if(!is_file($filepath)) {
		$filehash = array(
		"code"	=>	0,
		"msg"	=>	"文件不存在！"
		);
		$filehash = json_encode($filehash);
		echo $filehash;
		exit;
	}
	//如果文件不是项目路径
	if(!strstr($filepath,$thedir)){
		$filehash = array(
		"code"	=>	0,
		"msg"	=>	"目录不正确！"
		);
		$filehash = json_encode($filehash);
		echo $filehash;
		exit;
	}
	
	$md5 = md5_file($filepath);
	
	$sha1 = sha1_file($filepath);

	//计算文件hash
	$filehash = array(
		"code"	=>	1,
		"md5"	=>	$md5,
		"sha1"	=>	$sha1
	);
	
	$filehash = json_encode($filehash);
	echo $filehash;
?>