$(function(){
	if (is_mobile()) {
		//$('body').removeClass('left-side-collapsed');
		//$('.left-side').hide();
	}
	
	//左侧菜单显示
	/*
	var str = $('#permissionModel').val();
	if (str != "") {
		var str2 = $('#permissionModelc').val();
			
		var strArr = str.split(',');
		var strtmp;
		var str2Arr = str2.split(',');
		var str2tmp;
		var html = '';
		var html2= '';
	
		for(var i=0;i<strArr.length;i++) {
			strArrc =strArr[i].split('||');
			strtmp = strArrc[0].split('-');		
			var m = 0;	
			for(var j=0;j<str2Arr.length;j++) {
				str2Arrc =str2Arr[j].split('||');
				str2tmp = str2Arrc[0].split('-');			
				if (str2tmp[1]==strtmp[1]) {		
					m++;
					if (m == 1) { html2 += '<ul class="sub-menu-list">'; }			
					html2 += '<li><a href="/'+(str2Arrc[1].replace(/-/, '/'))+'"> '+str2tmp[0]+'</a></li>';
				}			
			}
			if (m >0) {
				html2 += '</ul>';
				html += '<li class="menu-list">';			
			} else {
				html += '<li>';
			}
			html += '<a href="'+(strArrc[1] == '#' ? '#' : '/'+strArrc[1].replace(/-/, '/'))+'"><i class="fa fa-'+strtmp[2]+'"></i> <span>'+strtmp[0]+'</span></a>';
			html+=html2
			html2 = '';
			html += '</li>';		
		}	
		//$('.js-left-nav').append(html);
	}
	
	
	//左边菜单加选中状态
	var pre = location.pathname;
	var qstr = pre.split('/');       
    if (qstr) {            
        var lefthref = '/'+qstr[1]+'/'+qstr[2];
        $('.sub-menu-list a').filter(function(){           
        	return $(this).attr('href') == lefthref;
        }).parent().addClass('active').parents('.menu-list').addClass('nav-active');
    };
	*/
	$('.js_checkboxAll').on('click', function(){
		var that = $(this);
		var chk = that.parent().prev('table').find('input[type="checkbox"]');
		if (that.is(':checked')) {
			chk.prop('checked', true);
		} else {
			chk.prop('checked', false);
		}
	});
	
	//顶部消息点击更新状态
	$('.js-msg-status').on('click', function(){
		var that = $(this);
		var id = that.attr('data-id');
		$.post('/system/message/ajax/status', {id:id},function(data){
						
		},'json');
	});
	
    $('#login-form').validate({
        ignore:'',
        rules : {
            username:{ required: true},
            password:{required: true}
        },
        messages : {
            username : {required: '请填写用户名'},
            password : {required: '请填写密码'}
        },
        submitHandler:function(form) {
            var url = '/login';
            $(form).ajaxSubmit({
                url:url,
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       setTimeout(function(){window.location.href="/"}, 1000);
                    } else {
                       setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 2000);
                    }
                }
            });
        }
    });
    
    $('.js-user-single').on('click', function(){
    	var that = $(this);
    	var status = that.attr('data-status')
    	var id = that.attr('data-id');
		$.post('/system/user/ajax/status', { status: status, id: id },function(data){
			dialogInfo(data.message)
			if (data.code) {
				that.attr('data-status', status == 2 ? 1 : 2).text(status == 2 ? '激活' : '禁用').parents('td').prev('td').text(status == 2 ? '禁用' : '激活');
			} else {
				
			}
			setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
		},'json');
    });    
   
    $('#userprofile-form').validate({
        ignore:'',        
		rules : {
			username:{required: true},
			role:{required: true},
        },
        messages : {
			username:{required: '请填写用户名'},
			role:{required: '请选择角色'}, 
        },
        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
						if (data.type) {
                       		setTimeout(function(){window.location.href="/system/user/show/"+data.id}, 1000);
						} else {							
							setTimeout(function(){window.location.href="/system/user/manage"}, 1000);
						}
                    } else {
                        setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });
	
	$('#userprofilepwd-form').validate({
        ignore:'',        
		 rules : {
            oldpwd :{ required:true},
            newpwd :{ required:true,rangelength:[6,15]},
            confpwd :{ required:true,equalTo:'#newpwd'}
        },
        messages : {
            oldpwd :{ required:"当前密码不能为空"},
            newpwd :{ required:"新密码不能为空",rangelength:  $.validator.format("密码长度{0}-{1}位")},
            confpwd :{ required:"确认密码不能为空",equalTo:'两次输入密码不一致'}
        },
        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {						
                       	setTimeout(function(){window.location.href="/system/user/show/"+data.id}, 1000);						
                    } else {
                        setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }					
                }
            });
        }
    });
	
	
	
	$('.js-search-username').on('keydown.autocomplete', function(){
        var ob = $(this);
		var obid = ob.attr('id');
		//alert(obid)
		var hideid;
		if (obid == 'accept-username') {
			hideid = '#acceptid';
		} else if (obid == 'cc-username') {
			hideid = '#ccid';	
		} else {
			hideid = '#userid';
		}
        $(this).autocomplete({
            source: "/system/user/ajax/search",
            minLength: 1,
            autoFocus: true,
            select: function(e, ui) {              
	            ob.val(ui.item.label).blur();
				$(hideid).val(ui.item.value);               
                return false;
            }
        })
    });
	
	
	$('#cc-username').on('click', function(){
		$('#acceptModal').modal('show');
	});
	$('.js-dialog-taskcc').on('click', function(){
		var that = $(this);		
		var cc = $('#cc-username');	
		var ccid = $('#ccid');
		cc.val('');
		ccid.val('');
		var ccnamestr,ccidstr
		$('.modal-body input[type="checkbox"]').each(function(i){
			if ($(this).is(':checked')) {
				ccnamestr = $(this).attr('data-name')+','+cc.val();		
				cc.val($.unique(ccnamestr.split(',').sort()).join(',').replace(/,$/gi,''));
				
				ccidstr = $(this).attr('data-value')+','+ccid.val();		
				ccid.val($.unique(ccidstr.split(',').sort()).join(',').replace(/,$/gi,''));
			}
		 });
		
		$('#acceptModal').modal('hide');
	});
	
	
	
	//多图上传	
	$('#uploadMulti-form').validate({        
        ignore:'',
        rules : {
            uploadFiles:{ required : true}
        },
        messages : {
            uploadFiles:{required : '请选择上传图片，可以多选'}
        },
        submitHandler:function(form) {
			var id = $('input[name="id"]').val();
			var url = '/uploadmulti';
            $(form).ajaxSubmit({
                url:url,
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code == 1) {						
						setTimeout(function(){window.location.href='/album/manage?filter=me'}, 1000);
                    } else {
                        setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });
	
	
	
	$('#permission-btn-new').on('click', function(){
		var str = '', model = '', modelc = '';
		
		$('input[name="permission[]"]:checked').each(function(i){
			str += $(this).val() + ',';
			if ($(this).parents('li').parents('li').attr('data-pmodel')) {
				model += $(this).parents('li').parents('li').attr('data-pmodel')+',';
				
				if ($(this).parents('div').attr('data-cmodel')) {
					modelc += $(this).parents('div').attr('data-cmodel')+',';
				}
			}
			
		});
		if (str == '') {
			return false;
		}
		
		model = uniqueString(model).toString();	
		model = model.substring(0, model.length-1);
		
		modelc = uniqueString(modelc).toString();		
		modelc = modelc.substring(0, modelc.length-1);
		
		str = str.substring(0, str.length-1)
		var groupid = $('#groupid').val();
		var url = '/system/role/permission/'+groupid;
		$.post(url, { groupid: groupid, permission: str, model:model, modelc:modelc },function(data){
			dialogInfo(data.message)
			if (data.code) {
				
			} else {
				
			}
			setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
		},'json');
	});
	
	$('#permission-btn').on('click', function(){
		var str = '', model = '', modelc = '';
		
		$('input[name="permission[]"]:checked').each(function(i){
			str += $(this).val() + ',';
			if ($(this).parents('li').parents('li').attr('data-pmodel')) {
				model += $(this).parents('li').parents('li').attr('data-pmodel')+',';
				
				if ($(this).parents('div').attr('data-cmodel')) {
					modelc += $(this).parents('div').attr('data-cmodel')+',';
				}
			}
			
		});
		if (str == '') {
			return false;
		}
		
		model = uniqueString(model).toString();	
		model = model.substring(0, model.length-1);
		
		modelc = uniqueString(modelc).toString();		
		modelc = modelc.substring(0, modelc.length-1);
		
		str = str.substring(0, str.length-1)
		var userid = $('#userid').val();
		var url = '/user/permission/'+userid;
		$.post(url, { userid: userid, permission: str, model:model, modelc:modelc },function(data){
			dialogInfo(data.message)
			if (data.code) {
				
			} else {
				
			}
			setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
		},'json');
	});
	
	
	
	
	//用户管理
	$('.js-user-delete').on('click', function(){
		var that = $(this);
		var id = that.attr('data-id');

		layer.confirm('您确定要删除吗？', {
			btn: ['确定','取消'] //按钮
			,title:"提示"
		}, function(index){
			layer.close(index);
			
			$.post('/system/user/ajax/delete', {ids:id},function(data){
				dialogInfo(data.message)
				if (data.code) {
					setTimeout(function(){ window.location.reload() }, 1000);
				} else {
					setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
				}
				
			},'json');

		}, function(){

		});
	});

	//用户密码重置
	$('.js-passwd-reset').on('click', function(){
		var that = $(this);
		var id = that.attr('data-id');

		layer.confirm('您确定要重置密码吗？', {
			btn: ['确定','取消'] //按钮
			,title:"提示"
		}, function(index){
			layer.close(index);
			
			$.post('/system/user/ajax/reset_passwd', {id:id},function(data){
				dialogInfo(data.message)
				if (data.code) {
					setTimeout(function(){ window.location.reload() }, 10000);
				} else {
					setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
				}
				
			},'json');

		}, function(){

		});
	});

	//角色权限
	$('#group-form').validate({
        ignore:'',		    
		rules : {
			'name':{required:true},
			'summary':{required:true}
        },
        messages : {		
			'name':{required:'请填写角色名称'},
			'summary':{required:'请填写角色描述'}
        },
        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
						setTimeout(function(){ window.location.href='/system/role/manage'; }, 1000);
                    } else {
                       setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }															
                }
            });
        }
    });
	
	$('.js-group-delete').on('click', function(){
		var that = $(this);
		var id = that.attr('data-id');

		layer.confirm('您确定要删除吗？', {
			btn: ['确定','取消'] //按钮
			,title:"提示"
		}, function(index){
			layer.close(index);
			
			$.post('/system/role/ajax/delete', {ids:id},function(data){
				dialogInfo(data.message)
				if (data.code) {
					setTimeout(function(){ window.location.reload() }, 1000);
				} else {
					setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
				}
				
			},'json');

		}, function(){

		});
		
	});
	
	$('#permission-form').validate({
        ignore:'',		    
		rules : {
			'name':{required:true},
			'ename':{required:true}
        },
        messages : {		
			'name':{required:'请填写名称'},
			'ename':{required:'请填写英文名称'}
        },
        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
						setTimeout(function(){ window.location.href='/system/permission/manage'; }, 1000);
                    } else {
                       setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }															
                }
            });
        }
    });
	$('.js-permission-delete').on('click', function(){
		var that = $(this);
		var id = that.attr('data-id');

		layer.confirm('您确定要删除吗？', {
			btn: ['确定','取消'] //按钮
			,title:"提示"
		}, function(index){
			layer.close(index);
			
			$.post('/system/permission/ajax/delete', {ids:id},function(data){
				dialogInfo(data.message)
				if (data.code) {
					setTimeout(function(){ window.location.reload() }, 1000);
				} else {
					setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
				}
				
			},'json');

		}, function(){

		});
	});
	
	$('#group-user-form').validate({
        ignore:'',		    
		rules : {
			username:{required: true}
        },
        messages : {
			username:{required: '请填写姓名'} 
        },
        submitHandler:function(form) {
            $(form).ajaxSubmit({
                type:'POST',
                dataType:'json',
                success:function(data) {
                    dialogInfo(data.message)
                    if (data.code) {
                       setTimeout(function(){window.location.href="/group/user/"+$('#groupid').val()}, 1000);
                    } else {
                       setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
                    }
                }
            });
        }
    });
	$('.js-group-user-single').on('click', function(){
    	var that = $(this);
		var testid = that.attr('data-id');	
		$.post('/system/group/user/ajax/delete', {id:testid},function(data){
			dialogInfo(data.message)
			if (data.code) {
				that.parents('tr').remove();
			} else {
				
			}
			setTimeout(function(){ $('#dialogInfo').modal('hide'); }, 1000);
		},'json');
    });
	
});

