package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EventDetail struct {
	Action string `json:"action"`
}

func startNatGateway(ctx context.Context) error {
	// Create AWS session
	sess := session.Must(session.NewSession())

	// Create EC2 service client
	svc := ec2.New(sess)

	// Specify the Nat Gateway ID to start
	natGatewayID := "YOUR_NAT_GATEWAY_ID"

	// Start the Nat Gateway
	_, err := svc.StartNatGatewayWithContext(ctx, &ec2.StartNatGatewayInput{
		NatGatewayId: aws.String(natGatewayID),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Nat Gateway %s started successfully\n", natGatewayID)
	return nil
}

func stopNatGateway(ctx context.Context) error {
	// Create AWS session
	sess := session.Must(session.NewSession())

	// Create EC2 service client
	svc := ec2.New(sess)

	// Specify the Nat Gateway ID to stop
	natGatewayID := "YOUR_NAT_GATEWAY_ID"

	// Stop the Nat Gateway
	_, err := svc.StopNatGatewayWithContext(ctx, &ec2.StopNatGatewayInput{
		NatGatewayId: aws.String(natGatewayID),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Nat Gateway %s stopped successfully\n", natGatewayID)
	return nil
}

func handler(ctx context.Context, event EventDetail) error {
	// Extract action from event detail
	action := event.Action

	// Perform action based on event
	switch action {
	case "start":
		return startNatGateway(ctx)
	case "stop":
		return stopNatGateway(ctx)
	default:
		return fmt.Errorf("unsupported action: %s", action)
	}
}

func main() {
	lambda.Start(handler)
}