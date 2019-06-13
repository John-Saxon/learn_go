package main

import (
	"github.com/colinmarc/hdfs"
	"github.com/oushu/gocommon/lavalog"
)

func main() {
	// err := os.etenv("HADOOP_CONF_DIR", "./")
	config := hdfs.LoadHadoopConf("./")
	nameNodes, err := config.Namenodes()
	if err != nil {
		lavalog.Error.Println(err)
		return
	}
	options := hdfs.ClientOptions{
		Addresses: nameNodes,
		User:      "hdfs",
	}
	client, err := hdfs.NewClient(options)
	if err != nil {
		lavalog.Error.Println(err)
		return
	}

	err = client.CopyToRemote("./test.csv", "/test.csv")
	if err != nil {
		lavalog.Error.Println(err)
		return
	}

	err = client.CopyToLocal("/test.csv", "/testToLocal.csv")
	if err != nil {
		lavalog.Error.Println(err)
		return
	}

	return
}
