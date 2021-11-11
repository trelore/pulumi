// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// TokenRequest requests a token for a given service account.
type TokenRequest struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec TokenRequestSpecOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the token can be authenticated.
	Status TokenRequestStatusPtrOutput `pulumi:"status"`
}

// NewTokenRequest registers a new resource with the given unique name, arguments, and options.
func NewTokenRequest(ctx *pulumi.Context,
	name string, args *TokenRequestArgs, opts ...pulumi.ResourceOption) (*TokenRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("authentication.k8s.io/v1")
	args.Kind = pulumi.StringPtr("TokenRequest")
	var resource TokenRequest
	err := ctx.RegisterResource("kubernetes:authentication.k8s.io/v1:TokenRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTokenRequest gets an existing TokenRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTokenRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenRequestState, opts ...pulumi.ResourceOption) (*TokenRequest, error) {
	var resource TokenRequest
	err := ctx.ReadResource("kubernetes:authentication.k8s.io/v1:TokenRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TokenRequest resources.
type tokenRequestState struct {
}

type TokenRequestState struct {
}

func (TokenRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenRequestState)(nil)).Elem()
}

type tokenRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated
	Spec TokenRequestSpec `pulumi:"spec"`
}

// The set of arguments for constructing a TokenRequest resource.
type TokenRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated
	Spec TokenRequestSpecInput
}

func (TokenRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenRequestArgs)(nil)).Elem()
}

type TokenRequestInput interface {
	pulumi.Input

	ToTokenRequestOutput() TokenRequestOutput
	ToTokenRequestOutputWithContext(ctx context.Context) TokenRequestOutput
}

func (*TokenRequest) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenRequest)(nil))
}

func (i *TokenRequest) ToTokenRequestOutput() TokenRequestOutput {
	return i.ToTokenRequestOutputWithContext(context.Background())
}

func (i *TokenRequest) ToTokenRequestOutputWithContext(ctx context.Context) TokenRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenRequestOutput)
}

func (i *TokenRequest) ToTokenRequestPtrOutput() TokenRequestPtrOutput {
	return i.ToTokenRequestPtrOutputWithContext(context.Background())
}

func (i *TokenRequest) ToTokenRequestPtrOutputWithContext(ctx context.Context) TokenRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenRequestPtrOutput)
}

type TokenRequestPtrInput interface {
	pulumi.Input

	ToTokenRequestPtrOutput() TokenRequestPtrOutput
	ToTokenRequestPtrOutputWithContext(ctx context.Context) TokenRequestPtrOutput
}

type tokenRequestPtrType TokenRequestArgs

func (*tokenRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenRequest)(nil))
}

func (i *tokenRequestPtrType) ToTokenRequestPtrOutput() TokenRequestPtrOutput {
	return i.ToTokenRequestPtrOutputWithContext(context.Background())
}

func (i *tokenRequestPtrType) ToTokenRequestPtrOutputWithContext(ctx context.Context) TokenRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenRequestPtrOutput)
}

// TokenRequestArrayInput is an input type that accepts TokenRequestArray and TokenRequestArrayOutput values.
// You can construct a concrete instance of `TokenRequestArrayInput` via:
//
//          TokenRequestArray{ TokenRequestArgs{...} }
type TokenRequestArrayInput interface {
	pulumi.Input

	ToTokenRequestArrayOutput() TokenRequestArrayOutput
	ToTokenRequestArrayOutputWithContext(context.Context) TokenRequestArrayOutput
}

type TokenRequestArray []TokenRequestInput

func (TokenRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TokenRequest)(nil)).Elem()
}

func (i TokenRequestArray) ToTokenRequestArrayOutput() TokenRequestArrayOutput {
	return i.ToTokenRequestArrayOutputWithContext(context.Background())
}

func (i TokenRequestArray) ToTokenRequestArrayOutputWithContext(ctx context.Context) TokenRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenRequestArrayOutput)
}

