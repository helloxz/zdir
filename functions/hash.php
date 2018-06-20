<?php
	//获取文件路径
	$filepath = $_POST['file'];

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

	$md5 = md5_file($filepath);
	$sha1 = sha1_file($filepath);

	//计算文件hash
	$filehash = array(
		"code"	=>	1,
		"path"	=>	$filepath,
		"md5"	=>	$md5,
		"sha1"	=>	$sha1
	);
	$filehash = json_encode($filehash);
	echo $filehash;
?>