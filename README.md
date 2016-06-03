# awsm
> AWS iMproved

## Intro
**awsm** is an alternative interface for Amazon Web Services. It is intended to streamline many of the tasks involved with setting up and scaling infrastructure across multiple AWS Regions.

## Origin
This application is the fourth rewrite of one of the tools I have built to make my duties as a one man Infrastructure / DevOps team of Salon as agile as a much larger team. Each implementation has incorporated things I have learned during my career, with the motivation of promoting the use of more secure and redundant infrastructure designs by providing tools that automate many of the tasks involved with setting up and maintaining complex AWS infrastructures.

## But what about containers? ##
Sometimes your current needs require you to build or migrate and scale monolith applications like a CMS on Amazon Web Services, with a very small team and limited resources. There weren't any tools that provided the automation and simplicity we needed, this tool is a result of that.

## Installation

## Commands

## Roadmap
* ~~dashboard - "Launch the awsm Dashboard GUI"~~
* attachVolume - "Attach an AWS EBS Volume"
* copyImage - "Copy an AWS Machine Image to another region"
* copySnapshot - "Copy an AWS EBS Snapshot to another region"
* createAddress - "Create an AWS Elastic IP Address (for use in a VPC or EC2-Classic)"
* createAutoScaleGroup - "Create an AWS AutoScaling Group"
* ~~createIAMUser - "Create an IAM User"~~
* createImage - "Create an AWS Machine Image from a running instance"
* createLaunchConfiguration - "Create an AWS AutoScaling Launch Configuration"
* ~~createSimpleDBDomain - "Create an AWS SimpleDB Domain"~~
* createSnapshot - "Create an AWS EBS snapshot of a volume"
* createVolume - "Create an AWS EBS volume (from a class snapshot or blank)"
* ~~createVpc - "Create an AWS VPC"~~
* ~~createSubnet - "Create an AWS VPC Subnet"~~
* deleteAutoScaleGroup - "Delete an AWS AutoScaling Group"
* ~~deleteIAMUser - "Delete an AWS Machine Image"~~
* deleteImage - "Delete an AWS Machine Image"
* deleteLaunchConfiguration - "Delete an AWS AutoScaling Launch Configuration"
* deleteSnapshot - "Delete an AWS EBS Snapshot"
* ~~deleteSimpleDBDomains - "Delete an AWS SimpleDB Domain"~~
* deleteVolume - "Delete an AWS EBS Volume"
* ~~deleteSubnets - "Delete AWS VPC Subnets"~~
* ~~deleteVpcs - "Delete AWS VPCs"~~
* detachVolume - "Detach an AWS EBS Volume"
* stopInstances - "Stop AWS instance(s)"
* pauseInstances - "Pause AWS instance(s)"
* killInstances - "Kill AWS instance(s)"
* launchInstance - "Launch an EC2 instance"
* ~~listAddresses  - "Lists all AWS Elastic IP Addresses"~~
* ~~listAlarms  - "Lists all CloudWatch Alarms"~~
* ~~listAutoScaleGroups  - "Lists all AutoScale Groups"~~
* ~~listIAMUsers - "Lists all IAM Users"~~
* ~~listImages  - "Lists all AWS Machine Images owned by us"~~
* ~~listInstances  - "Lists all AWS EC2 Instances"~~
* ~~listLaunchConfigurations  - "List all Launch Configurations"~~
* ~~listLoadBalancers  - "Lists all Elastic Load Balancers"~~
* ~~listScalingPolicies  - "Lists all Scaling Policies"~~
* ~~listSecurityGroups  - "Lists all Security Groups"~~
* ~~listSnapshots  - "Lists all AWS EBS Snapshots"~~
* ~~listSubnets  - "Lists all AWS Subnets"~~
* ~~listSimpleDBDomains - "Lists all AWS SimpleDB Domains"~~
* ~~listVolumes  - "Lists all AWS EBS Volumes"~~
* ~~listVpcs  - "Lists all AWS Vpcs"~~
* resumeProcesses - "Resume all autoscaling processes on a specific autoscaling group"
* runCommand - "Run a command on a set of instances"
* suspendProcesses - "Stop all autoscaling processes on a specific autoscaling group"
* updateAutoScaleGroup - "Update an AWS AutoScaling Group"