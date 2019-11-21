// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetAttendeeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK attendee ID.
	//
	// AttendeeId is a required field
	AttendeeId *string `location:"uri" locationName:"attendeeId" type:"string" required:"true"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAttendeeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAttendeeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAttendeeInput"}

	if s.AttendeeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttendeeId"))
	}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAttendeeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AttendeeId != nil {
		v := *s.AttendeeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "attendeeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetAttendeeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK attendee information.
	Attendee *Attendee `type:"structure"`
}

// String returns the string representation
func (s GetAttendeeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetAttendeeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attendee != nil {
		v := s.Attendee

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Attendee", v, metadata)
	}
	return nil
}

const opGetAttendee = "GetAttendee"

// GetAttendeeRequest returns a request value for making API operation for
// Amazon Chime.
//
// Gets the Amazon Chime SDK attendee details for a specified meeting ID and
// attendee ID. For more information about the Amazon Chime SDK, see Using the
// Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide.
//
//    // Example sending a request using GetAttendeeRequest.
//    req := client.GetAttendeeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetAttendee
func (c *Client) GetAttendeeRequest(input *GetAttendeeInput) GetAttendeeRequest {
	op := &aws.Operation{
		Name:       opGetAttendee,
		HTTPMethod: "GET",
		HTTPPath:   "/meetings/{meetingId}/attendees/{attendeeId}",
	}

	if input == nil {
		input = &GetAttendeeInput{}
	}

	req := c.newRequest(op, input, &GetAttendeeOutput{})
	return GetAttendeeRequest{Request: req, Input: input, Copy: c.GetAttendeeRequest}
}

// GetAttendeeRequest is the request type for the
// GetAttendee API operation.
type GetAttendeeRequest struct {
	*aws.Request
	Input *GetAttendeeInput
	Copy  func(*GetAttendeeInput) GetAttendeeRequest
}

// Send marshals and sends the GetAttendee API request.
func (r GetAttendeeRequest) Send(ctx context.Context) (*GetAttendeeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAttendeeResponse{
		GetAttendeeOutput: r.Request.Data.(*GetAttendeeOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAttendeeResponse is the response type for the
// GetAttendee API operation.
type GetAttendeeResponse struct {
	*GetAttendeeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAttendee request.
func (r *GetAttendeeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