function dialogInfo(msg) {
	$('#dialogInfo').remove();
	var html = '';
	html = '<div class="modal fade" id="dialogInfo" tabindex="-1" role="dialog" aria-labelledby="dialogInfoTitle">';
  	html += '<div class="modal-dialog" role="document">';
    html += '<div class="modal-content">';
    html += '<div class="modal-header">';
    html += '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>';
    html += '<h4 class="modal-title" id="dialogInfoTitle">OPMS提示</h4>';
    html += ' </div>';
    html += '<div class="modal-body">';
    html += '<p>'+msg+'</p>';
    html += '</div>';
    //html += '<div class="modal-footer">';
    //html += ' <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>';
    //html += ' <button type="button" class="btn btn-primary">Send message</button>';
    //html += '</div>';
    html += '</div>';
  	html += '</div>';
	html += '</div>';
	$('body').append(html);
	$('#dialogInfo').modal('show')  
}

function dialogAlbum(id, title, summary, status) {
	$('#dialogAlbum').remove();
	var html = '';
	html += '<div class="modal fade in" id="dialogAlbum" tabindex="-1" role="dialog" aria-labelledby="dialogAlbumTitle">';
    html += '  <div class="modal-dialog" role="document">';
	html += '   <form id="album-form">';
    html += '    <div class="modal-content">';
    html += '      <div class="modal-header">';
    html += '        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>';
    html += '        <h4 class="modal-title" id="dialogAlbumTitle">编辑'+title+'</h4>';
    html += '      </div>';
    html += '      <div class="modal-body">';
    html += '          <div class="form-group">';
    html += '            <label for="recipient-name" class="control-label">标题:</label>';
    html += '           <input type="text" class="form-control" name="title" value="'+title+'">';
    html += '          </div>';
	
    html += '          <div class="form-group">';
    html += '            <label for="message-text" class="control-label">说明:</label>';
    html += '            <textarea class="form-control" name="summary">'+summary+'</textarea>';
    html += '          </div>';
	
	
	html += '<div class="form-group">';
	html += '<label class="radio-inline">';
  	html += '<input type="radio" name="status" value="0" '+(status == 0 ? 'checked' : '')+'> 屏蔽';
	html += '</label>';
	html += '<label class="radio-inline">';
	html += '<input type="radio" name="status" value="1" '+(status == 1 ? 'checked' : '')+'> 正常';
	html += '</label>';
	html += '          </div>';
	
	
	
    html += '      </div>';
    html += '      <div class="modal-footer"><input type="hidden" name="id" value="'+id+'">';
    html += '        <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>';
    html += '        <button type="button" class="btn btn-primary js-album-submit">提交</button>';
    html += '      </div>';
    html += '    </div>';
	html += '   </form>';
    html += '  </div>';
    html += '</div>';
	$('body').append(html);
	$('#dialogAlbum').modal('show');
}

