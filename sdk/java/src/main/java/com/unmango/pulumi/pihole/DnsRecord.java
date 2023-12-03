// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.pulumi.pihole;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.pulumi.pihole.DnsRecordArgs;
import com.unmango.pulumi.pihole.Utilities;
import com.unmango.pulumi.pihole.inputs.DnsRecordState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a Pi-hole DNS record
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.pihole.DnsRecord;
 * import com.pulumi.pihole.DnsRecordArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var record = new DnsRecord(&#34;record&#34;, DnsRecordArgs.builder()        
 *             .domain(&#34;foo.com&#34;)
 *             .ip(&#34;127.0.0.1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import pihole:index/dnsRecord:DnsRecord record foo.com
 * ```
 * 
 */
@ResourceType(type="pihole:index/dnsRecord:DnsRecord")
public class DnsRecord extends com.pulumi.resources.CustomResource {
    /**
     * DNS record domain
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return DNS record domain
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * IP address to route traffic to from the DNS record domain
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return IP address to route traffic to from the DNS record domain
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DnsRecord(String name) {
        this(name, DnsRecordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DnsRecord(String name, DnsRecordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DnsRecord(String name, DnsRecordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("pihole:index/dnsRecord:DnsRecord", name, args == null ? DnsRecordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DnsRecord(String name, Output<String> id, @Nullable DnsRecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("pihole:index/dnsRecord:DnsRecord", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DnsRecord get(String name, Output<String> id, @Nullable DnsRecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DnsRecord(name, id, state, options);
    }
}
