package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManagedKafkaConnectionSpec contains credentials and connection parameters to Managed Kafka
type ManagedKafkaConnectionSpec struct {
	BootstrapServer BootstrapServerSpec `json:"bootstrapServer"`
	Credentials     CredentialsSpec     `json:"credentials"`
}

// BootstrapServerSpec contains server host information that can be used to connecto the Managed Kafka
type BootstrapServerSpec struct {
	// Host full host to Managed Kafka Service including port
	Host string `json:"host,omitempty"`
}

// CredentialType enumeration for types of credentails
type CredentialType string

const (
	// ClientCredentials ... type
	ClientCredentials CredentialType = "ClientCredentials"
)

// CredentialsSpec specification containing various formats of credentials
type CredentialsSpec struct {
	// Type of the credential format. For example "ClientCredentials"
	Kind CredentialType `json:"kind,omitempty"`
	// ClientID Represents username in client credentials
	CientID string `json:"clientID,omitempty"`
	// ClientID Represents password in client credentials
	ClientSecret string `json:"clientSecret,omitempty"`
}

// ManagedKafkaConnectionStatus defines the observed state of ManagedKafkaConnection
// type ManagedKafkaConnectionStatus struct {

// }

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// ManagedKafkaConnection schema
type ManagedKafkaConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ManagedKafkaConnectionSpec `json:"spec,omitempty"`
	// Status ManagedKafkaConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedKafkaConnectionList contains a list of ManagedKafkaConnection
type ManagedKafkaConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKafkaConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManagedKafkaConnection{}, &ManagedKafkaConnectionList{})
}
