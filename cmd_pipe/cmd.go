package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

type ImportanceRank struct {
	Rank       int     `json:"rank"`
	Importance float64 `json:"importance"`
}

// type FeatureImportance struct{

// }

func main() {
	fmt.Println("main")
	// TestUnmarshal()
	TestMain()
}

func TestUnmarshal() {
	str := `{"medinc": {"rank": 1, "importance": 0.4998692773106829}, "houseage": {"rank": 6, "importance": 0.03394551899684777}, "averooms": {"rank": 5, "importance": 0.05089803153592726}, "avebedrms": {"rank": 8, "importance": 0.0223154425839744}, "population": {"rank": 7, "importance": 0.028584765285710245}, "aveoccup": {"rank": 3, "importance": 0.09833268250245587}, "latitude": {"rank": 4, "importance": 0.07233666566391424}, "longitude": {"rank": 2, "importance": 0.19371761612048735}}`
	fm := make(map[string]ImportanceRank)
	err := json.Unmarshal([]byte(str), &fm)
	if err != nil {
		fmt.Println("Failed to Unmarshal:", err)
	}
	fmt.Println("After Unmarshal:", fm)
}

func TestMain() {
	// elec partial
	// config := "{\"hawq\":{\"master\":\"localhost\",\"standby\":\"\",\"port\":5432,\"user\":\"johnsaxon\",\"password\":\"test\",\"database\":\"postgres\"},\"data_config\":{\"schema\":\"public\",\"table\":\"elec\",\"features\":[{\"name\":\"id\",\"type\":\"categorical\",\"cates\":[\"2019-01-20\",\"2019-03-01\",\"2019-04-20\",\"2018-09-10\",\"2018-12-01\",\"2019-01-01\",\"2019-02-20\",\"2019-04-10\",\"2018-09-20\",\"2018-10-01\",\"2019-02-01\",\"2018-10-20\",\"2018-11-10\",\"2018-12-10\",\"2018-12-20\",\"2019-01-10\",\"2019-03-10\",\"2019-04-01\",\"2018-10-10\",\"2018-11-01\",\"2018-11-20\",\"2019-02-10\",\"2019-03-20\",\"2018-09-01\"]},{\"name\":\"stat_date\",\"type\":\"n\",\"cates\":[]},{\"name\":\"meter_id\",\"type\":\"n\",\"cates\":[]},{\"name\":\"energy_mean\",\"type\":\"n\",\"cates\":[]},{\"name\":\"energy_max\",\"type\":\"n\",\"cates\":[]},{\"name\":\"energy_min\",\"type\":\"n\",\"cates\":[]},{\"name\":\"energy_sum\",\"type\":\"n\",\"cates\":[]},{\"name\":\"energy_std\",\"type\":\"n\",\"cates\":[]},{\"name\":\"power_mean\",\"type\":\"n\",\"cates\":[]},{\"name\":\"power_max\",\"type\":\"n\",\"cates\":[]},{\"name\":\"power_min\",\"type\":\"n\",\"cates\":[]},{\"name\":\"power_std\",\"type\":\"n\",\"cates\":[]},{\"name\":\"cur_mean\",\"type\":\"n\",\"cates\":[]},{\"name\":\"cur_max\",\"type\":\"n\",\"cates\":[]},{\"name\":\"cur_min\",\"type\":\"n\",\"cates\":[]},{\"name\":\"cur_std\",\"type\":\"n\",\"cates\":[]},{\"name\":\"vol_mean\",\"type\":\"n\",\"cates\":[]},{\"name\":\"vol_max\",\"type\":\"n\",\"cates\":[]},{\"name\":\"vol_min\",\"type\":\"n\",\"cates\":[]},{\"name\":\"vol_std\",\"type\":\"n\",\"cates\":[]},{\"name\":\"x\",\"type\":\"n\",\"cates\":[]},{\"name\":\"avg_h8\",\"type\":\"n\",\"cates\":[]},{\"name\":\"avg_t_8\",\"type\":\"n\",\"cates\":[]},{\"name\":\"avg_ws_h\",\"type\":\"n\",\"cates\":[]},{\"name\":\"avg_wd_h\",\"type\":\"n\",\"cates\":[]},{\"name\":\"max_h8\",\"type\":\"n\",\"cates\":[]},{\"name\":\"max_t_8\",\"type\":\"n\",\"cates\":[]},{\"name\":\"max_ws_h\",\"type\":\"n\",\"cates\":[]},{\"name\":\"min_h8\",\"type\":\"n\",\"cates\":[]},{\"name\":\"min_t_8\",\"type\":\"n\",\"cates\":[]},{\"name\":\"min_ws_h\",\"type\":\"n\",\"cates\":[]},{\"name\":\"avg_irradiance\",\"type\":\"n\",\"cates\":[]},{\"name\":\"max_irradiance\",\"type\":\"n\",\"cates\":[]},{\"name\":\"min_irradiance\",\"type\":\"n\",\"cates\":[]}],\"label\":{\"name\":\"load\",\"type\":\"n\",\"cates\":[]}},\"task\":{\"type\":1,\"algorithm\":1,\"warm_start\":true,\"estimators\":3,\"incre\":1,\"batch\":50000}}"
	config := "{\"hawq\":{\"master\":\"localhost\",\"standby\":\"\",\"port\":5432,\"user\":\"johnsaxon\",\"password\":\"test\",\"database\":\"postgres\"},\"data_config\":{\"schema\":\"public\",\"table\":\"housing_predict\",\"features\":[{\"name\":\"medinc\",\"cates\":[],\"type\":\"numerical\"},{\"name\":\"houseage\",\"cates\":[],\"type\":\"numerical\"},{\"name\":\"averooms\",\"cates\":[],\"type\":\"numerical\"},{\"name\":\"avebedrms\",\"cates\":[],\"type\":\"numerical\"},{\"name\":\"population\",\"cates\":[],\"type\":\"numerical\"},{\"name\":\"aveoccup\",\"cates\":[],\"type\":\"numerical\"},{\"name\":\"latitude\",\"cates\":[],\"type\":\"numerical\"},{\"name\":\"longitude\",\"cates\":[],\"type\":\"numerical\"}],\"label\":{\"name\":\"price\"}},\"task\":{\"type\":2,\"algorithm\":1,\"warm_start\":false,\"estimators\":3,\"incre\":1,\"batch\":1000}}"
	// housing sheer
	cmd := exec.Command("/anaconda3/bin/python3", "/Users/johnsaxon/go/src/github.com/oushu-io/littleboy/feature-importance/feature_importance.py")

	// get output
	var out bytes.Buffer
	cmd.Stdout = &out

	stdIn, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Failed to bind StdinPipe:", err.Error())
		return
	}
	go func() {
		defer stdIn.Close()
		_, err := io.WriteString(stdIn, config)
		if err != nil {
			fmt.Println("write config:", err)
		}
	}()

	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// set error config
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		stderr.Close()
		// stdout.Close()
	}()

	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start cmd:", err.Error())
	}
	pid := cmd.Process.Pid
	fmt.Println("pid is:", pid)

	// read stderr
	errMsg := []byte{}
	r := bufio.NewReader(stderr)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			fmt.Println("Failed to read err:", err)
			break
		} else {
			errMsg = append(errMsg, line...)
			errMsg = append(errMsg, '\n')
		}
	}
	fmt.Println("STDERR: ", string(errMsg))

	// read stdout
	// outMsg := []byte{}
	// out_r := bufio.NewReader(stdout)
	// for {
	// 	line, _, err := out_r.ReadLine()
	// 	if err != nil {
	// 		fmt.Println("Failed to read err:", err)
	// 		break
	// 	} else {
	// 		errMsg = append(outMsg, line...)
	// 		errMsg = append(outMsg, '\n')
	// 	}
	// }
	// fmt.Println("STDOUT: ", string(outMsg))

	if err := cmd.Wait(); err != nil {
		fmt.Println("Failed to wait cmd:", err)
		return
	}

	fmt.Println("Success!")
	res := out.String()
	fmt.Println(res)
	fm := make(map[string]ImportanceRank)

	err = json.Unmarshal(out.Bytes(), &fm)
	if err != nil {
		fmt.Println("Failed to Unmarshal:", err)
	}
	fmt.Println("After Unmarshal:", fm)
}
