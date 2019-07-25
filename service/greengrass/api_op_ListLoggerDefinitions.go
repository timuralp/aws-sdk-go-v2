// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListLoggerDefinitionsRequest
type ListLoggerDefinitionsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *string `location:"querystring" locationName:"MaxResults" type:"string"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListLoggerDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLoggerDefinitionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListLoggerDefinitionsResponse
type ListLoggerDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	Definitions []DefinitionInformation `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListLoggerDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListLoggerDefinitionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Definitions != nil {
		v := s.Definitions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Definitions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListLoggerDefinitions = "ListLoggerDefinitions"

// ListLoggerDefinitionsRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves a list of logger definitions.
//
//    // Example sending a request using ListLoggerDefinitionsRequest.
//    req := client.ListLoggerDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/ListLoggerDefinitions
func (c *Client) ListLoggerDefinitionsRequest(input *ListLoggerDefinitionsInput) ListLoggerDefinitionsRequest {
	op := &aws.Operation{
		Name:       opListLoggerDefinitions,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/loggers",
	}

	if input == nil {
		input = &ListLoggerDefinitionsInput{}
	}

	req := c.newRequest(op, input, &ListLoggerDefinitionsOutput{})
	return ListLoggerDefinitionsRequest{Request: req, Input: input, Copy: c.ListLoggerDefinitionsRequest}
}

// ListLoggerDefinitionsRequest is the request type for the
// ListLoggerDefinitions API operation.
type ListLoggerDefinitionsRequest struct {
	*aws.Request
	Input *ListLoggerDefinitionsInput
	Copy  func(*ListLoggerDefinitionsInput) ListLoggerDefinitionsRequest
}

// Send marshals and sends the ListLoggerDefinitions API request.
func (r ListLoggerDefinitionsRequest) Send(ctx context.Context) (*ListLoggerDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListLoggerDefinitionsResponse{
		ListLoggerDefinitionsOutput: r.Request.Data.(*ListLoggerDefinitionsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListLoggerDefinitionsResponse is the response type for the
// ListLoggerDefinitions API operation.
type ListLoggerDefinitionsResponse struct {
	*ListLoggerDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListLoggerDefinitions request.
func (r *ListLoggerDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}