// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type RegisterUserInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the user is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The email address of the user that you want to register.
	//
	// Email is a required field
	Email *string `type:"string" required:"true"`

	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn *string `type:"string"`

	// Amazon QuickSight supports several ways of managing the identity of users.
	// This parameter accepts two values:
	//
	//    * IAM: A user whose identity maps to an existing IAM user or role.
	//
	//    * QUICKSIGHT: A user whose identity is owned and managed internally by
	//    Amazon QuickSight.
	//
	// IdentityType is a required field
	IdentityType IdentityType `type:"string" required:"true" enum:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`

	// You need to use this parameter only when you register one or more users using
	// an assumed IAM role. You don't need to provide the session name for other
	// scenarios, for example when you are registering an IAM user or an Amazon
	// QuickSight user. You can register multiple users using the same IAM role
	// if each user has a different session name. For more information on assuming
	// IAM roles, see assume-role (https://docs.aws.example.com/cli/latest/reference/sts/assume-role.html)
	// in the AWS CLI Reference.
	SessionName *string `min:"2" type:"string"`

	// The Amazon QuickSight user name that you want to create for the user you
	// are registering.
	UserName *string `min:"1" type:"string"`

	// The Amazon QuickSight role for the user. The user role can be one of the
	// following:
	//
	//    * READER: A user who has read-only access to dashboards.
	//
	//    * AUTHOR: A user who can create data sources, datasets, analyses, and
	//    dashboards.
	//
	//    * ADMIN: A user who is an author, who can also manage Amazon QuickSight
	//    settings.
	//
	//    * RESTRICTED_READER: This role isn't currently available for use.
	//
	//    * RESTRICTED_AUTHOR: This role isn't currently available for use.
	//
	// UserRole is a required field
	UserRole UserRole `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s RegisterUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterUserInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.Email == nil {
		invalidParams.Add(aws.NewErrParamRequired("Email"))
	}
	if len(s.IdentityType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("IdentityType"))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}
	if s.SessionName != nil && len(*s.SessionName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("SessionName", 2))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}
	if len(s.UserRole) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("UserRole"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Email != nil {
		v := *s.Email

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Email", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IamArn != nil {
		v := *s.IamArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IamArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.IdentityType) > 0 {
		v := s.IdentityType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.SessionName != nil {
		v := *s.SessionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SessionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserName != nil {
		v := *s.UserName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UserName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.UserRole) > 0 {
		v := s.UserRole

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UserRole", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RegisterUserOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The user name.
	User *User `type:"structure"`

	// The URL the user visits to complete registration and provide a password.
	// This is returned only for users with an identity type of QUICKSIGHT.
	UserInvitationUrl *string `type:"string"`
}

// String returns the string representation
func (s RegisterUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RegisterUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.User != nil {
		v := s.User

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "User", v, metadata)
	}
	if s.UserInvitationUrl != nil {
		v := *s.UserInvitationUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UserInvitationUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opRegisterUser = "RegisterUser"

// RegisterUserRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Creates an Amazon QuickSight user, whose identity is associated with the
// AWS Identity and Access Management (IAM) identity or role specified in the
// request.
//
// CLI Sample:
//
// aws quicksight register-user -\-aws-account-id=111122223333 -\-namespace=default
// -\-email=pat@example.com -\-identity-type=IAM -\-user-role=AUTHOR -\-iam-arn=arn:aws:iam::111122223333:user/Pat
//
//    // Example sending a request using RegisterUserRequest.
//    req := client.RegisterUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/RegisterUser
func (c *Client) RegisterUserRequest(input *RegisterUserInput) RegisterUserRequest {
	op := &aws.Operation{
		Name:       opRegisterUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users",
	}

	if input == nil {
		input = &RegisterUserInput{}
	}

	req := c.newRequest(op, input, &RegisterUserOutput{})
	return RegisterUserRequest{Request: req, Input: input, Copy: c.RegisterUserRequest}
}

// RegisterUserRequest is the request type for the
// RegisterUser API operation.
type RegisterUserRequest struct {
	*aws.Request
	Input *RegisterUserInput
	Copy  func(*RegisterUserInput) RegisterUserRequest
}

// Send marshals and sends the RegisterUser API request.
func (r RegisterUserRequest) Send(ctx context.Context) (*RegisterUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterUserResponse{
		RegisterUserOutput: r.Request.Data.(*RegisterUserOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterUserResponse is the response type for the
// RegisterUser API operation.
type RegisterUserResponse struct {
	*RegisterUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterUser request.
func (r *RegisterUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
