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
    public static class GetAwsKmsConnector
    {
        /// <summary>
        /// Datasource for looking up an AWS KMS connector.
        /// </summary>
        public static Task<GetAwsKmsConnectorResult> InvokeAsync(GetAwsKmsConnectorArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAwsKmsConnectorResult>("harness:platform/getAwsKmsConnector:getAwsKmsConnector", args ?? new GetAwsKmsConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource for looking up an AWS KMS connector.
        /// </summary>
        public static Output<GetAwsKmsConnectorResult> Invoke(GetAwsKmsConnectorInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAwsKmsConnectorResult>("harness:platform/getAwsKmsConnector:getAwsKmsConnector", args ?? new GetAwsKmsConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAwsKmsConnectorArgs : global::Pulumi.InvokeArgs
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

        public GetAwsKmsConnectorArgs()
        {
        }
        public static new GetAwsKmsConnectorArgs Empty => new GetAwsKmsConnectorArgs();
    }

    public sealed class GetAwsKmsConnectorInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetAwsKmsConnectorInvokeArgs()
        {
        }
        public static new GetAwsKmsConnectorInvokeArgs Empty => new GetAwsKmsConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetAwsKmsConnectorResult
    {
        /// <summary>
        /// A reference to the Harness secret containing the ARN of the AWS KMS.
        /// </summary>
        public readonly string ArnRef;
        /// <summary>
        /// The credentials to use for connecting to aws.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAwsKmsConnectorCredentialResult> Credentials;
        /// <summary>
        /// Connect using only the delegates which have these tags.
        /// </summary>
        public readonly ImmutableArray<string> DelegateSelectors;
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
        /// The AWS region where the AWS Secret Manager is.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Tags to associate with the resource. Tags should be in the form `name:value`.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetAwsKmsConnectorResult(
            string arnRef,

            ImmutableArray<Outputs.GetAwsKmsConnectorCredentialResult> credentials,

            ImmutableArray<string> delegateSelectors,

            string description,

            string id,

            string? identifier,

            string? name,

            string? orgId,

            string? projectId,

            string region,

            ImmutableArray<string> tags)
        {
            ArnRef = arnRef;
            Credentials = credentials;
            DelegateSelectors = delegateSelectors;
            Description = description;
            Id = id;
            Identifier = identifier;
            Name = name;
            OrgId = orgId;
            ProjectId = projectId;
            Region = region;
            Tags = tags;
        }
    }
}
