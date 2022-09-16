// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGithubConnectorCredentialSsh {
    private String sshKeyRef;

    private GetGithubConnectorCredentialSsh() {}
    public String sshKeyRef() {
        return this.sshKeyRef;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGithubConnectorCredentialSsh defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String sshKeyRef;
        public Builder() {}
        public Builder(GetGithubConnectorCredentialSsh defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.sshKeyRef = defaults.sshKeyRef;
        }

        @CustomType.Setter
        public Builder sshKeyRef(String sshKeyRef) {
            this.sshKeyRef = Objects.requireNonNull(sshKeyRef);
            return this;
        }
        public GetGithubConnectorCredentialSsh build() {
            final var o = new GetGithubConnectorCredentialSsh();
            o.sshKeyRef = sshKeyRef;
            return o;
        }
    }
}