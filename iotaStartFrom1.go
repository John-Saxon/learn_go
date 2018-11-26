package main

import (
	"fmt"
)

const (
	DNN = iota + 1
	KMeans
	RandomForest
	BoostedTree
	SVM
)

func main() {

	fmt.Println(DNN,
		KMeans,
		RandomForest,
		BoostedTree,
		SVM)
}
