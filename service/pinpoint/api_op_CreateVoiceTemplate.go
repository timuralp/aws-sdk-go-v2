// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateVoiceTemplateInput struct {
	_ struct{} `type:"structure" payload:"VoiceTemplateRequest"`

	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"template-name" type:"string" required:"true"`

	// Specifies the content and settings for a message template that can be used
	// in messages that are sent through the voice channel.
	//
	// VoiceTemplateRequest is a required field
	VoiceTemplateRequest *VoiceTemplateRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVoiceTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVoiceTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVoiceTemplateInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if s.VoiceTemplateRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceTemplateRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVoiceTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "template-name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VoiceTemplateRequest != nil {
		v := s.VoiceTemplateRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "VoiceTemplateRequest", v, metadata)
	}
	return nil
}

type CreateVoiceTemplateOutput struct {
	_ struct{} `type:"structure" payload:"CreateTemplateMessageBody"`

	// Provides information about a request to create a message template.
	//
	// CreateTemplateMessageBody is a required field
	CreateTemplateMessageBody *CreateTemplateMessageBody `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVoiceTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateVoiceTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreateTemplateMessageBody != nil {
		v := s.CreateTemplateMessageBody

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CreateTemplateMessageBody", v, metadata)
	}
	return nil
}

const opCreateVoiceTemplate = "CreateVoiceTemplate"

// CreateVoiceTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Creates a message template that you can use in messages that are sent through
// the voice channel.
//
//    // Example sending a request using CreateVoiceTemplateRequest.
//    req := client.CreateVoiceTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/CreateVoiceTemplate
func (c *Client) CreateVoiceTemplateRequest(input *CreateVoiceTemplateInput) CreateVoiceTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateVoiceTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/templates/{template-name}/voice",
	}

	if input == nil {
		input = &CreateVoiceTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateVoiceTemplateOutput{})
	return CreateVoiceTemplateRequest{Request: req, Input: input, Copy: c.CreateVoiceTemplateRequest}
}

// CreateVoiceTemplateRequest is the request type for the
// CreateVoiceTemplate API operation.
type CreateVoiceTemplateRequest struct {
	*aws.Request
	Input *CreateVoiceTemplateInput
	Copy  func(*CreateVoiceTemplateInput) CreateVoiceTemplateRequest
}

// Send marshals and sends the CreateVoiceTemplate API request.
func (r CreateVoiceTemplateRequest) Send(ctx context.Context) (*CreateVoiceTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVoiceTemplateResponse{
		CreateVoiceTemplateOutput: r.Request.Data.(*CreateVoiceTemplateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVoiceTemplateResponse is the response type for the
// CreateVoiceTemplate API operation.
type CreateVoiceTemplateResponse struct {
	*CreateVoiceTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVoiceTemplate request.
func (r *CreateVoiceTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
