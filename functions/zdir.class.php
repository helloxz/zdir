<?php
	class Zdir{
		//获取.exe文件版本号
		function exe_version($filepath){
			//判断文件
			$this->checkfile($filepath);
			//判断文件后缀
			$suffix = $this->suffix($filepath);
			$status = 0;
			
			//设置支持的后缀
			$support = array(
				"exe"
			);
			//遍历后缀
			foreach( $support as  $value )
			{
				if($suffix == $value){
					$status = 1;
					break;
				}
			}
			if($status != 1){
				echo '不支持的文件格式！';
				exit;
			}
			$strFileContent = file_get_contents($filepath);
		    if($strFileContent)
		    {
		        //$strTagBefore = 'F\0i\0l\0e\0V\0e\0r\0s\0i\0o\0n\0\0\0\0\0';        // 如果使用这行，读取的是 FileVersion
		        $strTagBefore = 'P\0r\0o\0d\0u\0c\0t\0V\0e\0r\0s\0i\0o\0n\0\0';    // 如果使用这行，读取的是 ProductVersion
		        $strTagAfter = '\0\0';
		        if (preg_match("/$strTagBefore(.*?)$strTagAfter/", $strFileContent, $arrMatches))
		        {
		            if(count($arrMatches) == 2) 
		            {
		                $fileversion = str_replace("\0", '', $arrMatches[1]);
		            }
		        }
		    }
	    	echo $fileversion;
		}
		//文件图标
		function ico($suffix){
			//根据不同后缀显示不同图标
		    switch ( $suffix )
		    {
			    //音频文件
		    	case 'mp3':
		    	case 'wma':
		    	case 'wav':
		    	case 'ape':
		    	case 'flac':
		    		$ico = "fa fa-music";
		    		break;
		    	case 'pdf':
	    			$ico = "fa fa-file-pdf-o";
		    		break;
		    	case 'doc':
		    	case 'docx':
		    		$ico = "fa fa-file-word-o";
					break;
				case 'ppt':
				case 'pptx':
					$ico = "fa fa-file-powerpoint-o";
					break;
		    	case 'xls':
		    	case 'xlsx':
		    		$ico = "fa fa-file-excel-o";
		    		break;
		    	//图片文件
		    	case 'jpg':
		    	case 'png':
		    	case 'gif':
		    	case 'jpeg':
		    	case 'bmp':
		    		$ico = "fa fa-file-image-o";
		    		break;
		    	//压缩包
		    	case 'zip':
		    	case 'rar':
		    	case 'gz':
		    	case '7z':
		    		$ico = "fa fa-file-archive-o";
		    		break;
		    	//windows软件
		    	case 'exe':
		    		$ico = "fa fa-windows";
		    		break;
		    	case 'apk':
		    		$ico = "fa fa-android";
		    		break;
		    	case 'deb':
		    		$ico = "fa fa-linux";
		    		break;
		    	case 'mp4':
		    	case 'm3u8':
		    	case 'flv':
		    	case 'rm':
		    	case 'rmvb':
		    	case 'mkv':
		    	case 'avi':
		    		$ico = "fa fa-file-video-o";
		    		break;
		    	case 'py':
		    	case 'sh':
		    	case 'c':
		    	case 'cpp':
		    	case 'go':
		    		$ico = "fa fa-file-code-o";
		    		break;
		    	default:
		    		$ico = "fa fa-file-text-o";
		    		break;
		    }
		    return $ico;
		}
		//删除某个文件
		function delfile($password,$config,$filepath){
			$myip = $this->getIP();

			//遍历配置的IP
			foreach( $config['allowip'] as $ip )
			{
				$reip = strstr($myip,$ip);
				if($ip == '0.0.0.0'){
					$reip = true;
					break;
				}
				//如果已经找到结果
				elseif($reip){
					break;
				}
			}
			//对返回进行判断
			if(!$reip){
				$redata = array(
					"code"		=>	0,
					"msg"		=>	"IP不在允许范围内！"
				);
				$redata = json_encode($redata);
				echo $redata;
				exit;
			}
			//对文件进行判断
			$filepath = $this->checkfile($filepath);
			//判断密码
			if($config['password'] != $password){
				$redata = array(
					"code"		=>	0,
					"msg"		=>	"密码错误！"
				);
				$redata = json_encode($redata);
				echo $redata;
				exit;
			}
			//执行删除文件
			unlink($filepath);
			//返回json数据
			$redata = array(
				"code"		=>	1,
				"msg"		=>	"已删除"
			);
			$redata = json_encode($redata);
			echo $redata;
			exit;
		}
		//获取访客真实IP
		function getIP() { 
		    if (getenv('HTTP_CLIENT_IP')) { 
		    $ip = getenv('HTTP_CLIENT_IP'); 
		    } 
		    elseif (getenv('HTTP_X_FORWARDED_FOR')) { 
		    $ip = getenv('HTTP_X_FORWARDED_FOR'); 
		    } 
		    elseif (getenv('HTTP_X_FORWARDED')) { 
		    $ip = getenv('HTTP_X_FORWARDED'); 
		    } 
		    elseif (getenv('HTTP_FORWARDED_FOR')) { 
		    $ip = getenv('HTTP_FORWARDED_FOR'); 

		    } 
		    elseif (getenv('HTTP_FORWARDED')) { 
		    $ip = getenv('HTTP_FORWARDED'); 
		    } 
		    else { 
		    $ip = $_SERVER['REMOTE_ADDR']; 
		    } 
		    return $ip; 
	    } 
	    //验证文件是否是当前目录
	    function checkfile($filepath){
		    //获取当前路径
			$thedir = __DIR__;
			$thedir = str_replace("\\","/",$thedir);
			$thedir = str_replace("/functions","",$thedir);
			//将/zdir替换为空
			$thedir = preg_replace("/\/zdir$/",'',$thedir);
			#$thedir = str_replace("");

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
			//如果载入上级目录
			if(stripos($filepath,'../')){
				$filehash = array(
				"code"	=>	0,
				"msg"	=>	"非法请求！"
				);
				$filehash = json_encode($filehash);
				echo $filehash;
				exit;
			}
			return $filepath;
	    }
	    //判断是否是mp4
		function video($filepath){
			//echo $filepath;
			//对文件进行判断
			//$filepath = $this->checkfile($filepath);
			//获取文件后缀
			$suffix = explode(".",$filepath);
			$suffix = end($suffix);
			$suffix = strtolower($suffix);

			//允许播放的类型
			$type_arr = array('mp4','m3u8','ts');
			$re_type = gettype(array_search($suffix,$type_arr));
			
			if( $re_type === 'integer' ){
				return true;
			}
			else{
				return false;
			}
		}
		//判断是否是音乐
		public function music($filepath){
			//echo $filepath;
			//对文件进行判断
			//$filepath = $this->checkfile($filepath);
			//获取文件后缀
			$suffix = explode(".",$filepath);
			$suffix = end($suffix);
			$suffix = strtolower($suffix);

			//允许播放的类型
			$type_arr = array('mp3','wav','flac','ape');
			$re_type = gettype(array_search($suffix,$type_arr));
			
			if( $re_type === 'integer' ){
				return true;
			}
			else{
				return false;
			}
		}
		//判断是否是.exe文件
		function is_exe($filepath){
			//echo $filepath;
			//对文件进行判断
			//$filepath = $this->checkfile($filepath);
			//获取文件后缀
			$suffix = explode(".",$filepath);
			$suffix = end($suffix);
			$suffix = strtolower($suffix);

			//允许播放的类型
			$type_arr = array('exe');
			$re_type = gettype(array_search($suffix,$type_arr));
			
			if( $re_type === 'integer' ){
				return true;
			}
			else{
				return false;
			}
		}
		//判断是否是office文档
		function office($filepath){
			$suffix = explode(".",$filepath);
			$suffix = end($suffix);
			$suffix = strtolower($suffix);

			switch ($suffix) {
				case 'doc':
				case 'docx':
				case 'xls':
				case 'xlsx':
				case 'ppt':
				case 'pptx':
					return true;
					break;
				default:
					return false;
					break;
			}
		}
		//获取文件后缀
		function suffix($filepath){
			//获取文件后缀
			$suffix = explode(".",$filepath);
			$suffix = end($suffix);
			$suffix = strtolower($suffix);

			return $suffix;
		}
		//如果是指定后缀，显示查看按钮
		function is_text($filepath){
			$suffix = $this->suffix($filepath);
			//设置支持的后缀
			$support = array(
				"txt",
				"py",
				"sh",
				"conf",
				"go",
				"c",
				"cpp"
			);
			$status = false;
			foreach( $support as $value )
			{
				if($suffix == $value){
					$status = true;
					break;
				}
			}
			return $status;
		}
		//文本查看器
		function vtext($filepath){
			//判断文件
			$this->checkfile($filepath);
			//判断文件后缀
			$suffix = $this->suffix($filepath);
			$status = 0;
			
			//设置支持的后缀
			$support = array(
				"txt",
				"py",
				"sh",
				"conf",
				"go",
				"c",
				"cpp"
			);
			//遍历后缀
			foreach( $support as  $value )
			{
				if($suffix == $value){
					$status = 1;
					break;
				}
			}
			if($status != 1){
				echo '不支持的文本格式！';
				exit;
			}
			//打开文件
			$content = file_get_contents($filepath);
			//@$content = iconv('GB2312', 'UTF-8', $content);
			//$content = file_get_contents($filepath);
			$coding = mb_detect_encoding($content,"UTF-8,GBK,GB2312");
			//如果不是UTF-8编码就转换为UTF-8
			if($coding != 'UTF-8'){
				@$content = iconv('GB2312', 'UTF-8', $content);
			}
			$content = str_replace("<","&lt;",$content);
			$content = str_replace(">","&gt;",$content);

			return $content;
		}
		//markdown查看器
		function viewmd($filepath){
			$filepath = con_coding($filepath,FALSE);
			//判断文件
			$this->checkfile($filepath);
			//获取文件后缀
			$suffix = $this->suffix($filepath);

			if($suffix == 'md'){
				@$content = file_get_contents($filepath) or die('文件不存在！');
				return $content;
			}
			else{
				echo '不支持的文件后缀';
				exit;
			}
		}
		//域名切换按钮
		function https(){
			//获取当前主机名
			$server = $_SERVER['SERVER_NAME'];
			if($server == 'soft.xiaoz.org'){
				echo '<a href = "https://wget.ovh/"><i class="fa fa-expeditedssl" aria-hidden="true"></i> HTTPS</a>';
			}
			else{
				echo '<a href = "http://soft.xiaoz.org/"><i class="fa fa-globe" aria-hidden="true"></i> HTTP</a>';
			}
		}
	}
	//预览pdf
	function viewpdf($filepath){
		//对文件进行判断
		$filepath = $this->checkfile($filepath);
		$file = fopen($filepath,"r");
		fclose($file);
		return $file;
	}
	function con_coding($str,$type = TRUE){
		$os = PHP_OS;
		//如果是Windows系统则转换编码
		if( stristr($os,'WINNT') ){
			if($type === TRUE) {
				//GB2312转UTF-8
				$str = iconv('gb2312' , 'utf-8' , $str ) ? iconv('gb2312' , 'utf-8' , $str ) : $str;
			}
			else if($type === FALSE) {
				//UTF-8转GB2312
				@$str = iconv('utf-8' , 'gb2312' , $str ) ? iconv('utf-8' , 'gb2312' , $str ) : $str;
			}
			
			//echo 'dsdsd';
		}
		return $str;
	}

	$zdir = new Zdir;
?>