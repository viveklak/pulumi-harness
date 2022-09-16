# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DelegateApprovalArgs', 'DelegateApproval']

@pulumi.input_type
class DelegateApprovalArgs:
    def __init__(__self__, *,
                 approve: pulumi.Input[bool],
                 delegate_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a DelegateApproval resource.
        :param pulumi.Input[bool] approve: Whether or not to approve the delegate.
        :param pulumi.Input[str] delegate_id: The id of the delegate.
        """
        pulumi.set(__self__, "approve", approve)
        pulumi.set(__self__, "delegate_id", delegate_id)

    @property
    @pulumi.getter
    def approve(self) -> pulumi.Input[bool]:
        """
        Whether or not to approve the delegate.
        """
        return pulumi.get(self, "approve")

    @approve.setter
    def approve(self, value: pulumi.Input[bool]):
        pulumi.set(self, "approve", value)

    @property
    @pulumi.getter(name="delegateId")
    def delegate_id(self) -> pulumi.Input[str]:
        """
        The id of the delegate.
        """
        return pulumi.get(self, "delegate_id")

    @delegate_id.setter
    def delegate_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "delegate_id", value)


@pulumi.input_type
class _DelegateApprovalState:
    def __init__(__self__, *,
                 approve: Optional[pulumi.Input[bool]] = None,
                 delegate_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DelegateApproval resources.
        :param pulumi.Input[bool] approve: Whether or not to approve the delegate.
        :param pulumi.Input[str] delegate_id: The id of the delegate.
        :param pulumi.Input[str] status: The status of the delegate.
        """
        if approve is not None:
            pulumi.set(__self__, "approve", approve)
        if delegate_id is not None:
            pulumi.set(__self__, "delegate_id", delegate_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def approve(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to approve the delegate.
        """
        return pulumi.get(self, "approve")

    @approve.setter
    def approve(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "approve", value)

    @property
    @pulumi.getter(name="delegateId")
    def delegate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the delegate.
        """
        return pulumi.get(self, "delegate_id")

    @delegate_id.setter
    def delegate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delegate_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the delegate.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class DelegateApproval(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approve: Optional[pulumi.Input[bool]] = None,
                 delegate_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource for approving or rejecting delegates.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_harness as harness
        import pulumi_harness as harness

        test_delegate = harness.get_delegate(name="my-delegate")
        test_delegate_approval = harness.DelegateApproval("testDelegateApproval",
            delegate_id=test_delegate.id,
            approve=True)
        ```

        ## Import

        # Import the status of the delegate approval.

        ```sh
         $ pulumi import harness:index/delegateApproval:DelegateApproval example <delegate_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] approve: Whether or not to approve the delegate.
        :param pulumi.Input[str] delegate_id: The id of the delegate.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DelegateApprovalArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for approving or rejecting delegates.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_harness as harness
        import pulumi_harness as harness

        test_delegate = harness.get_delegate(name="my-delegate")
        test_delegate_approval = harness.DelegateApproval("testDelegateApproval",
            delegate_id=test_delegate.id,
            approve=True)
        ```

        ## Import

        # Import the status of the delegate approval.

        ```sh
         $ pulumi import harness:index/delegateApproval:DelegateApproval example <delegate_id>
        ```

        :param str resource_name: The name of the resource.
        :param DelegateApprovalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DelegateApprovalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approve: Optional[pulumi.Input[bool]] = None,
                 delegate_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DelegateApprovalArgs.__new__(DelegateApprovalArgs)

            if approve is None and not opts.urn:
                raise TypeError("Missing required property 'approve'")
            __props__.__dict__["approve"] = approve
            if delegate_id is None and not opts.urn:
                raise TypeError("Missing required property 'delegate_id'")
            __props__.__dict__["delegate_id"] = delegate_id
            __props__.__dict__["status"] = None
        super(DelegateApproval, __self__).__init__(
            'harness:index/delegateApproval:DelegateApproval',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            approve: Optional[pulumi.Input[bool]] = None,
            delegate_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'DelegateApproval':
        """
        Get an existing DelegateApproval resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] approve: Whether or not to approve the delegate.
        :param pulumi.Input[str] delegate_id: The id of the delegate.
        :param pulumi.Input[str] status: The status of the delegate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DelegateApprovalState.__new__(_DelegateApprovalState)

        __props__.__dict__["approve"] = approve
        __props__.__dict__["delegate_id"] = delegate_id
        __props__.__dict__["status"] = status
        return DelegateApproval(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def approve(self) -> pulumi.Output[bool]:
        """
        Whether or not to approve the delegate.
        """
        return pulumi.get(self, "approve")

    @property
    @pulumi.getter(name="delegateId")
    def delegate_id(self) -> pulumi.Output[str]:
        """
        The id of the delegate.
        """
        return pulumi.get(self, "delegate_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the delegate.
        """
        return pulumi.get(self, "status")
