// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package timestream

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_timestream_table", tableResource)
}

// tableResource returns the Terraform awscc_timestream_table resource.
// This Terraform resource corresponds to the CloudFormation AWS::Timestream::Table resource.
func tableResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatabaseName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the database which the table to be created belongs to.",
		//	  "pattern": "^[a-zA-Z0-9_.-]{3,256}$",
		//	  "type": "string"
		//	}
		"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the database which the table to be created belongs to.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_.-]{3,256}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MagneticStoreWriteProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The properties that determine whether magnetic store writes are enabled.",
		//	  "properties": {
		//	    "EnableMagneticStoreWrites": {
		//	      "description": "Boolean flag indicating whether magnetic store writes are enabled.",
		//	      "type": "boolean"
		//	    },
		//	    "MagneticStoreRejectedDataLocation": {
		//	      "additionalProperties": false,
		//	      "description": "Location to store information about records that were asynchronously rejected during magnetic store writes.",
		//	      "properties": {
		//	        "S3Configuration": {
		//	          "additionalProperties": false,
		//	          "description": "S3 configuration for location to store rejections from magnetic store writes",
		//	          "properties": {
		//	            "BucketName": {
		//	              "description": "The bucket name used to store the data.",
		//	              "type": "string"
		//	            },
		//	            "EncryptionOption": {
		//	              "description": "Either SSE_KMS or SSE_S3.",
		//	              "type": "string"
		//	            },
		//	            "KmsKeyId": {
		//	              "description": "Must be provided if SSE_KMS is specified as the encryption option",
		//	              "type": "string"
		//	            },
		//	            "ObjectKeyPrefix": {
		//	              "description": "String used to prefix all data in the bucket.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "EncryptionOption",
		//	            "BucketName"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "EnableMagneticStoreWrites"
		//	  ],
		//	  "type": "object"
		//	}
		"magnetic_store_write_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableMagneticStoreWrites
				"enable_magnetic_store_writes": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Boolean flag indicating whether magnetic store writes are enabled.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: MagneticStoreRejectedDataLocation
				"magnetic_store_rejected_data_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: S3Configuration
						"s3_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BucketName
								"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The bucket name used to store the data.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: EncryptionOption
								"encryption_option": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Either SSE_KMS or SSE_S3.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: KmsKeyId
								"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Must be provided if SSE_KMS is specified as the encryption option",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ObjectKeyPrefix
								"object_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "String used to prefix all data in the bucket.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "S3 configuration for location to store rejections from magnetic store writes",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Location to store information about records that were asynchronously rejected during magnetic store writes.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The properties that determine whether magnetic store writes are enabled.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The table name exposed as a read-only attribute.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The table name exposed as a read-only attribute.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RetentionProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The retention duration of the memory store and the magnetic store.",
		//	  "properties": {
		//	    "MagneticStoreRetentionPeriodInDays": {
		//	      "description": "The duration for which data must be stored in the magnetic store.",
		//	      "type": "string"
		//	    },
		//	    "MemoryStoreRetentionPeriodInHours": {
		//	      "description": "The duration for which data must be stored in the memory store.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"retention_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MagneticStoreRetentionPeriodInDays
				"magnetic_store_retention_period_in_days": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The duration for which data must be stored in the magnetic store.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MemoryStoreRetentionPeriodInHours
				"memory_store_retention_period_in_hours": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The duration for which data must be stored in the memory store.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The retention duration of the memory store and the magnetic store.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Schema
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A Schema specifies the expected data model of the table.",
		//	  "properties": {
		//	    "CompositePartitionKey": {
		//	      "description": "A list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed.",
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "An attribute used in partitioning data in a table. There are two types of partition keys: dimension keys and measure keys. A dimension key partitions data on a dimension name, while a measure key partitions data on the measure name.",
		//	        "properties": {
		//	          "EnforcementInRecord": {
		//	            "description": "The level of enforcement for the specification of a dimension key in ingested records. Options are REQUIRED (dimension key must be specified) and OPTIONAL (dimension key does not have to be specified).",
		//	            "enum": [
		//	              "REQUIRED",
		//	              "OPTIONAL"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "description": "The name of the attribute used for a dimension key.",
		//	            "maxLength": 2048,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Type": {
		//	            "description": "The type of the partition key. Options are DIMENSION (dimension key) and MEASURE (measure key).",
		//	            "enum": [
		//	              "DIMENSION",
		//	              "MEASURE"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Type"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schema": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CompositePartitionKey
				"composite_partition_key": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: EnforcementInRecord
							"enforcement_in_record": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The level of enforcement for the specification of a dimension key in ingested records. Options are REQUIRED (dimension key must be specified) and OPTIONAL (dimension key does not have to be specified).",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"REQUIRED",
										"OPTIONAL",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the attribute used for a dimension key.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 2048),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The type of the partition key. Options are DIMENSION (dimension key) and MEASURE (measure key).",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"DIMENSION",
										"MEASURE",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "A list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A Schema specifies the expected data model of the table.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TableName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.",
		//	  "pattern": "^[a-zA-Z0-9_.-]{3,256}$",
		//	  "type": "string"
		//	}
		"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_.-]{3,256}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "You can use the Resource Tags property to apply tags to resources, which can help you identify and categorize those resources.",
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The AWS::Timestream::Table resource creates a Timestream Table.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Timestream::Table").WithTerraformTypeName("awscc_timestream_table")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                   "Arn",
		"bucket_name":                           "BucketName",
		"composite_partition_key":               "CompositePartitionKey",
		"database_name":                         "DatabaseName",
		"enable_magnetic_store_writes":          "EnableMagneticStoreWrites",
		"encryption_option":                     "EncryptionOption",
		"enforcement_in_record":                 "EnforcementInRecord",
		"key":                                   "Key",
		"kms_key_id":                            "KmsKeyId",
		"magnetic_store_rejected_data_location": "MagneticStoreRejectedDataLocation",
		"magnetic_store_retention_period_in_days": "MagneticStoreRetentionPeriodInDays",
		"magnetic_store_write_properties":         "MagneticStoreWriteProperties",
		"memory_store_retention_period_in_hours":  "MemoryStoreRetentionPeriodInHours",
		"name":                                    "Name",
		"object_key_prefix":                       "ObjectKeyPrefix",
		"retention_properties":                    "RetentionProperties",
		"s3_configuration":                        "S3Configuration",
		"schema":                                  "Schema",
		"table_name":                              "TableName",
		"tags":                                    "Tags",
		"type":                                    "Type",
		"value":                                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
