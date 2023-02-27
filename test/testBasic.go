package main

import (
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	StudentId int
	ClassId   int
	ClassName string
}

func test(t *testing.T) {

}
