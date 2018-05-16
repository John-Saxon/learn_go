package main

import "fmt"

type Books struct{
	title string
	author string
	subject string 
	book_id int 
}

func main(){
	var Book1 Books 
	var Book2 Books 
	
	Book1.title="Go 语言"
	Book1.author = "zxx"
	Book1.subject  = "subject:"
	Book1.book_id = 123131


	Book2.title="python 语言"
	Book1.author = "zxzx"
	Book1.subject  = "subject2"
	Book1.book_id = 1212

	fmt.Printf("Book 1 title:%s\n", Book1.title)
	fmt.Printf("Book 1 title:%s\n", Book2.title)

	printBook(Book1)
}
