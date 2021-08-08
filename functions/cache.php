<?php
	error_reporting(E_ALL^E_NOTICE^E_WARNING^E_DEPRECATED);
	//载入zdir类
	@$del = $_GET['del'];
	//缓存文件夹路径
	if ( $config['thedir'] == '' ){
		$cachefile = __DIR__."/caches/indexes.html";
	}
	else{
		$cachefile = $config['thedir']."/zdir/functions/caches/indexes.html";
	}
	
	//echo $cachefile;
	//获取文件修改时间
	@$ftime = filemtime($cachefile);
	@$cachetime = date('Y-m-d H:i:s',$ftime);
	(int)@$ftime = date('YmdH',$ftime);
	(int)$thetime = date('YmdH',time());

	//计算时差
	$diff = $thetime - $ftime;
	//删除缓存文件
	if($del == 'cache') {
		unlink($cachefile);
		header("location:./?c=cache");
		exit;
	}
	//判断缓存文件是否存在
	if((!file_exists($cachefile)) || ($diff > 24)){
		$url = get_url();
		$url = $url."?c=indexes";
		$url = str_replace("\\","/",$url);
		$curl = curl_init($url);


	    curl_setopt($curl, CURLOPT_USERAGENT, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36");
	    curl_setopt($curl, CURLOPT_FAILONERROR, true);
	    curl_setopt($curl, CURLOPT_FOLLOWLOCATION, true);
	    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
	    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, false);
	    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, false);
	    #设置超时时间，最小为1s（可选）
	    #curl_setopt($curl , CURLOPT_TIMEOUT, 1);

	    $html = curl_exec($curl);
	    curl_close($curl);
	    
		#索引数据写入文件
		file_put_contents($cachefile,$html);
		#读取索引
		$cache = @file_get_contents($cachefile) or die("Unable to open file!");
	}
	else{
		$cache = file_get_contents($cachefile);
	}

	$cache = con_coding($cache);

	//获取页面URL
	function get_url(){
		$port = $_SERVER["SERVER_PORT"];
		//对端口进行判断
		switch ( $port )
		{
			case 80:
				$protocol = "http://";
				$port = '';
				break;	
			case 443:
				$protocol = "https://";
				$port = '';
				break;
			default:
				$protocol = "http://";
				$port = "";
				break;
		}
		
		//或如URI
		$uri =  $_SERVER["REQUEST_URI"];
		$uri = dirname($uri);
		$uri = str_replace("\\",'/',$uri);
		//二级目录
		if($uri != '/'){
			$uri = $uri.'/';
		}
		
		//$uri = str_replace("cache.php","",$uri);
		//如果主机名是localhost，则获取localhost
		if( $_SERVER['SERVER_NAME'] == 'localhost' ) {
			$domain = $protocol.$_SERVER['SERVER_NAME'].$port.$uri;
		}
		else {
			$domain = $protocol.$_SERVER['HTTP_HOST'].$port.$uri;
		}
		
		return $domain;
	}
?>

<!--载入页头-->
<?php include_once( './template/header.php' ); ?>

<!--中间内容部分-->
<div class="layui-container">
	<div class="layui-row">
		<div class="layui-col-lg10 layui-col-md-offset1">
			<div id="cache">
				
			
			<?php
				echo "<h1>文件索引 - 该数据缓存于$cachetime</h1><br />";
				echo "<a href='/?c=cache&del=cache' class='layui-btn layui-btn-primary layui-border-blue'>清除缓存</a>";
				echo $cache;
			?>
			</div>
		</div>
	</div>
</div>
<!--中间内容部分END-->

<!--载入页脚-->
<?php include_once( './template/footer.php' ); ?>