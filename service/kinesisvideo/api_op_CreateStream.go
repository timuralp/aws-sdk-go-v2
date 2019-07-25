// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/CreateStreamInput
type CreateStreamInput struct {
	_ struct{} `type:"structure"`

	// The number of hours that you want to retain the data in the stream. Kinesis
	// Video Streams retains the data in a data store that is associated with the
	// stream.
	//
	// The default value is 0, indicating that the stream does not persist data.
	//
	// When the DataRetentionInHours value is 0, consumers can still consume the
	// fragments that remain in the service host buffer, which has a retention time
	// limit of 5 minutes and a retention memory limit of 200 MB. Fragments are
	// removed from the buffer when either limit is reached.
	DataRetentionInHours *int64 `type:"integer"`

	// The name of the device that is writing to the stream.
	//
	// In the current implementation, Kinesis Video Streams does not use this name.
	DeviceName *string `min:"1" type:"string"`

	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis
	// Video Streams to use to encrypt stream data.
	//
	// If no key ID is specified, the default, Kinesis Video-managed key (aws/kinesisvideo)
	// is used.
	//
	// For more information, see DescribeKey (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters).
	KmsKeyId *string `min:"1" type:"string"`

	// The media type of the stream. Consumers of the stream can use this information
	// when processing the stream. For more information about media types, see Media
	// Types (http://www.iana.org/assignments/media-types/media-types.xhtml). If
	// you choose to specify the MediaType, see Naming Requirements (https://tools.ietf.org/html/rfc6838#section-4.2)
	// for guidelines.
	//
	// This parameter is optional; the default value is null (or empty in JSON).
	MediaType *string `min:"1" type:"string"`

	// A name for the stream that you are creating.
	//
	// The stream name is an identifier for the stream, and must be unique for each
	// account and region.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`

	// A list of tags to associate with the specified stream. Each tag is a key-value
	// pair (the value is optional).
	Tags map[string]string `min:"1" type:"map"`
}

// String returns the string representation
func (s CreateStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStreamInput"}
	if s.DeviceName != nil && len(*s.DeviceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceName", 1))
	}
	if s.KmsKeyId != nil && len(*s.KmsKeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KmsKeyId", 1))
	}
	if s.MediaType != nil && len(*s.MediaType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MediaType", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateStreamInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DataRetentionInHours != nil {
		v := *s.DataRetentionInHours

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataRetentionInHours", protocol.Int64Value(v), metadata)
	}
	if s.DeviceName != nil {
		v := *s.DeviceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeviceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KmsKeyId != nil {
		v := *s.KmsKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KmsKeyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MediaType != nil {
		v := *s.MediaType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MediaType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/CreateStreamOutput
type CreateStreamOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the stream.
	StreamARN *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateStreamOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.StreamARN != nil {
		v := *s.StreamARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateStream = "CreateStream"

// CreateStreamRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams.
//
// Creates a new Kinesis video stream.
//
// When you create a new stream, Kinesis Video Streams assigns it a version
// number. When you change the stream's metadata, Kinesis Video Streams updates
// the version.
//
// CreateStream is an asynchronous operation.
//
// For information about how the service works, see How it Works (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/how-it-works.html).
//
// You must have permissions for the KinesisVideo:CreateStream action.
//
//    // Example sending a request using CreateStreamRequest.
//    req := client.CreateStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisvideo-2017-09-30/CreateStream
func (c *Client) CreateStreamRequest(input *CreateStreamInput) CreateStreamRequest {
	op := &aws.Operation{
		Name:       opCreateStream,
		HTTPMethod: "POST",
		HTTPPath:   "/createStream",
	}

	if input == nil {
		input = &CreateStreamInput{}
	}

	req := c.newRequest(op, input, &CreateStreamOutput{})
	return CreateStreamRequest{Request: req, Input: input, Copy: c.CreateStreamRequest}
}

// CreateStreamRequest is the request type for the
// CreateStream API operation.
type CreateStreamRequest struct {
	*aws.Request
	Input *CreateStreamInput
	Copy  func(*CreateStreamInput) CreateStreamRequest
}

// Send marshals and sends the CreateStream API request.
func (r CreateStreamRequest) Send(ctx context.Context) (*CreateStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStreamResponse{
		CreateStreamOutput: r.Request.Data.(*CreateStreamOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStreamResponse is the response type for the
// CreateStream API operation.
type CreateStreamResponse struct {
	*CreateStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStream request.
func (r *CreateStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}