package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/models/response"
	"os"
)

const file = "models/response/err_code.json"

type ErrMap map[response.ErrorCode]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
	fmt.Println(errMap[response.SettingsError])
}
