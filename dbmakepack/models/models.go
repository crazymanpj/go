package models

import(
	"time"
)

type Product struct{
	id            int
	name          string
	description   string
	addtime       time.Time
	status        int
}

func init(){
}