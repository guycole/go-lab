package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {

	// need service account

	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-2"), config.WithSharedConfigProfile("terraform_funny"))
	if err != nil {
		log.Fatal(err)
	}

	client := sts.NewFromConfig(cfg)
	identity, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(identity)
	fmt.Printf("Account: %s, Arn: %s", aws.ToString(identity.Account), aws.ToString(identity.Arn))

	log.Println("graceful exit")
}
