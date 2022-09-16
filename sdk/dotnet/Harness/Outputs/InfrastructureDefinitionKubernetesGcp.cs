// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.Outputs
{

    [OutputType]
    public sealed class InfrastructureDefinitionKubernetesGcp
    {
        /// <summary>
        /// The name of the cloud provider to connect with.
        /// </summary>
        public readonly string CloudProviderName;
        /// <summary>
        /// The name of the cluster being deployed to.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// The namespace in Kubernetes to deploy to.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// The naming convention of the release.
        /// </summary>
        public readonly string ReleaseName;

        [OutputConstructor]
        private InfrastructureDefinitionKubernetesGcp(
            string cloudProviderName,

            string clusterName,

            string @namespace,

            string releaseName)
        {
            CloudProviderName = cloudProviderName;
            ClusterName = clusterName;
            Namespace = @namespace;
            ReleaseName = releaseName;
        }
    }
}