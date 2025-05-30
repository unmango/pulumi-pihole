// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.pulumi.pihole;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.unmango.pulumi.pihole.CnameRecordArgs;
import com.unmango.pulumi.pihole.Utilities;
import com.unmango.pulumi.pihole.inputs.CnameRecordState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a Pi-hole CNAME record
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.pihole.CnameRecord;
 * import com.pulumi.pihole.CnameRecordArgs;
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
 *         var record = new CnameRecord(&#34;record&#34;, CnameRecordArgs.builder()
 *             .domain(&#34;foo.com&#34;)
 *             .target(&#34;bar.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import pihole:index/cnameRecord:CnameRecord record foo.com
 * ```
 * 
 */
@ResourceType(type="pihole:index/cnameRecord:CnameRecord")
public class CnameRecord extends com.pulumi.resources.CustomResource {
    /**
     * Domain to create a CNAME record for
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return Domain to create a CNAME record for
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * Value of the CNAME record where traffic will be directed to from the configured domain value
     * 
     */
    @Export(name="target", refs={String.class}, tree="[0]")
    private Output<String> target;

    /**
     * @return Value of the CNAME record where traffic will be directed to from the configured domain value
     * 
     */
    public Output<String> target() {
        return this.target;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CnameRecord(String name) {
        this(name, CnameRecordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CnameRecord(String name, CnameRecordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CnameRecord(String name, CnameRecordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("pihole:index/cnameRecord:CnameRecord", name, args == null ? CnameRecordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CnameRecord(String name, Output<String> id, @Nullable CnameRecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("pihole:index/cnameRecord:CnameRecord", name, state, makeResourceOptions(options, id));
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
    public static CnameRecord get(String name, Output<String> id, @Nullable CnameRecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CnameRecord(name, id, state, options);
    }
}
