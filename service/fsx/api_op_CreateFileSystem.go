// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request object used to create a new Amazon FSx file system.
type CreateFileSystemInput struct {
	_ struct{} `type:"structure"`

	// (Optional) A string of up to 64 ASCII characters that Amazon FSx uses to
	// ensure idempotent creation. This string is automatically filled on your behalf
	// when you use the AWS Command Line Interface (AWS CLI) or an AWS SDK.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The type of Amazon FSx file system to create.
	//
	// FileSystemType is a required field
	FileSystemType FileSystemType `type:"string" required:"true" enum:"true"`

	// The ID of your AWS Key Management Service (AWS KMS) key. This ID is used
	// to encrypt the data in your file system at rest. For more information, see
	// Encrypt (https://docs.aws.amazon.com/kms/latest/APIReference/API_Encrypt.html)
	// in the AWS Key Management Service API Reference.
	KmsKeyId *string `min:"1" type:"string"`

	// The Lustre configuration for the file system being created. This value is
	// required if FileSystemType is set to LUSTRE.
	LustreConfiguration *CreateFileSystemLustreConfiguration `type:"structure"`

	// A list of IDs specifying the security groups to apply to all network interfaces
	// created for file system access. This list isn't returned in later requests
	// to describe the file system.
	SecurityGroupIds []string `type:"list"`

	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB - 65,536 GiB.
	//
	// For Lustre file systems, valid values are 1,200, 2,400, 3,600, then continuing
	// in increments of 3600 GiB.
	//
	// StorageCapacity is a required field
	StorageCapacity *int64 `min:"1" type:"integer" required:"true"`

	// Specifies the IDs of the subnets that the file system will be accessible
	// from. For Windows MULTI_AZ_1 file system deployment types, provide exactly
	// two subnet IDs, one for the preferred file server and one for the standy
	// file server. You specify one of these subnets as the preferred subnet using
	// the WindowsConfiguration > PreferredSubnetID property.
	//
	// For Windows SINGLE_AZ_1 file system deployment types and Lustre file systems,
	// provide exactly one subnet ID. The file server is launched in that subnet's
	// Availability Zone.
	//
	// SubnetIds is a required field
	SubnetIds []string `type:"list" required:"true"`

	// The tags to apply to the file system being created. The key value of the
	// Name tag appears in the console as the file system name.
	Tags []Tag `min:"1" type:"list"`

	// The Microsoft Windows configuration for the file system being created. This
	// value is required if FileSystemType is set to WINDOWS.
	WindowsConfiguration *CreateFileSystemWindowsConfiguration `type:"structure"`
}

// String returns the string representation
func (s CreateFileSystemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFileSystemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFileSystemInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if len(s.FileSystemType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemType"))
	}
	if s.KmsKeyId != nil && len(*s.KmsKeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KmsKeyId", 1))
	}

	if s.StorageCapacity == nil {
		invalidParams.Add(aws.NewErrParamRequired("StorageCapacity"))
	}
	if s.StorageCapacity != nil && *s.StorageCapacity < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("StorageCapacity", 1))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.LustreConfiguration != nil {
		if err := s.LustreConfiguration.Validate(); err != nil {
			invalidParams.AddNested("LustreConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.WindowsConfiguration != nil {
		if err := s.WindowsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("WindowsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response object returned after the file system is created.
type CreateFileSystemOutput struct {
	_ struct{} `type:"structure"`

	// The configuration of the file system that was created.
	FileSystem *FileSystem `type:"structure"`
}

// String returns the string representation
func (s CreateFileSystemOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateFileSystem = "CreateFileSystem"

// CreateFileSystemRequest returns a request value for making API operation for
// Amazon FSx.
//
// Creates a new, empty Amazon FSx file system.
//
// If a file system with the specified client request token exists and the parameters
// match, CreateFileSystem returns the description of the existing file system.
// If a file system specified client request token exists and the parameters
// don't match, this call returns IncompatibleParameterError. If a file system
// with the specified client request token doesn't exist, CreateFileSystem does
// the following:
//
//    * Creates a new, empty Amazon FSx file system with an assigned ID, and
//    an initial lifecycle state of CREATING.
//
//    * Returns the description of the file system.
//
// This operation requires a client request token in the request that Amazon
// FSx uses to ensure idempotent creation. This means that calling the operation
// multiple times with the same client request token has no effect. By using
// the idempotent operation, you can retry a CreateFileSystem operation without
// the risk of creating an extra file system. This approach can be useful when
// an initial call fails in a way that makes it unclear whether a file system
// was created. Examples are if a transport level timeout occurred, or your
// connection was reset. If you use the same client request token and the initial
// call created a file system, the client receives success as long as the parameters
// are the same.
//
// The CreateFileSystem call returns while the file system's lifecycle state
// is still CREATING. You can check the file-system creation status by calling
// the DescribeFileSystems operation, which returns the file system state along
// with other information.
//
//    // Example sending a request using CreateFileSystemRequest.
//    req := client.CreateFileSystemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fsx-2018-03-01/CreateFileSystem
func (c *Client) CreateFileSystemRequest(input *CreateFileSystemInput) CreateFileSystemRequest {
	op := &aws.Operation{
		Name:       opCreateFileSystem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFileSystemInput{}
	}

	req := c.newRequest(op, input, &CreateFileSystemOutput{})
	return CreateFileSystemRequest{Request: req, Input: input, Copy: c.CreateFileSystemRequest}
}

// CreateFileSystemRequest is the request type for the
// CreateFileSystem API operation.
type CreateFileSystemRequest struct {
	*aws.Request
	Input *CreateFileSystemInput
	Copy  func(*CreateFileSystemInput) CreateFileSystemRequest
}

// Send marshals and sends the CreateFileSystem API request.
func (r CreateFileSystemRequest) Send(ctx context.Context) (*CreateFileSystemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFileSystemResponse{
		CreateFileSystemOutput: r.Request.Data.(*CreateFileSystemOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFileSystemResponse is the response type for the
// CreateFileSystem API operation.
type CreateFileSystemResponse struct {
	*CreateFileSystemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFileSystem request.
func (r *CreateFileSystemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
