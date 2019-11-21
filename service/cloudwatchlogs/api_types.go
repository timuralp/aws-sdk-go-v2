// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Represents a cross-account destination that receives subscription log events.
type Destination struct {
	_ struct{} `type:"structure"`

	// An IAM policy document that governs which AWS accounts can create subscription
	// filters against this destination.
	AccessPolicy *string `locationName:"accessPolicy" min:"1" type:"string"`

	// The ARN of this destination.
	Arn *string `locationName:"arn" type:"string"`

	// The creation time of the destination, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`

	// The name of the destination.
	DestinationName *string `locationName:"destinationName" min:"1" type:"string"`

	// A role for impersonation, used when delivering log events to the target.
	RoleArn *string `locationName:"roleArn" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the physical target to where the log events
	// are delivered (for example, a Kinesis stream).
	TargetArn *string `locationName:"targetArn" min:"1" type:"string"`
}

// String returns the string representation
func (s Destination) String() string {
	return awsutil.Prettify(s)
}

// Represents an export task.
type ExportTask struct {
	_ struct{} `type:"structure"`

	// The name of Amazon S3 bucket to which the log data was exported.
	Destination *string `locationName:"destination" min:"1" type:"string"`

	// The prefix that was used as the start of Amazon S3 key for every object exported.
	DestinationPrefix *string `locationName:"destinationPrefix" type:"string"`

	// Execution info about the export task.
	ExecutionInfo *ExportTaskExecutionInfo `locationName:"executionInfo" type:"structure"`

	// The start time, expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC. Events with a timestamp before this time are not exported.
	From *int64 `locationName:"from" type:"long"`

	// The name of the log group from which logs data was exported.
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string"`

	// The status of the export task.
	Status *ExportTaskStatus `locationName:"status" type:"structure"`

	// The ID of the export task.
	TaskId *string `locationName:"taskId" min:"1" type:"string"`

	// The name of the export task.
	TaskName *string `locationName:"taskName" min:"1" type:"string"`

	// The end time, expressed as the number of milliseconds after Jan 1, 1970 00:00:00
	// UTC. Events with a timestamp later than this time are not exported.
	To *int64 `locationName:"to" type:"long"`
}

// String returns the string representation
func (s ExportTask) String() string {
	return awsutil.Prettify(s)
}

// Represents the status of an export task.
type ExportTaskExecutionInfo struct {
	_ struct{} `type:"structure"`

	// The completion time of the export task, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CompletionTime *int64 `locationName:"completionTime" type:"long"`

	// The creation time of the export task, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`
}

// String returns the string representation
func (s ExportTaskExecutionInfo) String() string {
	return awsutil.Prettify(s)
}

// Represents the status of an export task.
type ExportTaskStatus struct {
	_ struct{} `type:"structure"`

	// The status code of the export task.
	Code ExportTaskStatusCode `locationName:"code" type:"string" enum:"true"`

	// The status message related to the status code.
	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ExportTaskStatus) String() string {
	return awsutil.Prettify(s)
}

// Represents a matched event.
type FilteredLogEvent struct {
	_ struct{} `type:"structure"`

	// The ID of the event.
	EventId *string `locationName:"eventId" type:"string"`

	// The time the event was ingested, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	IngestionTime *int64 `locationName:"ingestionTime" type:"long"`

	// The name of the log stream to which this event belongs.
	LogStreamName *string `locationName:"logStreamName" min:"1" type:"string"`

	// The data contained in the log event.
	Message *string `locationName:"message" min:"1" type:"string"`

	// The time the event occurred, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC.
	Timestamp *int64 `locationName:"timestamp" type:"long"`
}

// String returns the string representation
func (s FilteredLogEvent) String() string {
	return awsutil.Prettify(s)
}

// Represents a log event, which is a record of activity that was recorded by
// the application or resource being monitored.
type InputLogEvent struct {
	_ struct{} `type:"structure"`

	// The raw event message.
	//
	// Message is a required field
	Message *string `locationName:"message" min:"1" type:"string" required:"true"`

	// The time the event occurred, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC.
	//
	// Timestamp is a required field
	Timestamp *int64 `locationName:"timestamp" type:"long" required:"true"`
}

