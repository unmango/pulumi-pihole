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
    public static class GetCnameRecords
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Pihole = Pulumi.Pihole;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var records = Pihole.GetCnameRecords.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCnameRecordsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCnameRecordsResult>("pihole:index/getCnameRecords:getCnameRecords", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Pihole = Pulumi.Pihole;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var records = Pihole.GetCnameRecords.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCnameRecordsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCnameRecordsResult>("pihole:index/getCnameRecords:getCnameRecords", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Pihole = Pulumi.Pihole;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var records = Pihole.GetCnameRecords.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCnameRecordsResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCnameRecordsResult>("pihole:index/getCnameRecords:getCnameRecords", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetCnameRecordsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of CNAME Pi-hole records
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCnameRecordsRecordResult> Records;

        [OutputConstructor]
        private GetCnameRecordsResult(
            string id,

            ImmutableArray<Outputs.GetCnameRecordsRecordResult> records)
        {
            Id = id;
            Records = records;
        }
    }
}
