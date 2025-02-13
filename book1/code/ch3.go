package main

// 3.1.1
//func main() {
//	var v1, v2 int
//	v3, v4 := f1(v1, v2)
//	println(&v1, &v2, &v3, &v4)
//	f2(v3)
//
//}
//
//func f1(a1, a2 int) (r1, r2 int) {
//	var l1, l2 int
//	println(&r2, &r1, &a2, &a1, &l1, &l2)
//	return
//}

//func f2(a1 int) {
//	println(&a1)
//}
// 3.1.3

type S struct {
	a int8
	b int64
	c int32
	d int16
}

func f1(s S) (r S) {
	println(&r.a, &r.b, &r.c, &r.d, &s.a, &s.b, &s.c, &s.d)
	return s
}

func f2(aa int8, bb int64, cc int32, dd int16) (ra int8, rb int64, rc int32, rd int16) {
	println(&ra, &rb, &rc, &rd, &aa, &bb, &cc, &dd)
	return
}

func main() {
	f1(S{})
	f2(0, 0, 0, 0)
}
