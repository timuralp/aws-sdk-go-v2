// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to configure the custom MAIL FROM domain for a verified identity.
type PutEmailIdentityMailFromAttributesInput struct {
	_ struct{} `type:"structure"`

	// The action that you want to take if the required MX record isn't found when
	// you send an email. When you set this value to UseDefaultValue, the mail is
	// sent using amazonses.com as the MAIL FROM domain. When you set this value
	// to RejectMessage, the Amazon SES API v2 returns a MailFromDomainNotVerified
	// error, and doesn't attempt to deliver the email.
	//
	// These behaviors are taken when the custom MAIL FROM domain configuration
	// is in the Pending, Failed, and TemporaryFailure states.
	BehaviorOnMxFailure BehaviorOnMxFailure `type:"string" enum:"true"`

	// The verified email identity that you want to set up the custom MAIL FROM
	// domain for.
	//
	// EmailIdentity is a required field
	EmailIdentity *string `location:"uri" locationName:"EmailIdentity" type:"string" required:"true"`

	// The custom MAIL FROM domain that you want the verified identity to use. The
	// MAIL FROM domain must meet the following criteria:
	//
	//    * It has to be a subdomain of the verified identity.
	//
	//    * It can't be used to receive email.
	//
	//    * It can't be used in a "From" address if the MAIL FROM domain is a destination
	//    for feedback forwarding emails.
	MailFromDomain *string `type:"string"`
}

// String returns the string representation
func (s PutEmailIdentityMailFromAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEmailIdentityMailFromAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEmailIdentityMailFromAttributesInput"}

	if s.EmailIdentity == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailIdentity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEmailIdentityMailFromAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.BehaviorOnMxFailure) > 0 {
		v := s.BehaviorOnMxFailure

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BehaviorOnMxFailure", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.MailFromDomain != nil {
		v := *s.MailFromDomain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MailFromDomain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EmailIdentity != nil {
		v := *s.EmailIdentity

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EmailIdentity", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutEmailIdentityMailFromAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutEmailIdentityMailFromAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEmailIdentityMailFromAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutEmailIdentityMailFromAttributes = "PutEmailIdentityMailFromAttributes"

// PutEmailIdentityMailFromAttributesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Used to enable or disable the custom Mail-From domain configuration for an
// email identity.
//
//    // Example sending a request using PutEmailIdentityMailFromAttributesRequest.
//    req := client.PutEmailIdentityMailFromAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/PutEmailIdentityMailFromAttributes
func (c *Client) PutEmailIdentityMailFromAttributesRequest(input *PutEmailIdentityMailFromAttributesInput) PutEmailIdentityMailFromAttributesRequest {
	op := &aws.Operation{
		Name:       opPutEmailIdentityMailFromAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v2/email/identities/{EmailIdentity}/mail-from",
	}

	if input == nil {
		input = &PutEmailIdentityMailFromAttributesInput{}
	}

	req := c.newRequest(op, input, &PutEmailIdentityMailFromAttributesOutput{})
	return PutEmailIdentityMailFromAttributesRequest{Request: req, Input: input, Copy: c.PutEmailIdentityMailFromAttributesRequest}
}

// PutEmailIdentityMailFromAttributesRequest is the request type for the
// PutEmailIdentityMailFromAttributes API operation.
type PutEmailIdentityMailFromAttributesRequest struct {
	*aws.Request
	Input *PutEmailIdentityMailFromAttributesInput
	Copy  func(*PutEmailIdentityMailFromAttributesInput) PutEmailIdentityMailFromAttributesRequest
}

// Send marshals and sends the PutEmailIdentityMailFromAttributes API request.
func (r PutEmailIdentityMailFromAttributesRequest) Send(ctx context.Context) (*PutEmailIdentityMailFromAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEmailIdentityMailFromAttributesResponse{
		PutEmailIdentityMailFromAttributesOutput: r.Request.Data.(*PutEmailIdentityMailFromAttributesOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEmailIdentityMailFromAttributesResponse is the response type for the
// PutEmailIdentityMailFromAttributes API operation.
type PutEmailIdentityMailFromAttributesResponse struct {
	*PutEmailIdentityMailFromAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEmailIdentityMailFromAttributes request.
func (r *PutEmailIdentityMailFromAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
