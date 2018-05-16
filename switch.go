package main 


func main(){
	grade:= "B"
 	marks:=90

	switch marks{
	case 90: grade = "A"
	case 80:grade = "B"
	case 50,60,70:grade="C"
	default: grade="D"
	}

	switch{
		case grade =="A":
			println("A")
		case grade =="B":println("B")
		case grade=="C":
			println("C")
		case grade=="D":
			println("D")
		default:
			println("E")
	}

}
