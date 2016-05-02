package aws

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/murdinc/cli"
)

type ScalingPolicies []ScalingPolicy

type ScalingPolicy struct {
	Name               string
	AdjustmentType     string
	Adjustment         string
	Cooldown           string
	AutoScaleGroupName string
	Alarms             string
	Region             string
}

func GetScalingPolicies() (*ScalingPolicies, error) {
	var wg sync.WaitGroup

	spList := new(ScalingPolicies)
	regions := GetRegionList()

	for _, region := range regions {
		wg.Add(1)

		go func(region *ec2.Region) {
			defer wg.Done()
			err := GetRegionScalingPolicies(region.RegionName, spList)
			if err != nil {
				cli.ShowErrorMessage("Error gathering ScalingPolicy list", err.Error())
			}
		}(region)
	}
	wg.Wait()

	return spList, nil
}

func GetRegionScalingPolicies(region *string, spList *ScalingPolicies) error {
	svc := autoscaling.New(session.New(&aws.Config{Region: region}))
	result, err := svc.DescribePolicies(&autoscaling.DescribePoliciesInput{})

	if err != nil {
		return err
	}

	sp := make(ScalingPolicies, len(result.ScalingPolicies))
	for i, scalingpolicy := range result.ScalingPolicies {
		sp[i] = ScalingPolicy{
			Name:               aws.StringValue(scalingpolicy.PolicyName),
			AdjustmentType:     aws.StringValue(scalingpolicy.AdjustmentType),
			Adjustment:         string(aws.Int64Value(scalingpolicy.ScalingAdjustment)),
			Cooldown:           string(aws.Int64Value(scalingpolicy.Cooldown)),
			AutoScaleGroupName: aws.StringValue(scalingpolicy.AutoScalingGroupName),
			//Alarms:             strings.Join(aws.StringValueSlice(scalingpolicy.Alarms), ","), // TODO
			Region: *region,
		}
	}
	*spList = append(*spList, sp[:]...)

	return nil
}

func (i *ScalingPolicies) PrintTable() {
	collumns := []string{"Name", "Adjustment Type", "Adjustment", "Cooldown", "AutoScaling Group Name", "Alarms", "Region"}

	rows := make([][]string, len(*i))
	for index, val := range *i {
		rows[index] = []string{
			val.Name,
			val.AdjustmentType,
			val.Adjustment,
			val.Cooldown,
			val.AutoScaleGroupName,
			val.Alarms,
			val.Region,
		}
	}

	printTable(collumns, rows)
}