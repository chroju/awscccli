package awscc

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestNew(t *testing.T) {
	type args struct {
		profile string
		region  string
	}
	tests := []struct {
		name    string
		args    args
		want    AWSCCManager
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.profile, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awsccManager_ListTypes(t *testing.T) {
	types := make([]*string, len(mockResources))
	var i int
	for key := range mockResources {
		types[i] = aws.String(key)
		i++
	}
	tests := []struct {
		name    string
		want    []*string
		wantErr bool
	}{
		{
			name:    "execute",
			want:    types,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMockAWSCCManager()
			got, err := m.ListTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("awsccManager.ListTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("awsccManager.ListTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awsccManager_ListResources(t *testing.T) {
	vpcs := make(map[*string]*string)
	for _, v := range mockResources["AWS::EC2::VPC"] {
		vpcs[v.Identifier] = v.Properties
	}
	type args struct {
		typeName string
	}
	tests := []struct {
		name    string
		args    args
		want    map[*string]*string
		wantErr bool
	}{
		{
			name:    "exist type",
			args:    args{typeName: "AWS::EC2::VPC"},
			want:    vpcs,
			wantErr: false,
		},
		{
			name:    "not exist type",
			args:    args{typeName: "AWS::EC2::NotExistType"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMockAWSCCManager()
			got, err := m.ListResources(tt.args.typeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("awsccManager.ListResources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("awsccManager.ListResources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awsccManager_GetResource(t *testing.T) {
	type args struct {
		typeName   string
		Identifier string
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		{
			name:    "exist resource",
			args:    args{typeName: "AWS::EC2::InternetGateway", Identifier: "igw-2exxxxxx"},
			want:    aws.String("{\"InternetGatewayId\":\"igw-2exxxxxx\"}"),
			wantErr: false,
		},
		{
			name:    "not exist resource",
			args:    args{typeName: "AWS::EC2::InternetGateway", Identifier: "igw-3cxxxxxx"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not exist type",
			args:    args{typeName: "AWS::EC2::NotExistType", Identifier: "igw-3cxxxxxx"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMockAWSCCManager()
			got, err := m.GetResource(tt.args.typeName, tt.args.Identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("awsccManager.GetResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("awsccManager.GetResource() = %v, want %v", got, tt.want)
			}
		})
	}
}