// String returns the string representation
func (s InputLogEvent) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InputLogEvent) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InputLogEvent"}

	if s.Message == nil {
		invalidParams.Add(aws.NewErrParamRequired("Message"))
	}
	if s.Message != nil && len(*s.Message) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Message", 1))
	}

	if s.Timestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("Timestamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a log group.
type LogGroup struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the log group.
	Arn *string `locationName:"arn" type:"string"`

	// The creation time of the log group, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`

	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The name of the log group.
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string"`

	// The number of metric filters.
	MetricFilterCount *int64 `locationName:"metricFilterCount" type:"integer"`

	// The number of days to retain the log events in the specified log group. Possible
	// values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731,
	// 1827, and 3653.
	RetentionInDays *int64 `locationName:"retentionInDays" type:"integer"`

	// The number of bytes stored.
	StoredBytes *int64 `locationName:"storedBytes" type:"long"`
}

// String returns the string representation
func (s LogGroup) String() string {
	return awsutil.Prettify(s)
}

// The fields contained in log events found by a GetLogGroupFields operation,
// along with the percentage of queried log events in which each field appears.
type LogGroupField struct {
	_ struct{} `type:"structure"`

	// The name of a log field.
	Name *string `locationName:"name" type:"string"`

	// The percentage of log events queried that contained the field.
	Percent *int64 `locationName:"percent" type:"integer"`
}

// String returns the string representation
func (s LogGroupField) String() string {
	return awsutil.Prettify(s)
}

// Represents a log stream, which is a sequence of log events from a single
// emitter of logs.
type LogStream struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the log stream.
	Arn *string `locationName:"arn" type:"string"`

	// The creation time of the stream, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`

	// The time of the first event, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC.
	FirstEventTimestamp *int64 `locationName:"firstEventTimestamp" type:"long"`

	// The time of the most recent log event in the log stream in CloudWatch Logs.
	// This number is expressed as the number of milliseconds after Jan 1, 1970
	// 00:00:00 UTC. The lastEventTime value updates on an eventual consistency
	// basis. It typically updates in less than an hour from ingestion, but may
	// take longer in some rare situations.
	LastEventTimestamp *int64 `locationName:"lastEventTimestamp" type:"long"`

	// The ingestion time, expressed as the number of milliseconds after Jan 1,
	// 1970 00:00:00 UTC.
	LastIngestionTime *int64 `locationName:"lastIngestionTime" type:"long"`

	// The name of the log stream.
	LogStreamName *string `locationName:"logStreamName" min:"1" type:"string"`

	// The number of bytes stored.
	//
	// IMPORTANT:On June 17, 2019, this parameter was deprecated for log streams,
	// and is always reported as zero. This change applies only to log streams.
	// The storedBytes parameter for log groups is not affected.
	StoredBytes *int64 `locationName:"storedBytes" deprecated:"true" type:"long"`

	// The sequence token.
	UploadSequenceToken *string `locationName:"uploadSequenceToken" min:"1" type:"string"`
}

// String returns the string representation
func (s LogStream) String() string {
	return awsutil.Prettify(s)
}

// Metric filters express how CloudWatch Logs would extract metric observations
// from ingested log events and transform them into metric data in a CloudWatch
// metric.
type MetricFilter struct {
	_ struct{} `type:"structure"`

	// The creation time of the metric filter, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`

	// The name of the metric filter.
	FilterName *string `locationName:"filterName" min:"1" type:"string"`

	// A symbolic description of how CloudWatch Logs should interpret the data in
	// each log event. For example, a log event may contain timestamps, IP addresses,
	// strings, and so on. You use the filter pattern to specify what to look for
	// in the log event message.
	FilterPattern *string `locationName:"filterPattern" type:"string"`

	// The name of the log group.
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string"`

	// The metric transformations.
	MetricTransformations []MetricTransformation `locationName:"metricTransformations" min:"1" type:"list"`
}