// TokenRequestMapInput is an input type that accepts TokenRequestMap and TokenRequestMapOutput values.
// You can construct a concrete instance of `TokenRequestMapInput` via:
//
//          TokenRequestMap{ "key": TokenRequestArgs{...} }
type TokenRequestMapInput interface {
	pulumi.Input

	ToTokenRequestMapOutput() TokenRequestMapOutput
	ToTokenRequestMapOutputWithContext(context.Context) TokenRequestMapOutput
}

type TokenRequestMap map[string]TokenRequestInput

func (TokenRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TokenRequest)(nil)).Elem()
}

func (i TokenRequestMap) ToTokenRequestMapOutput() TokenRequestMapOutput {
	return i.ToTokenRequestMapOutputWithContext(context.Background())
}

func (i TokenRequestMap) ToTokenRequestMapOutputWithContext(ctx context.Context) TokenRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenRequestMapOutput)
}

type TokenRequestOutput struct{ *pulumi.OutputState }

func (TokenRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenRequest)(nil))
}

func (o TokenRequestOutput) ToTokenRequestOutput() TokenRequestOutput {
	return o
}

func (o TokenRequestOutput) ToTokenRequestOutputWithContext(ctx context.Context) TokenRequestOutput {
	return o
}

func (o TokenRequestOutput) ToTokenRequestPtrOutput() TokenRequestPtrOutput {
	return o.ToTokenRequestPtrOutputWithContext(context.Background())
}

func (o TokenRequestOutput) ToTokenRequestPtrOutputWithContext(ctx context.Context) TokenRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenRequest) *TokenRequest {
		return &v
	}).(TokenRequestPtrOutput)
}

type TokenRequestPtrOutput struct{ *pulumi.OutputState }

func (TokenRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenRequest)(nil))
}

func (o TokenRequestPtrOutput) ToTokenRequestPtrOutput() TokenRequestPtrOutput {
	return o
}

func (o TokenRequestPtrOutput) ToTokenRequestPtrOutputWithContext(ctx context.Context) TokenRequestPtrOutput {
	return o
}

func (o TokenRequestPtrOutput) Elem() TokenRequestOutput {
	return o.ApplyT(func(v *TokenRequest) TokenRequest {
		if v != nil {
			return *v
		}
		var ret TokenRequest
		return ret
	}).(TokenRequestOutput)
}

type TokenRequestArrayOutput struct{ *pulumi.OutputState }

func (TokenRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenRequest)(nil))
}

func (o TokenRequestArrayOutput) ToTokenRequestArrayOutput() TokenRequestArrayOutput {
	return o
}

func (o TokenRequestArrayOutput) ToTokenRequestArrayOutputWithContext(ctx context.Context) TokenRequestArrayOutput {
	return o
}

func (o TokenRequestArrayOutput) Index(i pulumi.IntInput) TokenRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenRequest {
		return vs[0].([]TokenRequest)[vs[1].(int)]
	}).(TokenRequestOutput)
}

type TokenRequestMapOutput struct{ *pulumi.OutputState }

func (TokenRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TokenRequest)(nil))
}

func (o TokenRequestMapOutput) ToTokenRequestMapOutput() TokenRequestMapOutput {
	return o
}

func (o TokenRequestMapOutput) ToTokenRequestMapOutputWithContext(ctx context.Context) TokenRequestMapOutput {
	return o
}

func (o TokenRequestMapOutput) MapIndex(k pulumi.StringInput) TokenRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TokenRequest {
		return vs[0].(map[string]TokenRequest)[vs[1].(string)]
	}).(TokenRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TokenRequestInput)(nil)).Elem(), &TokenRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenRequestPtrInput)(nil)).Elem(), &TokenRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenRequestArrayInput)(nil)).Elem(), TokenRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenRequestMapInput)(nil)).Elem(), TokenRequestMap{})
	pulumi.RegisterOutputType(TokenRequestOutput{})
	pulumi.RegisterOutputType(TokenRequestPtrOutput{})
	pulumi.RegisterOutputType(TokenRequestArrayOutput{})
	pulumi.RegisterOutputType(TokenRequestMapOutput{})
}