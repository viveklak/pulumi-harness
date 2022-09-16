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

__all__ = [
    'GetAwsCCConnectorResult',
    'AwaitableGetAwsCCConnectorResult',
    'get_aws_cc_connector',
    'get_aws_cc_connector_output',
]

@pulumi.output_type
class GetAwsCCConnectorResult:
    """
    A collection of values returned by getAwsCCConnector.
    """
    def __init__(__self__, account_id=None, cross_account_accesses=None, description=None, features_enableds=None, id=None, identifier=None, name=None, org_id=None, project_id=None, report_name=None, s3_bucket=None, tags=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if cross_account_accesses and not isinstance(cross_account_accesses, list):
            raise TypeError("Expected argument 'cross_account_accesses' to be a list")
        pulumi.set(__self__, "cross_account_accesses", cross_account_accesses)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if features_enableds and not isinstance(features_enableds, list):
            raise TypeError("Expected argument 'features_enableds' to be a list")
        pulumi.set(__self__, "features_enableds", features_enableds)
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
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if report_name and not isinstance(report_name, str):
            raise TypeError("Expected argument 'report_name' to be a str")
        pulumi.set(__self__, "report_name", report_name)
        if s3_bucket and not isinstance(s3_bucket, str):
            raise TypeError("Expected argument 's3_bucket' to be a str")
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The AWS account id.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="crossAccountAccesses")
    def cross_account_accesses(self) -> Sequence['outputs.GetAwsCCConnectorCrossAccountAccessResult']:
        """
        Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.
        """
        return pulumi.get(self, "cross_account_accesses")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="featuresEnableds")
    def features_enableds(self) -> Sequence[str]:
        """
        The features enabled for the connector. Valid values are BILLING, OPTIMIZATION, VISIBILITY.
        """
        return pulumi.get(self, "features_enableds")

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        Unique identifier of the project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="reportName")
    def report_name(self) -> str:
        """
        The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.
        """
        return pulumi.get(self, "report_name")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> str:
        """
        The name of s3 bucket.
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags to associate with the resource. Tags should be in the form `name:value`.
        """
        return pulumi.get(self, "tags")


class AwaitableGetAwsCCConnectorResult(GetAwsCCConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAwsCCConnectorResult(
            account_id=self.account_id,
            cross_account_accesses=self.cross_account_accesses,
            description=self.description,
            features_enableds=self.features_enableds,
            id=self.id,
            identifier=self.identifier,
            name=self.name,
            org_id=self.org_id,
            project_id=self.project_id,
            report_name=self.report_name,
            s3_bucket=self.s3_bucket,
            tags=self.tags)


def get_aws_cc_connector(identifier: Optional[str] = None,
                         name: Optional[str] = None,
                         org_id: Optional[str] = None,
                         project_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAwsCCConnectorResult:
    """
    Datasource for looking up an AWS Cloud Cost connector.


    :param str identifier: Unique identifier of the resource.
    :param str name: Name of the resource.
    :param str org_id: Unique identifier of the organization.
    :param str project_id: Unique identifier of the project.
    """
    __args__ = dict()
    __args__['identifier'] = identifier
    __args__['name'] = name
    __args__['orgId'] = org_id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('harness:platform/getAwsCCConnector:getAwsCCConnector', __args__, opts=opts, typ=GetAwsCCConnectorResult).value

    return AwaitableGetAwsCCConnectorResult(
        account_id=__ret__.account_id,
        cross_account_accesses=__ret__.cross_account_accesses,
        description=__ret__.description,
        features_enableds=__ret__.features_enableds,
        id=__ret__.id,
        identifier=__ret__.identifier,
        name=__ret__.name,
        org_id=__ret__.org_id,
        project_id=__ret__.project_id,
        report_name=__ret__.report_name,
        s3_bucket=__ret__.s3_bucket,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_aws_cc_connector)
def get_aws_cc_connector_output(identifier: Optional[pulumi.Input[Optional[str]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                org_id: Optional[pulumi.Input[Optional[str]]] = None,
                                project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAwsCCConnectorResult]:
    """
    Datasource for looking up an AWS Cloud Cost connector.


    :param str identifier: Unique identifier of the resource.
    :param str name: Name of the resource.
    :param str org_id: Unique identifier of the organization.
    :param str project_id: Unique identifier of the project.
    """
    ...