// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResourceGroupResourceFilterResourceAttributeFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResourceGroupResourceFilterResourceAttributeFilterArgs Empty = new ResourceGroupResourceFilterResourceAttributeFilterArgs();

    @Import(name="attributeName")
    private @Nullable Output<String> attributeName;

    public Optional<Output<String>> attributeName() {
        return Optional.ofNullable(this.attributeName);
    }

    @Import(name="attributeValues")
    private @Nullable Output<List<String>> attributeValues;

    public Optional<Output<List<String>>> attributeValues() {
        return Optional.ofNullable(this.attributeValues);
    }

    private ResourceGroupResourceFilterResourceAttributeFilterArgs() {}

    private ResourceGroupResourceFilterResourceAttributeFilterArgs(ResourceGroupResourceFilterResourceAttributeFilterArgs $) {
        this.attributeName = $.attributeName;
        this.attributeValues = $.attributeValues;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceGroupResourceFilterResourceAttributeFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceGroupResourceFilterResourceAttributeFilterArgs $;

        public Builder() {
            $ = new ResourceGroupResourceFilterResourceAttributeFilterArgs();
        }

        public Builder(ResourceGroupResourceFilterResourceAttributeFilterArgs defaults) {
            $ = new ResourceGroupResourceFilterResourceAttributeFilterArgs(Objects.requireNonNull(defaults));
        }

        public Builder attributeName(@Nullable Output<String> attributeName) {
            $.attributeName = attributeName;
            return this;
        }

        public Builder attributeName(String attributeName) {
            return attributeName(Output.of(attributeName));
        }

        public Builder attributeValues(@Nullable Output<List<String>> attributeValues) {
            $.attributeValues = attributeValues;
            return this;
        }

        public Builder attributeValues(List<String> attributeValues) {
            return attributeValues(Output.of(attributeValues));
        }

        public Builder attributeValues(String... attributeValues) {
            return attributeValues(List.of(attributeValues));
        }

        public ResourceGroupResourceFilterResourceAttributeFilterArgs build() {
            return $;
        }
    }

}