// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.Platform.Outputs
{

    [OutputType]
    public sealed class BitbucketConnectorApiAuthentication
    {
        /// <summary>
        /// Personal access token for interacting with the BitBucket api.
        /// </summary>
        public readonly string TokenRef;
        /// <summary>
        /// The username used for connecting to the api.
        /// </summary>
        public readonly string? Username;
        /// <summary>
        /// The name of the Harness secret containing the username.
        /// </summary>
        public readonly string? UsernameRef;

        [OutputConstructor]
        private BitbucketConnectorApiAuthentication(
            string tokenRef,

            string? username,

            string? usernameRef)
        {
            TokenRef = tokenRef;
            Username = username;
            UsernameRef = usernameRef;
        }
    }
}