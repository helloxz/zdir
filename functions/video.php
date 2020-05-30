<?php
	//获取视频播放地址
	@$url = $_GET['url'];
	$url = con_coding($url,FALSE);
	if( $config['thedir'] != '' ){
		$url = str_replace("./","../",$url);
	}
	
	//判断文件是否存在
	if(!file_exists($url)){
		echo '视频文件不存在！';
		exit;
	}
	$url = con_coding($url);
?>
<!DOCTYPE html>
<html lang="en" xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta charset="utf-8" />
	<title>视频播放器</title>
	<meta name="generator" content="EverEdit" />
	<meta name="author" content="" />
	<meta name="keywords" content="" />
	<meta name="description" content="" />
	<link rel="stylesheet" href="./static/dplayer/DPlayer.min.css" type="" media=""/>
</head>
<body>
	<!--视频播放容器-->
	<div id="dplayer"></div>
	<!--视频播放容器END-->
	<!--<video id="my-video" class="video-js" controls preload="auto" width="1280" height="720" data-setup="{}">
    	<source src="<?php echo $url; ?>">
	<?php
	if (file_exists(str_replace(".mp4", ".vtt", $url)))
	{
	?>
	<track src="<?php echo str_replace(".mp4", ".vtt", $url); ?>" label="English" kind="subtitles" srclang="en" default>
	<?php
	}
	?>
  	</video>-->
	<script src="./static/hls.min.js"></script>
	<script src = "./static/dplayer/DPlayer.min.js"></script>
	<script type="text/javascript">
		const dp = new DPlayer({
	    container: document.getElementById('dplayer'),
	    video: {
	        url: '<?php echo $url; ?>'
	    }
	});
	</script>
</body>
</html>
