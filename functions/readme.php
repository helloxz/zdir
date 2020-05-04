<?php
	include_once("./zdir.class.php");
	include_once("./Parsedown.php");
	$Parsedown = new Parsedown();
	@$file = $_GET['file'];

	$content = $zdir->viewmd($file);
	$content = $Parsedown->text($content);
	$content = str_replace('[x]','<input type="checkbox" checked>',$content);
	$content = str_replace('[ ]','<input type="checkbox">',$content);
?>
<!DOCTYPE html>
<html lang="zh-cmn-Hans" xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta charset="utf-8" />
	<title>MDtoHTML</title>
	<meta name="generator" content="EverEdit" />
	<meta name="author" content="" />
	<meta name="keywords" content="MDtoHTML,markdown" />
	<meta name="description" content="MDtoHTML快速将Markdown文件转换为HTML" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel='stylesheet' href='https://libs.xiaoz.top/highlight.js/9.12.0/styles/github.min.css'>
	<link rel="stylesheet" href="https://libs.xiaoz.top/layui-v2.2.5/layui/css/layui.css">
	<link rel="stylesheet" href="../static/md.css">
</head>
<body style = "background-color:#FFFFFF;">
	<div class="md-html" style = "margin-right:1em;margin-top:-1em;box-shadow:none;">
		<?php echo $content; ?>
	</div>
	<div class="footer">
		<hr />
		<p>&copy;2020 本文档使用 <a href="https://markdown.win/" target = "_blank">MDtoHTML</a> 构建 | The author <a href="https://www.xiaoz.me/" target = "_blank" title = "小z博客">xiaoz</a></p>
	</div>
	<script src = 'https://libs.xiaoz.top/highlight.js/9.12.0/highlight.min.js'></script>
	<script src = 'https://libs.xiaoz.top/jquery/2.2.4/jquery.min.js'></script>
	<script>hljs.initHighlightingOnLoad();</script>
	<script>
		$(document).ready(function(){
			//获取第一个H1作为标题
			var h1 = $("h1").text();
			//如果获取到了H1
			if(h1) {
				$("title").text(h1 + " - MDtoHTML");
			}
		});
	</script>
</body>
</html>
