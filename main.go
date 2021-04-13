package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func main() {
	ec2svc := ec2.New(session.New())
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending")},
			},
		},
	}
	resp, err := ec2svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
		log.Fatal(err.Error())
	}
	fmt.Println("Deleting instances... ")
	for idx, res := range resp.Reservations {
		fmt.Println("  > Reservation Id", *res.ReservationId, " Num Instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			fmt.Println("    - Stopping Instance ID: ", *inst.InstanceId)
			input := &ec2.StopInstancesInput{
				InstanceIds: []*string{
					aws.String(*inst.InstanceId),
				},
			}
			_, err := ec2svc.StopInstances(input)
			if err != nil {
				fmt.Println("there was an error stopping instances", err.Error())
				log.Fatal(err.Error())
			}

		}
	}
}

