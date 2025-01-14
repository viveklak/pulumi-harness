# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetRolesResult',
    'AwaitableGetRolesResult',
    'get_roles',
    'get_roles_output',
]

@pulumi.output_type
class GetRolesResult:
    """
    A collection of values returned by getRoles.
    """
    def __init__(__self__, allowed_scope_levels=None, description=None, id=None, identifier=None, name=None, org_id=None, permissions=None, project_id=None, tags=None):
        if allowed_scope_levels and not isinstance(allowed_scope_levels, list):
            raise TypeError("Expected argument 'allowed_scope_levels' to be a list")
        pulumi.set(__self__, "allowed_scope_levels", allowed_scope_levels)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identifier and not isinstance(identifier, str):
            raise TypeError("Expected argument 'identifier' to be a str")
        pulumi.set(__self__, "identifier", identifier)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="allowedScopeLevels")
    def allowed_scope_levels(self) -> Sequence[str]:
        """
        The scope levels at which this role can be used
        """
        return pulumi.get(self, "allowed_scope_levels")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identifier(self) -> Optional[str]:
        """
        Unique identifier of the resource.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[str]:
        """
        Unique identifier of the organization.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def permissions(self) -> Optional[Sequence[str]]:
        """
        List of the permission identifiers
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        Unique identifier of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        return pulumi.get(self, "tags")


class AwaitableGetRolesResult(GetRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRolesResult(
            allowed_scope_levels=self.allowed_scope_levels,
            description=self.description,
            id=self.id,
            identifier=self.identifier,
            name=self.name,
            org_id=self.org_id,
            permissions=self.permissions,
            project_id=self.project_id,
            tags=self.tags)


def get_roles(allowed_scope_levels: Optional[Sequence[str]] = None,
              identifier: Optional[str] = None,
              name: Optional[str] = None,
              org_id: Optional[str] = None,
              permissions: Optional[Sequence[str]] = None,
              project_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRolesResult:
    """
    Data source for retrieving roles

    ## Example Usage

    ```python
    import pulumi
    import pulumi_harness as harness

    example = harness.platform.get_roles(identifier="identifier",
        org_id="org_id",
        project_id="project_id")
    ```


    :param Sequence[str] allowed_scope_levels: The scope levels at which this role can be used
    :param str identifier: Unique identifier of the resource.
    :param str name: Name of the resource.
    :param str org_id: Unique identifier of the organization.
    :param Sequence[str] permissions: List of the permission identifiers
    :param str project_id: Unique identifier of the project.
    """
    __args__ = dict()
    __args__['allowedScopeLevels'] = allowed_scope_levels
    __args__['identifier'] = identifier
    __args__['name'] = name
    __args__['orgId'] = org_id
    __args__['permissions'] = permissions
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('harness:platform/getRoles:getRoles', __args__, opts=opts, typ=GetRolesResult).value

    return AwaitableGetRolesResult(
        allowed_scope_levels=__ret__.allowed_scope_levels,
        description=__ret__.description,
        id=__ret__.id,
        identifier=__ret__.identifier,
        name=__ret__.name,
        org_id=__ret__.org_id,
        permissions=__ret__.permissions,
        project_id=__ret__.project_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_roles)
def get_roles_output(allowed_scope_levels: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     identifier: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     org_id: Optional[pulumi.Input[Optional[str]]] = None,
                     permissions: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     project_id: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRolesResult]:
    """
    Data source for retrieving roles

    ## Example Usage

    ```python
    import pulumi
    import pulumi_harness as harness

    example = harness.platform.get_roles(identifier="identifier",
        org_id="org_id",
        project_id="project_id")
    ```


    :param Sequence[str] allowed_scope_levels: The scope levels at which this role can be used
    :param str identifier: Unique identifier of the resource.
    :param str name: Name of the resource.
    :param str org_id: Unique identifier of the organization.
    :param Sequence[str] permissions: List of the permission identifiers
    :param str project_id: Unique identifier of the project.
    """
    ...
