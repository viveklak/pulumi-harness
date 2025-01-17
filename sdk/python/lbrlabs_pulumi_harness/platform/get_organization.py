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
    'GetOrganizationResult',
    'AwaitableGetOrganizationResult',
    'get_organization',
    'get_organization_output',
]

@pulumi.output_type
class GetOrganizationResult:
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, description=None, id=None, identifier=None, name=None, tags=None):
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
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        return pulumi.get(self, "tags")


class AwaitableGetOrganizationResult(GetOrganizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationResult(
            description=self.description,
            id=self.id,
            identifier=self.identifier,
            name=self.name,
            tags=self.tags)


def get_organization(identifier: Optional[str] = None,
                     name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationResult:
    """
    Data source for retrieving a Harness organization


    :param str identifier: Unique identifier of the resource.
    :param str name: Name of the resource.
    """
    __args__ = dict()
    __args__['identifier'] = identifier
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('harness:platform/getOrganization:getOrganization', __args__, opts=opts, typ=GetOrganizationResult).value

    return AwaitableGetOrganizationResult(
        description=__ret__.description,
        id=__ret__.id,
        identifier=__ret__.identifier,
        name=__ret__.name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_organization)
def get_organization_output(identifier: Optional[pulumi.Input[Optional[str]]] = None,
                            name: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationResult]:
    """
    Data source for retrieving a Harness organization


    :param str identifier: Unique identifier of the resource.
    :param str name: Name of the resource.
    """
    ...
