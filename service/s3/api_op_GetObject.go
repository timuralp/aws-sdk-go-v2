// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetObjectInput struct {
	_ struct{} `type:"structure"`

	// The bucket name containing the object.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Return the object only if its entity tag (ETag) is the same as the one specified,
	// otherwise return a 412 (precondition failed).
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`

	// Return the object only if it has been modified since the specified time,
	// otherwise return a 304 (not modified).
	IfModifiedSince *time.Time `location:"header" locationName:"If-Modified-Since" type:"timestamp"`

	// Return the object only if its entity tag (ETag) is different from the one
	// specified, otherwise return a 304 (not modified).
	IfNoneMatch *string `location:"header" locationName:"If-None-Match" type:"string"`

	// Return the object only if it has not been modified since the specified time,
	// otherwise return a 412 (precondition failed).
	IfUnmodifiedSince *time.Time `location:"header" locationName:"If-Unmodified-Since" type:"timestamp"`

	// Key of the object to get.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Part number of the object being read. This is a positive integer between
	// 1 and 10,000. Effectively performs a 'ranged' GET request for the part specified.
	// Useful for downloading just a part of an object.
	PartNumber *int64 `location:"querystring" locationName:"partNumber" type:"integer"`

	// Downloads the specified range bytes of an object. For more information about
	// the HTTP Range header, go to http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35.
	Range *string `location:"header" locationName:"Range" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Sets the Cache-Control header of the response.
	ResponseCacheControl *string `location:"querystring" locationName:"response-cache-control" type:"string"`

	// Sets the Content-Disposition header of the response
	ResponseContentDisposition *string `location:"querystring" locationName:"response-content-disposition" type:"string"`

	// Sets the Content-Encoding header of the response.
	ResponseContentEncoding *string `location:"querystring" locationName:"response-content-encoding" type:"string"`

	// Sets the Content-Language header of the response.
	ResponseContentLanguage *string `location:"querystring" locationName:"response-content-language" type:"string"`

	// Sets the Content-Type header of the response.
	ResponseContentType *string `location:"querystring" locationName:"response-content-type" type:"string"`

	// Sets the Expires header of the response.
	ResponseExpires *time.Time `location:"querystring" locationName:"response-expires" type:"timestamp"`

	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// does not store the encryption key. The key must be appropriate for use with
	// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// VersionId used to reference a specific version of the object.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *GetObjectInput) getSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IfMatch != nil {
		v := *s.IfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Match", protocol.StringValue(v), metadata)
	}
	if s.IfModifiedSince != nil {
		v := *s.IfModifiedSince

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Modified-Since",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.IfNoneMatch != nil {
		v := *s.IfNoneMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-None-Match", protocol.StringValue(v), metadata)
	}
	if s.IfUnmodifiedSince != nil {
		v := *s.IfUnmodifiedSince

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Unmodified-Since",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.Range != nil {
		v := *s.Range

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Range", protocol.StringValue(v), metadata)
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
	if s.PartNumber != nil {
		v := *s.PartNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "partNumber", protocol.Int64Value(v), metadata)
	}
	if s.ResponseCacheControl != nil {
		v := *s.ResponseCacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-cache-control", protocol.StringValue(v), metadata)
	}
	if s.ResponseContentDisposition != nil {
		v := *s.ResponseContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-content-disposition", protocol.StringValue(v), metadata)
	}
	if s.ResponseContentEncoding != nil {
		v := *s.ResponseContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-content-encoding", protocol.StringValue(v), metadata)
	}
	if s.ResponseContentLanguage != nil {
		v := *s.ResponseContentLanguage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-content-language", protocol.StringValue(v), metadata)
	}
	if s.ResponseContentType != nil {
		v := *s.ResponseContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-content-type", protocol.StringValue(v), metadata)
	}
	if s.ResponseExpires != nil {
		v := *s.ResponseExpires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "response-expires",
			protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetObjectOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// Indicates that a range of bytes was specifed.
	AcceptRanges *string `location:"header" locationName:"accept-ranges" type:"string"`

	// Object data.
	Body io.ReadCloser `type:"blob"`

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// Size of the body in bytes.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`

	// The portion of the object returned in the response.
	ContentRange *string `location:"header" locationName:"Content-Range" type:"string"`

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool `location:"header" locationName:"x-amz-delete-marker" type:"boolean"`

	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *string `location:"header" locationName:"Expires" type:"string"`

	// Last modified date of the object
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP,
	// you can create metadata whose values are not legal HTTP headers.
	MissingMeta *int64 `location:"header" locationName:"x-amz-missing-meta" type:"integer"`

	// Indicates whether this object has an active legal hold. This field is only
	// returned if you have permission to view an object's legal hold status.
	ObjectLockLegalHoldStatus ObjectLockLegalHoldStatus `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"true"`

	// The Object Lock mode currently in place for this object.
	ObjectLockMode ObjectLockMode `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"true"`

	// The date and time when this object's Object Lock will expire.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// The count of parts this object has.
	PartsCount *int64 `location:"header" locationName:"x-amz-mp-parts-count" type:"integer"`

	// Amazon S3 can return this if your request involves a bucket that is either
	// a source or destination in a replication rule.
	ReplicationStatus ReplicationStatus `location:"header" locationName:"x-amz-replication-status" type:"string" enum:"true"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// Provides information about object restoration operation and expiration time
	// of the restored object copy.
	Restore *string `location:"header" locationName:"x-amz-restore" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) customer
	// master key (CMK) that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// Provides storage class information of the object. Amazon S3 returns this
	// header for all objects except for Standard storage class objects.
	StorageClass StorageClass `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"true"`

	// The number of tags, if any, on the object.
	TagCount *int64 `location:"header" locationName:"x-amz-tagging-count" type:"integer"`

	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s GetObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AcceptRanges != nil {
		v := *s.AcceptRanges

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "accept-ranges", protocol.StringValue(v), metadata)
	}
	if s.CacheControl != nil {
		v := *s.CacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Cache-Control", protocol.StringValue(v), metadata)
	}
	if s.ContentDisposition != nil {
		v := *s.ContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Disposition", protocol.StringValue(v), metadata)
	}
	if s.ContentEncoding != nil {
		v := *s.ContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Encoding", protocol.StringValue(v), metadata)
	}
	if s.ContentLanguage != nil {
		v := *s.ContentLanguage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Language", protocol.StringValue(v), metadata)
	}
	if s.ContentLength != nil {
		v := *s.ContentLength

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Length", protocol.Int64Value(v), metadata)
	}
	if s.ContentRange != nil {
		v := *s.ContentRange

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Range", protocol.StringValue(v), metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue(v), metadata)
	}
	if s.DeleteMarker != nil {
		v := *s.DeleteMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-delete-marker", protocol.BoolValue(v), metadata)
	}
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Expiration != nil {
		v := *s.Expiration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-expiration", protocol.StringValue(v), metadata)
	}
	if s.Expires != nil {
		v := *s.Expires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Expires", protocol.StringValue(v), metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Last-Modified",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.MissingMeta != nil {
		v := *s.MissingMeta

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-missing-meta", protocol.Int64Value(v), metadata)
	}
	if len(s.ObjectLockLegalHoldStatus) > 0 {
		v := s.ObjectLockLegalHoldStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-legal-hold", v, metadata)
	}
	if len(s.ObjectLockMode) > 0 {
		v := s.ObjectLockMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-mode", v, metadata)
	}
	if s.ObjectLockRetainUntilDate != nil {
		v := *s.ObjectLockRetainUntilDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-retain-until-date",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	if s.PartsCount != nil {
		v := *s.PartsCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-mp-parts-count", protocol.Int64Value(v), metadata)
	}
	if len(s.ReplicationStatus) > 0 {
		v := s.ReplicationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-replication-status", v, metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.Restore != nil {
		v := *s.Restore

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-restore", protocol.StringValue(v), metadata)
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
	if len(s.StorageClass) > 0 {
		v := s.StorageClass

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-storage-class", v, metadata)
	}
	if s.TagCount != nil {
		v := *s.TagCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-tagging-count", protocol.Int64Value(v), metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	if s.WebsiteRedirectLocation != nil {
		v := *s.WebsiteRedirectLocation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-website-redirect-location", protocol.StringValue(v), metadata)
	}
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.HeadersTarget, "x-amz-meta-", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.StringValue(v1))
		}
		ms0.End()

	}
	// Skipping Body Output type's body not valid.
	return nil
}

