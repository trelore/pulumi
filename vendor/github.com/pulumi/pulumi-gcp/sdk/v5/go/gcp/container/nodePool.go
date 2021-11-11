// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a node pool in a Google Kubernetes Engine (GKE) cluster separately from
// the cluster control plane. For more information see [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
// and [the API reference](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters.nodePools).
//
// ## Example Usage
// ### Using A Separately Managed Node Pool (Recommended)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := serviceAccount.NewAccount(ctx, "_default", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("service-account-id"),
// 			DisplayName: pulumi.String("Service Account"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primary, err := container.NewCluster(ctx, "primary", &container.ClusterArgs{
// 			Location:              pulumi.String("us-central1"),
// 			RemoveDefaultNodePool: pulumi.Bool(true),
// 			InitialNodeCount:      pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = container.NewNodePool(ctx, "primaryPreemptibleNodes", &container.NodePoolArgs{
// 			Cluster:   primary.ID(),
// 			NodeCount: pulumi.Int(1),
// 			NodeConfig: &container.NodePoolNodeConfigArgs{
// 				Preemptible:    pulumi.Bool(true),
// 				MachineType:    pulumi.String("e2-medium"),
// 				ServiceAccount: _default.Email,
// 				OauthScopes: pulumi.StringArray{
// 					pulumi.String("https://www.googleapis.com/auth/cloud-platform"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Node pools can be imported using the `project`, `location`, `cluster` and `name`. If the project is omitted, the project value in the provider configuration will be used. Examples
//
// ```sh
//  $ pulumi import gcp:container/nodePool:NodePool mainpool my-gcp-project/us-east1-a/my-cluster/main-pool
// ```
//
// ```sh
//  $ pulumi import gcp:container/nodePool:NodePool mainpool us-east1/my-cluster/main-pool
// ```
type NodePool struct {
	pulumi.CustomResourceState

	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	Autoscaling NodePoolAutoscalingPtrOutput `pulumi:"autoscaling"`
	// The cluster to create the node pool for. Cluster must be present in `location` provided for clusters. May be specified in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}` or as just the name of the cluster.
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// The initial number of nodes for the pool. In
	// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
	// this will force recreation of the resource. WARNING: Resizing your node pool manually
	// may change this value in your existing cluster, which will trigger destruction
	// and recreation on the next provider run (to rectify the discrepancy).  If you don't
	// need this value, don't set it.  If you do need it, you can use a lifecycle block to
	// ignore subsqeuent changes to this field.
	InitialNodeCount pulumi.IntOutput `pulumi:"initialNodeCount"`
	// The resource URLs of the managed instance groups associated with this node pool.
	InstanceGroupUrls pulumi.StringArrayOutput `pulumi:"instanceGroupUrls"`
	// The location (region or zone) of the cluster.
	Location pulumi.StringOutput `pulumi:"location"`
	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	Management NodePoolManagementOutput `pulumi:"management"`
	// The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	// for more information.
	MaxPodsPerNode pulumi.IntOutput `pulumi:"maxPodsPerNode"`
	// The name of the node pool. If left blank, the provider will
	// auto-generate a unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name for the node pool beginning
	// with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The network configuration of the pool. See
	// container.Cluster for schema.
	NetworkConfig NodePoolNetworkConfigOutput `pulumi:"networkConfig"`
	// Parameters used in creating the node pool. See
	// container.Cluster for schema.
	NodeConfig NodePoolNodeConfigOutput `pulumi:"nodeConfig"`
	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside `autoscaling`.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// The list of zones in which the node pool's nodes should be located. Nodes must
	// be in the region of their regional cluster or in the same region as their
	// cluster's zone for zonal clusters. If unspecified, the cluster-level
	// `nodeLocations` will be used.
	NodeLocations pulumi.StringArrayOutput `pulumi:"nodeLocations"`
	Operation     pulumi.StringOutput      `pulumi:"operation"`
	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Specify node upgrade settings to change how many nodes GKE attempts to
	// upgrade at once. The number of nodes upgraded simultaneously is the sum of `maxSurge` and `maxUnavailable`.
	// The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
	UpgradeSettings NodePoolUpgradeSettingsOutput `pulumi:"upgradeSettings"`
	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and `autoUpgrade` are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as the provider will see spurious diffs
	// when fuzzy versions are used. See the `container.getEngineVersions` data source's
	// `versionPrefix` field to approximate fuzzy versions in a provider-compatible way.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewNodePool registers a new resource with the given unique name, arguments, and options.
func NewNodePool(ctx *pulumi.Context,
	name string, args *NodePoolArgs, opts ...pulumi.ResourceOption) (*NodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	var resource NodePool
	err := ctx.RegisterResource("gcp:container/nodePool:NodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodePool gets an existing NodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodePoolState, opts ...pulumi.ResourceOption) (*NodePool, error) {
	var resource NodePool
	err := ctx.ReadResource("gcp:container/nodePool:NodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodePool resources.
type nodePoolState struct {
	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	Autoscaling *NodePoolAutoscaling `pulumi:"autoscaling"`
	// The cluster to create the node pool for. Cluster must be present in `location` provided for clusters. May be specified in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}` or as just the name of the cluster.
	Cluster *string `pulumi:"cluster"`
	// The initial number of nodes for the pool. In
	// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
	// this will force recreation of the resource. WARNING: Resizing your node pool manually
	// may change this value in your existing cluster, which will trigger destruction
	// and recreation on the next provider run (to rectify the discrepancy).  If you don't
	// need this value, don't set it.  If you do need it, you can use a lifecycle block to
	// ignore subsqeuent changes to this field.
	InitialNodeCount *int `pulumi:"initialNodeCount"`
	// The resource URLs of the managed instance groups associated with this node pool.
	InstanceGroupUrls []string `pulumi:"instanceGroupUrls"`
	// The location (region or zone) of the cluster.
	Location *string `pulumi:"location"`
	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	Management *NodePoolManagement `pulumi:"management"`
	// The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	// for more information.
	MaxPodsPerNode *int `pulumi:"maxPodsPerNode"`
	// The name of the node pool. If left blank, the provider will
	// auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name for the node pool beginning
	// with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The network configuration of the pool. See
	// container.Cluster for schema.
	NetworkConfig *NodePoolNetworkConfig `pulumi:"networkConfig"`
	// Parameters used in creating the node pool. See
	// container.Cluster for schema.
	NodeConfig *NodePoolNodeConfig `pulumi:"nodeConfig"`
	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside `autoscaling`.
	NodeCount *int `pulumi:"nodeCount"`
	// The list of zones in which the node pool's nodes should be located. Nodes must
	// be in the region of their regional cluster or in the same region as their
	// cluster's zone for zonal clusters. If unspecified, the cluster-level
	// `nodeLocations` will be used.
	NodeLocations []string `pulumi:"nodeLocations"`
	Operation     *string  `pulumi:"operation"`
	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	Project *string `pulumi:"project"`
	// Specify node upgrade settings to change how many nodes GKE attempts to
	// upgrade at once. The number of nodes upgraded simultaneously is the sum of `maxSurge` and `maxUnavailable`.
	// The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
	UpgradeSettings *NodePoolUpgradeSettings `pulumi:"upgradeSettings"`
	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and `autoUpgrade` are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as the provider will see spurious diffs
	// when fuzzy versions are used. See the `container.getEngineVersions` data source's
	// `versionPrefix` field to approximate fuzzy versions in a provider-compatible way.
	Version *string `pulumi:"version"`
}

type NodePoolState struct {
	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	Autoscaling NodePoolAutoscalingPtrInput
	// The cluster to create the node pool for. Cluster must be present in `location` provided for clusters. May be specified in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}` or as just the name of the cluster.
	Cluster pulumi.StringPtrInput
	// The initial number of nodes for the pool. In
	// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
	// this will force recreation of the resource. WARNING: Resizing your node pool manually
	// may change this value in your existing cluster, which will trigger destruction
	// and recreation on the next provider run (to rectify the discrepancy).  If you don't
	// need this value, don't set it.  If you do need it, you can use a lifecycle block to
	// ignore subsqeuent changes to this field.
	InitialNodeCount pulumi.IntPtrInput
	// The resource URLs of the managed instance groups associated with this node pool.
	InstanceGroupUrls pulumi.StringArrayInput
	// The location (region or zone) of the cluster.
	Location pulumi.StringPtrInput
	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	Management NodePoolManagementPtrInput
	// The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	// for more information.
	MaxPodsPerNode pulumi.IntPtrInput
	// The name of the node pool. If left blank, the provider will
	// auto-generate a unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name for the node pool beginning
	// with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The network configuration of the pool. See
	// container.Cluster for schema.
	NetworkConfig NodePoolNetworkConfigPtrInput
	// Parameters used in creating the node pool. See
	// container.Cluster for schema.
	NodeConfig NodePoolNodeConfigPtrInput
	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside `autoscaling`.
	NodeCount pulumi.IntPtrInput
	// The list of zones in which the node pool's nodes should be located. Nodes must
	// be in the region of their regional cluster or in the same region as their
	// cluster's zone for zonal clusters. If unspecified, the cluster-level
	// `nodeLocations` will be used.
	NodeLocations pulumi.StringArrayInput
	Operation     pulumi.StringPtrInput
	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	Project pulumi.StringPtrInput
	// Specify node upgrade settings to change how many nodes GKE attempts to
	// upgrade at once. The number of nodes upgraded simultaneously is the sum of `maxSurge` and `maxUnavailable`.
	// The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
	UpgradeSettings NodePoolUpgradeSettingsPtrInput
	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and `autoUpgrade` are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as the provider will see spurious diffs
	// when fuzzy versions are used. See the `container.getEngineVersions` data source's
	// `versionPrefix` field to approximate fuzzy versions in a provider-compatible way.
	Version pulumi.StringPtrInput
}

func (NodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePoolState)(nil)).Elem()
}

type nodePoolArgs struct {
	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	Autoscaling *NodePoolAutoscaling `pulumi:"autoscaling"`
	// The cluster to create the node pool for. Cluster must be present in `location` provided for clusters. May be specified in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}` or as just the name of the cluster.
	Cluster string `pulumi:"cluster"`
	// The initial number of nodes for the pool. In
	// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
	// this will force recreation of the resource. WARNING: Resizing your node pool manually
	// may change this value in your existing cluster, which will trigger destruction
	// and recreation on the next provider run (to rectify the discrepancy).  If you don't
	// need this value, don't set it.  If you do need it, you can use a lifecycle block to
	// ignore subsqeuent changes to this field.
	InitialNodeCount *int `pulumi:"initialNodeCount"`
	// The location (region or zone) of the cluster.
	Location *string `pulumi:"location"`
	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	Management *NodePoolManagement `pulumi:"management"`
	// The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	// for more information.
	MaxPodsPerNode *int `pulumi:"maxPodsPerNode"`
	// The name of the node pool. If left blank, the provider will
	// auto-generate a unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name for the node pool beginning
	// with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The network configuration of the pool. See
	// container.Cluster for schema.
	NetworkConfig *NodePoolNetworkConfig `pulumi:"networkConfig"`
	// Parameters used in creating the node pool. See
	// container.Cluster for schema.
	NodeConfig *NodePoolNodeConfig `pulumi:"nodeConfig"`
	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside `autoscaling`.
	NodeCount *int `pulumi:"nodeCount"`
	// The list of zones in which the node pool's nodes should be located. Nodes must
	// be in the region of their regional cluster or in the same region as their
	// cluster's zone for zonal clusters. If unspecified, the cluster-level
	// `nodeLocations` will be used.
	NodeLocations []string `pulumi:"nodeLocations"`
	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	Project *string `pulumi:"project"`
	// Specify node upgrade settings to change how many nodes GKE attempts to
	// upgrade at once. The number of nodes upgraded simultaneously is the sum of `maxSurge` and `maxUnavailable`.
	// The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
	UpgradeSettings *NodePoolUpgradeSettings `pulumi:"upgradeSettings"`
	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and `autoUpgrade` are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as the provider will see spurious diffs
	// when fuzzy versions are used. See the `container.getEngineVersions` data source's
	// `versionPrefix` field to approximate fuzzy versions in a provider-compatible way.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a NodePool resource.
type NodePoolArgs struct {
	// Configuration required by cluster autoscaler to adjust
	// the size of the node pool to the current cluster usage. Structure is documented below.
	Autoscaling NodePoolAutoscalingPtrInput
	// The cluster to create the node pool for. Cluster must be present in `location` provided for clusters. May be specified in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}` or as just the name of the cluster.
	Cluster pulumi.StringInput
	// The initial number of nodes for the pool. In
	// regional or multi-zonal clusters, this is the number of nodes per zone. Changing
	// this will force recreation of the resource. WARNING: Resizing your node pool manually
	// may change this value in your existing cluster, which will trigger destruction
	// and recreation on the next provider run (to rectify the discrepancy).  If you don't
	// need this value, don't set it.  If you do need it, you can use a lifecycle block to
	// ignore subsqeuent changes to this field.
	InitialNodeCount pulumi.IntPtrInput
	// The location (region or zone) of the cluster.
	Location pulumi.StringPtrInput
	// Node management configuration, wherein auto-repair and
	// auto-upgrade is configured. Structure is documented below.
	Management NodePoolManagementPtrInput
	// The maximum number of pods per node in this node pool.
	// Note that this does not work on node pools which are "route-based" - that is, node
	// pools belonging to clusters that do not have IP Aliasing enabled.
	// See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
	// for more information.
	MaxPodsPerNode pulumi.IntPtrInput
	// The name of the node pool. If left blank, the provider will
	// auto-generate a unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name for the node pool beginning
	// with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The network configuration of the pool. See
	// container.Cluster for schema.
	NetworkConfig NodePoolNetworkConfigPtrInput
	// Parameters used in creating the node pool. See
	// container.Cluster for schema.
	NodeConfig NodePoolNodeConfigPtrInput
	// The number of nodes per instance group. This field can be used to
	// update the number of nodes per instance group but should not be used alongside `autoscaling`.
	NodeCount pulumi.IntPtrInput
	// The list of zones in which the node pool's nodes should be located. Nodes must
	// be in the region of their regional cluster or in the same region as their
	// cluster's zone for zonal clusters. If unspecified, the cluster-level
	// `nodeLocations` will be used.
	NodeLocations pulumi.StringArrayInput
	// The ID of the project in which to create the node pool. If blank,
	// the provider-configured project will be used.
	Project pulumi.StringPtrInput
	// Specify node upgrade settings to change how many nodes GKE attempts to
	// upgrade at once. The number of nodes upgraded simultaneously is the sum of `maxSurge` and `maxUnavailable`.
	// The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
	UpgradeSettings NodePoolUpgradeSettingsPtrInput
	// The Kubernetes version for the nodes in this pool. Note that if this field
	// and `autoUpgrade` are both specified, they will fight each other for what the node version should
	// be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
	// recommended that you specify explicit versions as the provider will see spurious diffs
	// when fuzzy versions are used. See the `container.getEngineVersions` data source's
	// `versionPrefix` field to approximate fuzzy versions in a provider-compatible way.
	Version pulumi.StringPtrInput
}

func (NodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodePoolArgs)(nil)).Elem()
}

type NodePoolInput interface {
	pulumi.Input

	ToNodePoolOutput() NodePoolOutput
	ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput
}

func (*NodePool) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePool)(nil))
}

