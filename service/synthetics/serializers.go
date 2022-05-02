// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateCanary struct {
}

func (*awsRestjson1_serializeOpCreateCanary) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateCanary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateCanaryInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canary")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateCanaryInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCreateCanaryInput(v *CreateCanaryInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateCanaryInput(v *CreateCanaryInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ArtifactConfig != nil {
		ok := object.Key("ArtifactConfig")
		if err := awsRestjson1_serializeDocumentArtifactConfigInput(v.ArtifactConfig, ok); err != nil {
			return err
		}
	}

	if v.ArtifactS3Location != nil {
		ok := object.Key("ArtifactS3Location")
		ok.String(*v.ArtifactS3Location)
	}

	if v.Code != nil {
		ok := object.Key("Code")
		if err := awsRestjson1_serializeDocumentCanaryCodeInput(v.Code, ok); err != nil {
			return err
		}
	}

	if v.ExecutionRoleArn != nil {
		ok := object.Key("ExecutionRoleArn")
		ok.String(*v.ExecutionRoleArn)
	}

	if v.FailureRetentionPeriodInDays != nil {
		ok := object.Key("FailureRetentionPeriodInDays")
		ok.Integer(*v.FailureRetentionPeriodInDays)
	}

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.RunConfig != nil {
		ok := object.Key("RunConfig")
		if err := awsRestjson1_serializeDocumentCanaryRunConfigInput(v.RunConfig, ok); err != nil {
			return err
		}
	}

	if v.RuntimeVersion != nil {
		ok := object.Key("RuntimeVersion")
		ok.String(*v.RuntimeVersion)
	}

	if v.Schedule != nil {
		ok := object.Key("Schedule")
		if err := awsRestjson1_serializeDocumentCanaryScheduleInput(v.Schedule, ok); err != nil {
			return err
		}
	}

	if v.SuccessRetentionPeriodInDays != nil {
		ok := object.Key("SuccessRetentionPeriodInDays")
		ok.Integer(*v.SuccessRetentionPeriodInDays)
	}

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	if v.VpcConfig != nil {
		ok := object.Key("VpcConfig")
		if err := awsRestjson1_serializeDocumentVpcConfigInput(v.VpcConfig, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDeleteCanary struct {
}

func (*awsRestjson1_serializeOpDeleteCanary) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteCanary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteCanaryInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canary/{Name}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteCanaryInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteCanaryInput(v *DeleteCanaryInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.DeleteLambda {
		encoder.SetQuery("deleteLambda").Boolean(v.DeleteLambda)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDescribeCanaries struct {
}

func (*awsRestjson1_serializeOpDescribeCanaries) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeCanaries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeCanariesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canaries")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentDescribeCanariesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeCanariesInput(v *DescribeCanariesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDescribeCanariesInput(v *DescribeCanariesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.Names != nil {
		ok := object.Key("Names")
		if err := awsRestjson1_serializeDocumentDescribeCanariesNameFilter(v.Names, ok); err != nil {
			return err
		}
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpDescribeCanariesLastRun struct {
}

func (*awsRestjson1_serializeOpDescribeCanariesLastRun) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeCanariesLastRun) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeCanariesLastRunInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canaries/last-run")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentDescribeCanariesLastRunInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeCanariesLastRunInput(v *DescribeCanariesLastRunInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDescribeCanariesLastRunInput(v *DescribeCanariesLastRunInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.Names != nil {
		ok := object.Key("Names")
		if err := awsRestjson1_serializeDocumentDescribeCanariesLastRunNameFilter(v.Names, ok); err != nil {
			return err
		}
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpDescribeRuntimeVersions struct {
}

func (*awsRestjson1_serializeOpDescribeRuntimeVersions) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeRuntimeVersions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeRuntimeVersionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/runtime-versions")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentDescribeRuntimeVersionsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeRuntimeVersionsInput(v *DescribeRuntimeVersionsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentDescribeRuntimeVersionsInput(v *DescribeRuntimeVersionsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpGetCanary struct {
}

func (*awsRestjson1_serializeOpGetCanary) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetCanary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetCanaryInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canary/{Name}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetCanaryInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetCanaryInput(v *GetCanaryInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetCanaryRuns struct {
}

func (*awsRestjson1_serializeOpGetCanaryRuns) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetCanaryRuns) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetCanaryRunsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canary/{Name}/runs")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetCanaryRunsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentGetCanaryRunsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetCanaryRunsInput(v *GetCanaryRunsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentGetCanaryRunsInput(v *GetCanaryRunsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpListTagsForResource struct {
}

func (*awsRestjson1_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{ResourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member ResourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("ResourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStartCanary struct {
}

func (*awsRestjson1_serializeOpStartCanary) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartCanary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartCanaryInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canary/{Name}/start")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsStartCanaryInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStartCanaryInput(v *StartCanaryInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStopCanary struct {
}

func (*awsRestjson1_serializeOpStopCanary) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStopCanary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StopCanaryInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canary/{Name}/stop")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsStopCanaryInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStopCanaryInput(v *StopCanaryInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpTagResource struct {
}

func (*awsRestjson1_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{ResourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsTagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member ResourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("ResourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Tags != nil {
		ok := object.Key("Tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpUntagResource struct {
}

func (*awsRestjson1_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{ResourceArn}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUntagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil || len(*v.ResourceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member ResourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if err := encoder.SetURI("ResourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	if v.TagKeys != nil {
		for i := range v.TagKeys {
			encoder.AddQuery("tagKeys").String(v.TagKeys[i])
		}
	}

	return nil
}

type awsRestjson1_serializeOpUpdateCanary struct {
}

func (*awsRestjson1_serializeOpUpdateCanary) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateCanary) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateCanaryInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/canary/{Name}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "PATCH"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateCanaryInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateCanaryInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUpdateCanaryInput(v *UpdateCanaryInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Name == nil || len(*v.Name) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member Name must not be empty")}
	}
	if v.Name != nil {
		if err := encoder.SetURI("Name").String(*v.Name); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateCanaryInput(v *UpdateCanaryInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ArtifactConfig != nil {
		ok := object.Key("ArtifactConfig")
		if err := awsRestjson1_serializeDocumentArtifactConfigInput(v.ArtifactConfig, ok); err != nil {
			return err
		}
	}

	if v.ArtifactS3Location != nil {
		ok := object.Key("ArtifactS3Location")
		ok.String(*v.ArtifactS3Location)
	}

	if v.Code != nil {
		ok := object.Key("Code")
		if err := awsRestjson1_serializeDocumentCanaryCodeInput(v.Code, ok); err != nil {
			return err
		}
	}

	if v.ExecutionRoleArn != nil {
		ok := object.Key("ExecutionRoleArn")
		ok.String(*v.ExecutionRoleArn)
	}

	if v.FailureRetentionPeriodInDays != nil {
		ok := object.Key("FailureRetentionPeriodInDays")
		ok.Integer(*v.FailureRetentionPeriodInDays)
	}

	if v.RunConfig != nil {
		ok := object.Key("RunConfig")
		if err := awsRestjson1_serializeDocumentCanaryRunConfigInput(v.RunConfig, ok); err != nil {
			return err
		}
	}

	if v.RuntimeVersion != nil {
		ok := object.Key("RuntimeVersion")
		ok.String(*v.RuntimeVersion)
	}

	if v.Schedule != nil {
		ok := object.Key("Schedule")
		if err := awsRestjson1_serializeDocumentCanaryScheduleInput(v.Schedule, ok); err != nil {
			return err
		}
	}

	if v.SuccessRetentionPeriodInDays != nil {
		ok := object.Key("SuccessRetentionPeriodInDays")
		ok.Integer(*v.SuccessRetentionPeriodInDays)
	}

	if v.VisualReference != nil {
		ok := object.Key("VisualReference")
		if err := awsRestjson1_serializeDocumentVisualReferenceInput(v.VisualReference, ok); err != nil {
			return err
		}
	}

	if v.VpcConfig != nil {
		ok := object.Key("VpcConfig")
		if err := awsRestjson1_serializeDocumentVpcConfigInput(v.VpcConfig, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentArtifactConfigInput(v *types.ArtifactConfigInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.S3Encryption != nil {
		ok := object.Key("S3Encryption")
		if err := awsRestjson1_serializeDocumentS3EncryptionConfig(v.S3Encryption, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentBaseScreenshot(v *types.BaseScreenshot, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.IgnoreCoordinates != nil {
		ok := object.Key("IgnoreCoordinates")
		if err := awsRestjson1_serializeDocumentBaseScreenshotIgnoreCoordinates(v.IgnoreCoordinates, ok); err != nil {
			return err
		}
	}

	if v.ScreenshotName != nil {
		ok := object.Key("ScreenshotName")
		ok.String(*v.ScreenshotName)
	}

	return nil
}

func awsRestjson1_serializeDocumentBaseScreenshotIgnoreCoordinates(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentBaseScreenshots(v []types.BaseScreenshot, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentBaseScreenshot(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentCanaryCodeInput(v *types.CanaryCodeInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Handler != nil {
		ok := object.Key("Handler")
		ok.String(*v.Handler)
	}

	if v.S3Bucket != nil {
		ok := object.Key("S3Bucket")
		ok.String(*v.S3Bucket)
	}

	if v.S3Key != nil {
		ok := object.Key("S3Key")
		ok.String(*v.S3Key)
	}

	if v.S3Version != nil {
		ok := object.Key("S3Version")
		ok.String(*v.S3Version)
	}

	if v.ZipFile != nil {
		ok := object.Key("ZipFile")
		ok.Base64EncodeBytes(v.ZipFile)
	}

	return nil
}

func awsRestjson1_serializeDocumentCanaryRunConfigInput(v *types.CanaryRunConfigInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ActiveTracing != nil {
		ok := object.Key("ActiveTracing")
		ok.Boolean(*v.ActiveTracing)
	}

	if v.EnvironmentVariables != nil {
		ok := object.Key("EnvironmentVariables")
		if err := awsRestjson1_serializeDocumentEnvironmentVariablesMap(v.EnvironmentVariables, ok); err != nil {
			return err
		}
	}

	if v.MemoryInMB != nil {
		ok := object.Key("MemoryInMB")
		ok.Integer(*v.MemoryInMB)
	}

	if v.TimeoutInSeconds != nil {
		ok := object.Key("TimeoutInSeconds")
		ok.Integer(*v.TimeoutInSeconds)
	}

	return nil
}

func awsRestjson1_serializeDocumentCanaryScheduleInput(v *types.CanaryScheduleInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DurationInSeconds != nil {
		ok := object.Key("DurationInSeconds")
		ok.Long(*v.DurationInSeconds)
	}

	if v.Expression != nil {
		ok := object.Key("Expression")
		ok.String(*v.Expression)
	}

	return nil
}

func awsRestjson1_serializeDocumentDescribeCanariesLastRunNameFilter(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentDescribeCanariesNameFilter(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentEnvironmentVariablesMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentS3EncryptionConfig(v *types.S3EncryptionConfig, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.EncryptionMode) > 0 {
		ok := object.Key("EncryptionMode")
		ok.String(string(v.EncryptionMode))
	}

	if v.KmsKeyArn != nil {
		ok := object.Key("KmsKeyArn")
		ok.String(*v.KmsKeyArn)
	}

	return nil
}

func awsRestjson1_serializeDocumentSecurityGroupIds(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentSubnetIds(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentTagMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentVisualReferenceInput(v *types.VisualReferenceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BaseCanaryRunId != nil {
		ok := object.Key("BaseCanaryRunId")
		ok.String(*v.BaseCanaryRunId)
	}

	if v.BaseScreenshots != nil {
		ok := object.Key("BaseScreenshots")
		if err := awsRestjson1_serializeDocumentBaseScreenshots(v.BaseScreenshots, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentVpcConfigInput(v *types.VpcConfigInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SecurityGroupIds != nil {
		ok := object.Key("SecurityGroupIds")
		if err := awsRestjson1_serializeDocumentSecurityGroupIds(v.SecurityGroupIds, ok); err != nil {
			return err
		}
	}

	if v.SubnetIds != nil {
		ok := object.Key("SubnetIds")
		if err := awsRestjson1_serializeDocumentSubnetIds(v.SubnetIds, ok); err != nil {
			return err
		}
	}

	return nil
}
