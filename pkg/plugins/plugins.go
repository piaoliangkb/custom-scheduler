package plugins

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

type SamplePlugin struct{}

var (
	_ framework.PreFilterPlugin = &SamplePlugin{}
	_ framework.FilterPlugin    = &SamplePlugin{}
)

// plugin-name
const Name = "myscheduler"

func (s *SamplePlugin) Name() string {
	return Name
}

func (s *SamplePlugin) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("This is custom scheduler stage of: PreFilter")
	klog.V(3).Infof("prefilter pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *SamplePlugin) PreFilterExtensions() framework.PreFilterExtensions {
	klog.V(3).Infof("This is custom scheduler stage of: PreFilter using PreFilterExtensions")
	return nil
}

func (s *SamplePlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("This is custom scheduler stage of: Filter")
	if pod.Name != "nginx" {
		return framework.NewStatus(framework.Unschedulable, "only pod name 'nginx' is allowed")
	}
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, node.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

// release-1.19 pkg/scheduler/framework/runtime/registry.go
//type PluginFactory = func(configuration *runtime.Unknown, f FrameworkHandle) (Plugin, error)
func New(_ runtime.Object, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &SamplePlugin{}, nil
}