func (i *NodePool) ToNodePoolOutput() NodePoolOutput {
	return i.ToNodePoolOutputWithContext(context.Background())
}

func (i *NodePool) ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolOutput)
}

func (i *NodePool) ToNodePoolPtrOutput() NodePoolPtrOutput {
	return i.ToNodePoolPtrOutputWithContext(context.Background())
}

func (i *NodePool) ToNodePoolPtrOutputWithContext(ctx context.Context) NodePoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolPtrOutput)
}

type NodePoolPtrInput interface {
	pulumi.Input

	ToNodePoolPtrOutput() NodePoolPtrOutput
	ToNodePoolPtrOutputWithContext(ctx context.Context) NodePoolPtrOutput
}

type nodePoolPtrType NodePoolArgs

func (*nodePoolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodePool)(nil))
}

func (i *nodePoolPtrType) ToNodePoolPtrOutput() NodePoolPtrOutput {
	return i.ToNodePoolPtrOutputWithContext(context.Background())
}

func (i *nodePoolPtrType) ToNodePoolPtrOutputWithContext(ctx context.Context) NodePoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolPtrOutput)
}

// NodePoolArrayInput is an input type that accepts NodePoolArray and NodePoolArrayOutput values.
// You can construct a concrete instance of `NodePoolArrayInput` via:
//
//          NodePoolArray{ NodePoolArgs{...} }
type NodePoolArrayInput interface {
	pulumi.Input

	ToNodePoolArrayOutput() NodePoolArrayOutput
	ToNodePoolArrayOutputWithContext(context.Context) NodePoolArrayOutput
}

