package main

import "testing"

func TestHitungLuasJajarGenjang(t *testing.T) {
	jg := JajarGenjang{Alas: 5, Tinggi: 8}
	expected := 40.0
	actual := jg.HitungLuas()
	if actual != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, actual)
	}
}

func TestHitungLuasPersegi(t *testing.T) {
	p := Persegi{Sisi: 4}
	expected := 16.0
	actual := p.HitungLuas()
	if actual != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, actual)
	}
}

func TestHitungLuasSegitiga(t *testing.T) {
	s := Segitiga{Alas: 6, Tinggi: 9}
	expected := 27.0
	actual := s.HitungLuas()
	if actual != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, actual)
	}
}

func TestHitungLuasLingkaran(t *testing.T) {
	l := Lingkaran{JariJari: 5}
	expected := 78.54 // 5^2 * pi
	actual := l.HitungLuas()
	if actual != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, actual)
	}
}
