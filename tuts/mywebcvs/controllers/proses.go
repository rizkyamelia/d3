package controllers

import (
	"amelia/gomvc"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ProsesController struct {
	gomvc.Controller
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (c *ProsesController) Project1(rw http.ResponseWriter, req *http.Request) {
	// req.ParseForm()

	val1 := req.FormValue("param01")
	val2 := req.FormValue("param02")
	hasil := 0

	a, _ := strconv.Atoi(val1)
	b, _ := strconv.Atoi(val2)

	var jumlah int
	var x [36]int
	var k int
	k = 35
	a = 0
	b = 4000000
	jumlah = 0

	for i := a; i <= k; i++ {
		if i < 2 {
			x[i] = i + 1
		} else {
			x[i] = x[i-2] + x[i-1]
		}
	}

	for i := a; i <= k; i++ {
		if x[i] <= b && x[i]%2 == 0 {
			jumlah += x[i]
			hasil = jumlah
		}
	}

	fmt.Println(x)
	fmt.Printf("%d", hasil)

	s := "x :"

	//membuat file dat.txt:
	j := " "

	for f := 0; f < k; f++ {
		g := strconv.Itoa(x[f])
		j = j + "	" + g
	}

	z := s + j + "\n\nhasil : " + strconv.Itoa(hasil)

	d1 := []byte(z)

	name := "output/data/" + strconv.Itoa(a) + "_" + strconv.Itoa(b) + "_" + strconv.Itoa(hasil) + "_" + ".txt"

	fmt.Println(name)

	err := ioutil.WriteFile(name, d1, 0644)
	check(err)

	c.ServeJson(hasil)
}
