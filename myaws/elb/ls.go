package elb

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/pkg/errors"

	"github.com/minamijoyo/myaws/myaws"
)

// Ls describes ELBs.
func Ls(client *myaws.Client) error {
	params := &elb.DescribeLoadBalancersInput{}

	response, err := client.ELB.DescribeLoadBalancers(params)
	if err != nil {
		return errors.Wrap(err, "DescribeLoadBalancers failed:")
	}

	for _, lb := range response.LoadBalancerDescriptions {
		fmt.Println(formatLoadBalancer(lb))
	}

	return nil
}

func formatLoadBalancer(lb *elb.LoadBalancerDescription) string {
	output := []string{
		*lb.LoadBalancerName,
	}

	return strings.Join(output[:], "\t")
}
