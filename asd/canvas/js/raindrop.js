var c = document.getElementById("canvas-club");
var ctx = c.getContext("2d");//获取canvas上下文
var w = c.width = window.innerWidth;
var h = c.height = window.innerHeight;//设置canvas宽、高
var clearColor = 'rgba(12, 0, 1, .3)';//画板背景,注意最后的透明度0.1 这是产生叠加效果的基础
 console.log(clearColor)
function random(min, max) {
    return Math.random() * (max - min) + min;
}
 
function RainDrop(){}
//雨滴对象 这是绘制雨滴动画的关键
RainDrop.prototype = {
    init:function(){
        this.x =  random(0, w);//雨滴的位置x
        this.y = 0;//雨滴的位置y
        this.color = 'hsl(131, 10%, 90%)';//雨滴颜色 长方形的填充色
        this.vy = random(2, 7);//雨滴下落速度
        this.hit = random(h * .6, h * .9);//下落的最大值
        this.size = 2;//长方形宽度
    },
    draw:function(){
        if (this.y < this.hit) {
            ctx.fillStyle = this.color;
            ctx.fillRect(this.x, this.y, this.size, this.size * 5);//绘制长方形，通过多次叠加长方形，形成雨滴下落效果
        }
        this.update();//更新位置
    },
    update:function(){
        if(this.y < this.hit){
            this.y += this.vy;//未达到底部，增加雨滴y坐标
        }else{
            this.init();
        }
    }
};
 
function resize(){
    w = c.width = window.innerWidth;
    h = c.height = window.innerHeight;
}
 
//初始化一个雨滴
var r = new RainDrop();
r.init();
 
function canim() {
    ctx.fillStyle = clearColor;//每一帧都填充背景色
    ctx.fillRect(0,0,w,h);//填充背景色，注意不要用clearRect，否则会清空前面的雨滴，导致不能产生叠加的效果
    r.draw();//绘制雨滴
    anim()
    requestAnimationFrame(canim);//控制动画帧
}
 
window.addEventListener("resize", resize);
//启动动画
canim();