type NodePoolArray []NodePoolInput

func (NodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodePool)(nil)).Elem()
}

func (i NodePoolArray) ToNodePoolArrayOutput() NodePoolArrayOutput {
	return i.ToNodePoolArrayOutputWithContext(context.Background())
}

func (i NodePoolArray) ToNodePoolArrayOutputWithContext(ctx context.Context) NodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolArrayOutput)
}

// NodePoolMapInput is an input type that accepts NodePoolMap and NodePoolMapOutput values.
// You can construct a concrete instance of `NodePoolMapInput` via:
//
//          NodePoolMap{ "key": NodePoolArgs{...} }
type NodePoolMapInput interface {
	pulumi.Input

	ToNodePoolMapOutput() NodePoolMapOutput
	ToNodePoolMapOutputWithContext(context.Context) NodePoolMapOutput
}

type NodePoolMap map[string]NodePoolInput

func (NodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodePool)(nil)).Elem()
}

func (i NodePoolMap) ToNodePoolMapOutput() NodePoolMapOutput {
	return i.ToNodePoolMapOutputWithContext(context.Background())
}

func (i NodePoolMap) ToNodePoolMapOutputWithContext(ctx context.Context) NodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodePoolMapOutput)
}

type NodePoolOutput struct{ *pulumi.OutputState }

