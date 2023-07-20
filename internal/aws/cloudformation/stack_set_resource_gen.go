// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_cloudformation_stack_set", stackSetResource)
}

// stackSetResource returns the Terraform awscc_cloudformation_stack_set resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFormation::StackSet resource.
func stackSetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdministrationRoleARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"administration_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AutoDeployment
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.",
		//	  "properties": {
		//	    "Enabled": {
		//	      "description": "If set to true, StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions. If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.",
		//	      "type": "boolean"
		//	    },
		//	    "RetainStacksOnAccountRemoval": {
		//	      "description": "If set to true, stack resources are retained when an account is removed from a target organization or OU. If set to false, stack resources are deleted. Specify only if Enabled is set to True.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"auto_deployment": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If set to true, StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions. If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RetainStacksOnAccountRemoval
				"retain_stacks_on_account_removal": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If set to true, stack resources are retained when an account is removed from a target organization or OU. If set to false, stack resources are deleted. Specify only if Enabled is set to True.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CallAs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the AWS account that you are acting from. By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN.",
		//	  "enum": [
		//	    "SELF",
		//	    "DELEGATED_ADMIN"
		//	  ],
		//	  "type": "string"
		//	}
		"call_as": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the AWS account that you are acting from. By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SELF",
					"DELEGATED_ADMIN",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// CallAs is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Capabilities
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "CAPABILITY_IAM",
		//	      "CAPABILITY_NAMED_IAM",
		//	      "CAPABILITY_AUTO_EXPAND"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"capabilities": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"CAPABILITY_IAM",
						"CAPABILITY_NAMED_IAM",
						"CAPABILITY_AUTO_EXPAND",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the stack set. You can use the description to identify the stack set's purpose or other important information.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the stack set. You can use the description to identify the stack set's purpose or other important information.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ExecutionRoleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"execution_role_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManagedExecution
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.",
		//	  "properties": {
		//	    "Active": {
		//	      "description": "When true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"managed_execution": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Active
				"active": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "When true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OperationPreferences
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The user-specified preferences for how AWS CloudFormation performs a stack set operation.",
		//	  "properties": {
		//	    "FailureToleranceCount": {
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "FailureTolerancePercentage": {
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "MaxConcurrentCount": {
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "MaxConcurrentPercentage": {
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "RegionConcurrencyType": {
		//	      "description": "The concurrency type of deploying StackSets operations in regions, could be in parallel or one region at a time",
		//	      "enum": [
		//	        "SEQUENTIAL",
		//	        "PARALLEL"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "RegionOrder": {
		//	      "items": {
		//	        "pattern": "^[a-zA-Z0-9-]{1,128}$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"operation_preferences": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FailureToleranceCount
				"failure_tolerance_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FailureTolerancePercentage
				"failure_tolerance_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 100),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MaxConcurrentCount
				"max_concurrent_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MaxConcurrentPercentage
				"max_concurrent_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 100),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RegionConcurrencyType
				"region_concurrency_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The concurrency type of deploying StackSets operations in regions, could be in parallel or one region at a time",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"SEQUENTIAL",
							"PARALLEL",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RegionOrder
				"region_order": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-]{1,128}$"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The user-specified preferences for how AWS CloudFormation performs a stack set operation.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// OperationPreferences is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Parameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The input parameters for the stack set template.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ParameterKey": {
		//	        "description": "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
		//	        "type": "string"
		//	      },
		//	      "ParameterValue": {
		//	        "description": "The input value associated with the parameter.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ParameterKey",
		//	      "ParameterValue"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"parameters": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ParameterKey
					"parameter_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: ParameterValue
					"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The input value associated with the parameter.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The input parameters for the stack set template.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PermissionModel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified.",
		//	  "enum": [
		//	    "SERVICE_MANAGED",
		//	    "SELF_MANAGED"
		//	  ],
		//	  "type": "string"
		//	}
		"permission_model": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SERVICE_MANAGED",
					"SELF_MANAGED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StackInstancesGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A group of stack instances with parameters in some specific accounts and regions.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Stack instances in some specific accounts and Regions.",
		//	    "properties": {
		//	      "DeploymentTargets": {
		//	        "additionalProperties": false,
		//	        "description": " The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.",
		//	        "properties": {
		//	          "AccountFilterType": {
		//	            "description": "The filter type you want to apply on organizational units and accounts.",
		//	            "enum": [
		//	              "NONE",
		//	              "UNION",
		//	              "INTERSECTION",
		//	              "DIFFERENCE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Accounts": {
		//	            "description": "AWS accounts that you want to create stack instances in the specified Region(s) for.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "description": "AWS account that you want to create stack instances in the specified Region(s) for.",
		//	              "pattern": "^[0-9]{12}$",
		//	              "type": "string"
		//	            },
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "OrganizationalUnitIds": {
		//	            "description": "The organization root ID or organizational unit (OU) IDs to which StackSets deploys.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "pattern": "^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$",
		//	              "type": "string"
		//	            },
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ParameterOverrides": {
		//	        "description": "A list of stack set parameters whose values you want to override in the selected stack instances.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ParameterKey": {
		//	              "description": "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
		//	              "type": "string"
		//	            },
		//	            "ParameterValue": {
		//	              "description": "The input value associated with the parameter.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "ParameterKey",
		//	            "ParameterValue"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "Regions": {
		//	        "description": "The names of one or more Regions where you want to create stack instances using the specified AWS account(s).",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "pattern": "^[a-zA-Z0-9-]{1,128}$",
		//	          "type": "string"
		//	        },
		//	        "minItems": 1,
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      }
		//	    },
		//	    "required": [
		//	      "DeploymentTargets",
		//	      "Regions"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"stack_instances_group": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DeploymentTargets
					"deployment_targets": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AccountFilterType
							"account_filter_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The filter type you want to apply on organizational units and accounts.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"NONE",
										"UNION",
										"INTERSECTION",
										"DIFFERENCE",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Accounts
							"accounts": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "AWS accounts that you want to create stack instances in the specified Region(s) for.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.Set{ /*START VALIDATORS*/
									setvalidator.SizeAtLeast(1),
									setvalidator.ValueStringsAre(
										stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{12}$"), ""),
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
									setplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: OrganizationalUnitIds
							"organizational_unit_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "The organization root ID or organizational unit (OU) IDs to which StackSets deploys.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.Set{ /*START VALIDATORS*/
									setvalidator.SizeAtLeast(1),
									setvalidator.ValueStringsAre(
										stringvalidator.RegexMatches(regexp.MustCompile("^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$"), ""),
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
									setplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: " The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: ParameterOverrides
					"parameter_overrides": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ParameterKey
								"parameter_key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: ParameterValue
								"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The input value associated with the parameter.",
									Required:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "A list of stack set parameters whose values you want to override in the selected stack instances.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
							setplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Regions
					"regions": schema.SetAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The names of one or more Regions where you want to create stack instances using the specified AWS account(s).",
						Required:    true,
						Validators: []validator.Set{ /*START VALIDATORS*/
							setvalidator.SizeAtLeast(1),
							setvalidator.ValueStringsAre(
								stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-]{1,128}$"), ""),
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A group of stack instances with parameters in some specific accounts and regions.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// StackInstancesGroup is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: StackSetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the stack set that you're creating.",
		//	  "type": "string"
		//	}
		"stack_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the stack set that you're creating.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StackSetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name to associate with the stack set. The name must be unique in the Region where you create your stack set.",
		//	  "maxLength": 128,
		//	  "pattern": "^[a-zA-Z][a-zA-Z0-9\\-]{0,127}$",
		//	  "type": "string"
		//	}
		"stack_set_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name to associate with the stack set. The name must be unique in the Region where you create your stack set.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9\\-]{0,127}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Tag type enables you to specify a key-value pair that can be used to store information about an AWS CloudFormation StackSet.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A string used to identify this tag. You can specify a maximum of 127 characters for a tag key.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string used to identify this tag. You can specify a maximum of 127 characters for a tag key.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TemplateBody
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.",
		//	  "maxLength": 51200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"template_body": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 51200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TemplateURL
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Location of file containing the template body. The URL must point to a template (max size: 460,800 bytes) that is located in an Amazon S3 bucket.",
		//	  "maxLength": 5120,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"template_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Location of file containing the template body. The URL must point to a template (max size: 460,800 bytes) that is located in an Amazon S3 bucket.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 5120),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// TemplateURL is a write-only property.
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
		Description: "StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::StackSet").WithTerraformTypeName("awscc_cloudformation_stack_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_filter_type":              "AccountFilterType",
		"accounts":                         "Accounts",
		"active":                           "Active",
		"administration_role_arn":          "AdministrationRoleARN",
		"auto_deployment":                  "AutoDeployment",
		"call_as":                          "CallAs",
		"capabilities":                     "Capabilities",
		"deployment_targets":               "DeploymentTargets",
		"description":                      "Description",
		"enabled":                          "Enabled",
		"execution_role_name":              "ExecutionRoleName",
		"failure_tolerance_count":          "FailureToleranceCount",
		"failure_tolerance_percentage":     "FailureTolerancePercentage",
		"key":                              "Key",
		"managed_execution":                "ManagedExecution",
		"max_concurrent_count":             "MaxConcurrentCount",
		"max_concurrent_percentage":        "MaxConcurrentPercentage",
		"operation_preferences":            "OperationPreferences",
		"organizational_unit_ids":          "OrganizationalUnitIds",
		"parameter_key":                    "ParameterKey",
		"parameter_overrides":              "ParameterOverrides",
		"parameter_value":                  "ParameterValue",
		"parameters":                       "Parameters",
		"permission_model":                 "PermissionModel",
		"region_concurrency_type":          "RegionConcurrencyType",
		"region_order":                     "RegionOrder",
		"regions":                          "Regions",
		"retain_stacks_on_account_removal": "RetainStacksOnAccountRemoval",
		"stack_instances_group":            "StackInstancesGroup",
		"stack_set_id":                     "StackSetId",
		"stack_set_name":                   "StackSetName",
		"tags":                             "Tags",
		"template_body":                    "TemplateBody",
		"template_url":                     "TemplateURL",
		"value":                            "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/TemplateURL",
		"/properties/OperationPreferences",
		"/properties/StackInstancesGroup",
		"/properties/CallAs",
	})
	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(2160)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
