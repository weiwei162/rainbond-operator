package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InstallMode is the mode of Rainbond cluster installation
type InstallMode string

const (
	// InstallationModeFullOnline means some Rainbond images are from the specified image repository, some are from the installation package.
	InstallationModeWithPackage InstallMode = "WithPackage"
	// InstallationModeFullOnline means all Rainbond images are from the specified image repository, not the installation package.
	InstallationModeWithoutPackage InstallMode = "WithoutPackage"
)

type ImageHub struct {
	Domain    string `json:"domain,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

// Database defines the connection information of database.
type Database struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// EtcdConfig defines the configuration of etcd client.
type EtcdConfig struct {
	// Endpoints is a list of URLs.
	Endpoints []string `json:"endpoints,omitempty"`
	// Whether to use tls to connect to etcd
	UseTLS bool `json:"useTLS,omitempty"`
	// Secret to mount to read certificate files for tls.
	CertSecret metav1.LabelSelector `json:"selector,omitempty"`
}

// RainbondClusterSpec defines the desired state of RainbondCluster
type RainbondClusterSpec struct {
	// Domain name of the image repository which Rainbond is installed
	// Default goodrain.me
	// +optional
	RainbondImageRepositoryDomain string `json:"rainbondImageRepositoryHost,omitempty"`
	// Suffix of component default domain name
	SuffixHTTPHost string `json:"suffixHTTPHost,omitempty"`
	// Ingress IP addresses of rbd-gateway. If not specified,
	// the IP of the node where the rbd-gateway is located will be used.
	GatewayIngressIPs []string `json:"gatewayIngressIPs,omitempty"`
	// Information about the node where the gateway is located.
	// If not specified, the gateway will run on nodes where all ports do not conflict.
	GatewayNodes []NodeAvailPorts `json:"gatewayNodes,omitempty"`
	// InstallMode is the mode of Rainbond cluster installation.
	InstallMode InstallMode `json:"installMode,omitempty"`

	ImageHub *ImageHub `json:"imageHub,omitempty"`
	// the storage class that rainbond component will be used.
	// rainbond-operator will create one if StorageClassName is empty
	StorageClassName string `json:"storageClassName,omitempty"`
	// the region database information that rainbond component will be used.
	// rainbond-operator will create one if DBInfo is empty
	RegionDatabase *Database `json:"regionDatabase,omitempty"`
	// the ui database information that rainbond component will be used.
	// rainbond-operator will create one if DBInfo is empty
	UIDatabase *Database `json:"uiDatabase,omitempty"`
	// the etcd connection information that rainbond component will be used.
	// rainbond-operator will create one if EtcdConfig is empty
	EtcdConfig *EtcdConfig `json:"etcdConfig,omitempty"`
}

// RainbondClusterPhase is a label for the condition of a rainbondcluster at the current time.
type RainbondClusterPhase string

// These are the valid statuses of rainbondcluster.
const (
	// RainbondClusterWaiting -
	RainbondClusterWaiting RainbondClusterPhase = "Waiting"
	// RainbondClusterPreparing -
	RainbondClusterPreparing RainbondClusterPhase = "Preparing"
	// RainbondClusterPackageProcessing means the installation package is being processed.
	RainbondClusterPackageProcessing RainbondClusterPhase = "PackageProcessing"
	// RainbondClusterRunning means all of the rainbond components has been created.
	// And at least one component is not ready.
	RainbondClusterPending RainbondClusterPhase = "Pending"
	// RainbondClusterRunning means all of the rainbond components has been created.
	// For each component controller(eg. deploy, sts, ds), at least one Pod is already Ready.
	RainbondClusterRunning RainbondClusterPhase = "Running"
)

// RainbondClusterConditionType is a valid value for RainbondClusterConditionType.Type
type RainbondClusterConditionType string

// These are valid conditions of rainbondcluster.
const (
	// StorageReady indicates whether the storage is ready.
	StorageReady RainbondClusterConditionType = "StorageReady"
	// ImageRepositoryReady indicates whether the image repository is ready.
	ImageRepositoryInstalled RainbondClusterConditionType = "ImageRepositoryInstalled"
	// PackageExtracted indicates whether the installation package has been decompressed.
	PackageExtracted RainbondClusterConditionType = "PackageExtracted"
	// ImagesLoaded means that all images from the installation package has been loaded successfully.
	ImagesLoaded RainbondClusterConditionType = "ImagesLoaded"
	// ImagesPushed means that all images from the installation package has been pushed successfully.
	ImagesPushed RainbondClusterConditionType = "ImagesPushed"
)

type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means rainbond operator
// can't decide if a resource is in the condition or not.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// RainbondClusterCondition contains details for the current condition of this rainbondcluster.
type RainbondClusterCondition struct {
	// Type is the type of the condition.
	Type RainbondClusterConditionType `json:"type"`
	// Status is the status of the condition.
	Status ConditionStatus `json:"status"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime *metav1.Time `json:"lastProbeTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

type NodeAvailPorts struct {
	NodeName string `json:"nodeName,omitempty"`
	NodeIP   string `json:"nodeIP,omitempty"`
	Ports    []int  `json:"ports,omitempty"`
}

type StorageClass struct {
	Name        string `json:"name"`
	Provisioner string `json:"provisioner"`
}

type ControllerStatus struct {
	Name          string `json:"name,omitempty"`
	Replicas      int32  `json:"replicas,omitempty"`
	ReadyReplicas int32  `json:"readyReplicas,omitempty"`
}

// RainbondClusterStatus defines the observed state of RainbondCluster
type RainbondClusterStatus struct {
	// Rainbond cluster phase
	Phase      RainbondClusterPhase       `json:"phase,omitempty"`
	Conditions []RainbondClusterCondition `json:"conditions,omitempty"`
	// A human readable message indicating details about why the pod is in this condition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`
	// A brief CamelCase message indicating details about why the pod is in this state.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`

	// +optional
	NodeAvailPorts []*NodeAvailPorts `json:"NodeAvailPorts,omitempty"`
	// List of existing StorageClasses in the cluster
	// +optional
	StorageClasses []*StorageClass `json:"storageClasses,omitempty"`
	// Destination path of the installation package extraction.
	PkgDestPath string `json:"pkgDestPath"`
	// A list of controller statuses associated with rbdcomponent.
	ControllerStatues []*ControllerStatus `json:"controllerStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RainbondCluster is the Schema for the rainbondclusters API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=rainbondclusters,scope=Namespaced
type RainbondCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RainbondClusterSpec    `json:"spec,omitempty"`
	Status *RainbondClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RainbondClusterList contains a list of RainbondCluster
type RainbondClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RainbondCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RainbondCluster{}, &RainbondClusterList{})
}