<?php
	@$admin = $_GET['admin'];
	/**
	获取需要读取的目录
	*/
	//如果没有设置thedir，则默认读取当前路径
	if( $config['thedir'] == '' ) {
		$thedir = __DIR__;
		$thedir = str_replace("\\",'/',$thedir);
		$thedir = str_replace("functions",'',$thedir);
	}
	else{
		$thedir = $config['thedir'];
	}
	
	$i = 0;

	//获取目录
	$dir = $_GET['dir'];
	$dir = con_coding($dir,FALSE);
	$dir = strip_tags($dir);
	$dir = str_replace("\\","/",$dir);
	$rel_path = $thedir."/".$dir;
	
	//获取markdown文件地址
	//判断是否是首页
	function is_home(){
		$dir = $GLOBALS['dir'];
		if(empty($dir)){
			return TRUE;
		}
		else{
			return FALSE;
		}
	}
	
	
	//echo $readme;
	//对目录进行过滤
	if((stripos($dir,'./') === 0) || (stripos($dir,'../')) || (stripos($dir,'../') === 0) || (stripos($dir,'..') === 0) || (stripos($dir,'..'))){
		echo '非法请求！';
		exit;
	}
	//分割字符串
	$navigation = explode("/",$dir);

	if(($dir == '') || (!isset($dir))) {
		$listdir = scandir($thedir);
		$readme = $thedir.'/README.md';
	}
	//如果目录不存在
	else if(!is_dir($rel_path)){
		echo '目录不存在，3s后返回首页！';
		header("Refresh:3;url=/");
		exit;
	}
	else{
		$listdir = scandir($thedir."/".$dir);
		$readme = $thedir."/".$dir.'/README.md';
	}
	//遍历目录和文件，并进行排序，文件夹排前面
	$newdir = array();
	$newfile = array();
	foreach( $listdir as $value )
	{
		//如果参数为空
		if(!isset($dir)){
			$tmp_path = $thedir;
		}
		if(isset($dir)){
			$tmp_path = $thedir.'/'.$dir.'/'.$value;
		}
		$tmp_path = str_replace("///","/",$tmp_path);
		//echo $tmp_path."<br />";
		//如果是文件夹
		if(is_dir($tmp_path)){
			array_push($newdir,$value);
		}
		else{
			array_push($newfile,$value);
		}
	}
	//两个数组顺序合并
	$listdir = array_merge($newdir,$newfile);
	//返回数组的差集
	$listdir = array_diff($listdir,$ignore);
	//如果是首页，隐藏..
	if(is_home()){
		$listdir = array_diff($listdir,['..']);
	}
	$readme = str_replace('\\','/',$readme);
	//计算上级目录
	function updir($dir){
		//分割目录
		$dirarr = explode("/",$dir);
		$dirnum = count($dirarr);
		
		#var_dump($dirarr);
		if($dirnum == 2) {
			$updir = '/';
		}
		else{
			$updir = '';
			for ( $i=1; $i < ($dirnum - 1); $i++ )
			{ 
				$next = $i + 1;
				$updir = $updir.'/'.$dirarr[$i];
				
			}
			$updir = '/?dir='.$updir;
		}
		return $updir;
	}
	#echo updir($dir);
	$updir = updir($dir);

?>
<?php
	//载入页头
	include_once("./template/header.php")
