// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Pulumi.Pihole.Outputs
{

    [OutputType]
    public sealed class GetDnsRecordsRecordResult
    {
        /// <summary>
        /// DNS record domain
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// IP address where traffic is routed to from the DNS record domain
        /// </summary>
        public readonly string Ip;

        [OutputConstructor]
        private GetDnsRecordsRecordResult(
            string domain,

            string ip)
        {
            Domain = domain;
            Ip = ip;
        }
    }
}
