package services

import (
	"context"
	"go-aws-playground/common"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

func WriteToSSM(paramStorePath string, paramStoreValue string) error {
	ssmClient := common.GetSSMClient()
	input := &ssm.PutParameterInput{
		Name:      aws.String(paramStorePath),
		Value:     aws.String(paramStoreValue),
		Type:      types.ParameterTypeString,
		Overwrite: aws.Bool(true),
	}
	log.Default().Println("Calling SSM PutParameter")
	_, err := ssmClient.PutParameter(context.Background(), input)
	if err != nil {
		log.Default().Fatal("Unable to write to SSM: ", err.Error())
		return err
	}
	return nil

}

func ReadFromSSM(paramStorePath string) (string, error) {
	ssmClient := common.GetSSMClient()
	input := &ssm.GetParameterInput{
		Name: &paramStorePath,
	}
	log.Default().Println("Calling SSM GetParameter")
	result, err := ssmClient.GetParameter(context.Background(), input)
	if err != nil {
		log.Default().Fatal("Unable to read from SSM: ", err.Error())
		return "", err
	}
	return *result.Parameter.Value, nil
}