func (NodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodePool)(nil))
}

func (o NodePoolOutput) ToNodePoolOutput() NodePoolOutput {
	return o
}

func (o NodePoolOutput) ToNodePoolOutputWithContext(ctx context.Context) NodePoolOutput {
	return o
}

func (o NodePoolOutput) ToNodePoolPtrOutput() NodePoolPtrOutput {
	return o.ToNodePoolPtrOutputWithContext(context.Background())
}

func (o NodePoolOutput) ToNodePoolPtrOutputWithContext(ctx context.Context) NodePoolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NodePool) *NodePool {
		return &v
	}).(NodePoolPtrOutput)
}

type NodePoolPtrOutput struct{ *pulumi.OutputState }

func (NodePoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodePool)(nil))
}

func (o NodePoolPtrOutput) ToNodePoolPtrOutput() NodePoolPtrOutput {
	return o
}

func (o NodePoolPtrOutput) ToNodePoolPtrOutputWithContext(ctx context.Context) NodePoolPtrOutput {
	return o
}

func (o NodePoolPtrOutput) Elem() NodePoolOutput {
	return o.ApplyT(func(v *NodePool) NodePool {
		if v != nil {
			return *v
		}
		var ret NodePool
		return ret
	}).(NodePoolOutput)
}

type NodePoolArrayOutput struct{ *pulumi.OutputState }

