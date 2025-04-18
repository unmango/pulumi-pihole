// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.pulumi.pihole.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDnsRecordsRecord {
    /**
     * @return DNS record domain
     * 
     */
    private String domain;
    /**
     * @return IP address where traffic is routed to from the DNS record domain
     * 
     */
    private String ip;

    private GetDnsRecordsRecord() {}
    /**
     * @return DNS record domain
     * 
     */
    public String domain() {
        return this.domain;
    }
    /**
     * @return IP address where traffic is routed to from the DNS record domain
     * 
     */
    public String ip() {
        return this.ip;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDnsRecordsRecord defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String domain;
        private String ip;
        public Builder() {}
        public Builder(GetDnsRecordsRecord defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domain = defaults.domain;
    	      this.ip = defaults.ip;
        }

        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            this.ip = Objects.requireNonNull(ip);
            return this;
        }
        public GetDnsRecordsRecord build() {
            final var _resultValue = new GetDnsRecordsRecord();
            _resultValue.domain = domain;
            _resultValue.ip = ip;
            return _resultValue;
        }
    }
}
