// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudtrail

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cloudtrail_trail", trailResourceType)
}

// trailResourceType returns the Terraform awscc_cloudtrail_trail resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudTrail::Trail resource type.
func trailResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"cloudwatch_logs_log_group_arn": {
			// Property: CloudWatchLogsLogGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.",
			//   "type": "string"
			// }
			Description: "Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.",
			Type:        types.StringType,
			Optional:    true,
		},
		"cloudwatch_logs_role_arn": {
			// Property: CloudWatchLogsRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.",
			//   "type": "string"
			// }
			Description: "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.",
			Type:        types.StringType,
			Optional:    true,
		},
		"enable_log_file_validation": {
			// Property: EnableLogFileValidation
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether log file validation is enabled. The default is false.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether log file validation is enabled. The default is false.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"event_selectors": {
			// Property: EventSelectors
			// CloudFormation resource type schema:
			// {
			//   "description": "Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The type of email sending events to publish to the event destination.",
			//     "properties": {
			//       "DataResources": {
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "CloudTrail supports data event logging for Amazon S3 objects and AWS Lambda functions. You can specify up to 250 resources for an individual event selector, but the total number of data resources cannot exceed 250 across all event selectors in a trail. This limit does not apply if you configure resource logging for all data events.",
			//           "properties": {
			//             "Type": {
			//               "description": "The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.",
			//               "type": "string"
			//             },
			//             "Values": {
			//               "description": "An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array",
			//               "uniqueItems": true
			//             }
			//           },
			//           "required": [
			//             "Type"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "ExcludeManagementEventSources": {
			//         "description": "An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing \"kms.amazonaws.com\". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.",
			//         "insertionOrder": false,
			//         "items": {
			//           "type": "string"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "IncludeManagementEvents": {
			//         "description": "Specify if you want your event selector to include management events for your trail.",
			//         "type": "boolean"
			//       },
			//       "ReadWriteType": {
			//         "description": "Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.",
			//         "enum": [
			//           "All",
			//           "ReadOnly",
			//           "WriteOnly"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 5,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"data_resources": {
						// Property: DataResources
						Attributes: providertypes.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"type": {
									// Property: Type
									Description: "The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.",
									Type:        types.StringType,
									Required:    true,
								},
								"values": {
									// Property: Values
									Description: "An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.",
									Type:        providertypes.SetType{ElemType: types.StringType},
									Optional:    true,
								},
							},
							providertypes.SetNestedAttributesOptions{},
						),
						Optional: true,
					},
					"exclude_management_event_sources": {
						// Property: ExcludeManagementEventSources
						Description: "An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing \"kms.amazonaws.com\". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.",
						Type:        providertypes.SetType{ElemType: types.StringType},
						Optional:    true,
					},
					"include_management_events": {
						// Property: IncludeManagementEvents
						Description: "Specify if you want your event selector to include management events for your trail.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"read_write_type": {
						// Property: ReadWriteType
						Description: "Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"All",
								"ReadOnly",
								"WriteOnly",
							}),
						},
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 5,
				},
			),
			Optional: true,
		},
		"include_global_service_events": {
			// Property: IncludeGlobalServiceEvents
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the trail is publishing events from global services such as IAM to the log files.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the trail is publishing events from global services such as IAM to the log files.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"insight_selectors": {
			// Property: InsightSelectors
			// CloudFormation resource type schema:
			// {
			//   "description": "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A string that contains insight types that are logged on a trail.",
			//     "properties": {
			//       "InsightType": {
			//         "description": "The type of insight to log on a trail.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"insight_type": {
						// Property: InsightType
						Description: "The type of insight to log on a trail.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"is_logging": {
			// Property: IsLogging
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether the CloudTrail is currently logging AWS API calls.",
			//   "type": "boolean"
			// }
			Description: "Whether the CloudTrail is currently logging AWS API calls.",
			Type:        types.BoolType,
			Required:    true,
		},
		"is_multi_region_trail": {
			// Property: IsMultiRegionTrail
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"is_organization_trail": {
			// Property: IsOrganizationTrail
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.",
			//   "type": "boolean"
			// }
			Description: "Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"kms_key_id": {
			// Property: KMSKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
			//   "type": "string"
			// }
			Description: "Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.",
			Type:        types.StringType,
			Optional:    true,
		},
		"s3_bucket_name": {
			// Property: S3BucketName
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.",
			//   "type": "string"
			// }
			Description: "Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.",
			Type:        types.StringType,
			Required:    true,
		},
		"s3_key_prefix": {
			// Property: S3KeyPrefix
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.",
			//   "maxLength": 200,
			//   "type": "string"
			// }
			Description: "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 200),
			},
		},
		"sns_topic_arn": {
			// Property: SnsTopicArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"sns_topic_name": {
			// Property: SnsTopicName
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 256),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "An arbitrary set of tags (key-value pairs) for this trail.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"trail_name": {
			// Property: TrailName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(3, 128),
			},
			// TrailName is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket. A maximum of five trails can exist in a region, irrespective of the region in which they were created.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudTrail::Trail").WithTerraformTypeName("awscc_cloudtrail_trail")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                              "Arn",
		"cloudwatch_logs_log_group_arn":    "CloudWatchLogsLogGroupArn",
		"cloudwatch_logs_role_arn":         "CloudWatchLogsRoleArn",
		"data_resources":                   "DataResources",
		"enable_log_file_validation":       "EnableLogFileValidation",
		"event_selectors":                  "EventSelectors",
		"exclude_management_event_sources": "ExcludeManagementEventSources",
		"include_global_service_events":    "IncludeGlobalServiceEvents",
		"include_management_events":        "IncludeManagementEvents",
		"insight_selectors":                "InsightSelectors",
		"insight_type":                     "InsightType",
		"is_logging":                       "IsLogging",
		"is_multi_region_trail":            "IsMultiRegionTrail",
		"is_organization_trail":            "IsOrganizationTrail",
		"key":                              "Key",
		"kms_key_id":                       "KMSKeyId",
		"read_write_type":                  "ReadWriteType",
		"s3_bucket_name":                   "S3BucketName",
		"s3_key_prefix":                    "S3KeyPrefix",
		"sns_topic_arn":                    "SnsTopicArn",
		"sns_topic_name":                   "SnsTopicName",
		"tags":                             "Tags",
		"trail_name":                       "TrailName",
		"type":                             "Type",
		"value":                            "Value",
		"values":                           "Values",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_cloudtrail_trail", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
