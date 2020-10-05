// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets the retention settings for the specified Amazon Chime Enterprise account.
// For more information about retention settings, see Managing Chat Retention
// Policies (https://docs.aws.amazon.com/chime/latest/ag/chat-retention.html) in
// the Amazon Chime Administration Guide.
func (c *Client) GetRetentionSettings(ctx context.Context, params *GetRetentionSettingsInput, optFns ...func(*Options)) (*GetRetentionSettingsOutput, error) {
	stack := middleware.NewStack("GetRetentionSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetRetentionSettingsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetRetentionSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRetentionSettings(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetRetentionSettings",
			Err:           err,
		}
	}
	out := result.(*GetRetentionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRetentionSettingsInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string
}

type GetRetentionSettingsOutput struct {

	// The timestamp representing the time at which the specified items are permanently
	// deleted, in ISO 8601 format.
	InitiateDeletionTimestamp *time.Time

	// The retention settings.
	RetentionSettings *types.RetentionSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetRetentionSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetRetentionSettings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRetentionSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRetentionSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetRetentionSettings",
	}
}