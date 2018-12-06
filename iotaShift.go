package main

import (
	"fmt"
)

const (
	ClassificationMask = 1 << iota
	RegressionMask
	ClusteringMask
)

const (
	ClassificationMask1 = 1 << (2 - iota)
	RegressionMask1
	ClusteringMask1
)

func main() {
	fmt.Println(ClassificationMask, RegressionMask, ClusteringMask)
	fmt.Println(ClassificationMask1, RegressionMask1, ClusteringMask1)
}
