// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Pulumi.Pihole
{
    /// <summary>
    /// Manages a Pi-hole DNS record
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Pihole = UnMango.Pulumi.Pihole;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var record = new Pihole.DnsRecord("record", new()
    ///     {
    ///         Domain = "foo.com",
    ///         Ip = "127.0.0.1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import pihole:index/dnsRecord:DnsRecord record foo.com
    /// ```
    /// </summary>
    [PiholeResourceType("pihole:index/dnsRecord:DnsRecord")]
    public partial class DnsRecord : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DNS record domain
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// IP address to route traffic to from the DNS record domain
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;


        /// <summary>
        /// Create a DnsRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsRecord(string name, DnsRecordArgs args, CustomResourceOptions? options = null)
            : base("pihole:index/dnsRecord:DnsRecord", name, args ?? new DnsRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsRecord(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
            : base("pihole:index/dnsRecord:DnsRecord", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/unmango/pulumi-pihole",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DnsRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsRecord Get(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsRecord(name, id, state, options);
        }
    }

    public sealed class DnsRecordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS record domain
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// IP address to route traffic to from the DNS record domain
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        public DnsRecordArgs()
        {
        }
        public static new DnsRecordArgs Empty => new DnsRecordArgs();
    }

    public sealed class DnsRecordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS record domain
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// IP address to route traffic to from the DNS record domain
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public DnsRecordState()
        {
        }
        public static new DnsRecordState Empty => new DnsRecordState();
    }
}
