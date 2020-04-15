package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20191121_163140 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20191121_163140{}
	m.Created = "20191121_163140"

	migration.Register("User_20191121_163140", m)
}

// Run the migrations
func (m *User_20191121_163140) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *User_20191121_163140) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
