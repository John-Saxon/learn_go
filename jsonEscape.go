package main

import (
	"fmt"

	"github.com/oushu/gocommon/other"
)

func main() {
	pss := `-a 2 -S "[{"person":{"p_personid":"32985348844221"}},{"organisation":{"o_organisationid":"3013"}}]"`
	fmt.Println(other.EncodeJson(pss))
	fmt.Println(pss)

	fmt.Println(fmt.Sprintf("-n 1 %s ", other.EncodeJson(pss)))
}
