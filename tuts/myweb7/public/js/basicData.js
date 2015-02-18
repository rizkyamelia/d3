var svg = d3.select("body").append("svg").attr({
	width: window.innerWidth,
	height: window.innerHeight
});

var padding	 	= 10;
var radius 		= 4;
var data 		= d3.range(500);

// linear utk membuat kurva linear : d3.scale.linear, 
var cxScale = d3.scale.pow().exponent(2) //y = mx^2 + b
	.domain([0, d3.max(data)])
	.range([padding + radius/2, window.innerWidth - padding - radius/2]);

var cyScale = d3.scale.linear()
	.domain([0, d3.max(data)])
	.range([window.innerHeight - padding - radius/2, padding + radius/2]);

svg.selectAll("circle")
	.data(data)
	.enter()
	.append("circle")
	.attr({
		cx: cxScale,
		cy: cyScale,
		r: radius
	});