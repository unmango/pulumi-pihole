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
    public static class GetDomains
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Pihole = Pulumi.Pihole;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Pihole.GetDomains.Invoke();
        /// 
        ///     var denied = Pihole.GetDomains.Invoke(new()
        ///     {
        ///         Type = "deny",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDomainsResult> InvokeAsync(GetDomainsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDomainsResult>("pihole:index/getDomains:getDomains", args ?? new GetDomainsArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Pihole = Pulumi.Pihole;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var all = Pihole.GetDomains.Invoke();
        /// 
        ///     var denied = Pihole.GetDomains.Invoke(new()
        ///     {
        ///         Type = "deny",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDomainsResult> Invoke(GetDomainsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDomainsResult>("pihole:index/getDomains:getDomains", args ?? new GetDomainsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetDomainsArgs()
        {
        }
        public static new GetDomainsArgs Empty => new GetDomainsArgs();
    }

    public sealed class GetDomainsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetDomainsInvokeArgs()
        {
        }
        public static new GetDomainsInvokeArgs Empty => new GetDomainsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDomainsResult
    {
        /// <summary>
        /// Domains
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainResult> Domains;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetDomainsResult(
            ImmutableArray<Outputs.GetDomainsDomainResult> domains,

            string id,

            string? type)
        {
            Domains = domains;
            Id = id;
            Type = type;
        }
    }
}