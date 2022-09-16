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

    public sealed class SshCredentialSshAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Inline SSH authentication configuration. Only ond of `passphrase_secret_id` or `ssh_key_file_id` should be used
        /// </summary>
        [Input("inlineSsh")]
        public Input<Inputs.SshCredentialSshAuthenticationInlineSshArgs>? InlineSsh { get; set; }

        /// <summary>
        /// The port to connect to
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Server password authentication configuration
        /// </summary>
        [Input("serverPassword")]
        public Input<Inputs.SshCredentialSshAuthenticationServerPasswordArgs>? ServerPassword { get; set; }

        /// <summary>
        /// Use ssh key file for authentication
        /// </summary>
        [Input("sshKeyFile")]
        public Input<Inputs.SshCredentialSshAuthenticationSshKeyFileArgs>? SshKeyFile { get; set; }

        /// <summary>
        /// The username to use when connecting to ssh
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SshCredentialSshAuthenticationArgs()
        {
        }
        public static new SshCredentialSshAuthenticationArgs Empty => new SshCredentialSshAuthenticationArgs();
    }
}