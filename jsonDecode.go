package main

import (
	"encoding/json"
	"fmt"
)

type ManualTrainConfig struct {
	TrainBatchSize           []int       `json:"batch_size,omitempty"` // common
	NumEpochs                []int       `json:"num_epochs,omitempty"`
	TrainSteps               []int       `json:"train_steps"`
	HiddenUnits              [][]int     `json:"hidden_layers,omitempty"` // dnn
	LearningRate             []float32   `json:"learning_rate,omitempty"`
	Optimizer                []string    `json:"optimizer,omitempty"`
	RandomSearch             []float32   `json:"random_search,omitempty"`
	DistanceMetric           []string    `json:"distance_metric,omitempty"`              //kmeans
	InitialClusters          []string    `json:"initial_clusters,omitempty"`             // must be random or kmeans_plus_plus
	KmeansPlusPlusNumRetries []int64     `json:"kmeans_plus_plus_num_retries,omitempty"` //  a heuristic is used to sample O(log(num_to_sample))
	NCluster                 []int       `json:"n_cluster,omitempty"`
	RandomSeed               []int       `json:"random_seed,omitempty"`
	MaxDepth                 []int       `json:"max_depth,omitempty"` // boosted tree
	MinNodeWeight            []int       `json:"min_node_weight,omitempty"`
	NBatchesPerLayer         []int       `json:"n_batches_per_layer"`
	NTrees                   []int       `json:"n_trees,omitempty"`
	Regularization           [][]float32 `json:"regularization"`
	TreeComplexity           []int       `json:"tree_complexity,omitempty"`
	ApplyRff                 []int       `json:"apply_rff_mapping,omitempty"` // svm
	MapDim                   []int       `json:"map_dim"`
	MapStddev                []float32   `json:"map_stddev"`
	NumTrees                 []int       `json:"num_trees"` //random forest
	MaxNodes                 []int       `json:"max_nodes"`
}

func main() {
	str := `{"1":{"batch_size":[50],"hidden_layers":[[64,64,64]],"learning_rate":[0.1],"num_epochs":[1],"optimizer":["adam"],"random_search":[1],"regularization":[[0,0]],"train_steps":[10000]},"2":{"batch_size":[50],"hidden_layers":[[64,64,64]],"learning_rate":[0.1],"num_epochs":[1],"optimizer":["adam"],"random_search":[1],"regularization":[[0,0]],"train_steps":[10000]},"3":{"batch_size":[50],"distance_metric":["cosine"],"initial_clusters":["random"],"kmeans_plus_plus_num_retries":[2],"n_cluster":[2],"num_epochs":[1],"random_seed":[0],"train_steps":[10000]},"4":{"batch_size":[50],"max_nodes":[20],"num_epochs":[1],"num_trees":[10],"random_search":[1],"train_steps":[10000]},"5":{"batch_size":[50],"learning_rate":[0.1],"max_depth":[6],"min_node_weight":[0],"n_batches_per_layer":[100],"n_trees":[100],"num_epochs":[1],"regularization":[[0,0]],"train_steps":[10000],"tree_complexity":[0]},"6":{"apply_rff_mapping":[1],"batch_size":[50],"learning_rate":[0.1],"map_dim":[15000],"map_stddev":[1],"num_epochs":[1],"train_steps":[10000]}}`
	// aaa := `{"1":{"batch_size":[50],"num_epochs":[1],"train_steps":[10000],"hidden_layers":[[64,64,64]],"learning_rate":[0.1],"optimizer":["adam"],"random_search":[1],"n_batches_per_layer":null,"regularization":[[0,0]],"map_dim":null,"map_stddev":null},"2":{"batch_size":[50],"num_epochs":[1],"train_steps":[10000],"hidden_layers":[[64,64,64]],"learning_rate":[0.1],"optimizer":["adam"],"random_search":[1],"n_batches_per_layer":null,"regularization":[[0,0]],"map_dim":null,"map_stddev":null},"3":{"batch_size":[50],"num_epochs":[1],"train_steps":[10000],"distance_metric":["cosine"],"initial_clusters":["random"],"kmeans_plus_plus_num_retries":[2],"n_cluster":[2],"random_seed":[0],"n_batches_per_layer":null,"regularization":null,"map_dim":null,"map_stddev":null},"4":{"batch_size":[50],"num_epochs":[1],"train_steps":[10000],"random_search":[1],"n_batches_per_layer":null,"regularization":null,"map_dim":null,"map_stddev":null},"5":{"batch_size":[50],"num_epochs":[1],"train_steps":[10000],"learning_rate":[0.1],"max_depth":[6],"min_node_weight":[0],"n_batches_per_layer":[100],"n_trees":[100],"regularization":[[0,0]],"tree_complexity":[0],"map_dim":null,"map_stddev":null},"6":{"batch_size":[50],"num_epochs":[1],"train_steps":[10000],"learning_rate":[0.1],"n_batches_per_layer":null,"regularization":null,"apply_rff_mapping":[1],"map_dim":[15000],"map_stddev":[1]}}`
	var gsConfig = make(map[int]ManualTrainConfig)
	err := json.Unmarshal([]byte(str), &gsConfig)
	if err != nil {
		fmt.Println(err)
	}
	a, _ := json.Marshal(gsConfig)
	fmt.Println(string(a))
}
