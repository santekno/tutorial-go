package main

import (
	"testing"
)

func TestExerciseOne(t *testing.T) {
	test1 := ExerciseOne("madam")
	expect1 := false

	if test1 != expect1 {
		t.Error("Output tidak sesuai!")
	}
}

func BenchmarkExerciseOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExerciseOne("madam")
	}
}
