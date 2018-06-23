layui.use(['layer','element'], function(){
    var layer = layui.layer;
    var element = layui.element;
})
$(document).ready(function(){
	//隐藏.
	$("#id1").remove();
	//如果是首页，隐藏..
	uri = window.location.search;
	if(uri == '') {
		$("#id2").remove();
	}
});

protocol = window.location.protocol;		//获取协议
host = window.location.host;				//获取主机
pageurl = protocol + '//' + host + '/';

//复制按钮
function copy(url){
	url = url.replace("./","");
	//重组url
	protocol = window.location.protocol;		//获取协议
	host = window.location.host;				//获取主机
	url = protocol + '//' + host + '/' + url;

	//获取文件后缀
	var index1=url.lastIndexOf(".");
	var index2=url.length;
	var suffix=url.substring(index1+1,index2);

	switch(suffix){
		case 'js':
			url = "<script src = '" + url + "'></script>";
			break;
		case 'css':
			url = "<link rel='stylesheet' href='" + url + "'>";
		default:
			//如果是图片
			if((suffix == 'jpg') || (suffix == 'jpeg') || (suffix == 'gif') || (suffix == 'bmp') || (suffix == 'png')){
				url = "<img src = '" + url + "' />";
			}
			else{
				url = url;
			}
		break;
	}
	
	
	var copy = new clipBoard(document.getElementById('list'), {
        beforeCopy: function() {
            
        },
        copy: function() {
            return url;
        },
        afterCopy: function() {
			layer.msg('复制成功！');
        }
    });
}

//查看markdown文件
function viewmd(url){
	url = url.replace("./","");
	//重组url
	protocol = window.location.protocol;		//获取协议
	host = window.location.host;				//获取主机
	url = protocol + '//' + host + '/' + url;
	url = 'https://markdown.win/api.php?url=' + url;
	layer.open({
	  	type: 2, 
	  	area: ['80%', '80%'],
	  	content: url //这里content是一个普通的String
	});
}

//计算文件hash
function filehash(name,path){
	var file = path;
	
	//alert(file);
	$.post("./functions/hash.php",{file:file},function(data,status){
		var fileinfo = eval('(' + data + ')');
		if(fileinfo.code == 1){
			layer.open({
  				title:name,
  				area: ['400px', 'auto'],
			  	content: '<b>md5: </b>' + fileinfo.md5 + '<br /><b>sha1: </b>' + fileinfo.sha1
			});  
		}
		else{
			layer.msg(fileinfo.msg); 
		}
	});
}

//显示图片
function showimg(id,url){
	var imgid = "imgid" + id;
	//获取上一个id
	var upid = id - 1;
	//获取下一个ID
	var dnid = id + 1;
	//隐藏图片
	$("#show" + upid).hide();
	$("#show" + dnid).hide();
	//显示图片
	$("#" + imgid).attr('src',url);
	$("#show" + id).show(); 
}
//隐藏图片
function hideimg(id){
	var upid = id - 1;
	//获取下一个ID
	var dnid = id + 1;
	$("#show" + id).hide();
	$("#show" + upid).hide();
	$("#show" + dnid).hide();
}

//显示二维码
function qrcode(name,url){
	url = url.replace("./","");
	//重组url
	protocol = window.location.protocol;		//获取协议
	host = window.location.host;				//获取主机
	url = protocol + '//' + host + '/' + url;

	//二维码接口
	qrcodeapi = "https://sapi.k780.com/?app=qr.get&level=L&size=5&data=" + url;
	var qrimg = "<center><img src = '" + qrcodeapi + "' /></center>";
	layer.open({
		type: 1,
		area: '200px',
	  	title: name,
	  	content: qrimg
	});   
	
}