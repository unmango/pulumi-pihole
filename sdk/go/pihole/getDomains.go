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
//			_, err := pihole.GetDomains(ctx, &pihole.GetDomainsArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = pihole.GetDomains(ctx, &pihole.GetDomainsArgs{
//				Type: pulumi.StringRef("deny"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDomains(ctx *pulumi.Context, args *GetDomainsArgs, opts ...pulumi.InvokeOption) (*GetDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDomainsResult
	err := ctx.Invoke("pihole:index/getDomains:getDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomains.
type GetDomainsArgs struct {
	// Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getDomains.
type GetDomainsResult struct {
	// Domains
	Domains []GetDomainsDomain `pulumi:"domains"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
	Type *string `pulumi:"type"`
}

func GetDomainsOutput(ctx *pulumi.Context, args GetDomainsOutputArgs, opts ...pulumi.InvokeOption) GetDomainsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDomainsResultOutput, error) {
			args := v.(GetDomainsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("pihole:index/getDomains:getDomains", args, GetDomainsResultOutput{}, options).(GetDomainsResultOutput), nil
		}).(GetDomainsResultOutput)
}

// A collection of arguments for invoking getDomains.
type GetDomainsOutputArgs struct {
	// Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getDomains.
type GetDomainsResultOutput struct{ *pulumi.OutputState }

func (GetDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsResult)(nil)).Elem()
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutput() GetDomainsResultOutput {
	return o
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutputWithContext(ctx context.Context) GetDomainsResultOutput {
	return o
}

// Domains
func (o GetDomainsResultOutput) Domains() GetDomainsDomainArrayOutput {
	return o.ApplyT(func(v GetDomainsResult) []GetDomainsDomain { return v.Domains }).(GetDomainsDomainArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
func (o GetDomainsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainsResultOutput{})
}