// String returns the string representation
func (s MetricFilter) String() string {
	return awsutil.Prettify(s)
}

// Represents a matched event.
type MetricFilterMatchRecord struct {
	_ struct{} `type:"structure"`

	// The raw event data.
	EventMessage *string `locationName:"eventMessage" min:"1" type:"string"`

	// The event number.
	EventNumber *int64 `locationName:"eventNumber" type:"long"`

	// The values extracted from the event data by the filter.
	ExtractedValues map[string]string `locationName:"extractedValues" type:"map"`
}

// String returns the string representation
func (s MetricFilterMatchRecord) String() string {
	return awsutil.Prettify(s)
}

// Indicates how to transform ingested log events to metric data in a CloudWatch
// metric.
type MetricTransformation struct {
	_ struct{} `type:"structure"`

	// (Optional) The value to emit when a filter pattern does not match a log event.
	// This value can be null.
	DefaultValue *float64 `locationName:"defaultValue" type:"double"`

	// The name of the CloudWatch metric.
	//
	// MetricName is a required field
	MetricName *string `locationName:"metricName" type:"string" required:"true"`

	// The namespace of the CloudWatch metric.
	//
	// MetricNamespace is a required field
	MetricNamespace *string `locationName:"metricNamespace" type:"string" required:"true"`

	// The value to publish to the CloudWatch metric when a filter pattern matches
	// a log event.
	//
	// MetricValue is a required field
	MetricValue *string `locationName:"metricValue" type:"string" required:"true"`
}

// String returns the string representation
func (s MetricTransformation) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricTransformation) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricTransformation"}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.MetricNamespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricNamespace"))
	}

	if s.MetricValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricValue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a log event.
type OutputLogEvent struct {
	_ struct{} `type:"structure"`

	// The time the event was ingested, expressed as the number of milliseconds
	// after Jan 1, 1970 00:00:00 UTC.
	IngestionTime *int64 `locationName:"ingestionTime" type:"long"`

	// The data contained in the log event.
	Message *string `locationName:"message" min:"1" type:"string"`

	// The time the event occurred, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC.
	Timestamp *int64 `locationName:"timestamp" type:"long"`
}

// String returns the string representation
func (s OutputLogEvent) String() string {
	return awsutil.Prettify(s)
}

// Reserved.
type QueryCompileError struct {
	_ struct{} `type:"structure"`

	// Reserved.
	Location *QueryCompileErrorLocation `locationName:"location" type:"structure"`

	// Reserved.
	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s QueryCompileError) String() string {
	return awsutil.Prettify(s)
}

// Reserved.
type QueryCompileErrorLocation struct {
	_ struct{} `type:"structure"`

	// Reserved.
	EndCharOffset *int64 `locationName:"endCharOffset" type:"integer"`

	// Reserved.
	StartCharOffset *int64 `locationName:"startCharOffset" type:"integer"`
}

// String returns the string representation
func (s QueryCompileErrorLocation) String() string {
	return awsutil.Prettify(s)
}

// Information about one CloudWatch Logs Insights query that matches the request
// in a DescribeQueries operation.
type QueryInfo struct {
	_ struct{} `type:"structure"`

	// The date and time that this query was created.
	CreateTime *int64 `locationName:"createTime" type:"long"`

	// The name of the log group scanned by this query.
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string"`

	// The unique ID number of this query.
	QueryId *string `locationName:"queryId" type:"string"`

	// The query string used in this query.
	QueryString *string `locationName:"queryString" type:"string"`

	// The status of this query. Possible values are Cancelled, Complete, Failed,
	// Running, Scheduled, and Unknown.
	Status QueryStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s QueryInfo) String() string {
	return awsutil.Prettify(s)
}

