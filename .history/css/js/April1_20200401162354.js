
        Think.setValue("pid", {$info.pid|default = 0});
        //导航高亮
        highlight_subnav('{:U('ProductSKU/skushow')}');

        $(".attradd").on("click",function(){
        node = "<div class='form-item-attr'><label class='item-label'>属性参数:<span class='check-tips'></span></label><div class='controls'><input type='text' class='text input-large' name='attrval[]' ><span class='attrdel'>-</span></div></div>"
        $(this).parent().parent().after(node)
        })
        $(document).on('click','.attrdel',function(){
         var del=   $(this).parent().parent()
         $(del).remove()
        })

        $(".addsku").on("change",function(){
        	var cateid = $(".addsku option:selected").val()


        	$.ajax({
			   type: "POST",
			   url: "{:U('getskuAttr')}",
			   data: "cateid="+cateid,
			   success: function(msg){
			   	data = JSON.parse(msg)
			   	$(".form-item-attr-append").remove()
			   	$(data).each(function(i){

			   		if(data[i]['type_option']==1){
			   			node = Setext(data[i]['attrbute_name'])
			   			$(".form-item").prev().append(node)
			   		}
			   		if(data[i]['type_option']==2){
			   			node = Setradio(data[i]['id'],data[i]['attrbute_name'],data[i]['value'])
			   			$(".form-item").prev().append(node)
			   		}
			   		if(data[i]['type_option']==3){
			   			node = Setcheckbox(data[i]['attrbute_name'],data[i]['value'])
			   			$(".form-item").prev().append(node)
			   		}
			   	})
			     //console.log(data)
			   }
			});
        })

        //设置描述框
        function Setext(attrname){
        	var str = ""
        	var text = "<div class='form-item-attr-append'><label class='item-label'>"+attrname+"<span class='check-tips'>（属性描述）</span> </label><div class='controls'><input type='text' class='text input-large' name='"+attrname+"'  ></label></div></div>"
        	
        	
        	return text
        }


        //设置单选框		属性id   属性名称  属性值
        function Setradio(id,attrname,attrvalue){
        	var str = ""
        	$(attrvalue).each(function(i){
        		str += "<input type='radio' name='"+attrname+"' value='"+attrvalue[i]['attrbute_value']+"'>"+attrvalue[i]['attrbute_value']
        	})
        	checkbox = "<div class='form-item-attr-append'><label class='item-label'>"+attrname+"<span class='check-tips'></span></label><div class='controls'>"+str+"</div></div>"
        	return checkbox
        }


        //设置多选框		属性名称  属性值
        function Setcheckbox(attrname,attrvalue){
        	var str = ""
        	$(attrvalue).each(function(i){
        		str += "<input type='checkbox' name='"+i+"|"+attrname+"' value='"+attrvalue[i]['attrbute_value']+"'>"+attrvalue[i]['attrbute_value']
        	})
        	checkbox = "<div class='form-item-attr-append'><label class='item-label'>"+attrname+"<span class='check-tips'></span></label><div class='controls'>"+str+"</div></div>"
        	return checkbox
        }

        
