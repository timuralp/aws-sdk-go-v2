// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/UploadPartRequest
type UploadPartInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Object data.
	Body io.ReadSeeker `type:"blob"`

	// Name of the bucket to which the multipart upload was initiated.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Size of the body in bytes. This parameter is useful when the size of the
	// body cannot be determined automatically.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The base64-encoded 128-bit MD5 digest of the part data. This parameter is
	// auto-populated when using the command from the CLI. This parameted is required
	// if object lock parameters are specified.
	ContentMD5 *string `location:"header" locationName:"Content-MD5" type:"string"`

	// Object key for which the multipart upload was initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Part number of part being uploaded. This is a positive integer between 1
	// and 10,000.
	//
	// PartNumber is a required field
	PartNumber *int64 `location:"querystring" locationName:"partNumber" type:"integer" required:"true"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// does not store the encryption key. The key must be appropriate for use with
	// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header. This must be the same encryption key specified in the initiate multipart
	// upload request.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// Upload ID identifying the multipart upload whose part is being uploaded.
	//
	// UploadId is a required field
	UploadId *string `location:"querystring" locationName:"uploadId" type:"string" required:"true"`
}

// String returns the string representation
func (s UploadPartInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadPartInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadPartInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.PartNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartNumber"))
	}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *UploadPartInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *UploadPartInput) getSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadPartInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ContentLength != nil {
		v := *s.ContentLength

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Length", protocol.Int64Value(v), metadata)
	}
	if s.ContentMD5 != nil {
		v := *s.ContentMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-MD5", protocol.StringValue(v), metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKey != nil {
		v := *s.SSECustomerKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.Body != nil {
		v := s.Body

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Body", protocol.ReadSeekerStream{V: v}, metadata)
	}
	if s.PartNumber != nil {
		v := *s.PartNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "partNumber", protocol.Int64Value(v), metadata)
	}
	if s.UploadId != nil {
		v := *s.UploadId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "uploadId", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/UploadPartOutput
type UploadPartOutput struct {
	_ struct{} `type:"structure"`

	// Entity tag for the uploaded object.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`
}

// String returns the string representation
func (s UploadPartOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UploadPartOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	return nil
}

const opUploadPart = "UploadPart"

// UploadPartRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Uploads a part in a multipart upload.
//
// Note: After you initiate multipart upload and upload one or more parts, you
// must either complete or abort multipart upload in order to stop getting charged
// for storage of the uploaded parts. Only after you either complete or abort
// multipart upload, Amazon S3 frees up the parts storage and stops charging
// you for the parts storage.
//
//    // Example sending a request using UploadPartRequest.
//    req := client.UploadPartRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/UploadPart
func (c *Client) UploadPartRequest(input *UploadPartInput) UploadPartRequest {
	op := &aws.Operation{
		Name:       opUploadPart,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &UploadPartInput{}
	}

	req := c.newRequest(op, input, &UploadPartOutput{})
	return UploadPartRequest{Request: req, Input: input, Copy: c.UploadPartRequest}
}

// UploadPartRequest is the request type for the
// UploadPart API operation.
type UploadPartRequest struct {
	*aws.Request
	Input *UploadPartInput
	Copy  func(*UploadPartInput) UploadPartRequest
}

// Send marshals and sends the UploadPart API request.
func (r UploadPartRequest) Send(ctx context.Context) (*UploadPartResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UploadPartResponse{
		UploadPartOutput: r.Request.Data.(*UploadPartOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UploadPartResponse is the response type for the
// UploadPart API operation.
type UploadPartResponse struct {
	*UploadPartOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UploadPart request.
func (r *UploadPartResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}