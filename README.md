awscccli
========

`awscccli` (AWS Cloud Control API CLI) is a simple style AWS command line tool with [AWS Cloud Control API](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/what-is-cloudcontrolapi.html). `awscccli` supports only read actions (list, get).

Install
-------

### Homebrew

```bash
$ brew install chroju/tap/awscccli
```

### Download binary

Download the latest binary from [here](https://github.com/chroju/awscccli/releases) and place it in the some directory specified by `$PATH`.

Authentication
--------------

`awscccli` requires your AWS IAM authentication. The same authentication method as [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) is available. Tools like [aws-vault](https://github.com/99designs/aws-vault) can be used as well.

```
# with command line options
$ acc --profile YOUR_PROFILE

# with aws-vault
$ aws-vault exec YOUR_PROFILE -- acc
```

Usage
-----

### types

```bash
$ acc types | grep VPC
AWS::EC2::LocalGatewayRouteTableVPCAssociation
AWS::EC2::VPC
AWS::EC2::VPCCidrBlock
AWS::EC2::VPCDHCPOptionsAssociation
AWS::EC2::VPCEndpoint
AWS::EC2::VPCEndpointConnectionNotification
AWS::EC2::VPCEndpointService
AWS::EC2::VPCEndpointServicePermissions
AWS::EC2::VPCGatewayAttachment
AWS::EC2::VPCPeeringConnection
```

### list

```bash
$ acc list AWS::EC2::VPC
vpc-1fxxxxxx
vpc-8cxxxxxx
vpc-7fxxxxxx
```

### get

```bash
$ acc get AWS::EC2::VPC vpc-1fb7907b --format yaml
CidrBlock: 10.0.0.0/16
CidrBlockAssociations:
    - vpc-cidr-assoc-XXXXXXXX
DefaultNetworkAcl: acl-XXXXXXXX
DefaultSecurityGroup: sg-XXXXXXXX
EnableDnsHostnames: false
EnableDnsSupport: true
InstanceTenancy: default
Ipv6CidrBlocks: []
Tags:
    - Key: Name
      Value: dev
VpcId: vpc-1fXXXXXX
```

LICENSE
-------

MIT
