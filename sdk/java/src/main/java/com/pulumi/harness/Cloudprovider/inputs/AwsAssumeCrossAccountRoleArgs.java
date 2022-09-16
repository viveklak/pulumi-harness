// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.cloudprovider.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AwsAssumeCrossAccountRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final AwsAssumeCrossAccountRoleArgs Empty = new AwsAssumeCrossAccountRoleArgs();

    /**
     * If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.
     * 
     */
    @Import(name="externalId")
    private @Nullable Output<String> externalId;

    /**
     * @return If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.
     * 
     */
    public Optional<Output<String>> externalId() {
        return Optional.ofNullable(this.externalId);
    }

    /**
     * This is an IAM role in the target deployment AWS account.
     * 
     */
    @Import(name="roleArn", required=true)
    private Output<String> roleArn;

    /**
     * @return This is an IAM role in the target deployment AWS account.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }

    private AwsAssumeCrossAccountRoleArgs() {}

    private AwsAssumeCrossAccountRoleArgs(AwsAssumeCrossAccountRoleArgs $) {
        this.externalId = $.externalId;
        this.roleArn = $.roleArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AwsAssumeCrossAccountRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AwsAssumeCrossAccountRoleArgs $;

        public Builder() {
            $ = new AwsAssumeCrossAccountRoleArgs();
        }

        public Builder(AwsAssumeCrossAccountRoleArgs defaults) {
            $ = new AwsAssumeCrossAccountRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param externalId If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param roleArn This is an IAM role in the target deployment AWS account.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn This is an IAM role in the target deployment AWS account.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        public AwsAssumeCrossAccountRoleArgs build() {
            $.roleArn = Objects.requireNonNull($.roleArn, "expected parameter 'roleArn' to be non-null");
            return $;
        }
    }

}