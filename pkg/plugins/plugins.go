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

type stateData struct {
	data string
}

func (s *stateData) Clone() framework.StateData {
	copy := &stateData{
		data: s.data,
	}
	return copy
}

// PreFilter print logs and return Success in this extension point
func (s *SamplePlugin) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("This is custom scheduler stage of: PreFilter")
	klog.V(3).Infof("prefilter pod: %v", pod.Name)

	podjson, _ := json.Marshal(pod)
	resp, err := http.Post(
		"http://192.168.229.1:8080/print",
		"application/json",
		strings.NewReader(string(podjson)),
	)
	// If http post error, let this pod unscheduable
	if err != nil {
		klog.V(3).Infof("Error: %v", err)
		return framework.NewStatus(framework.Unschedulable, "Cant send pod json to remote server, schedule terminated")
	}

	// This body should return nodes info that will use in future extension points
	body, _ := ioutil.ReadAll(resp.Body)
	klog.V(3).Infof("Get response from remote: %s", body)
	// Write response data to CycleState storage
	state.Write(framework.StateKey(pod.Name), &stateData{data: string(body)})
	return framework.NewStatus(framework.Success, "Pod info send to remote server, get essential nodes info")
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

	// Read data from CycleState storage
	if v, e := state.Read(framework.StateKey(pod.Name)); e == nil {
		if data, ok := v.(*stateData); ok {
			klog.V(3).Infof("Get data from PreFilter extension point: %s", data)
		}
	}

    // Check whether pod name is nginx
    // If not, return status with a message notify that
    // this pod name is not nginx
	if pod.Name != "nginx" {
		return framework.NewStatus(framework.Success, "pod name not 'nginx' is allowed")
	}
	klog.V(3).Infof("filter pod: %v, node: %v", pod.Name, node.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

// New create a new SamplePlugin instance
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &SamplePlugin{}, nil
}
