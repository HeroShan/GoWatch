var c = document.getElementById("canvas-club");
var ctx = c.getContext("2d");//获取canvas上下文
var w = c.width = window.innerWidth;
var h = c.height = window.innerHeight;//设置canvas宽、高
var clearColor = 'rgba(0, 2, 0, .1)';//画板背景,注意最后的透明度0.1 这是产生叠加效果的基础
 
function random(min, max) {
    return Math.random() * (max - min) + min;
}
 
function Rippling(){}
//涟漪对象 这是涟漪动画的主要部分
Rippling.prototype = {
    init:function(){
        this.x = random(0,w);//涟漪x坐标
        this.y = random(h * .8, h * .9);//涟漪y坐标
        this.w = 2;//椭圆形涟漪宽
        this.h = 1;//椭圆涟漪高
        this.vw = 3;//宽度增长速度
        this.vh = 1;//高度增长速度
        this.a = 1;//透明度
        this.va = .70;//涟漪消失的渐变速度
    },
    draw:function(){
        ctx.beginPath();
        ctx.moveTo(this.x, this.y - this.h / 2);
        //绘制右弧线
        ctx.bezierCurveTo(
            this.x + this.w / 2, this.y - this.h / 2,
            this.x + this.w / 2, this.y + this.h / 2,
            this.x, this.y + this.h / 2);
        //绘制左弧线
        ctx.bezierCurveTo(
            this.x - this.w / 2, this.y + this.h / 2,
            this.x - this.w / 2, this.y - this.h / 2,
            this.x, this.y - this.h / 2);
         
        ctx.strokeStyle = 'hsla(180, 100%, 50%, '+this.a+')';
        ctx.stroke();
        ctx.closePath();
        this.update();//更新坐标
    },
    update:function(){
        if(this.a > .03){
            this.w += this.vw;//宽度增长
            this.h += this.vh;//高度增长
            if(this.w > 100){
                this.a *= this.va;//当宽度超过100，涟漪逐渐变淡消失
                this.vw *= .150;//宽度增长变缓慢
                this.vh *= .150;//高度增长变缓慢
            }
        } else {
            this.init();
        }
 
    }
};
 
function resize(){
    w = c.width = window.innerWidth;
    h = c.height = window.innerHeight;
}
 
//初始化一个涟漪
var r = new Rippling();
r.init();
 
function anim() {
    ctx.fillStyle = clearColor;
    ctx.fillRect(0,0,w,h);
    r.draw();
    requestAnimationFrame(anim);
}
 
window.addEventListener("resize", resize);
//启动动画
anim();