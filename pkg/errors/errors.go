package errors

import (
	"fmt"
	"log"
)

type Mssage string

func (e Mssage) Error(){
	panic(e)
}

func (e Mssage) Print(){
	fmt.Println(e)
}

func (e Mssage) Log(){
	log.Fatal(e)
}