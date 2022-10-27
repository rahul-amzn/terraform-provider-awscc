// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package fsx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_fsx_data_repository_association", dataRepositoryAssociationDataSource)
}

// dataRepositoryAssociationDataSource returns the Terraform awscc_fsx_data_repository_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::FSx::DataRepositoryAssociation resource.
func dataRepositoryAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"association_id": {
			// Property: AssociationId
			// CloudFormation resource type schema:
			// {
			//   "description": "The system-generated, unique ID of the data repository association.",
			//   "type": "string"
			// }
			Description: "The system-generated, unique ID of the data repository association.",
			Type:        types.StringType,
			Computed:    true,
		},
		"batch_import_meta_data_on_create": {
			// Property: BatchImportMetaDataOnCreate
			// CloudFormation resource type schema:
			// {
			//   "description": "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to true.",
			//   "type": "boolean"
			// }
			Description: "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to true.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"data_repository_path": {
			// Property: DataRepositoryPath
			// CloudFormation resource type schema:
			// {
			//   "description": "The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format s3://myBucket/myPrefix/ . This path specifies where in the S3 data repository files will be imported from or exported to.",
			//   "type": "string"
			// }
			Description: "The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format s3://myBucket/myPrefix/ . This path specifies where in the S3 data repository files will be imported from or exported to.",
			Type:        types.StringType,
			Computed:    true,
		},
		"file_system_id": {
			// Property: FileSystemId
			// CloudFormation resource type schema:
			// {
			//   "description": "The globally unique ID of the file system, assigned by Amazon FSx.",
			//   "type": "string"
			// }
			Description: "The globally unique ID of the file system, assigned by Amazon FSx.",
			Type:        types.StringType,
			Computed:    true,
		},
		"file_system_path": {
			// Property: FileSystemPath
			// CloudFormation resource type schema:
			// {
			//   "description": "This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.",
			//   "type": "string"
			// }
			Description: "This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.",
			Type:        types.StringType,
			Computed:    true,
		},
		"imported_file_chunk_size": {
			// Property: ImportedFileChunkSize
			// CloudFormation resource type schema:
			// {
			//   "description": "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.",
			//   "type": "integer"
			// }
			Description: "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"resource_arn": {
			// Property: ResourceARN
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify Amazon Web Services resources. We require an ARN when you need to specify a resource unambiguously across all of Amazon Web Services. For more information, see Amazon Resource Names (ARNs) in the Amazon Web Services General Reference.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify Amazon Web Services resources. We require an ARN when you need to specify a resource unambiguously across all of Amazon Web Services. For more information, see Amazon Resource Names (ARNs) in the Amazon Web Services General Reference.",
			Type:        types.StringType,
			Computed:    true,
		},
		"s3": {
			// Property: S3
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.",
			//   "properties": {
			//     "AutoExportPolicy": {
			//       "additionalProperties": false,
			//       "description": "Specifies the type of updated objects (new, changed, deleted) that will be automatically exported from your file system to the linked S3 bucket.",
			//       "properties": {
			//         "Events": {
			//           "insertionOrder": false,
			//           "items": {
			//             "enum": [
			//               "NEW",
			//               "CHANGED",
			//               "DELETED"
			//             ],
			//             "type": "string"
			//           },
			//           "maxItems": 3,
			//           "type": "array",
			//           "uniqueItems": true
			//         }
			//       },
			//       "required": [
			//         "Events"
			//       ],
			//       "type": "object"
			//     },
			//     "AutoImportPolicy": {
			//       "additionalProperties": false,
			//       "description": "Specifies the type of updated objects (new, changed, deleted) that will be automatically imported from the linked S3 bucket to your file system.",
			//       "properties": {
			//         "Events": {
			//           "insertionOrder": false,
			//           "items": {
			//             "enum": [
			//               "NEW",
			//               "CHANGED",
			//               "DELETED"
			//             ],
			//             "type": "string"
			//           },
			//           "maxItems": 3,
			//           "type": "array",
			//           "uniqueItems": true
			//         }
			//       },
			//       "required": [
			//         "Events"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"auto_export_policy": {
						// Property: AutoExportPolicy
						Description: "Specifies the type of updated objects (new, changed, deleted) that will be automatically exported from your file system to the linked S3 bucket.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"events": {
									// Property: Events
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"auto_import_policy": {
						// Property: AutoImportPolicy
						Description: "Specifies the type of updated objects (new, changed, deleted) that will be automatically imported from the linked S3 bucket to your file system.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"events": {
									// Property: Events
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of Tag values, with a maximum of 50 elements.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of Tag values, with a maximum of 50 elements.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
				},
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
		Description: "Data Source schema for AWS::FSx::DataRepositoryAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FSx::DataRepositoryAssociation").WithTerraformTypeName("awscc_fsx_data_repository_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id":                   "AssociationId",
		"auto_export_policy":               "AutoExportPolicy",
		"auto_import_policy":               "AutoImportPolicy",
		"batch_import_meta_data_on_create": "BatchImportMetaDataOnCreate",
		"data_repository_path":             "DataRepositoryPath",
		"events":                           "Events",
		"file_system_id":                   "FileSystemId",
		"file_system_path":                 "FileSystemPath",
		"imported_file_chunk_size":         "ImportedFileChunkSize",
		"key":                              "Key",
		"resource_arn":                     "ResourceARN",
		"s3":                               "S3",
		"tags":                             "Tags",
		"value":                            "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
