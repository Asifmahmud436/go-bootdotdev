package main

type rect struct{
	width int;
	height int;
}

func (r rect) area() int{
	return  r.width * r.height;
}

var r = rect{
	width: 20,
	height: 40,
}

