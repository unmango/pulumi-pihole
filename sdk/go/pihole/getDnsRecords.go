// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pihole

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unmango/pulumi-pihole/sdk/go/pihole/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/unmango/pulumi-pihole/sdk/go/pihole"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pihole.GetDnsRecords(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDnsRecords(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetDnsRecordsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDnsRecordsResult
	err := ctx.Invoke("pihole:index/getDnsRecords:getDnsRecords", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getDnsRecords.
type GetDnsRecordsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of Pi-hole DNS records
	Records []GetDnsRecordsRecord `pulumi:"records"`
}

func GetDnsRecordsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetDnsRecordsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetDnsRecordsResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("pihole:index/getDnsRecords:getDnsRecords", nil, GetDnsRecordsResultOutput{}, options).(GetDnsRecordsResultOutput), nil
	}).(GetDnsRecordsResultOutput)
}

// A collection of values returned by getDnsRecords.
type GetDnsRecordsResultOutput struct{ *pulumi.OutputState }

func (GetDnsRecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsRecordsResult)(nil)).Elem()
}

func (o GetDnsRecordsResultOutput) ToGetDnsRecordsResultOutput() GetDnsRecordsResultOutput {
	return o
}

func (o GetDnsRecordsResultOutput) ToGetDnsRecordsResultOutputWithContext(ctx context.Context) GetDnsRecordsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDnsRecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsRecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of Pi-hole DNS records
func (o GetDnsRecordsResultOutput) Records() GetDnsRecordsRecordArrayOutput {
	return o.ApplyT(func(v GetDnsRecordsResult) []GetDnsRecordsRecord { return v.Records }).(GetDnsRecordsRecordArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsRecordsResultOutput{})
}