?>
    <!--面包屑导航-->
	<div id="navigation" class = "layui-hide-xs">
		<div class="layui-container">
			<div class="layui-row">
				<!--滚动消息-->
				<div id = "msg" class="layui-col-lg12">
					<i class="layui-icon layui-icon-notice" style="color: #FF5722;font-weight:bold;"></i> 
					<span id = "msg-content"></span>
				</div>
				<!--滚动消息END-->
				<div class="layui-col-lg12">
					<p>
						当前位置：<a href="./">首页</a> 
						<!--遍历导航-->
						<?php foreach( $navigation as $menu )
						{
							$menu = con_coding($menu);
							$remenu = $remenu.'/'.$menu;
							
							if($remenu == '/'){
								$remenu = $menu;
							}
						?>
						
						<a href="./?dir=<?php echo $remenu; ?>"><?php echo $menu; ?></a> / 
						<?php } ?>
					</p>
				</div>
				<!--使用说明-->
				<?php 
				//判断readme文件是否存在
					$readme_dir = $fullpath = $thedir.'/'.$dir.'/'.$showdir;
					if( is_file($readme_dir.'/readme.md') || (is_file($readme_dir.'/README.md')) )	{	
					$readme = con_coding($readme);
				?>
				<div class="layui-col-lg12" style = "margin-top:1em;">
					<div class="layui-collapse">
					  <div class="layui-colla-item">
					    <h2 class="layui-colla-title">使用说明（必看）</h2>
					    <div class="layui-colla-content">
						    <iframe src="<?php echo './?c=readme&file='.$readme; ?>" width="100%" height="600px" name="" frameborder = "0" align="middle"></iframe>
					    </div>
					  </div>
					</div>
				</div>
				<?php } ?>
				<!--使用说明END-->
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
					    <col width="560">
					    <col width="100">
					    <col width="160">
					    <col width="160">
					    <col>
					  </colgroup>
					  <thead>
					    <tr>
					      <th>文件名</th>
					      <th class = "layui-hide-xs"></th>
					      <th class = "layui-hide-xs">修改时间</th>
					      <th>文件大小</th>
					      <th class = "layui-hide-xs layui-hide-sm layui-show-md-block">操作</th>
					    </tr> 
					  </thead>
					  <tbody>
					    <?php foreach( $listdir as $showdir ) {
						    //防止中文乱码
						    //$showdir = con_coding($showdir);
						    $fullpath = $thedir.'/'.$dir.'/'.$showdir;
						    //去掉多余的斜杠
						    $fullpath = str_replace("\\","\/",$fullpath);
						    $fullpath = str_replace("//","/",$fullpath);
						    $fullpath = str_replace("//","/",$fullpath);
						    //$fullpath = con_coding($fullpath);
						    
						    //var_dump($fullpath);
						    //获取文件修改时间
						    $ctime = filemtime($fullpath);
						    $ctime = date("Y-m-d H:i",$ctime);

						    //搜索忽略的目录，如果包含.php 一并排除
						    if( strripos($showdir,".php") ) {
							    continue;
						    }
						    //判读文件是否是目录,当前路径 + 获取到的路径 + 遍历后的目录
						    if(is_dir($fullpath)){
							    $suffix = '';
							    //设置上级目录
							    if($showdir == '..'){
								    $url = $updir;
							    }
							    else{
								    $url = "./?dir=".$dir.'/'.$showdir;
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

							    //根据不同后缀显示不同图标
							    $ico = $zdir->ico($suffix);
							    

							    //获取文件大小
							    $fsize = filesize($fullpath);
							    $fsize = ceil ($fsize / 1024);
							    if($fsize >= 1024) {
								    $fsize = $fsize / 1024;
								    if( $fsize >= 1024 ) {
									    $fsize = $fsize / 1024;
									    $fsize = round($fsize,2).'Gb';
								    }
								    else{
									    $fsize = round($fsize,2).'Mb';
								    }
							    }
							    else{
								    $fsize = $fsize.'Kb';
							    }
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
								    $url = "./?dir=".$dir.'/'.$showdir;
							    }
							    
							    $ico = "fa fa-folder-open";
							    $fsize = '-';
							    $type = 'dir';
							}
						    $i++;
						?>
						
					    <tr id = "id<?php echo $i; ?>">
						    <td>
							    <?php
							    	$showdir = con_coding($showdir);
							    	$fullpath = con_coding($fullpath);
							    	//echo $fullpath;
							    	$url = con_coding($url);
							    ?>
							    <!--判断文件是否是图片-->
							    <?php if(($suffix == 'jpg') || ($suffix == 'jpeg') || ($suffix == 'png') || ($suffix == 'gif') || ($suffix == 'bmp')){

							   	?>
							   	<a class = "fname1" href="<?php echo $url ?>" id = "url<?php echo $i; ?>" onmouseover = "showimg(<?php echo $i; ?>,'<?php echo $url; ?>')" onmouseout = "hideimg(<?php echo $i; ?>)"><i class="<?php echo $ico; ?>"></i> <?php echo $showdir; ?></a>
							   	<div class = "showimg" id = "show<?php echo $i; ?>"><img src="" id = "imgid<?php echo $i; ?>"></div>
							   	<!--如果是.exe文件-->
							   	<?php }elseif($zdir->is_exe($fullpath)){ ?>
								<a class = "fname1" href="<?php echo $url ?>" id = "url<?php echo $i; ?>"><i class="<?php echo $ico; ?>"></i> <?php echo $showdir; ?></a>
							   	<!--.exe文件END-->
							   	<?php }else{ ?>
							    <a class = "fname1" href="<?php echo $url ?>" id = "url<?php echo $i; ?>"><i class="<?php echo $ico; ?>"></i> <?php echo $showdir; ?></a>
							    <?php } ?>
						    </td>
							<!-- 查看HASH例 -->
						    <td id = "info" class = "layui-hide-xs">
							    <!--如果是文件-->
							    <?php if($type == 'file'){ ?>
									<a href="javascript:;" title = "查看文件hash" onclick = "filehash('<?php echo $showdir; ?>','<?php echo $fullpath; ?>')"><i class="fa fa-info-circle" aria-hidden="true"></i></a>
									<a href="javascript:;" onclick = "qrcode('<?php echo $showdir; ?>','<?php echo $url; ?>')" title = "显示二维码"><i class="fa fa-qrcode" aria-hidden="true"></i></a>
							    <?php } ?>
						    </td>
						    <td class = "layui-hide-xs"><?php echo $ctime; ?></td>
						    <td><?php echo $fsize; ?></td>
							<!-- 操作例 -->
						    <td class = "layui-hide-xs">
								<div class = "layui-hide-sm layui-show-md-block">
									<!--复制链接-->
									<?php if($fsize != '-'){ ?>
									<a href="javascript:;" class = "layui-btn layui-btn-xs layui-btn-normal" title = "复制链接" onclick = "copy('<?php echo $url ?>')"><i class="fa fa-copy"></i></a>
									<a download href="<?php echo $url ?>" class = "layui-btn layui-btn-xs layui-btn-normal" title = "点击下载"><i class="fa fa-download"></i></a>
									<?php } ?>
									<!--如果是音乐文件-->
									<?php if( $zdir->music($url) ) { ?>
									<a class = "layui-btn layui-btn-xs layui-btn-normal" title = "点此播放" href="javascript:;" onclick = "music('<?php echo $url ?>')"><i class="fa fa-play-circle"></i></a>
									<?php } ?>
									<!--音乐文件END-->
									<!--如果是markdown文件-->
									<?php if(($suffix == 'md') && ($suffix != null)){ ?>
									<a href="javascript:;" class = "layui-btn layui-btn-xs layui-btn-normal" onclick = "newmd('<?php echo $fullpath; ?>')" title = "点击查看"><i class="fa fa-eye"></i></a> 
									<?php }else if( $zdir->video($url) ){ ?>
									<a class = "layui-btn layui-btn-xs layui-btn-normal" title = "点此播放" href="javascript:;" onclick = "video('<?php echo $url ?>')"><i class="fa fa-play-circle"></i></a>
									<!--文本查看器-->
									<?php }
									else if( $zdir->is_text($url) ){ ?>
									<a class = "layui-btn layui-btn-xs layui-btn-normal" title = "点此查看" href="javascript:;" onclick = "viewtext('<?php echo $fullpath; ?>')"><i class="fa fa-eye"></i></a>
									<?php }else if( $zdir->office($url) ) { ?>
									<!--查看Office-->
									<a class = "layui-btn layui-btn-xs layui-btn-normal" href="javascript:;" title = "点此查看" onclick = "office('<?php echo $url ?>')"><i class="fa fa-eye"></i></a>
									<?php } ?>
								</div>
						    </td>
							<!-- 操作例END -->
					    </tr>
					    <?php } ?>
						
					  </tbody>
					</table>
		    	</div>
		  	</div>
		</div> 
	</div>
<?php
	//载入页脚
	include_once("./template/footer.php");
?>