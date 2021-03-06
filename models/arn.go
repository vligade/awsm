package models

// ARN represents an Amazon Resource Number
type ARN struct {
	Arn                  string `json:"arn"`
	Partition            string `json:"partition"`
	Service              string `json:"service"`
	Region               string `json:"region"`
	AccountID            string `json:"accountID"`
	PolicyID             string `json:"policyID"`
	GroupID              string `json:"groupID"`
	AutoScalingGroupName string `json:"autoScalingGroupName"`
	PolicyName           string `json:"policyName"`
	ProfileName          string `json:"profileName"`
	ResourceType         string `json:"resourceType"`
	Resource             string `json:"resource"`
}
