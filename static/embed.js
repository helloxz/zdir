layui.use(['layer','element'], function(){
    var layer = layui.layer;
    var element = layui.element;
})
$(document).ready(function(){
	//隐藏.
	//$("#id1").remove();
	//如果是首页，隐藏..
	uri = window.location.search;
	//console.log(uri);
	//隐藏首页..，现在不需要了
	// if(uri == '') {
	// 	$("#id1").remove();
	// }
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
	dir = window.location.pathname;				//获取目录
	dir = dir.replace("index.php","");
	
	url = protocol + '//' + host + dir + url;

	//console.log(url);

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
			layer.msg('链接已复制！');
        }
    });
}

//复制按钮
function scopy(url){
	url = url.replace("./","");
	//重组url
	protocol = window.location.protocol;		//获取协议
	host = window.location.host;				//获取主机
	dir = window.location.pathname;				//获取目录
	dir = dir.replace("index.php","");
	
	url = protocol + '//' + host + url;

	//console.log(url);

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
	  	area: ['100%', '100%'],
	  	content: url //这里content是一个普通的String
	});
}
//新版markdown查看器
function newmd(url){
	var url = "./?c=viewmd&file=" + url;
	layer.open({
		title:'MarkDown查看器',
	  	type: 2, 
	  	area: ['100%', '100%'],
	  	content: url //这里content是一个普通的String
	});
}
//播放视频文件
function video(url){
	
	var videourl = "./?c=video&url=" + url;
	//layer.msg(videourl);
	layer.open({
		title:false,
	  	type: 2, 
	  	area: ['68%', '78%'],
	  	content: videourl //这里content是一个普通的String
	});
}
//播放音频
function music(url){
	var videourl = "./?c=music&url=" + url;
	//layer.msg(videourl);
	layer.open({
		title:false,
	  	type: 2, 
	  	area: ['60%', '88px'],
	  	content: videourl //这里content是一个普通的String
	});
}

//查看文本文档
function viewtext(url){
	var url = "./?c=viewtext&file=" + url;
	layer.open({
		title:'Zdir文本查看器',
	  	type: 2, 
	  	area: ['100%', '100%'],
	  	content: url //这里content是一个普通的String
	});
}
//预览office
function office(url){
	//文件名
	filename = url;
	var uri = url.replace("./","/");
	uri = encodeURI(uri);
	//获取协议
	var protocol = window.location.protocol + '//';
	//获取主机
	var host = window.location.host;
	//获取页面目录，通常是二级目录的情况下
	var pathname = window.location.pathname;
	pathname = pathname.replace("index.php","");
	url = protocol + host + pathname + uri;
	var apiurl = "https://view.officeapps.live.com/op/view.aspx?src=" + url;
	layer.open({
		title:filename,
	  	type: 2, 
	  	area: ['100%', '100%'],
	  	content: apiurl //这里content是一个普通的String
	});
}
//预览PDF文件
//function viewpdf(filepath){
//	//重组url
//	protocol = window.location.protocol;		//获取协议
//	host = window.location.host;				//获取主机
//	url = protocol + '//' + host + '/' + url;
//	alert(url);
//}

//计算文件hash
function filehash(name,path){
	var file = path;
	
	//alert(file);
	$.post("./?c=hash",{file:file},function(data,status){
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
	qrcodeapi = "https://qr.png.pub/v1/?text=" + url;
	var qrimg = "<center class = 'qrcode'><img src = '" + qrcodeapi + "' /></center>";
	layer.open({
		type: 1,
		area: '230px',
	  	title: name,
	  	content: qrimg
	});   
}

//删除文件
function delfile(id,filename,filepath){
	id = "id" + id;
	layer.prompt({
		formType: 1,
		title: '请输入密码删除 - ' + filename
	}, function(value, index, elem){
		$.post('./functions/delfile.php',{filepath:filepath,password:value},function(data,status){
			var redata = eval('(' + data + ')');
			if(redata.code == 1){
				$("#" + id).remove();
				layer.msg(redata.msg + ' ' + filename);
			}
			else{
				layer.msg(redata.msg);
			}
		});
		layer.close(index);
	});
}