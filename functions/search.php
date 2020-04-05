<?php
	//$thedir = __DIR__;
	//$thedir = str_replace("\\","/",$thedir);
	
	//$thedir = str_replace("/functions","",$thedir);
	//当前站点的运行目录
	$thedir = $_SERVER['DOCUMENT_ROOT'];
	//echo $_SERVER['DOCUMENT_ROOT'];
	$html = file_get_contents("./functions/caches/indexes.html");

	$s = @$_GET['s'];
	$s = con_coding($s,FALSE);
	$s = trim($s);
	$s = strip_tags($s);
	$s = strtolower($s);

	$text = strip_tags($html);
	$arr = explode("\n",$text);
	$txt = array();
	foreach( $arr as $value )
	{
		if((stripos($value,$s)) || (stripos($value,$s) === 0)){
			//echo $value.'--';
			$value = trim($value);
			//$value = strtolower($value);
			//如果值带有.，说明是一个文件
			if(strpos($value,'.')){
				array_push($txt,$value);
			}
		}
	}
	$s = con_coding($s);
?>
<?php
	//载入页头
	include_once("./template/header.php")
?>
    <!--面包屑导航-->
	<div id="navigation" class = "layui-hide-xs">
		<div class="layui-container">
			<div class="layui-row">
				<div class="layui-col-lg12">
					<h2>"<?php echo $s; ?>"搜索结果</h2>
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
					    <col width="620">
					    <col width="220">
					    <col width="160">
					    <col>
					  </colgroup>
					  <thead>
					    <tr>
					      <th>文件名</th>
					      <th class = "layui-hide-xs">修改时间</th>
					      <th class = "layui-hide-xs">文件大小</th>
					      <th class = "layui-hide-xs">操作</th>
					    </tr> 
					  </thead>
					  <tbody>
						<?php foreach( $txt as $name )
						{
							//获取文件修改时间
						    $ctime = filemtime($thedir.$name);
						    $ctime = date("Y-m-d H:i",$ctime);
						    //获取文件大小
						    $fsize = filesize($thedir.$name);
						    $fsize = ceil ($fsize / 1024);
						    if($fsize >= 1024) {
							    $fsize = $fsize / 1024;
							    $fsize = round($fsize,2).'MB';
						    }
						    else{
							    $fsize = $fsize.'KB';
						    }
						    $name = con_coding($name);
						?>
					    <tr id = "id<?php echo $i; ?>">
						    <td>
							    <!--判断文件是否是图片-->
							    <?php if(($suffix == 'jpg') || ($suffix == 'jpeg') || ($suffix == 'png') || ($suffix == 'gif') || ($suffix == 'bmp')){

							   	?>
							   	<a href="<?php echo $name ?>" id = "url<?php echo $i; ?>" onmouseover = "showimg(<?php echo $i; ?>,'<?php echo $name; ?>')" onmouseout = "hideimg(<?php echo $i; ?>)"><i class="<?php echo $ico; ?>"></i> <?php echo $name; ?></a>
							   	<div class = "showimg" id = "show<?php echo $i; ?>"><img src="" id = "imgid<?php echo $i; ?>"></div>
							   	<?php }else{ ?>
							    <a href="<?php echo $name ?>" id = "url<?php echo $i; ?>"><i class="<?php echo $ico; ?>"></i> <?php echo $name; ?></a>
							    <?php } ?>
						    </td>
						    <td class = "layui-hide-xs"><?php echo $ctime; ?></td>
						    <td class = "layui-hide-xs"><?php echo $fsize; ?></td>
						    <td class = "layui-hide-xs">
							    <?php if($fsize != '-'){ ?>
								<a href="javascript:;" class = "layui-btn layui-btn-xs" onclick = "scopy('<?php echo $name ?>')">复制</a>
							    <?php } ?>
							    <!--如果是管理模式-->
							    <?php if((isset($admin)) && ($fsize != '-')) { ?>
									<a href="javascript:;" class = "layui-btn layui-btn-xs layui-btn-danger" onclick = "delfile(<?php echo $i; ?>,'<?php echo $name; ?>','<?php echo $fullpath; ?>')">删除</a>
							    <?php } ?>
							    <!--如果是markdown文件-->
							    <?php if(($suffix == 'md') && ($suffix != null)){ ?>
								&nbsp;&nbsp;<a href="javascript:;" onclick = "newmd('<?php echo $fullpath; ?>')" title = "点击查看"><i class="fa fa-eye fa-lg"></i></a> 
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
<?php
	//载入页脚
	include_once("./template/footer.php");
?>