package simplelog

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

// SimpleLog defined two extension points
// PreFilter, Filter
type SimpleLog struct{}

var (
	_ framework.PreFilterPlugin = &SimpleLog{}
	_ framework.FilterPlugin    = &SimpleLog{}
)

// Name of SimpleLog plugin: simplelog
const Name = "simplelog"

// Name function return name of this plugin
func (s *SimpleLog) Name() string {
	return Name
}

// PreFilter extension point print logs and return Success
func (s *SimpleLog) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("This is custom scheduler: [%v], stage of: [PreFilter]", s.Name())
	klog.V(3).Infof("prefilter pod: %v", pod.Name)
	klog.V(3).Infof("Finished custom scheduler: [%v] [PreFilter] extension point", s.Name())
	return framework.NewStatus(framework.Success, "")
}

// PreFilterExtensions reutrn nil
func (s *SimpleLog) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

// Filter extension point prints pod and node info and return Success
func (s *SimpleLog) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *nodeinfo.NodeInfo) *framework.Status {
	klog.V(3).Infof("This is custom scheduler: [%v], stage of: [Filter]", s.Name())
	klog.V(3).Infof("Filter pod: %v, node: %v", pod.Name, node.Node().Name)
	klog.V(3).Infof("Finished custom scheduler: [%v] [Filter] extension point", s.Name())
	return framework.NewStatus(framework.Success, "")
}

// New create a new SimpleLog instance
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &SimpleLog{}, nil
}
