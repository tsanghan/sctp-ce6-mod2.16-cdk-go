package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
)

type SctpCe6Mod216CdkGoStackProps struct {
	awscdk.StackProps
}

func NewSctpCe6Mod216CdkGoStack(scope constructs.Construct, id string, props *SctpCe6Mod216CdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("SctpCe6Mod216CdkGoQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })
	awss3.NewBucket(stack, jsii.String("tsanghan-ce6-Go-CdkBucket"), &awss3.BucketProps{
		Versioned: jsii.Bool(true),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		AutoDeleteObjects: jsii.Bool(true),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewSctpCe6Mod216CdkGoStack(app, "C6TsanghanGoS3CdkStack", &SctpCe6Mod216CdkGoStackProps{
		awscdk.StackProps{
			Description: jsii.String("tsanghan-ce6: This stack is synth by Go"),
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
	 Account: jsii.String(os.Getenv("AWS_ACCOUNT_ID")),
	 Region:  jsii.String(os.Getenv("AWS_DEFAULT_REGION")),
	}
}
