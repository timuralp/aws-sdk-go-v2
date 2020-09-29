module github.com/aws/aws-sdk-go-v2/service/internal/integrationtest

require (
	github.com/aws/aws-sdk-go-v2 v0.25.0
	github.com/aws/aws-sdk-go-v2/config v0.1.0
	github.com/aws/aws-sdk-go-v2/service/acm v0.1.0
	github.com/aws/aws-sdk-go-v2/service/apigateway v0.1.0
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v0.1.0
	github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice v0.1.0
	github.com/aws/aws-sdk-go-v2/service/appstream v0.1.0
	github.com/aws/aws-sdk-go-v2/service/athena v0.1.0
	github.com/aws/aws-sdk-go-v2/service/autoscaling v0.1.0
	github.com/aws/aws-sdk-go-v2/service/batch v0.1.0
	github.com/aws/aws-sdk-go-v2/service/cloudformation v0.1.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v0.1.0
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v0.1.0
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v0.1.0
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v0.1.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v0.1.0
	github.com/aws/aws-sdk-go-v2/service/codebuild v0.1.0
	github.com/aws/aws-sdk-go-v2/service/codecommit v0.1.0
	github.com/aws/aws-sdk-go-v2/service/codedeploy v0.1.0
	github.com/aws/aws-sdk-go-v2/service/codepipeline v0.1.0
	github.com/aws/aws-sdk-go-v2/service/codestar v0.1.0
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v0.1.0
	github.com/aws/aws-sdk-go-v2/service/configservice v0.1.0
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v0.1.0
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v0.1.0
	github.com/aws/aws-sdk-go-v2/service/devicefarm v0.1.0
	github.com/aws/aws-sdk-go-v2/service/directconnect v0.1.0
	github.com/aws/aws-sdk-go-v2/service/directoryservice v0.1.0
	github.com/aws/aws-sdk-go-v2/service/docdb v0.1.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v0.1.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v0.1.0
	github.com/aws/aws-sdk-go-v2/service/ecr v0.1.0
	github.com/aws/aws-sdk-go-v2/service/ecs v0.1.0
	github.com/aws/aws-sdk-go-v2/service/efs v0.1.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v0.1.0
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v0.1.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v0.1.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v0.1.0
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v0.1.0
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v0.1.0
	github.com/aws/aws-sdk-go-v2/service/emr v0.1.0
	github.com/aws/aws-sdk-go-v2/service/eventbridge v0.1.0
	github.com/aws/aws-sdk-go-v2/service/firehose v0.1.0
	github.com/aws/aws-sdk-go-v2/service/gamelift v0.1.0
	github.com/aws/aws-sdk-go-v2/service/glacier v0.1.0
	github.com/aws/aws-sdk-go-v2/service/glue v0.1.0
	github.com/aws/aws-sdk-go-v2/service/health v0.1.0
	github.com/aws/aws-sdk-go-v2/service/iam v0.1.0
	github.com/aws/aws-sdk-go-v2/service/inspector v0.1.0
	github.com/aws/aws-sdk-go-v2/service/iot v0.1.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v0.1.0
	github.com/aws/aws-sdk-go-v2/service/kms v0.1.0
	github.com/aws/aws-sdk-go-v2/service/lambda v0.1.0
	github.com/aws/aws-sdk-go-v2/service/lightsail v0.1.0
	github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics v0.1.0
	github.com/aws/aws-sdk-go-v2/service/neptune v0.1.0
	github.com/aws/aws-sdk-go-v2/service/opsworks v0.1.0
	github.com/aws/aws-sdk-go-v2/service/pinpointemail v0.1.0
	github.com/aws/aws-sdk-go-v2/service/polly v0.1.0
	github.com/aws/aws-sdk-go-v2/service/rds v0.1.0
	github.com/aws/aws-sdk-go-v2/service/redshift v0.1.0
	github.com/aws/aws-sdk-go-v2/service/rekognition v0.1.0
	github.com/aws/aws-sdk-go-v2/service/route53 v0.1.0
	github.com/aws/aws-sdk-go-v2/service/route53domains v0.1.0
	github.com/aws/aws-sdk-go-v2/service/route53resolver v0.1.0
	github.com/aws/aws-sdk-go-v2/service/s3 v0.1.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v0.1.0
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v0.1.0
	github.com/aws/aws-sdk-go-v2/service/ses v0.1.0
	github.com/aws/aws-sdk-go-v2/service/sfn v0.1.0
	github.com/aws/aws-sdk-go-v2/service/shield v0.1.0
	github.com/aws/aws-sdk-go-v2/service/sms v0.1.0
	github.com/aws/aws-sdk-go-v2/service/snowball v0.1.0
	github.com/aws/aws-sdk-go-v2/service/sns v0.1.0
	github.com/aws/aws-sdk-go-v2/service/sqs v0.1.0
	github.com/aws/aws-sdk-go-v2/service/ssm v0.1.0
	github.com/aws/aws-sdk-go-v2/service/sts v0.1.0
	github.com/aws/aws-sdk-go-v2/service/support v0.1.0
	github.com/aws/aws-sdk-go-v2/service/waf v0.1.0
	github.com/aws/aws-sdk-go-v2/service/wafregional v0.1.0
	github.com/aws/aws-sdk-go-v2/service/wafv2 v0.1.0
	github.com/aws/aws-sdk-go-v2/service/workspaces v0.1.0
	github.com/awslabs/smithy-go v0.1.0
	github.com/aws/aws-sdk-go-v2/credentials v0.1.0
	github.com/aws/aws-sdk-go-v2/ec2imds v0.1.0
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v0.1.0
)

go 1.15
