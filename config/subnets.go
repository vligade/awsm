package config

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/simpledb"
)

// SubnetClasses is a map of Subnet Classes
type SubnetClasses map[string]SubnetClass

// SubnetClass is a single Subnet Class
type SubnetClass struct {
	CIDR string `json:"cidr" awsmClass:"CIDR"`

	// INTERNET GATEWAY
	CreateInternetGateway              bool `json:"createInternetGateway" awsmClass:"Create Internet Gateway"`
	AddInternetGatewayToMainRouteTable bool `json:"addInternetGatewayToMainRouteTable" awsmClass:"Add Internet Gateway To Main Route Table"`
	AddInternetGatewayToNewRouteTable  bool `json:"addInternetGatewayToNewRouteTable" awsmClass:"Add Internet Gateway To New Route Table"`

	// NAT GATEWAY
	CreateNatGateway              bool `json:"createNatGateway" awsmClass:"Create NAT Gateway"`
	AddNatGatewayToMainRouteTable bool `json:"addNatGatewayToMainRouteTable" awsmClass:"Add NAT Gateway To Main Route Table"`
	AddNatGatewayToNewRouteTable  bool `json:"addNatGatewayToNewRouteTable" awsmClass:"Add NAT Gateway To New Route Table"`
}

// DefaultSubnetClasses returns the defauly Subnet Classes
func DefaultSubnetClasses() SubnetClasses {
	defaultSubnets := make(SubnetClasses)

	defaultSubnets["private"] = SubnetClass{
		CIDR: "/24",
	}

	defaultSubnets["public"] = SubnetClass{
		CIDR: "/24",
		CreateInternetGateway:             true,
		AddInternetGatewayToNewRouteTable: true,
		CreateNatGateway:                  true,
		AddNatGatewayToMainRouteTable:     true,
	}

	return defaultSubnets
}

// SaveSubnetClass reads unmarshals a byte slice and inserts it into the db
func SaveSubnetClass(className string, data []byte) (class SubnetClass, err error) {
	err = json.Unmarshal(data, &class)
	if err != nil {
		return
	}

	err = Insert("subnets", SubnetClasses{className: class})
	return
}

// LoadSubnetClass loads a Subnet Class by its name
func LoadSubnetClass(name string) (SubnetClass, error) {
	cfgs := make(SubnetClasses)
	item, err := GetItemByName("subnets", name)
	if err != nil {
		return cfgs[name], err
	}

	cfgs.Marshal([]*simpledb.Item{item})
	return cfgs[name], nil
}

// LoadAllSubnetClasses loads all Subnet Classes
func LoadAllSubnetClasses() (SubnetClasses, error) {
	cfgs := make(SubnetClasses)
	items, err := GetItemsByType("subnets")
	if err != nil {
		return cfgs, err
	}

	cfgs.Marshal(items)
	return cfgs, nil
}

// Marshal puts items from SimpleDB into a Subnet Class
func (c SubnetClasses) Marshal(items []*simpledb.Item) {
	for _, item := range items {
		name := strings.Replace(*item.Name, "subnets/", "", -1)
		cfg := new(SubnetClass)
		for _, attribute := range item.Attributes {

			val := *attribute.Value

			switch *attribute.Name {

			case "CIDR":
				cfg.CIDR = val

			case "CreateInternetGateway":
				cfg.CreateInternetGateway, _ = strconv.ParseBool(val)

			case "AddInternetGatewayToMainRouteTable":
				cfg.AddInternetGatewayToMainRouteTable, _ = strconv.ParseBool(val)

			case "AddInternetGatewayToNewRouteTable":
				cfg.AddInternetGatewayToNewRouteTable, _ = strconv.ParseBool(val)

			case "CreateNatGateway":
				cfg.CreateNatGateway, _ = strconv.ParseBool(val)

			case "AddNatGatewayToMainRouteTable":
				cfg.AddNatGatewayToMainRouteTable, _ = strconv.ParseBool(val)

			case "AddNatGatewayToNewRouteTable":
				cfg.AddNatGatewayToNewRouteTable, _ = strconv.ParseBool(val)

			}
		}
		c[name] = *cfg
	}
}