function workDay(started, ended) {
	var beginDate = new Date(started.replace(/-/g, "/"));  
	//结束日期  
	var endDate = new Date(ended.replace(/-/g, "/"));  
	//日期差值,即包含周六日、以天为单位的工时，86400000=1000*60*60*24.  
	var workDayVal = (endDate - beginDate)/86400000 + 1;  
	//工时的余数  
	var remainder = workDayVal % 7; 
	//工时向下取整的除数  
	var divisor = Math.floor(workDayVal / 7);  
	var weekendDay = 2 * divisor;  
	  
	//起始日期的星期，星期取值有（1,2,3,4,5,6,0）  
	var nextDay = beginDate.getDay();  
	//从起始日期的星期开始 遍历remainder天  
	for(var tempDay = remainder; tempDay>=1; tempDay--) {  
	    //第一天不用加1  
	    if(tempDay == remainder) {  
	        nextDay = nextDay + 0;  
	    } else if(tempDay != remainder) {  
	        nextDay = nextDay + 1;  
	    }  
	    //周日，变更为0  
	    if(nextDay == 7) {  
	        nextDay = 0;  
	    }  
	  
	    //周六日  
	    if(nextDay == 0 || nextDay == 6) {  
	        weekendDay = weekendDay + 1;  
	    }  
	}  
	//实际工时（天） = 起止日期差 - 周六日数目。  
	workDayVal = workDayVal - weekendDay; 
	return  workDayVal;
}

