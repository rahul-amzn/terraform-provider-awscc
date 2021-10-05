// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lightsail_disk", diskDataSourceType)
}

// diskDataSourceType returns the Terraform awscc_lightsail_disk data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Lightsail::Disk resource type.
func diskDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"add_ons": {
			// Property: AddOns
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of objects representing the add-ons to enable for the new instance.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A addon associate with a resource.",
			//     "properties": {
			//       "AddOnType": {
			//         "description": "The add-on type",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "AutoSnapshotAddOnRequest": {
			//         "additionalProperties": false,
			//         "description": "An object that represents additional parameters when enabling or modifying the automatic snapshot add-on",
			//         "properties": {
			//           "SnapshotTimeOfDay": {
			//             "description": "The daily time when an automatic snapshot will be created.",
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Status": {
			//         "description": "Status of the Addon",
			//         "enum": [
			//           "Enabling",
			//           "Disabling",
			//           "Enabled",
			//           "Terminating",
			//           "Terminated",
			//           "Disabled",
			//           "Failed"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "AddOnType"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "An array of objects representing the add-ons to enable for the new instance.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"add_on_type": {
						// Property: AddOnType
						Description: "The add-on type",
						Type:        types.StringType,
						Computed:    true,
					},
					"auto_snapshot_add_on_request": {
						// Property: AutoSnapshotAddOnRequest
						Description: "An object that represents additional parameters when enabling or modifying the automatic snapshot add-on",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"snapshot_time_of_day": {
									// Property: SnapshotTimeOfDay
									Description: "The daily time when an automatic snapshot will be created.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"status": {
						// Property: Status
						Description: "Status of the Addon",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"attached_to": {
			// Property: AttachedTo
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the attached Lightsail Instance",
			//   "type": "string"
			// }
			Description: "Name of the attached Lightsail Instance",
			Type:        types.StringType,
			Computed:    true,
		},
		"attachment_state": {
			// Property: AttachmentState
			// CloudFormation resource type schema:
			// {
			//   "description": "Attachment State of the Lightsail disk",
			//   "type": "string"
			// }
			Description: "Attachment State of the Lightsail disk",
			Type:        types.StringType,
			Computed:    true,
		},
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			// {
			//   "description": "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
			Type:        types.StringType,
			Computed:    true,
		},
		"disk_arn": {
			// Property: DiskArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"disk_name": {
			// Property: DiskName
			// CloudFormation resource type schema:
			// {
			//   "description": "The names to use for your new Lightsail disk.",
			//   "maxLength": 254,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The names to use for your new Lightsail disk.",
			Type:        types.StringType,
			Computed:    true,
		},
		"iops": {
			// Property: Iops
			// CloudFormation resource type schema:
			// {
			//   "description": "Iops of the Lightsail disk",
			//   "type": "integer"
			// }
			Description: "Iops of the Lightsail disk",
			Type:        types.NumberType,
			Computed:    true,
		},
		"is_attached": {
			// Property: IsAttached
			// CloudFormation resource type schema:
			// {
			//   "description": "Check is Disk is attached state",
			//   "type": "boolean"
			// }
			Description: "Check is Disk is attached state",
			Type:        types.BoolType,
			Computed:    true,
		},
		"location": {
			// Property: Location
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Location of a resource.",
			//   "properties": {
			//     "AvailabilityZone": {
			//       "description": "The Availability Zone in which to create your disk. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
			//       "type": "string"
			//     },
			//     "RegionName": {
			//       "description": "The Region Name in which to create your disk.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Location of a resource.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"availability_zone": {
						// Property: AvailabilityZone
						Description: "The Availability Zone in which to create your disk. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
						Type:        types.StringType,
						Computed:    true,
					},
					"region_name": {
						// Property: RegionName
						Description: "The Region Name in which to create your disk.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"path": {
			// Property: Path
			// CloudFormation resource type schema:
			// {
			//   "description": "Path of the  attached Disk",
			//   "type": "string"
			// }
			Description: "Path of the  attached Disk",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "Resource type of Lightsail instance.",
			//   "type": "string"
			// }
			Description: "Resource type of Lightsail instance.",
			Type:        types.StringType,
			Computed:    true,
		},
		"size_in_gb": {
			// Property: SizeInGb
			// CloudFormation resource type schema:
			// {
			//   "description": "Size of the Lightsail disk",
			//   "type": "integer"
			// }
			Description: "Size of the Lightsail disk",
			Type:        types.NumberType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "State of the Lightsail disk",
			//   "type": "string"
			// }
			Description: "State of the Lightsail disk",
			Type:        types.StringType,
			Computed:    true,
		},
		"support_code": {
			// Property: SupportCode
			// CloudFormation resource type schema:
			// {
			//   "description": "Support code to help identify any issues",
			//   "type": "string"
			// }
			Description: "Support code to help identify any issues",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Lightsail::Disk",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Disk").WithTerraformTypeName("awscc_lightsail_disk")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_on_type":                  "AddOnType",
		"add_ons":                      "AddOns",
		"attached_to":                  "AttachedTo",
		"attachment_state":             "AttachmentState",
		"auto_snapshot_add_on_request": "AutoSnapshotAddOnRequest",
		"availability_zone":            "AvailabilityZone",
		"disk_arn":                     "DiskArn",
		"disk_name":                    "DiskName",
		"iops":                         "Iops",
		"is_attached":                  "IsAttached",
		"key":                          "Key",
		"location":                     "Location",
		"path":                         "Path",
		"region_name":                  "RegionName",
		"resource_type":                "ResourceType",
		"size_in_gb":                   "SizeInGb",
		"snapshot_time_of_day":         "SnapshotTimeOfDay",
		"state":                        "State",
		"status":                       "Status",
		"support_code":                 "SupportCode",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
