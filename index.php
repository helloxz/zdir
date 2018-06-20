<?php
	error_reporting(E_ALL^E_NOTICE^E_WARNING^E_DEPRECATED);
	//载入配置文件
	include_once("./config.php");
	//获取当前目录
	$thedir = __DIR__;
	$i = 0;
	

	//获取目录
	$dir = $_GET['dir'];
	//分割字符串
	$navigation = explode("/",$dir);

	if(($dir == '') || (!isset($dir))) {
		$listdir = scandir($thedir);
	}
	else{
		$listdir = scandir($thedir."/".$dir);
	}

	//计算上级目录
	function updir($dir){
		//分割目录
		$dirarr = explode("/",$dir);
		$dirnum = count($dirarr);
		
		#var_dump($dirarr);
		if($dirnum == 2) {
			$updir = 'index.php';
		}
		else{
			$updir = '';
			for ( $i=1; $i < ($dirnum - 1); $i++ )
			{ 
				$next = $i + 1;
				$updir = $updir.'/'.$dirarr[$i];
				
			}
			$updir = 'index.php?dir='.$updir;
		}
		return $updir;
	}
	#echo updir($dir);
	$updir = updir($dir);
?>
<!DOCTYPE html>
<html lang="zh-cmn-Hans" xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta charset="utf-8" />
	<title><?php echo $siteinfo['title']; ?></title>
	<meta name="generator" content="EverEdit" />
	<meta name="author" content="xiaoz.me" />
	<meta name="keywords" content="<?php echo $siteinfo['keywords']; ?>" />
	<meta name="description" content="<?php echo $siteinfo['description']; ?>" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="shortcut icon" href="favicon.ico"  type="image/x-icon" />
	<link rel="stylesheet" href="./static/layui/css/layui.css">
	<link rel='stylesheet' href='./static/style.css'>
	<link rel="stylesheet" href="./static/font-awesome/css/font-awesome.min.css">
