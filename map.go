package main 

import "fmt"

func main(){
	var ccmap map[string]string
	ccmap = make(map[string]string)

	ccmap["france"] = "paris"
	
	ccmap["italy"] = "rome"
	ccmap["japan"]="tokyo"

	for country := range ccmap{
		fmt.Println(country,"capital is ", ccmap[country])

	}

	capital,ok := ccmap["usa"]
	if(ok){
		fmt.Println("usa's capital is ", capital)
	}else{
		fmt.Println("not found")
	}
}
