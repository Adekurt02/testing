package main

import (
	"fmt"
	"math"
)

// JajarGenjang struct untuk menyimpan sisi alas dan tinggi jajar genjang
type JajarGenjang struct {
	Alas, Tinggi float64
}

// Persegi struct untuk menyimpan panjang sisi persegi
type Persegi struct {
	Sisi float64
}

// Segitiga struct untuk menyimpan panjang alas dan tinggi segitiga
type Segitiga struct {
	Alas, Tinggi float64
}

// Lingkaran struct untuk menyimpan jari-jari lingkaran
type Lingkaran struct {
	JariJari float64
}

// HitungLuas mengembalikan luas jajar genjang
func (jg JajarGenjang) HitungLuas() float64 {
	return jg.Alas * jg.Tinggi
}

// HitungLuas mengembalikan luas persegi
func (p Persegi) HitungLuas() float64 {
	return math.Pow(p.Sisi, 2)
}

// HitungLuas mengembalikan luas segitiga
func (s Segitiga) HitungLuas() float64 {
	return 0.5 * s.Alas * s.Tinggi
}

// HitungLuas mengembalikan luas lingkaran
func (l Lingkaran) HitungLuas() float64 {
	return math.Pi * math.Pow(l.JariJari, 2)
}

func main() {
	jg := JajarGenjang{Alas: 5, Tinggi: 8}
	fmt.Printf("Luas jajar genjang: %.2f\n", jg.HitungLuas())

	p := Persegi{Sisi: 4}
	fmt.Printf("Luas persegi: %.2f\n", p.HitungLuas())

	s := Segitiga{Alas: 6, Tinggi: 9}
	fmt.Printf("Luas segitiga: %.2f\n", s.HitungLuas())

	l := Lingkaran{JariJari: 5}
	fmt.Printf("Luas lingkaran: %.2f\n", l.HitungLuas())
}
