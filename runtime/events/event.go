/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package events

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These constants define valid event severity values.
const (
	// EventSeverityTrace represents a trace event, usually
	// informing about actions taken during reconciliation.
	EventSeverityTrace string = "trace"
	// EventSeverityInfo represents an informational event, usually
	// informing about changes.
	EventSeverityInfo string = "info"
	// EventSeverityError represent an error event, usually a warning
	// that something goes wrong.
	EventSeverityError string = "error"
)

// Event is a report of an event issued by a controller.
// +kubebuilder:object:generate=true
type Event struct {
	// The object that this event is about.
	// +required
	InvolvedObject corev1.ObjectReference `json:"involvedObject"`

	// Severity type of this event (trace, info, error)
	// +kubebuilder:validation:Enum=trace,info;error
	// +required
	Severity string `json:"severity"`

	// The time at which this event was recorded.
	// +required
	Timestamp metav1.Time `json:"timestamp"`

	// A human-readable description of this event.
	// Maximum length 39,000 characters.
	// +kubebuilder:validation:MaxLength=39000
	// +required
	Message string `json:"message"`

	// A machine understandable string that gives the reason
	// for the transition into the object's current status.
	// +required
	Reason string `json:"reason"`

	// Metadata of this event, e.g. apply change set.
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`

	// Name of the controller that emitted this event, e.g. `source-controller`.
	// +required
	ReportingController string `json:"reportingController"`

	// ID of the controller instance, e.g. `source-controller-xyzf`.
	// +optional
	ReportingInstance string `json:"reportingInstance,omitempty"`
}
