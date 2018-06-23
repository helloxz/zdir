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
	}

	$zdir = new Zdir;
?>