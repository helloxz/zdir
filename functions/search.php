<?php
	$html = file_get_contents("./functions/caches/indexes.html");

	$s = @$_GET['s'];
	$s = trim($s);
	$s = strip_tags($s);

	$text = strip_tags($html);
	$arr = explode("\n",$text);
	$txt = array();
	foreach( $arr as $value )
	{
		if((strpos($value,$s)) || (strpos($value,$s) === 0)){
			//echo $value.'--';
			$value = trim($value);
			array_push($txt,$value);
		}
	}
	
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
					    <col width="1000">
					    <col>
					  </colgroup>
					  <thead>
					    <tr>
					      <th>文件名</th>
					      <th class = "layui-hide-xs">操作</th>
					    </tr> 
					  </thead>
					  <tbody>
						<?php foreach( $txt as $name )
						{
						
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