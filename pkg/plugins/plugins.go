package plugins

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

// SamplePlugin defined two extension points
// PreFilter, Filter
type SamplePlugin struct{}

var (
	_ framework.PreFilterPlugin = &SamplePlugin{}
	_ framework.FilterPlugin    = &SamplePlugin{}
)

// Name of SamplePlugin: myscheduler
const Name = "myscheduler"

// Name function Return Name of this plugin
func (s *SamplePlugin) Name() string {
	return Name
}

// PreFilter print logs and return Success in this extension point
func (s *SamplePlugin) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("This is custom scheduler stage of: PreFilter")
	klog.V(3).Infof("prefilter pod: %v, pod-namespace: %v, pod-spec-container: %v", pod.Name, pod.Namespace, pod.Spec.Containers)

	jsonstr, errjson := json.Marshal(pod)
	if errjson != nil {
		resp, err := http.Post(
			"http://192.168.229.1:8080/print",
			"application/json",
			strings.NewReader(string(jsonstr)),
		)
		if err != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			klog.V(3).Infof("Get response from remote: %s", body)
		} else {
			klog.V(3).Infof("Error: %v", err)
		}
		return framework.NewStatus(framework.Success, "Pod json construct successful and post to remote server, schedule ok")
	}

	klog.V(3).Infof("Json construction error: %v", errjson)
	return framework.NewStatus(framework.Unschedulable, "Pod json construct unsuccessful and not post to remote server, unscheduable")
}

// PreFilterExtensions reutrn nil
func (s *SamplePlugin) PreFilterExtensions() framework.PreFilterExtensions {
	klog.V(3).Infof("This is custom scheduler stage of: PreFilter using PreFilterExtensions")
	return nil
}

// Filter extension point checks whether pod name is nginx
// If not, set pod status: Uncheduable
func (s *SamplePlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *nodeinfo.NodeInfo) *framework.Status {
	klog.V(3).Infof("This is custom scheduler stage of: Filter")
	if pod.Name != "nginx" {
		return framework.NewStatus(framework.Unschedulable, "only pod name 'nginx' is allowed")
	}
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, node.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

// New create a new SamplePlugin instance
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &SamplePlugin{}, nil
}
