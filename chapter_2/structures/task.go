package main

type SimpleStruct struct {
	On          bool
	Ammo, Power *int
}

func Shoot(s *SimpleStruct) bool {
	if s.On == false || nil == s.Ammo || *s.Ammo <= 0 {
		return false
	}
	*s.Ammo--
	return true
}

func RideBike(s *SimpleStruct) bool {
	if s.On == false || nil == s.Power || *s.Power <= 0 {
		return false
	}
	*s.Power--
	return true
}