</head>
<body>
	<!--顶部导航栏-->
	<div class = "header">
        <div class = "layui-container">
            <div class = "layui-row">
                <div class = "layui-col-lg12">
                    <div class = "layui-hide-xs">
                        <ul class="layui-nav menu" lay-filter="">
                            <li class="layui-nav-item"><a href="./"><i class="fa fa-home fa-lg"></i> Zdir</a></li>
                            <li class="layui-nav-item"><a href="https://github.com/helloxz/zdir" target = "_blank" rel = "nofollow"><i class="fa fa-github fa-lg"></i> 源码</a></li>
                            <li class="layui-nav-item"><a href="https://www.xiaoz.me/archives/10465" target = "_blank"><i class="layui-icon">&#xe60b;</i> 关于</a></li>
                        </ul>
                    </div>
                    <div class = "layui-hide-lg">
                        <ul class="layui-nav menu" lay-filter="">
                            <li class="layui-nav-item"><a href="./"><i class="fa fa-home fa-lg"></i> Zdir</a></li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <!-- 顶部导航栏END -->
    <!--面包屑导航-->
	<div id="navigation" class = "layui-hide-xs">
		<div class="layui-container">
			<div class="layui-row">
				<div class="layui-col-lg12">
					<p>
						当前位置：<a href="./">首页</a> 
						<!--遍历导航-->
						<?php foreach( $navigation as $menu )
						{
							$remenu = $remenu.'/'.$menu;
							
							if($remenu == '/'){
								$remenu = $menu;
							}
						?>
						<a href="./index.php?dir=<?php echo $remenu; ?>"><?php echo $menu; ?></a> / 
						<?php } ?>
					</p>
				</div>
			</div>
		</div>
	</div>
    <!--面包屑导航END-->
	<!--遍历目录-->
	<div id="list">
		<div class="layui-container">
		  	<div class="layui-row">
		    	<div class="layui-col-lg12">
			    	<table class="layui-table" lay-skin="line">
					  	<colgroup>
					    <col width="400">
					    <col width="200">
					    <col width="200">
					    <col width="200">
					    <col>
					  </colgroup>
					  <thead>
					    <tr>
					      <th>文件名</th>
					      <th class = "layui-hide-xs"></th>
					      <th>修改时间</th>
					      <th>文件大小</th>
					      <th class = "layui-hide-xs">操作</th>
					    </tr> 
					  </thead>
					  <tbody>
					    <?php foreach( $listdir as $showdir ) {
						    //防止中文乱码
						    $showdir = iconv('gb2312' , 'utf-8' , $showdir );
						    //文件完整路径
						    $fullpath = $thedir.'/'.$dir.'/'.$showdir;
						    $fullpath = str_replace("\\","\/",$fullpath);
						    
						    //获取文件修改时间
						    $ctime = filemtime($fullpath);
						    $ctime = date("Y-m-d H:i",$ctime);

						    
						    //搜索忽略的目录
						    if(array_search($showdir,$ignore)) {
							    continue;
						    }
						    
						    //判读文件是否是目录,当前路径 + 获取到的路径 + 遍历后的目录
						    if(is_dir($thedir.'/'.$dir.'/'.$showdir)){
							    $suffix = '';
							    //设置上级目录
							    if($showdir == '..'){
								    $url = $updir;
							    }
							    else{
								    $url = "./index.php?dir=".$dir.'/'.$showdir;
							    }
							    
							    $ico = "fa fa-folder-open";
							    $fsize = '-';
							    //返回类型
							    $type = 'dir';
						    }
						    //如果是文件
						    if(is_file($fullpath)){
							    //获取文件后缀
						    	$suffix = explode(".",$showdir);
						    	$suffix = end($suffix);
						    	
						    	
							    $url = '.'.$dir.'/'.$showdir;
							    
							    $ico = "fa fa-file-text-o";

							    //获取文件大小
							    $fsize = filesize($fullpath);
							    $fsize = ceil ($fsize / 1024);
							    $fsize = $fsize.'kb';
							    $type = 'file';
							    #$info = "<a href = ''><i class='fa fa-info-circle' aria-hidden='true'></i></a>";
						    }
						    //其它情况，可能是中文目录
						    else{
							    $suffix = '';
							    //设置上级目录
							    if($showdir == '..'){
								    $url = $updir;
							    }
							    else{
								    $url = "./index.php?dir=".$dir.'/'.$showdir;
							    }
							    
							    $ico = "fa fa-folder-open";
							    $fsize = '-';
							    $type = 'dir';
						    }
						    $i++;
						?>
					    <tr id = "id<?php echo $i; ?>">
						    <td>
							    <a href="<?php echo $url ?>" id = "url<?php echo $i; ?>"><i class="<?php echo $ico; ?>"></i> <?php echo $showdir; ?></a>
						    </td>
						    <td id = "info" class = "layui-hide-xs">
							    <!--如果是readme.md-->
							    <?php if(($showdir == 'README.md') || ($showdir == 'readme.md')){ ?>
								<a class = "layui-btn layui-btn-xs" href="javascript:;" onclick = "viewmd('<?php echo $url ?>')" title = "点此查看使用说明">使用说明</a>
							    <?php } ?>
							    <!--如果是文件-->
							    <?php if($type == 'file'){ ?>
								<a href="javascript:;" title = "查看文件hash" onclick = "filehash('<?php echo $showdir; ?>','<?php echo $fullpath; ?>')"><i class="fa fa-info-circle" aria-hidden="true"></i></a>
							    <?php } ?>
						    </td>
						    <td><?php echo $ctime; ?></td>
						    <td><?php echo $fsize; ?></td>
						    <td class = "layui-hide-xs">
							    <?php if($fsize != '-'){ ?>
								<a href="javascript:;" class = "layui-btn layui-btn-xs" onclick = "copy('<?php echo $url ?>')">复制</a>
							    <?php } ?>
							    <!--如果是markdown文件-->
							    <?php if(($suffix == 'md') && ($suffix != null)){ ?>
								&nbsp;&nbsp;<a href="javascript:;" onclick = "viewmd('<?php echo $url ?>')" title = "点击查看"><i class="fa fa-eye fa-lg"></i></a> 
							    <?php } ?>
						    </td>
					    </tr>
					    <?php } ?>
					  </tbody>
					</table>
		    	</div>
		  	</div>
		</div> 
	</div>
	<!-- 底部 -->
 	<div class = "footer">
		<div class = "layui-container">
			<div class = "layui-row">
				<div class = "layui-col-lg12">
				Copyright © 2017-2018 Powered by <a href="https://github.com/helloxz/zdir" target = "_blank">Zdir</a> | Author <a href="https://www.xiaoz.me/" target = "_blank">xiaoz.me</a>
				</div>
			</div>
		</div>
	</div>
	<!-- 底部END -->
	
	<!--遍历目录END-->
	<script type="text/javascript" src = "https://libs.xiaoz.top/jquery/2.0.3/jquery-2.0.3.min.js"></script>
	<script type="text/javascript" src="./static/layui/layui.js"></script>
	<script type="text/javascript" src="./static/embed.js"></script>
	<script type="text/javascript" src="https://libs.xiaoz.top/clipBoard.js/clipBoard.min.js"></script>
</body>
</html>