$(document).ready(function(){
	msg();
});

//清空内容
function clear_msg(){
	$("#msg-content").empty();
}
function msg(){
	//生成0-5的随机数
	var num = parseInt(Math.random()*(4-0+1)+0,10);;
	
	var content = new Array();
	content[0] = 'Zdir v1.2已发布，建议尽快更新：<a href = "https://www.xiaoz.me/archives/12927" target = "_blank">https://www.xiaoz.me/archives/12927</a>';
	content[1] = 'UltraVPS(EU)美国洛杉矶机房，2核2G，3.8€欧/月：<a href = "https://www.xiaoz.me/archives/9633" target = "_blank">https://www.xiaoz.me/archives/9633</a>';
	content[2] = 'Vultr新用户注册送50$,16个机房可选，按时计费：<a href = "https://www.xiaoz.me/archives/12774" target = "_blank">https://www.xiaoz.me/archives/12774</a>';
	content[3] = 'CloudCone最新促销，低至$2.8/月：<a href = "https://www.xiaoz.me/archives/11183" target = "_blank">https://www.xiaoz.me/archives/11183</a>';
	content[4] = '老薛主机75折优惠码：<b>xiaoz25</b>，<a href = "https://dwz.ovh/laoxue" target = "_blank">点此使用</a>';
	$("#msg-content").append(content[num]);
	window.setTimeout("clear_msg()",9800);
	window.setTimeout("msg()",10000);
}