// Contains the number of log events scanned by the query, the number of log
// events that matched the query criteria, and the total number of bytes in
// the log events that were scanned.
type QueryStatistics struct {
	_ struct{} `type:"structure"`

	// The total number of bytes in the log events scanned during the query.
	BytesScanned *float64 `locationName:"bytesScanned" type:"double"`

	// The number of log events that matched the query string.
	RecordsMatched *float64 `locationName:"recordsMatched" type:"double"`

	// The total number of log events scanned during the query.
	RecordsScanned *float64 `locationName:"recordsScanned" type:"double"`
}

// String returns the string representation
func (s QueryStatistics) String() string {
	return awsutil.Prettify(s)
}

// Represents the rejected events.
type RejectedLogEventsInfo struct {
	_ struct{} `type:"structure"`

	// The expired log events.
	ExpiredLogEventEndIndex *int64 `locationName:"expiredLogEventEndIndex" type:"integer"`

	// The log events that are too new.
	TooNewLogEventStartIndex *int64 `locationName:"tooNewLogEventStartIndex" type:"integer"`

	// The log events that are too old.
	TooOldLogEventEndIndex *int64 `locationName:"tooOldLogEventEndIndex" type:"integer"`
}

// String returns the string representation
func (s RejectedLogEventsInfo) String() string {
	return awsutil.Prettify(s)
}

// A policy enabling one or more entities to put logs to a log group in this
// account.
type ResourcePolicy struct {
	_ struct{} `type:"structure"`

	// Timestamp showing when this policy was last updated, expressed as the number
	// of milliseconds after Jan 1, 1970 00:00:00 UTC.
	LastUpdatedTime *int64 `locationName:"lastUpdatedTime" type:"long"`

	// The details of the policy.
	PolicyDocument *string `locationName:"policyDocument" min:"1" type:"string"`

	// The name of the resource policy.
	PolicyName *string `locationName:"policyName" type:"string"`
}

// String returns the string representation
func (s ResourcePolicy) String() string {
	return awsutil.Prettify(s)
}

// Contains one field from one log event returned by a CloudWatch Logs Insights
// query, along with the value of that field.
type ResultField struct {
	_ struct{} `type:"structure"`

	// The log event field.
	Field *string `locationName:"field" type:"string"`

	// The value of this field.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s ResultField) String() string {
	return awsutil.Prettify(s)
}

// Represents the search status of a log stream.
type SearchedLogStream struct {
	_ struct{} `type:"structure"`

	// The name of the log stream.
	LogStreamName *string `locationName:"logStreamName" min:"1" type:"string"`

	// Indicates whether all the events in this log stream were searched.
	SearchedCompletely *bool `locationName:"searchedCompletely" type:"boolean"`
}

// String returns the string representation
func (s SearchedLogStream) String() string {
	return awsutil.Prettify(s)
}

// Represents a subscription filter.
type SubscriptionFilter struct {
	_ struct{} `type:"structure"`

	// The creation time of the subscription filter, expressed as the number of
	// milliseconds after Jan 1, 1970 00:00:00 UTC.
	CreationTime *int64 `locationName:"creationTime" type:"long"`

	// The Amazon Resource Name (ARN) of the destination.
	DestinationArn *string `locationName:"destinationArn" min:"1" type:"string"`

	// The method used to distribute log data to the destination, which can be either
	// random or grouped by log stream.
	Distribution Distribution `locationName:"distribution" type:"string" enum:"true"`

	// The name of the subscription filter.
	FilterName *string `locationName:"filterName" min:"1" type:"string"`

	// A symbolic description of how CloudWatch Logs should interpret the data in
	// each log event. For example, a log event may contain timestamps, IP addresses,
	// strings, and so on. You use the filter pattern to specify what to look for
	// in the log event message.
	FilterPattern *string `locationName:"filterPattern" type:"string"`

	// The name of the log group.
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string"`

	RoleArn *string `locationName:"roleArn" min:"1" type:"string"`
}

// String returns the string representation
func (s SubscriptionFilter) String() string {
	return awsutil.Prettify(s)
}
