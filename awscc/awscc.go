package awscc

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudcontrolapi"
	"github.com/aws/aws-sdk-go/service/cloudcontrolapi/cloudcontrolapiiface"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

// AWSCCManager is the wrapper of Cloud Control API.
type AWSCCManager interface {
	ListTypes() ([]*string, error)
	ListResources(string) ([]*string, error)
	GetResource(string, string) (*string, error)
}

type awsccManager struct {
	cloudControlAPI cloudcontrolapiiface.CloudControlApiAPI
	cloudFormation  cloudformationiface.CloudFormationAPI
}

// New returns a new AWSCCManager.
func New(profile, region string) (AWSCCManager, error) {
	config := &aws.Config{}
	if profile != "" {
		config.Credentials = credentials.NewSharedCredentials("", profile)
	}

	if region != "" {
		config.Region = aws.String(region)
	}

	sess := session.Must(session.NewSession(config))
	_, err := sess.Config.Credentials.Get()
	if err != nil {
		return nil, err
	}

	ccapi := cloudcontrolapi.New(sess)
	cf := cloudformation.New(sess)
	return &awsccManager{
		cloudControlAPI: ccapi,
		cloudFormation:  cf,
	}, nil
}

// ListTypes lists available resource types.
func (m *awsccManager) ListTypes() ([]*string, error) {
	input := &cloudformation.ListTypesInput{
		Type:       aws.String(cloudformation.RegistryTypeResource),
		Visibility: aws.String(cloudformation.VisibilityPublic),
		MaxResults: aws.Int64(100),
	}

	var typeSummaries []*cloudformation.TypeSummary
	for {
		output, err := m.cloudFormation.ListTypes(input)
		if err != nil {
			return nil, err
		}
		typeSummaries = append(typeSummaries, output.TypeSummaries...)
		if output.NextToken == nil {
			break
		}
		input.NextToken = output.NextToken
	}

	resp := make([]*string, len(typeSummaries))
	for i, v := range typeSummaries {
		resp[i] = v.TypeName
	}

	return resp, nil
}

// ListResources lists resources.
func (m *awsccManager) ListResources(typeName string) ([]*string, error) {
	input := &cloudcontrolapi.ListResourcesInput{
		TypeName:   aws.String(typeName),
		MaxResults: aws.Int64(100),
	}

	var resourceDescriptions []*cloudcontrolapi.ResourceDescription
	for {
		output, err := m.cloudControlAPI.ListResources(input)
		if err != nil {
			return nil, err
		}
		resourceDescriptions = append(resourceDescriptions, output.ResourceDescriptions...)
		if output.NextToken == nil {
			break
		}
		input.NextToken = output.NextToken
	}

	resp := make([]*string, len(resourceDescriptions))
	for i, v := range resourceDescriptions {
		resp[i] = v.Identifier
	}

	return resp, nil
}

// GetResource gets a resource.
func (m *awsccManager) GetResource(typeName, Identifier string) (*string, error) {
	input := &cloudcontrolapi.GetResourceInput{
		TypeName:   aws.String(typeName),
		Identifier: aws.String(Identifier),
	}

	output, err := m.cloudControlAPI.GetResource(input)
	if err != nil {
		return nil, err
	}

	return output.ResourceDescription.Properties, nil
}
