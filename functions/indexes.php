<?php
#设置超时时间
ini_set('max_execution_time', '600');
include_once( '../config.php' );

/**********************
一个简单的目录递归函数
第一种实现办法：用dir返回对象
***********************/
function tree($directory,$ignore) { 
    if(is_dir($directory)) {
        //返回一个 Directory 类实例
        $mydir = dir($directory);
        echo "<ul>\n";
        //从目录句柄中读取条目
        while($file = $mydir->read()) {
	        if(array_search($file,$ignore)) {
			    continue;
		    }
            if(is_dir("$directory/$file") && $file != "." && $file != "..") {
                echo "<li><i class='fa fa-folder-open'></i> <font color=\"#FF5722\"><b>$file</b></font></li>\n";
                //递归读取目录 
                tree("$directory/$file",$ignore);
            } elseif ($file != "." && $file != "..") {
	            $uri =  $_SERVER["REQUEST_URI"];
				$uri = dirname($uri);
				$uri = str_replace("/functions","",$uri);
				//echo $uri;
				//exit;
	            $filepath = "$directory/$file";
	            $url = $uri.'/'.$directory.'/'.$file;
                $url = str_replace("../","",$url);
                $url = str_replace("\/","/",$url);
                $url = str_replace("//","/",$url);
                $url = str_replace("//","/",$url);
                echo "<li><i class='fa fa-file-text-o'></i> <a href = '.$url' target = '_blank'>$url</a></li>\n";
            }

        }
        echo "</ul>\n";
        // 释放目录句柄
        $mydir->close();
    } else {
        echo $directory . '<br>';
    }

} 
//开始运行
//如果是子目录运行
if( $config['thedir'] != '' ) {
    tree('../',$ignore); 
}
//如果是顶级目录运行
else{
    tree('./',$ignore);
}