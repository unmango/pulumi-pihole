// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.pulumi.pihole.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DnsRecordState extends com.pulumi.resources.ResourceArgs {

    public static final DnsRecordState Empty = new DnsRecordState();

    /**
     * DNS record domain
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return DNS record domain
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * IP address to route traffic to from the DNS record domain
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return IP address to route traffic to from the DNS record domain
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    private DnsRecordState() {}

    private DnsRecordState(DnsRecordState $) {
        this.domain = $.domain;
        this.ip = $.ip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DnsRecordState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DnsRecordState $;

        public Builder() {
            $ = new DnsRecordState();
        }

        public Builder(DnsRecordState defaults) {
            $ = new DnsRecordState(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain DNS record domain
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain DNS record domain
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param ip IP address to route traffic to from the DNS record domain
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip IP address to route traffic to from the DNS record domain
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        public DnsRecordState build() {
            return $;
        }
    }

}
