// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package config

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_config_organization_conformance_pack", organizationConformancePackResource)
}

// organizationConformancePackResource returns the Terraform awscc_config_organization_conformance_pack resource.
// This Terraform resource corresponds to the CloudFormation AWS::Config::OrganizationConformancePack resource.
func organizationConformancePackResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConformancePackInputParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of ConformancePackInputParameter objects.",
		//	  "items": {
		//	    "description": "Input parameters in the form of key-value pairs for the conformance pack.",
		//	    "properties": {
		//	      "ParameterName": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "ParameterValue": {
		//	        "maxLength": 4096,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ParameterName",
		//	      "ParameterValue"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 60,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"conformance_pack_input_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ParameterName
					"parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 255),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: ParameterValue
					"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 4096),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of ConformancePackInputParameter objects.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 60),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliveryS3Bucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS Config stores intermediate files while processing conformance pack template.",
		//	  "maxLength": 63,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"delivery_s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AWS Config stores intermediate files while processing conformance pack template.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 63),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliveryS3KeyPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The prefix for the delivery S3 bucket.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"delivery_s3_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The prefix for the delivery S3 bucket.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ExcludedAccounts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "maxItems": 1000,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"excluded_accounts": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 1000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OrganizationConformancePackName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the organization conformance pack.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z][-a-zA-Z0-9]*",
		//	  "type": "string"
		//	}
		"organization_conformance_pack_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the organization conformance pack.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z][-a-zA-Z0-9]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TemplateBody
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A string containing full conformance pack template body.",
		//	  "maxLength": 51200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"template_body": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A string containing full conformance pack template body.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 51200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// TemplateBody is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: TemplateS3Uri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Location of file containing the template body.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": "s3://.*",
		//	  "type": "string"
		//	}
		"template_s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Location of file containing the template body.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile("s3://.*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// TemplateS3Uri is a write-only property.
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
		Description: "Resource schema for AWS::Config::OrganizationConformancePack.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::OrganizationConformancePack").WithTerraformTypeName("awscc_config_organization_conformance_pack")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"conformance_pack_input_parameters":  "ConformancePackInputParameters",
		"delivery_s3_bucket":                 "DeliveryS3Bucket",
		"delivery_s3_key_prefix":             "DeliveryS3KeyPrefix",
		"excluded_accounts":                  "ExcludedAccounts",
		"organization_conformance_pack_name": "OrganizationConformancePackName",
		"parameter_name":                     "ParameterName",
		"parameter_value":                    "ParameterValue",
		"template_body":                      "TemplateBody",
		"template_s3_uri":                    "TemplateS3Uri",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/TemplateBody",
		"/properties/TemplateS3Uri",
	})
	opts = opts.WithCreateTimeoutInMinutes(706).WithDeleteTimeoutInMinutes(706)

	opts = opts.WithUpdateTimeoutInMinutes(706)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
