// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class InfrastructureDefinitionDatacenterSsh {
    /**
     * @return The name of the cloud provider to connect with.
     * 
     */
    private String cloudProviderName;
    /**
     * @return The name of the SSH connection attributes to use.
     * 
     */
    private String hostConnectionAttributesName;
    /**
     * @return A list of hosts to deploy to.
     * 
     */
    private List<String> hostnames;

    private InfrastructureDefinitionDatacenterSsh() {}
    /**
     * @return The name of the cloud provider to connect with.
     * 
     */
    public String cloudProviderName() {
        return this.cloudProviderName;
    }
    /**
     * @return The name of the SSH connection attributes to use.
     * 
     */
    public String hostConnectionAttributesName() {
        return this.hostConnectionAttributesName;
    }
    /**
     * @return A list of hosts to deploy to.
     * 
     */
    public List<String> hostnames() {
        return this.hostnames;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InfrastructureDefinitionDatacenterSsh defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cloudProviderName;
        private String hostConnectionAttributesName;
        private List<String> hostnames;
        public Builder() {}
        public Builder(InfrastructureDefinitionDatacenterSsh defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cloudProviderName = defaults.cloudProviderName;
    	      this.hostConnectionAttributesName = defaults.hostConnectionAttributesName;
    	      this.hostnames = defaults.hostnames;
        }

        @CustomType.Setter
        public Builder cloudProviderName(String cloudProviderName) {
            this.cloudProviderName = Objects.requireNonNull(cloudProviderName);
            return this;
        }
        @CustomType.Setter
        public Builder hostConnectionAttributesName(String hostConnectionAttributesName) {
            this.hostConnectionAttributesName = Objects.requireNonNull(hostConnectionAttributesName);
            return this;
        }
        @CustomType.Setter
        public Builder hostnames(List<String> hostnames) {
            this.hostnames = Objects.requireNonNull(hostnames);
            return this;
        }
        public Builder hostnames(String... hostnames) {
            return hostnames(List.of(hostnames));
        }
        public InfrastructureDefinitionDatacenterSsh build() {
            final var o = new InfrastructureDefinitionDatacenterSsh();
            o.cloudProviderName = cloudProviderName;
            o.hostConnectionAttributesName = hostConnectionAttributesName;
            o.hostnames = hostnames;
            return o;
        }
    }
}