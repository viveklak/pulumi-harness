// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.harness.platform.outputs.ResourceGroupResourceFilterResourceAttributeFilter;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ResourceGroupResourceFilterResource {
    private @Nullable List<ResourceGroupResourceFilterResourceAttributeFilter> attributeFilters;
    private @Nullable List<String> identifiers;
    private String resourceType;

    private ResourceGroupResourceFilterResource() {}
    public List<ResourceGroupResourceFilterResourceAttributeFilter> attributeFilters() {
        return this.attributeFilters == null ? List.of() : this.attributeFilters;
    }
    public List<String> identifiers() {
        return this.identifiers == null ? List.of() : this.identifiers;
    }
    public String resourceType() {
        return this.resourceType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ResourceGroupResourceFilterResource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<ResourceGroupResourceFilterResourceAttributeFilter> attributeFilters;
        private @Nullable List<String> identifiers;
        private String resourceType;
        public Builder() {}
        public Builder(ResourceGroupResourceFilterResource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attributeFilters = defaults.attributeFilters;
    	      this.identifiers = defaults.identifiers;
    	      this.resourceType = defaults.resourceType;
        }

        @CustomType.Setter
        public Builder attributeFilters(@Nullable List<ResourceGroupResourceFilterResourceAttributeFilter> attributeFilters) {
            this.attributeFilters = attributeFilters;
            return this;
        }
        public Builder attributeFilters(ResourceGroupResourceFilterResourceAttributeFilter... attributeFilters) {
            return attributeFilters(List.of(attributeFilters));
        }
        @CustomType.Setter
        public Builder identifiers(@Nullable List<String> identifiers) {
            this.identifiers = identifiers;
            return this;
        }
        public Builder identifiers(String... identifiers) {
            return identifiers(List.of(identifiers));
        }
        @CustomType.Setter
        public Builder resourceType(String resourceType) {
            this.resourceType = Objects.requireNonNull(resourceType);
            return this;
        }
        public ResourceGroupResourceFilterResource build() {
            final var o = new ResourceGroupResourceFilterResource();
            o.attributeFilters = attributeFilters;
            o.identifiers = identifiers;
            o.resourceType = resourceType;
            return o;
        }
    }
}