<?php
	class Zdir{
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
				//如果已经找到结果
				if($reip){
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
			return $filepath;
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

	$zdir = new Zdir;
?>