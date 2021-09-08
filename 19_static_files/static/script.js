var canvas;
var ctx;
var running = 0;
var interval;

window.onload = function() {
	canvas = document.getElementById("mycanvas");
	ctx = canvas.getContext("2d");
}

createRect = () => {
	let x = 0;
	let y = 0;
	let side = 0;
	let color = "#";
	let colorpicker = "ABCDEF0123456789";
	x = Math.floor((Math.random()*400)+1);
	y = Math.floor((Math.random()*400)+1);
	side = Math.floor((Math.random()*80)+20);
	for(let i=0;i<6;i++) {
		let temp = Math.floor(Math.random()*16);
		color = color+colorpicker[temp];
	}
	ctx.fillStyle = color;
	ctx.fillRect(x,y,side,side);
}

clearCanvas = () => {
	ctx.clearRect(0,0,500,500);
}

startCanvas = () => {
	if(running) {
		running = 0;
		clearInterval(interval);
	} else {
		running = 1;
		interval = setInterval(createRect,200);
	}
}