func (NodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodePool)(nil))
}

func (o NodePoolArrayOutput) ToNodePoolArrayOutput() NodePoolArrayOutput {
	return o
}

func (o NodePoolArrayOutput) ToNodePoolArrayOutputWithContext(ctx context.Context) NodePoolArrayOutput {
	return o
}

func (o NodePoolArrayOutput) Index(i pulumi.IntInput) NodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodePool {
		return vs[0].([]NodePool)[vs[1].(int)]
	}).(NodePoolOutput)
}

type NodePoolMapOutput struct{ *pulumi.OutputState }

func (NodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NodePool)(nil))
}

func (o NodePoolMapOutput) ToNodePoolMapOutput() NodePoolMapOutput {
	return o
}

func (o NodePoolMapOutput) ToNodePoolMapOutputWithContext(ctx context.Context) NodePoolMapOutput {
	return o
}

func (o NodePoolMapOutput) MapIndex(k pulumi.StringInput) NodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NodePool {
		return vs[0].(map[string]NodePool)[vs[1].(string)]
	}).(NodePoolOutput)
}

func init() {
	pulumi.RegisterOutputType(NodePoolOutput{})
	pulumi.RegisterOutputType(NodePoolPtrOutput{})
	pulumi.RegisterOutputType(NodePoolArrayOutput{})
	pulumi.RegisterOutputType(NodePoolMapOutput{})
}