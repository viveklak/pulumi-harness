// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Harness.Inputs
{

    public sealed class SshCredentialKerberosAuthenticationTgtGenerationMethodGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("kerberosPasswordId")]
        public Input<string>? KerberosPasswordId { get; set; }

        [Input("keyTabFilePath")]
        public Input<string>? KeyTabFilePath { get; set; }

        public SshCredentialKerberosAuthenticationTgtGenerationMethodGetArgs()
        {
        }
        public static new SshCredentialKerberosAuthenticationTgtGenerationMethodGetArgs Empty => new SshCredentialKerberosAuthenticationTgtGenerationMethodGetArgs();
    }
}