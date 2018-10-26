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
	<link rel='stylesheet' href='./static/style.css?v=1.4'>
	<link rel="stylesheet" href="./static/font-awesome/css/font-awesome.min.css">
</head>
<body>
	<!--顶部导航栏-->
	<div class = "header">
        <div class = "layui-container">
            <div class = "layui-row">
                <div class = "layui-col-lg12">
	                <div class="logo">
		                <h1><a href="./"><i class="fa fa-sitemap" aria-hidden="true"></i> Zdir</a></h1>
	                </div>
                    <div class = "layui-hide-xs">
                        <ul class="layui-nav menu" lay-filter="">
                            <li class="layui-nav-item"><a href="./"><i class="fa fa-home" aria-hidden="true"></i> 首页</a></li>
                            <li class="layui-nav-item"><a href="./cache.php"><i class="fa fa-file-text" aria-hidden="true"></i> 文件索引</a></li>
                            <!--<li class="layui-nav-item"><?php $zdir->https(); ?></li>-->
                            <li class="layui-nav-item"><a href="https://github.com/helloxz/zdir" target = "_blank" rel = "nofollow"><i class="fa fa-code" aria-hidden="true"></i> 源码</a></li>
                            <li class="layui-nav-item"><a href="https://doc.xiaoz.me/#/zdir/" target = "_blank" rel = "nofollow"><i class="fa fa-file-text-o" aria-hidden="true"></i> 帮助文档</a></li>
                            <li class="layui-nav-item"><a href="https://www.xiaoz.me/api/pay/?name=%E8%8E%B7%E5%8F%96Zdir%E6%8D%90%E8%B5%A0%E7%89%88%20-%20%E6%82%A8%E7%9A%84%E5%9F%9F%E5%90%8D&price=20" target = "_blank" rel = "nofollow"><i class="fa fa-rmb" aria-hidden="true"></i> 捐赠</a></li>
                            <li class="layui-nav-item"><a href="https://www.xiaoz.me/archives/10465" target = "_blank"><i class="layui-icon">&#xe60b;</i> 关于</a></li>
                        </ul>
                    </div>
                    <!--<div class = "layui-hide-lg">
                        <ul class="layui-nav menu" lay-filter="">
                            <li class="layui-nav-item"><a href="./"><i class="fa fa-home fa-lg"></i> Zdir</a></li>
                        </ul>
                    </div>-->
                </div>
            </div>
        </div>
    </div>
    <!-- 顶部导航栏END -->