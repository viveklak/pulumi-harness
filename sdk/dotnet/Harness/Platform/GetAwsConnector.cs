// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.Platform
{
    public static class GetAwsConnector
    {
        /// <summary>
        /// Datasource for looking up an AWS connector.
        /// </summary>
        public static Task<GetAwsConnectorResult> InvokeAsync(GetAwsConnectorArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAwsConnectorResult>("harness:platform/getAwsConnector:getAwsConnector", args ?? new GetAwsConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource for looking up an AWS connector.
        /// </summary>
        public static Output<GetAwsConnectorResult> Invoke(GetAwsConnectorInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAwsConnectorResult>("harness:platform/getAwsConnector:getAwsConnector", args ?? new GetAwsConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAwsConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of the resource.
        /// </summary>
        [Input("identifier")]
        public string? Identifier { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Unique identifier of the organization.
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetAwsConnectorArgs()
        {
        }
        public static new GetAwsConnectorArgs Empty => new GetAwsConnectorArgs();
    }

    public sealed class GetAwsConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier of the resource.
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Unique identifier of the organization.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetAwsConnectorInvokeArgs()
        {
        }
        public static new GetAwsConnectorInvokeArgs Empty => new GetAwsConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetAwsConnectorResult
    {
        /// <summary>
        /// Select this option if you want to use one AWS account for the connection, but you want to deploy or build in a different AWS account. In this scenario, the AWS account used for AWS access in Credentials will assume the IAM role you specify in Cross-account role ARN setting. This option uses the AWS Security Token Service (STS) feature.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAwsConnectorCrossAccountAccessResult> CrossAccountAccesses;
        /// <summary>
        /// Description of the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Unique identifier of the resource.
        /// </summary>
        public readonly string? Identifier;
        /// <summary>
        /// Inherit credentials from the delegate.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAwsConnectorInheritFromDelegateResult> InheritFromDelegates;
        /// <summary>
        /// Use IAM role for service accounts.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAwsConnectorIrsaResult> Irsas;
        /// <summary>
        /// Use IAM role for service accounts.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAwsConnectorManualResult> Manuals;
        /// <summary>
        /// Name of the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Unique identifier of the organization.
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        public readonly string? ProjectId;
        /// <summary>
        /// Tags to associate with the resource. Tags should be in the form `name:value`.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetAwsConnectorResult(
            ImmutableArray<Outputs.GetAwsConnectorCrossAccountAccessResult> crossAccountAccesses,

            string description,

            string id,

            string? identifier,

            ImmutableArray<Outputs.GetAwsConnectorInheritFromDelegateResult> inheritFromDelegates,

            ImmutableArray<Outputs.GetAwsConnectorIrsaResult> irsas,

            ImmutableArray<Outputs.GetAwsConnectorManualResult> manuals,

            string? name,

            string? orgId,

            string? projectId,

            ImmutableArray<string> tags)
        {
            CrossAccountAccesses = crossAccountAccesses;
            Description = description;
            Id = id;
            Identifier = identifier;
            InheritFromDelegates = inheritFromDelegates;
            Irsas = irsas;
            Manuals = manuals;
            Name = name;
            OrgId = orgId;
            ProjectId = projectId;
            Tags = tags;
        }
    }
}
