// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDisassociateApprovalRuleTemplateFromRepositoriesInput struct {
	_ struct{} `type:"structure"`

	// The name of the template that you want to disassociate from one or more repositories.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`

	// The repository names that you want to disassociate from the approval rule
	// template.
	//
	// The length constraint limit is for each string in the array. The array itself
	// can be empty.
	//
	// RepositoryNames is a required field
	RepositoryNames []string `locationName:"repositoryNames" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDisassociateApprovalRuleTemplateFromRepositoriesInput"}

	if s.ApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateName"))
	}
	if s.ApprovalRuleTemplateName != nil && len(*s.ApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateName", 1))
	}

	if s.RepositoryNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput struct {
	_ struct{} `type:"structure"`

	// A list of repository names that have had their association with the template
	// removed.
	//
	// DisassociatedRepositoryNames is a required field
	DisassociatedRepositoryNames []string `locationName:"disassociatedRepositoryNames" type:"list" required:"true"`

	// A list of any errors that might have occurred while attempting to remove
	// the association between the template and the repositories.
	//
	// Errors is a required field
	Errors []BatchDisassociateApprovalRuleTemplateFromRepositoriesError `locationName:"errors" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDisassociateApprovalRuleTemplateFromRepositories = "BatchDisassociateApprovalRuleTemplateFromRepositories"

// BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Removes the association between an approval rule template and one or more
// specified repositories.
//
//    // Example sending a request using BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest.
//    req := client.BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/BatchDisassociateApprovalRuleTemplateFromRepositories
func (c *Client) BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest(input *BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest {
	op := &aws.Operation{
		Name:       opBatchDisassociateApprovalRuleTemplateFromRepositories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDisassociateApprovalRuleTemplateFromRepositoriesInput{}
	}

	req := c.newRequest(op, input, &BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput{})
	return BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest{Request: req, Input: input, Copy: c.BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest}
}

// BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest is the request type for the
// BatchDisassociateApprovalRuleTemplateFromRepositories API operation.
type BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest struct {
	*aws.Request
	Input *BatchDisassociateApprovalRuleTemplateFromRepositoriesInput
	Copy  func(*BatchDisassociateApprovalRuleTemplateFromRepositoriesInput) BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest
}

// Send marshals and sends the BatchDisassociateApprovalRuleTemplateFromRepositories API request.
func (r BatchDisassociateApprovalRuleTemplateFromRepositoriesRequest) Send(ctx context.Context) (*BatchDisassociateApprovalRuleTemplateFromRepositoriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDisassociateApprovalRuleTemplateFromRepositoriesResponse{
		BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput: r.Request.Data.(*BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDisassociateApprovalRuleTemplateFromRepositoriesResponse is the response type for the
// BatchDisassociateApprovalRuleTemplateFromRepositories API operation.
type BatchDisassociateApprovalRuleTemplateFromRepositoriesResponse struct {
	*BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDisassociateApprovalRuleTemplateFromRepositories request.
func (r *BatchDisassociateApprovalRuleTemplateFromRepositoriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
