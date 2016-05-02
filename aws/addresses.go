package aws

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/murdinc/cli"
)

type Addresses []Address

type Address struct {
	PublicIp   string
	PrivateIp  string
	Domain     string
	InstanceId string
	Region     string
}

func GetAddresses() (*Addresses, error) {
	var wg sync.WaitGroup

	ipList := new(Addresses)
	regions := GetRegionList()

	for _, region := range regions {
		wg.Add(1)

		go func(region *ec2.Region) {
			defer wg.Done()
			err := GetRegionAddresses(region.RegionName, ipList)
			if err != nil {
				cli.ShowErrorMessage("Error gathering address list", err.Error())
			}
		}(region)
	}
	wg.Wait()

	return ipList, nil
}

func GetRegionAddresses(region *string, adrList *Addresses) error {
	svc := ec2.New(session.New(&aws.Config{Region: region}))
	result, err := svc.DescribeAddresses(&ec2.DescribeAddressesInput{})

	if err != nil {
		return err
	}

	adr := make(Addresses, len(result.Addresses))
	for i, address := range result.Addresses {

		adr[i] = Address{
			PublicIp:   aws.StringValue(address.PublicIp),
			PrivateIp:  aws.StringValue(address.PrivateIpAddress),
			InstanceId: aws.StringValue(address.InstanceId),
			Domain:     aws.StringValue(address.Domain),
			Region:     fmt.Sprintf(*region),
		}
	}
	*adrList = append(*adrList, adr[:]...)

	return nil
}

func (i *Addresses) PrintTable() {
	collumns := []string{"Public IP", "Private IP", "Domain", "Instance Id", "Region"}

	rows := make([][]string, len(*i))
	for index, val := range *i {
		rows[index] = []string{
			val.PublicIp,
			val.PrivateIp,
			val.Domain,
			val.InstanceId,
			val.Region,
		}
	}

	printTable(collumns, rows)
}