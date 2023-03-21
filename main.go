package main

import (
	"go-aws-playground/services"
	"log"
)

func main() {
	paramStorePath := "/my/parameter/store/path"
	paramStoreValue := "my-parameter-value"

	error := services.WriteToSSM(paramStorePath, paramStoreValue)
	if error != nil {
		log.Default().Fatal("Unable to write to SSM" + error.Error())
	}
	ssmValue, err := services.ReadFromSSM(paramStorePath)
	if err != nil {
		log.Default().Fatal("Unable to read from SSM" + err.Error())
	}
	log.Default().Println("SSM value: " + ssmValue)
}
