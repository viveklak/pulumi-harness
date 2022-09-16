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
    public static class GetRoles
    {
        /// <summary>
        /// Data source for retrieving roles
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Harness = Pulumi.Harness;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Harness.Platform.GetRoles.Invoke(new()
        ///     {
        ///         Identifier = "identifier",
        ///         OrgId = "org_id",
        ///         ProjectId = "project_id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRolesResult> InvokeAsync(GetRolesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRolesResult>("harness:platform/getRoles:getRoles", args ?? new GetRolesArgs(), options.WithDefaults());

        /// <summary>
        /// Data source for retrieving roles
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Harness = Pulumi.Harness;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Harness.Platform.GetRoles.Invoke(new()
        ///     {
        ///         Identifier = "identifier",
        ///         OrgId = "org_id",
        ///         ProjectId = "project_id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRolesResult> Invoke(GetRolesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRolesResult>("harness:platform/getRoles:getRoles", args ?? new GetRolesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRolesArgs : global::Pulumi.InvokeArgs
    {
        [Input("allowedScopeLevels")]
        private List<string>? _allowedScopeLevels;

        /// <summary>
        /// The scope levels at which this role can be used
        /// </summary>
        public List<string> AllowedScopeLevels
        {
            get => _allowedScopeLevels ?? (_allowedScopeLevels = new List<string>());
            set => _allowedScopeLevels = value;
        }

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

        [Input("permissions")]
        private List<string>? _permissions;

        /// <summary>
        /// List of the permission identifiers
        /// </summary>
        public List<string> Permissions
        {
            get => _permissions ?? (_permissions = new List<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetRolesArgs()
        {
        }
        public static new GetRolesArgs Empty => new GetRolesArgs();
    }

    public sealed class GetRolesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("allowedScopeLevels")]
        private InputList<string>? _allowedScopeLevels;

        /// <summary>
        /// The scope levels at which this role can be used
        /// </summary>
        public InputList<string> AllowedScopeLevels
        {
            get => _allowedScopeLevels ?? (_allowedScopeLevels = new InputList<string>());
            set => _allowedScopeLevels = value;
        }

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

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// List of the permission identifiers
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetRolesInvokeArgs()
        {
        }
        public static new GetRolesInvokeArgs Empty => new GetRolesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRolesResult
    {
        /// <summary>
        /// The scope levels at which this role can be used
        /// </summary>
        public readonly ImmutableArray<string> AllowedScopeLevels;
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
        /// List of the permission identifiers
        /// </summary>
        public readonly ImmutableArray<string> Permissions;
        /// <summary>
        /// Unique identifier of the project.
        /// </summary>
        public readonly string? ProjectId;
        /// <summary>
        /// Tags to associate with the resource. Tags should be in the form `name:value`.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetRolesResult(
            ImmutableArray<string> allowedScopeLevels,

            string description,

            string id,

            string? identifier,

            string? name,

            string? orgId,

            ImmutableArray<string> permissions,

            string? projectId,

            ImmutableArray<string> tags)
        {
            AllowedScopeLevels = allowedScopeLevels;
            Description = description;
            Id = id;
            Identifier = identifier;
            Name = name;
            OrgId = orgId;
            Permissions = permissions;
            ProjectId = projectId;
            Tags = tags;
        }
    }
}