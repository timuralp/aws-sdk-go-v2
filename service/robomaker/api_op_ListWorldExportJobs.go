// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists world export jobs.
func (c *Client) ListWorldExportJobs(ctx context.Context, params *ListWorldExportJobsInput, optFns ...func(*Options)) (*ListWorldExportJobsOutput, error) {
	if params == nil {
		params = &ListWorldExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorldExportJobs", params, optFns, addOperationListWorldExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorldExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorldExportJobsInput struct {

	// Optional filters to limit results. You can use generationJobId and templateId.
	Filters []*types.Filter

	// When this parameter is used, ListWorldExportJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListWorldExportJobs
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter is not used, then ListWorldExportJobs returns up to 100
	// results and a nextToken value if applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorldExportJobs again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
}

type ListWorldExportJobsOutput struct {

	// Summary information for world export jobs.
	//
	// This member is required.
	WorldExportJobSummaries []*types.WorldExportJobSummary

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorldExportJobsRequest again and assign that
	// token to the request object's nextToken parameter. If there are no remaining
	// results, the previous response object's NextToken parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWorldExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorldExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorldExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWorldExportJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListWorldExportJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListWorldExportJobs",
	}
}