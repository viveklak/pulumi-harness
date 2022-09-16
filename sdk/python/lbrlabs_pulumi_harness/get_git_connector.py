# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetGitConnectorResult',
    'AwaitableGetGitConnectorResult',
    'get_git_connector',
    'get_git_connector_output',
]

@pulumi.output_type
class GetGitConnectorResult:
    """
    A collection of values returned by getGitConnector.
    """
    def __init__(__self__, branch=None, commit_details=None, created_at=None, delegate_selectors=None, generate_webhook_url=None, id=None, name=None, password_secret_id=None, ssh_setting_id=None, url=None, url_type=None, username=None, webhook_url=None):
        if branch and not isinstance(branch, str):
            raise TypeError("Expected argument 'branch' to be a str")
        pulumi.set(__self__, "branch", branch)
        if commit_details and not isinstance(commit_details, list):
            raise TypeError("Expected argument 'commit_details' to be a list")
        pulumi.set(__self__, "commit_details", commit_details)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if delegate_selectors and not isinstance(delegate_selectors, list):
            raise TypeError("Expected argument 'delegate_selectors' to be a list")
        pulumi.set(__self__, "delegate_selectors", delegate_selectors)
        if generate_webhook_url and not isinstance(generate_webhook_url, bool):
            raise TypeError("Expected argument 'generate_webhook_url' to be a bool")
        pulumi.set(__self__, "generate_webhook_url", generate_webhook_url)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if password_secret_id and not isinstance(password_secret_id, str):
            raise TypeError("Expected argument 'password_secret_id' to be a str")
        pulumi.set(__self__, "password_secret_id", password_secret_id)
        if ssh_setting_id and not isinstance(ssh_setting_id, str):
            raise TypeError("Expected argument 'ssh_setting_id' to be a str")
        pulumi.set(__self__, "ssh_setting_id", ssh_setting_id)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if url_type and not isinstance(url_type, str):
            raise TypeError("Expected argument 'url_type' to be a str")
        pulumi.set(__self__, "url_type", url_type)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if webhook_url and not isinstance(webhook_url, str):
            raise TypeError("Expected argument 'webhook_url' to be a str")
        pulumi.set(__self__, "webhook_url", webhook_url)

    @property
    @pulumi.getter
    def branch(self) -> str:
        """
        The branch of the git connector to use.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="commitDetails")
    def commit_details(self) -> Sequence['outputs.GetGitConnectorCommitDetailResult']:
        """
        Custom details to use when making commits using this git connector.
        """
        return pulumi.get(self, "commit_details")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The time the git connector was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="delegateSelectors")
    def delegate_selectors(self) -> Sequence[str]:
        """
        Delegate selectors to apply to this git connector.
        """
        return pulumi.get(self, "delegate_selectors")

    @property
    @pulumi.getter(name="generateWebhookUrl")
    def generate_webhook_url(self) -> bool:
        """
        Boolean indicating whether or not to generate a webhook url.
        """
        return pulumi.get(self, "generate_webhook_url")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Id of the git connector.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the git connector.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="passwordSecretId")
    def password_secret_id(self) -> str:
        """
        The id of the secret for connecting to the git repository.
        """
        return pulumi.get(self, "password_secret_id")

    @property
    @pulumi.getter(name="sshSettingId")
    def ssh_setting_id(self) -> str:
        """
        The id of the SSH secret to use.
        """
        return pulumi.get(self, "ssh_setting_id")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The url of the git repository or account/organization.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="urlType")
    def url_type(self) -> str:
        """
        The type of git url being used. Options are `ACCOUNT`, and `REPO`.
        """
        return pulumi.get(self, "url_type")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The name of the user used to connect to the git repository.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter(name="webhookUrl")
    def webhook_url(self) -> str:
        """
        The generated webhook url.
        """
        return pulumi.get(self, "webhook_url")


class AwaitableGetGitConnectorResult(GetGitConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGitConnectorResult(
            branch=self.branch,
            commit_details=self.commit_details,
            created_at=self.created_at,
            delegate_selectors=self.delegate_selectors,
            generate_webhook_url=self.generate_webhook_url,
            id=self.id,
            name=self.name,
            password_secret_id=self.password_secret_id,
            ssh_setting_id=self.ssh_setting_id,
            url=self.url,
            url_type=self.url_type,
            username=self.username,
            webhook_url=self.webhook_url)


def get_git_connector(id: Optional[str] = None,
                      name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGitConnectorResult:
    """
    Data source for retrieving a Harness application


    :param str id: Id of the git connector.
    :param str name: The name of the git connector.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('harness:index/getGitConnector:getGitConnector', __args__, opts=opts, typ=GetGitConnectorResult).value

    return AwaitableGetGitConnectorResult(
        branch=__ret__.branch,
        commit_details=__ret__.commit_details,
        created_at=__ret__.created_at,
        delegate_selectors=__ret__.delegate_selectors,
        generate_webhook_url=__ret__.generate_webhook_url,
        id=__ret__.id,
        name=__ret__.name,
        password_secret_id=__ret__.password_secret_id,
        ssh_setting_id=__ret__.ssh_setting_id,
        url=__ret__.url,
        url_type=__ret__.url_type,
        username=__ret__.username,
        webhook_url=__ret__.webhook_url)


@_utilities.lift_output_func(get_git_connector)
def get_git_connector_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGitConnectorResult]:
    """
    Data source for retrieving a Harness application


    :param str id: Id of the git connector.
    :param str name: The name of the git connector.
    """
    ...