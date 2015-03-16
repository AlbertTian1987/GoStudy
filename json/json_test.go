package json

import (
	"testing"
)

func Test_decode_user(t *testing.T) {
	t.Log(USER_JSON)
	var user User

	err := user.Decode_from_json(USER_JSON)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(user)
	}
}

func Test_decode_classSlice(t *testing.T) {
	t.Log(Class_JSON)
	var classes ClassSlice
	err := classes.Decode_from_json(Class_JSON)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(classes)
	}
}
func Test_decode_ComplexUser(t *testing.T) {
	t.Log(Complex_Json)
	var cpx ComplexUser
	err := cpx.Decode_from_json(Complex_Json)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(cpx)
	}
}

func Benchmark_Decode_ComplexUser(b *testing.B) {
	b.StopTimer()
	var user User
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		user.Decode_from_json(USER_JSON)
		b.Log(user)
	}
}

func Test_UnKnowJson(t *testing.T) {
	UnkonwJson()
}

func Test_EncodeJson(t *testing.T) {
	EncodeToJson()
}
