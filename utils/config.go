package utils

import (
	"blog/conf"
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
)

func InitConfig() {

	if os.Getenv("GIN_MODE") == "release" {
		conf.Enviroment	 = "production"
	} else {
		conf.Enviroment	 = "development"
	}

	curDir,_ 		:= os.Getwd()
	confPath		:= curDir + "/conf/" + conf.Enviroment + "/config.json"
	jsonData,err	:= ioutil.ReadFile(confPath)
	if err != nil {
		log.Println(err)
	}

	err	= json.Unmarshal(jsonData, &conf.LocalDB)

	if err != nil {
		log.Fatal(err)
	}
}