const opGetObject = "GetObject"

// GetObjectRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Retrieves objects from Amazon S3. To use GET, you must have READ access to
// the object. If you grant READ access to the anonymous user, you can return
// the object without using an authorization header.
//
// An Amazon S3 bucket has no directory hierarchy such as you would find in
// a typical computer file system. You can, however, create a logical hierarchy
// by using object key names that imply a folder structure. For example, instead
// of naming an object sample.jpg, you can name it photos/2006/February/sample.jpg.
//
// To get an object from such a logical hierarchy, specify the full key name
// for the object in the GET operation. For a virtual hosted-style request example,
// if you have the object photos/2006/February/sample.jpg, specify the resource
// as /photos/2006/February/sample.jpg. For a path-style request example, if
// you have the object photos/2006/February/sample.jpg in the bucket named examplebucket,
// specify the resource as /examplebucket/photos/2006/February/sample.jpg. For
// more information about request types, see HTTP Host Header Bucket Specification
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingSpecifyBucket).
//
// To distribute large files to many people, you can save bandwidth costs by
// using BitTorrent. For more information, see Amazon S3 Torrent (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3Torrent.html).
// For more information about returning the ACL of an object, see GetObjectAcl.
//
// If the object you are retrieving is stored in the GLACIER or DEEP_ARCHIVE
// storage classes, before you can retrieve the object you must first restore
// a copy using . Otherwise, this operation returns an InvalidObjectStateError
// error. For information about restoring archived objects, see Restoring Archived
// Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html).
//
// Encryption request headers, like x-amz-server-side-encryption, should not
// be sent for GET requests if your object uses server-side encryption with
// CMKs stored in AWS KMS (SSE-KMS) or server-side encryption with Amazon S3–managed
// encryption keys (SSE-S3). If your object does use these types of keys, you’ll
// get an HTTP 400 BadRequest error.
//
// If you encrypt an object by using server-side encryption with customer-provided
// encryption keys (SSE-C) when you store the object in Amazon S3, then when
// you GET the object, you must use the following headers:
//
//    * x-amz-server-side​-encryption​-customer-algorithm
//
//    * x-amz-server-side​-encryption​-customer-key
//
//    * x-amz-server-side​-encryption​-customer-key-MD5
//
// For more information about SSE-C, see Server-Side Encryption (Using Customer-Provided
// Encryption Keys) (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html).
//
// Assuming you have permission to read object tags (permission for the s3:GetObjectVersionTagging
// action), the response also returns the x-amz-tagging-count header that provides
// the count of number of tags associated with the object. You can use GetObjectTagging
// to retrieve the tag set associated with an object.
//
// Permissions
//
// You need the s3:GetObject permission for this operation. For more information,
// see Specifying Permissions in a Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html).
// If the object you request does not exist, the error Amazon S3 returns depends
// on whether you also have the s3:ListBucket permission.
//
//    * If you have the s3:ListBucket permission on the bucket, Amazon S3 will
//    return an HTTP status code 404 ("no such key") error.
//
//    * If you don’t have the s3:ListBucket permission, Amazon S3 will return
//    an HTTP status code 403 ("access denied") error.
//
// Versioning
//
// By default, the GET operation returns the current version of an object. To
// return a different version, use the versionId subresource.
//
// If the current version of the object is a delete marker, Amazon S3 behaves
// as if the object was deleted and includes x-amz-delete-marker: true in the
// response.
//
// For more information about versioning, see PutBucketVersioning.
//
// Overriding Response Header Values
//
// There are times when you want to override certain response header values
// in a GET response. For example, you might override the Content-Disposition
// response header value in your GET request.
//
// You can override values for a set of response headers using the following
// query parameters. These response header values are sent only on a successful
// request, that is, when status code 200 OK is returned. The set of headers
// you can override using these parameters is a subset of the headers that Amazon
// S3 accepts when you create an object. The response headers that you can override
// for the GET response are Content-Type, Content-Language, Expires, Cache-Control,
// Content-Disposition, and Content-Encoding. To override these header values
// in the GET response, you use the following request parameters.
//
// You must sign the request, either using an Authorization header or a presigned
// URL, when using these parameters. They cannot be used with an unsigned (anonymous)
// request.
//
//    * response-content-type
//
//    * response-content-language
//
//    * response-expires
//
//    * response-cache-control
//
//    * response-content-disposition
//
//    * response-content-encoding
//
// Additional Considerations about Request Headers
//
// If both of the If-Match and If-Unmodified-Since headers are present in the
// request as follows: If-Match condition evaluates to true, and; If-Unmodified-Since
// condition evaluates to false; then, S3 returns 200 OK and the data requested.
//
// If both of the If-None-Match and If-Modified-Since headers are present in
// the request as follows:If-None-Match condition evaluates to false, and; If-Modified-Since
// condition evaluates to true; then, S3 returns 304 Not Modified response code.
//
// For more information about conditional requests, see RFC 7232 (https://tools.ietf.org/html/rfc7232).
//
// The following operations are related to GetObject:
//
//    * ListBuckets
//
//    * GetObjectAcl
//
//    // Example sending a request using GetObjectRequest.
//    req := client.GetObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObject
func (c *Client) GetObjectRequest(input *GetObjectInput) GetObjectRequest {
	op := &aws.Operation{
		Name:       opGetObject,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &GetObjectInput{}
	}

	req := c.newRequest(op, input, &GetObjectOutput{})
	return GetObjectRequest{Request: req, Input: input, Copy: c.GetObjectRequest}
}

// GetObjectRequest is the request type for the
// GetObject API operation.
type GetObjectRequest struct {
	*aws.Request
	Input *GetObjectInput
	Copy  func(*GetObjectInput) GetObjectRequest
}

// Send marshals and sends the GetObject API request.
func (r GetObjectRequest) Send(ctx context.Context) (*GetObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectResponse{
		GetObjectOutput: r.Request.Data.(*GetObjectOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectResponse is the response type for the
// GetObject API operation.
type GetObjectResponse struct {
	*GetObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObject request.
func (r *GetObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
