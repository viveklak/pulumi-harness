# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['YamlConfigArgs', 'YamlConfig']

@pulumi.input_type
class YamlConfigArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 path: pulumi.Input[str],
                 app_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a YamlConfig resource.
        :param pulumi.Input[str] content: The raw YAML configuration.
        :param pulumi.Input[str] path: The path of the resource.
        :param pulumi.Input[str] app_id: The id of the application. This is required for all resources except global ones.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "path", path)
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        The raw YAML configuration.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The path of the resource.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the application. This is required for all resources except global ones.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)


@pulumi.input_type
class _YamlConfigState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering YamlConfig resources.
        :param pulumi.Input[str] app_id: The id of the application. This is required for all resources except global ones.
        :param pulumi.Input[str] content: The raw YAML configuration.
        :param pulumi.Input[str] name: The name of the resource.
        :param pulumi.Input[str] path: The path of the resource.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the application. This is required for all resources except global ones.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The raw YAML configuration.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The path of the resource.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)


class YamlConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource for creating a raw YAML configuration in Harness. Note: This works for all objects EXCEPT application objects. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `create_before_destroy = true` lifecycle setting.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_harness as harness

        test = harness.YamlConfig("test",
            content=\"\"\"harnessApiVersion: '1.0'
        type: KUBERNETES_CLUSTER
        delegateSelectors:
        - k8s
        skipValidation: true
        useKubernetesDelegate: true

        \"\"\",
            path="Setup/Cloud Providers/Kubernetes.yaml")
        ```

        ## Import

        Importing a global config only using the yaml path

        ```sh
         $ pulumi import harness:index/yamlConfig:YamlConfig k8s_cloudprovider "Setup/Cloud Providers/kubernetes.yaml"
        ```

         Importing a service which requires both the application id and the yaml path.

        ```sh
         $ pulumi import harness:index/yamlConfig:YamlConfig k8s_cloudprovider "Setup/Applications/MyApp/Services/MyService/Index.yaml:<APPLICATION_ID>"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The id of the application. This is required for all resources except global ones.
        :param pulumi.Input[str] content: The raw YAML configuration.
        :param pulumi.Input[str] path: The path of the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: YamlConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for creating a raw YAML configuration in Harness. Note: This works for all objects EXCEPT application objects. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `create_before_destroy = true` lifecycle setting.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_harness as harness

        test = harness.YamlConfig("test",
            content=\"\"\"harnessApiVersion: '1.0'
        type: KUBERNETES_CLUSTER
        delegateSelectors:
        - k8s
        skipValidation: true
        useKubernetesDelegate: true

        \"\"\",
            path="Setup/Cloud Providers/Kubernetes.yaml")
        ```

        ## Import

        Importing a global config only using the yaml path

        ```sh
         $ pulumi import harness:index/yamlConfig:YamlConfig k8s_cloudprovider "Setup/Cloud Providers/kubernetes.yaml"
        ```

         Importing a service which requires both the application id and the yaml path.

        ```sh
         $ pulumi import harness:index/yamlConfig:YamlConfig k8s_cloudprovider "Setup/Applications/MyApp/Services/MyService/Index.yaml:<APPLICATION_ID>"
        ```

        :param str resource_name: The name of the resource.
        :param YamlConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(YamlConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = YamlConfigArgs.__new__(YamlConfigArgs)

            __props__.__dict__["app_id"] = app_id
            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__.__dict__["content"] = content
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["name"] = None
        super(YamlConfig, __self__).__init__(
            'harness:index/yamlConfig:YamlConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None) -> 'YamlConfig':
        """
        Get an existing YamlConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The id of the application. This is required for all resources except global ones.
        :param pulumi.Input[str] content: The raw YAML configuration.
        :param pulumi.Input[str] name: The name of the resource.
        :param pulumi.Input[str] path: The path of the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _YamlConfigState.__new__(_YamlConfigState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["content"] = content
        __props__.__dict__["name"] = name
        __props__.__dict__["path"] = path
        return YamlConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id of the application. This is required for all resources except global ones.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The raw YAML configuration.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path of the resource.
        """
        return pulumi.get(self, "path")

