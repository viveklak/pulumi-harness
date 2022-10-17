# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ResourceGroupArgs', 'ResourceGroup']

@pulumi.input_type
class ResourceGroupArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[str],
                 identifier: pulumi.Input[str],
                 allowed_scope_levels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 included_scopes: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupIncludedScopeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_filters: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupResourceFilterArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ResourceGroup resource.
        :param pulumi.Input[str] account_id: Account Identifier of the account
        :param pulumi.Input[str] identifier: Unique identifier of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_scope_levels: The scope levels at which this resource group can be used
        :param pulumi.Input[str] color: Color of the environment.
        :param pulumi.Input[str] description: Description of the resource.
        :param pulumi.Input[Sequence[pulumi.Input['ResourceGroupIncludedScopeArgs']]] included_scopes: Included scopes
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] org_id: Unique identifier of the organization.
        :param pulumi.Input[str] project_id: Unique identifier of the project.
        :param pulumi.Input[Sequence[pulumi.Input['ResourceGroupResourceFilterArgs']]] resource_filters: Contains resource filter for a resource group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "identifier", identifier)
        if allowed_scope_levels is not None:
            pulumi.set(__self__, "allowed_scope_levels", allowed_scope_levels)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if included_scopes is not None:
            pulumi.set(__self__, "included_scopes", included_scopes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if resource_filters is not None:
            pulumi.set(__self__, "resource_filters", resource_filters)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[str]:
        """
        Account Identifier of the account
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Input[str]:
        """
        Unique identifier of the resource.
        """
        return pulumi.get(self, "identifier")

    @identifier.setter
    def identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "identifier", value)

    @property
    @pulumi.getter(name="allowedScopeLevels")
    def allowed_scope_levels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The scope levels at which this resource group can be used
        """
        return pulumi.get(self, "allowed_scope_levels")

    @allowed_scope_levels.setter
    def allowed_scope_levels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_scope_levels", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        """
        Color of the environment.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="includedScopes")
    def included_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupIncludedScopeArgs']]]]:
        """
        Included scopes
        """
        return pulumi.get(self, "included_scopes")

    @included_scopes.setter
    def included_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupIncludedScopeArgs']]]]):
        pulumi.set(self, "included_scopes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier of the organization.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="resourceFilters")
    def resource_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupResourceFilterArgs']]]]:
        """
        Contains resource filter for a resource group
        """
        return pulumi.get(self, "resource_filters")

    @resource_filters.setter
    def resource_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupResourceFilterArgs']]]]):
        pulumi.set(self, "resource_filters", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ResourceGroupState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 allowed_scope_levels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 included_scopes: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupIncludedScopeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_filters: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupResourceFilterArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ResourceGroup resources.
        :param pulumi.Input[str] account_id: Account Identifier of the account
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_scope_levels: The scope levels at which this resource group can be used
        :param pulumi.Input[str] color: Color of the environment.
        :param pulumi.Input[str] description: Description of the resource.
        :param pulumi.Input[str] identifier: Unique identifier of the resource.
        :param pulumi.Input[Sequence[pulumi.Input['ResourceGroupIncludedScopeArgs']]] included_scopes: Included scopes
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] org_id: Unique identifier of the organization.
        :param pulumi.Input[str] project_id: Unique identifier of the project.
        :param pulumi.Input[Sequence[pulumi.Input['ResourceGroupResourceFilterArgs']]] resource_filters: Contains resource filter for a resource group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if allowed_scope_levels is not None:
            pulumi.set(__self__, "allowed_scope_levels", allowed_scope_levels)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if identifier is not None:
            pulumi.set(__self__, "identifier", identifier)
        if included_scopes is not None:
            pulumi.set(__self__, "included_scopes", included_scopes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if resource_filters is not None:
            pulumi.set(__self__, "resource_filters", resource_filters)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Account Identifier of the account
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="allowedScopeLevels")
    def allowed_scope_levels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The scope levels at which this resource group can be used
        """
        return pulumi.get(self, "allowed_scope_levels")

    @allowed_scope_levels.setter
    def allowed_scope_levels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_scope_levels", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        """
        Color of the environment.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def identifier(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier of the resource.
        """
        return pulumi.get(self, "identifier")

    @identifier.setter
    def identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identifier", value)

    @property
    @pulumi.getter(name="includedScopes")
    def included_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupIncludedScopeArgs']]]]:
        """
        Included scopes
        """
        return pulumi.get(self, "included_scopes")

    @included_scopes.setter
    def included_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupIncludedScopeArgs']]]]):
        pulumi.set(self, "included_scopes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier of the organization.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier of the project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="resourceFilters")
    def resource_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupResourceFilterArgs']]]]:
        """
        Contains resource filter for a resource group
        """
        return pulumi.get(self, "resource_filters")

    @resource_filters.setter
    def resource_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResourceGroupResourceFilterArgs']]]]):
        pulumi.set(self, "resource_filters", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class ResourceGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 allowed_scope_levels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 included_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupIncludedScopeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupResourceFilterArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource for creating a Harness Resource Group

        ## Import

        Import using resource group id

        ```sh
         $ pulumi import harness:platform/resourceGroup:ResourceGroup example <resource_group_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: Account Identifier of the account
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_scope_levels: The scope levels at which this resource group can be used
        :param pulumi.Input[str] color: Color of the environment.
        :param pulumi.Input[str] description: Description of the resource.
        :param pulumi.Input[str] identifier: Unique identifier of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupIncludedScopeArgs']]]] included_scopes: Included scopes
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] org_id: Unique identifier of the organization.
        :param pulumi.Input[str] project_id: Unique identifier of the project.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupResourceFilterArgs']]]] resource_filters: Contains resource filter for a resource group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for creating a Harness Resource Group

        ## Import

        Import using resource group id

        ```sh
         $ pulumi import harness:platform/resourceGroup:ResourceGroup example <resource_group_id>
        ```

        :param str resource_name: The name of the resource.
        :param ResourceGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 allowed_scope_levels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 included_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupIncludedScopeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupResourceFilterArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceGroupArgs.__new__(ResourceGroupArgs)

            if account_id is None and not opts.urn:
                raise TypeError("Missing required property 'account_id'")
            __props__.__dict__["account_id"] = account_id
            __props__.__dict__["allowed_scope_levels"] = allowed_scope_levels
            __props__.__dict__["color"] = color
            __props__.__dict__["description"] = description
            if identifier is None and not opts.urn:
                raise TypeError("Missing required property 'identifier'")
            __props__.__dict__["identifier"] = identifier
            __props__.__dict__["included_scopes"] = included_scopes
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["resource_filters"] = resource_filters
            __props__.__dict__["tags"] = tags
        super(ResourceGroup, __self__).__init__(
            'harness:platform/resourceGroup:ResourceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            allowed_scope_levels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            color: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            included_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupIncludedScopeArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            resource_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupResourceFilterArgs']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ResourceGroup':
        """
        Get an existing ResourceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: Account Identifier of the account
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_scope_levels: The scope levels at which this resource group can be used
        :param pulumi.Input[str] color: Color of the environment.
        :param pulumi.Input[str] description: Description of the resource.
        :param pulumi.Input[str] identifier: Unique identifier of the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupIncludedScopeArgs']]]] included_scopes: Included scopes
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] org_id: Unique identifier of the organization.
        :param pulumi.Input[str] project_id: Unique identifier of the project.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResourceGroupResourceFilterArgs']]]] resource_filters: Contains resource filter for a resource group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceGroupState.__new__(_ResourceGroupState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["allowed_scope_levels"] = allowed_scope_levels
        __props__.__dict__["color"] = color
        __props__.__dict__["description"] = description
        __props__.__dict__["identifier"] = identifier
        __props__.__dict__["included_scopes"] = included_scopes
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["resource_filters"] = resource_filters
        __props__.__dict__["tags"] = tags
        return ResourceGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        Account Identifier of the account
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="allowedScopeLevels")
    def allowed_scope_levels(self) -> pulumi.Output[Sequence[str]]:
        """
        The scope levels at which this resource group can be used
        """
        return pulumi.get(self, "allowed_scope_levels")

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[str]:
        """
        Color of the environment.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[str]:
        """
        Unique identifier of the resource.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter(name="includedScopes")
    def included_scopes(self) -> pulumi.Output[Optional[Sequence['outputs.ResourceGroupIncludedScope']]]:
        """
        Included scopes
        """
        return pulumi.get(self, "included_scopes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        Unique identifier of the organization.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[str]]:
        """
        Unique identifier of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resourceFilters")
    def resource_filters(self) -> pulumi.Output[Optional[Sequence['outputs.ResourceGroupResourceFilter']]]:
        """
        Contains resource filter for a resource group
        """
        return pulumi.get(self, "resource_filters")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        return pulumi.get(self, "tags")

