package awscc

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudcontrolapi"
	"github.com/aws/aws-sdk-go/service/cloudcontrolapi/cloudcontrolapiiface"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

func NewMockAWSCCManager() AWSCCManager {
	return &awsccManager{
		cloudControlAPI: &mockCloudControlAPI{},
		cloudFormation:  &mockCloudFormation{},
	}
}

type mockCloudFormation struct {
	cloudformationiface.CloudFormationAPI
}

type mockCloudControlAPI struct {
	cloudcontrolapiiface.CloudControlApiAPI
}

var mockResources = map[string][]*cloudcontrolapi.ResourceDescription{
	"AWS::EC2::InternetGateway": {
		{
			Identifier: aws.String("igw-1fxxxxxx"),
			Properties: aws.String("{\"InternetGatewayId\":\"igw-1fxxxxxx\"}"),
		},
		{
			Identifier: aws.String("igw-2exxxxxx"),
			Properties: aws.String("{\"InternetGatewayId\":\"igw-2exxxxxx\"}"),
		},
	},
	"AWS::EC2::VPC": {
		{
			Identifier: aws.String("vpc-1fxxxxxx"),
			Properties: aws.String("{\"CidrBlock\":\"10.0.0.0/16\",\"CidrBlockAssociations\":[\"vpc-cidr-assoc-XXXXXXXX\"],\"InstanceTenancy\":\"default\",\"Tags\":[{\"Key\":\"Name\",\"Value\":\"dev\"}],\"VpcId\":\"vpc-1fxxxxxx\"}"),
		},
		{
			Identifier: aws.String("vpc-2exxxxxx"),
			Properties: aws.String("{\"CidrBlock\":\"10.0.0.0/16\",\"CidrBlockAssociations\":[\"vpc-cidr-assoc-XXXXXXXX\"],\"InstanceTenancy\":\"default\",\"Tags\":[{\"Key\":\"Name\",\"Value\":\"stg\"}],\"VpcId\":\"vpc-2exxxxxx\"}"),
		},
	},
}

func (m *mockCloudFormation) ListTypes(input *cloudformation.ListTypesInput) (*cloudformation.ListTypesOutput, error) {
	typeSummaries := make([]*cloudformation.TypeSummary, len(mockResources))
	var i int
	for key := range mockResources {
		typeSummaries[i] = &cloudformation.TypeSummary{
			TypeName: aws.String(key),
		}
		i++
	}
	return &cloudformation.ListTypesOutput{
		TypeSummaries: typeSummaries,
	}, nil
}

func (m *mockCloudControlAPI) ListResources(input *cloudcontrolapi.ListResourcesInput) (*cloudcontrolapi.ListResourcesOutput, error) {
	v, ok := mockResources[*input.TypeName]
	if !ok {
		return nil, fmt.Errorf(cloudcontrolapi.ErrCodeTypeNotFoundException)
	}
	return &cloudcontrolapi.ListResourcesOutput{
		ResourceDescriptions: v,
	}, nil
}

func (m *mockCloudControlAPI) GetResource(input *cloudcontrolapi.GetResourceInput) (*cloudcontrolapi.GetResourceOutput, error) {
	v, ok := mockResources[*input.TypeName]
	if !ok {
		return nil, fmt.Errorf(cloudcontrolapi.ErrCodeTypeNotFoundException)
	}

	var resp *cloudcontrolapi.ResourceDescription
	for _, resource := range v {
		if *resource.Identifier == *input.Identifier {
			resp = resource
			break
		}
	}
	if resp == nil {
		return nil, fmt.Errorf(cloudcontrolapi.ErrCodeResourceNotFoundException)
	}

	return &cloudcontrolapi.GetResourceOutput{
		ResourceDescription: resp,
	}, nil
}