function is_mobile() { 
	var regex_match = /(nokia|iphone|android|motorola|^mot-|softbank|foma|docomo|kddi|up.browser|up.link|htc|dopod|blazer|netfront|helio|hosin|huawei|novarra|CoolPad|webos|techfaith|palmsource|blackberry|alcatel|amoi|ktouch|nexian|samsung|^sam-|s[cg]h|^lge|ericsson|philips|sagem|wellcom|bunjalloo|maui|symbian|smartphone|midp|wap|phone|windows ce|iemobile|^spice|^bird|^zte-|longcos|pantech|gionee|^sie-|portalmmm|jigs browser|hiptop|^benq|haier|^lct|operas*mobi|opera*mini|320x320|240x320|176x220)/i; 
	var u = navigator.userAgent; 
	if (null == u) { 
		return true; 
	} 
	var result = regex_match.exec(u);
	if (null == result) { 
		return false 
	} else { 
		return true 
	} 
}

function myPrint(obj){  
    var newWindow=window.open("打印窗口","_blank");
    var docStr = obj.innerHTML;  
    newWindow.document.write(docStr);  
    newWindow.document.close();  
    newWindow.print();  
    newWindow.close();  
}

function uniqueString(str) {
	var strArr = str.split(',');  
	for(var i=0;strArr.length-1>i;i++){		
		for(var j=i+1;j<strArr.length;j++){		
			if(strArr[j]==strArr[i]){
				strArr.splice(j,1);
				j--; 
			} 
		} 
	}
	return strArr;
}

function addZero(v){
	return parseInt(v)<10 ? '0'+v : v;
}
	
function fixDate(time) {
	if (!document.all) {
		return new Date(time);
	}
	var arr = time.split(time.match(/\D+/g)[0]);
	return new Date(arr[0], arr[1] - 1, arr[2]);
}