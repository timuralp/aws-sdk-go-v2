// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to update the user pool.
type UpdateUserPoolInput struct {
	_ struct{} `type:"structure"`

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *AdminCreateUserConfigType `type:"structure"`

	// The attributes that are automatically verified when the Amazon Cognito service
	// makes a request to update user pools.
	AutoVerifiedAttributes []VerifiedAttributeType `type:"list"`

	// Device configuration.
	DeviceConfiguration *DeviceConfigurationType `type:"structure"`

	// Email configuration.
	EmailConfiguration *EmailConfigurationType `type:"structure"`

	// The contents of the email verification message.
	EmailVerificationMessage *string `min:"6" type:"string"`

	// The subject of the email verification message.
	EmailVerificationSubject *string `min:"1" type:"string"`

	// The AWS Lambda configuration information from the request to update the user
	// pool.
	LambdaConfig *LambdaConfigType `type:"structure"`

	// Can be one of the following values:
	//
	//    * OFF - MFA tokens are not required and cannot be specified during user
	//    registration.
	//
	//    * ON - MFA tokens are required for all user registrations. You can only
	//    specify required when you are initially creating a user pool.
	//
	//    * OPTIONAL - Users have the option when registering to create an MFA token.
	MfaConfiguration UserPoolMfaType `type:"string" enum:"true"`

	// A container with the policies you wish to update in a user pool.
	Policies *UserPoolPolicyType `type:"structure"`

	// The contents of the SMS authentication message.
	SmsAuthenticationMessage *string `min:"6" type:"string"`

	// SMS configuration.
	SmsConfiguration *SmsConfigurationType `type:"structure"`

	// A container with information about the SMS verification message.
	SmsVerificationMessage *string `min:"6" type:"string"`

	// Used to enable advanced security risk detection. Set the key AdvancedSecurityMode
	// to the value "AUDIT".
	UserPoolAddOns *UserPoolAddOnsType `type:"structure"`

	// The user pool ID for the user pool you want to update.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The tag keys and values to assign to the user pool. A tag is a label that
	// you can use to categorize and manage user pools in different ways, such as
	// by purpose, owner, environment, or other criteria.
	UserPoolTags map[string]string `type:"map"`

	// The template for verification messages.
	VerificationMessageTemplate *VerificationMessageTemplateType `type:"structure"`
}

// String returns the string representation
func (s UpdateUserPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserPoolInput"}
	if s.EmailVerificationMessage != nil && len(*s.EmailVerificationMessage) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("EmailVerificationMessage", 6))
	}
	if s.EmailVerificationSubject != nil && len(*s.EmailVerificationSubject) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EmailVerificationSubject", 1))
	}
	if s.SmsAuthenticationMessage != nil && len(*s.SmsAuthenticationMessage) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("SmsAuthenticationMessage", 6))
	}
	if s.SmsVerificationMessage != nil && len(*s.SmsVerificationMessage) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("SmsVerificationMessage", 6))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.AdminCreateUserConfig != nil {
		if err := s.AdminCreateUserConfig.Validate(); err != nil {
			invalidParams.AddNested("AdminCreateUserConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.EmailConfiguration != nil {
		if err := s.EmailConfiguration.Validate(); err != nil {
			invalidParams.AddNested("EmailConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.LambdaConfig != nil {
		if err := s.LambdaConfig.Validate(); err != nil {
			invalidParams.AddNested("LambdaConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			invalidParams.AddNested("Policies", err.(aws.ErrInvalidParams))
		}
	}
	if s.SmsConfiguration != nil {
		if err := s.SmsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("SmsConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.UserPoolAddOns != nil {
		if err := s.UserPoolAddOns.Validate(); err != nil {
			invalidParams.AddNested("UserPoolAddOns", err.(aws.ErrInvalidParams))
		}
	}
	if s.VerificationMessageTemplate != nil {
		if err := s.VerificationMessageTemplate.Validate(); err != nil {
			invalidParams.AddNested("VerificationMessageTemplate", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server when you make a request to update
// the user pool.
type UpdateUserPoolOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateUserPoolOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateUserPool = "UpdateUserPool"

// UpdateUserPoolRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Updates the specified user pool with the specified attributes. You can get
// a list of the current user pool settings with .
//
// If you don't provide a value for an attribute, it will be set to the default
// value.
//
//    // Example sending a request using UpdateUserPoolRequest.
//    req := client.UpdateUserPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/UpdateUserPool
func (c *Client) UpdateUserPoolRequest(input *UpdateUserPoolInput) UpdateUserPoolRequest {
	op := &aws.Operation{
		Name:       opUpdateUserPool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateUserPoolInput{}
	}

	req := c.newRequest(op, input, &UpdateUserPoolOutput{})
	return UpdateUserPoolRequest{Request: req, Input: input, Copy: c.UpdateUserPoolRequest}
}

// UpdateUserPoolRequest is the request type for the
// UpdateUserPool API operation.
type UpdateUserPoolRequest struct {
	*aws.Request
	Input *UpdateUserPoolInput
	Copy  func(*UpdateUserPoolInput) UpdateUserPoolRequest
}

// Send marshals and sends the UpdateUserPool API request.
func (r UpdateUserPoolRequest) Send(ctx context.Context) (*UpdateUserPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserPoolResponse{
		UpdateUserPoolOutput: r.Request.Data.(*UpdateUserPoolOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserPoolResponse is the response type for the
// UpdateUserPool API operation.
type UpdateUserPoolResponse struct {
	*UpdateUserPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserPool request.
func (r *UpdateUserPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
