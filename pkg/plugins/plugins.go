package plugins

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"
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

func (s *SamplePlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *nodeinfo.NodeInfo) *framework.Status {
	klog.V(3).Infof("This is custom scheduler stage of: Filter")
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, node.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

// func (s *Sample) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *nodeinfo.NodeInfo) *framework.Status {
// 	klog.V(3).Infof("filter pod: %v", pod.Name)
// 	return framework.NewStatus(framework.Success, "")
// }
//
// func (s *Sample) PreBind(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
// 	nodeInfo, err := s.handle.SnapshotSharedLister().NodeInfos().Get(nodeName)
// 	if err != nil {
// 		return framework.NewStatus(framework.Error, err.Error())
// 	}
// 	klog.V(3).Infof("prebind node info: %+v", nodeInfo.Node())
// 	return framework.NewStatus(framework.Success, "")
// }

// release-1.19 pkg/scheduler/framework/runtime/registry.go
//type PluginFactory = func(configuration *runtime.Unknown, f FrameworkHandle) (Plugin, error)
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &SamplePlugin{}, nil
}
