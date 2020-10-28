package httpreq

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

// HTTPReq defined two extension points
// PreFilter, Filter
type HTTPReq struct{}

var (
	_ framework.PreFilterPlugin = &HTTPReq{}
	_ framework.FilterPlugin    = &HTTPReq{}
)

// Name of HTTPReq plugin: httpreq
const Name = "httpreq"

// Name function return name of this plugin
func (h *HTTPReq) Name() string {
	return Name
}

type stateData struct {
	data string
}

func (h *stateData) Clone() framework.StateData {
	copy := &stateData{
		data: h.data,
	}
	return copy
}

// PreFilter print logs and send pod json to remote server and get returned result back
func (h *HTTPReq) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	klog.V(3).Infof("This is custom scheduler: [%v], stage: [PreFilter]", h.Name())
	klog.V(3).Infof("prefilter pod: %v", pod.Name)

	podjson, _ := json.Marshal(pod)
	resp, err := http.Post(
		"http://192.168.229.1:8080/print",
		"application/json",
		strings.NewReader(string(podjson)),
	)
	// If http post error, let this pod unscheduable
	if err != nil {
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
func (h *HTTPReq) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}

// Filter extension point get selected nodeinfo from framework.CycleState
// Check current node whether matches nodeinfo
func (h *HTTPReq) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *nodeinfo.NodeInfo) *framework.Status {
	klog.V(3).Infof("This is custom scheduler: [%v], stage: [Filter]", h.Name())

	var data *stateData
	var ok bool

	// Read data from CycleState storage
	if v, e := state.Read(framework.StateKey(pod.Name)); e == nil {
		if data, ok = v.(*stateData); ok {
			klog.V(3).Infof("Get data from PreFilter extension point: %s", data)
		}
	}

	// TODO: check data or data.data whether matches node.node.Name()
	if data.data != node.Node().Name {
		// TODO: change this to Unscheduable
		return framework.NewStatus(framework.Success, "Node name dosnot match, but schedule")
	}

	return framework.NewStatus(framework.Success, "")
}

// New creates a new HTTPReq instance
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &HTTPReq{}, nil
}
