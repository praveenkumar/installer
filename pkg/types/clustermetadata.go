package types

// ClusterMetadata contains information
// regarding the cluster that was created by installer.
type ClusterMetadata struct {
	ClusterName             string `json:"clusterName"`
	ClusterPlatformMetadata `json:",inline"`
}

// ClusterPlatformMetadata contains metadata for platfrom.
type ClusterPlatformMetadata struct {
	AWS       *ClusterAWSPlatformMetadata       `json:"aws,omitempty"`
	OpenStack *ClusterOpenStackPlatformMetadata `json:"openstack,omitempty"`
	Libvirt   *ClusterLibvirtPlatformMetadata   `json:"libvirt,omitempty"`
}

// Platform returns a string representation of the platform
// (e.g. "aws" if AWS is non-nil).  It returns an empty string if no
// platform is configured.
func (cpm *ClusterPlatformMetadata) Platform() string {
	if cpm == nil {
		return ""
	}
	if cpm.AWS != nil {
		return "aws"
	}
	if cpm.Libvirt != nil {
		return "libvirt"
	}
	return ""
}

// ClusterAWSPlatformMetadata contains AWS metadata.
type ClusterAWSPlatformMetadata struct {
	Region string `json:"region"`
	// Most AWS resources are tagged with these tags as identifier.
	Identifier map[string]string `json:"identifier"`
}

// ClusterOpenStackPlatformMetadata contains OpenStack metadata.
type ClusterOpenStackPlatformMetadata struct {
	Region string `json:"region"`
	// Most OpenStack resources are tagged with these tags as identifier.
	Identifier map[string]string `json:"identifier"`
}

// ClusterLibvirtPlatformMetadata contains libvirt metadata.
type ClusterLibvirtPlatformMetadata struct {
	URI string `json:"uri"`
}
