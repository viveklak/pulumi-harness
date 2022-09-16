// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetUsergroupNotificationConfig extends com.pulumi.resources.InvokeArgs {

    public static final GetUsergroupNotificationConfig Empty = new GetUsergroupNotificationConfig();

    /**
     * Group email
     * 
     */
    @Import(name="groupEmail", required=true)
    private String groupEmail;

    /**
     * @return Group email
     * 
     */
    public String groupEmail() {
        return this.groupEmail;
    }

    /**
     * Url of Microsoft teams webhook
     * 
     */
    @Import(name="microsoftTeamsWebhookUrl", required=true)
    private String microsoftTeamsWebhookUrl;

    /**
     * @return Url of Microsoft teams webhook
     * 
     */
    public String microsoftTeamsWebhookUrl() {
        return this.microsoftTeamsWebhookUrl;
    }

    /**
     * Pager duty key
     * 
     */
    @Import(name="pagerDutyKey", required=true)
    private String pagerDutyKey;

    /**
     * @return Pager duty key
     * 
     */
    public String pagerDutyKey() {
        return this.pagerDutyKey;
    }

    /**
     * Url of slack webhook
     * 
     */
    @Import(name="slackWebhookUrl", required=true)
    private String slackWebhookUrl;

    /**
     * @return Url of slack webhook
     * 
     */
    public String slackWebhookUrl() {
        return this.slackWebhookUrl;
    }

    /**
     * Can be one of EMAIL, SLACK, PAGERDUTY, MSTEAMS
     * 
     */
    @Import(name="type", required=true)
    private String type;

    /**
     * @return Can be one of EMAIL, SLACK, PAGERDUTY, MSTEAMS
     * 
     */
    public String type() {
        return this.type;
    }

    private GetUsergroupNotificationConfig() {}

    private GetUsergroupNotificationConfig(GetUsergroupNotificationConfig $) {
        this.groupEmail = $.groupEmail;
        this.microsoftTeamsWebhookUrl = $.microsoftTeamsWebhookUrl;
        this.pagerDutyKey = $.pagerDutyKey;
        this.slackWebhookUrl = $.slackWebhookUrl;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUsergroupNotificationConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUsergroupNotificationConfig $;

        public Builder() {
            $ = new GetUsergroupNotificationConfig();
        }

        public Builder(GetUsergroupNotificationConfig defaults) {
            $ = new GetUsergroupNotificationConfig(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupEmail Group email
         * 
         * @return builder
         * 
         */
        public Builder groupEmail(String groupEmail) {
            $.groupEmail = groupEmail;
            return this;
        }

        /**
         * @param microsoftTeamsWebhookUrl Url of Microsoft teams webhook
         * 
         * @return builder
         * 
         */
        public Builder microsoftTeamsWebhookUrl(String microsoftTeamsWebhookUrl) {
            $.microsoftTeamsWebhookUrl = microsoftTeamsWebhookUrl;
            return this;
        }

        /**
         * @param pagerDutyKey Pager duty key
         * 
         * @return builder
         * 
         */
        public Builder pagerDutyKey(String pagerDutyKey) {
            $.pagerDutyKey = pagerDutyKey;
            return this;
        }

        /**
         * @param slackWebhookUrl Url of slack webhook
         * 
         * @return builder
         * 
         */
        public Builder slackWebhookUrl(String slackWebhookUrl) {
            $.slackWebhookUrl = slackWebhookUrl;
            return this;
        }

        /**
         * @param type Can be one of EMAIL, SLACK, PAGERDUTY, MSTEAMS
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            $.type = type;
            return this;
        }

        public GetUsergroupNotificationConfig build() {
            $.groupEmail = Objects.requireNonNull($.groupEmail, "expected parameter 'groupEmail' to be non-null");
            $.microsoftTeamsWebhookUrl = Objects.requireNonNull($.microsoftTeamsWebhookUrl, "expected parameter 'microsoftTeamsWebhookUrl' to be non-null");
            $.pagerDutyKey = Objects.requireNonNull($.pagerDutyKey, "expected parameter 'pagerDutyKey' to be non-null");
            $.slackWebhookUrl = Objects.requireNonNull($.slackWebhookUrl, "expected parameter 'slackWebhookUrl' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}