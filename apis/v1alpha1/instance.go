// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InstanceSpec defines the desired state of Instance.
//
// Describes an instance.
type InstanceSpec struct {

	// The block device mapping, which defines the EBS volumes and instance store
	// volumes to attach to the instance at launch. For more information, see Block
	// device mappings (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
	// in the Amazon EC2 User Guide.
	BlockDeviceMappings []*BlockDeviceMapping `json:"blockDeviceMappings,omitempty"`
	// Information about the Capacity Reservation targeting option. If you do not
	// specify this parameter, the instance's Capacity Reservation preference defaults
	// to open, which enables it to run in any open Capacity Reservation that has
	// matching attributes (instance type, platform, Availability Zone, and tenancy).
	CapacityReservationSpecification *CapacityReservationSpecification `json:"capacityReservationSpecification,omitempty"`
	// The CPU options for the instance. For more information, see Optimize CPU
	// options (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html)
	// in the Amazon EC2 User Guide.
	CPUOptions *CPUOptionsRequest `json:"cpuOptions,omitempty"`
	// The credit option for CPU usage of the burstable performance instance. Valid
	// values are standard and unlimited. To change this attribute after launch,
	// use ModifyInstanceCreditSpecification (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html).
	// For more information, see Burstable performance instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
	// in the Amazon EC2 User Guide.
	//
	// Default: standard (T2 instances) or unlimited (T3/T3a/T4g instances)
	//
	// For T3 instances with host tenancy, only standard is supported.
	CreditSpecification *CreditSpecificationRequest `json:"creditSpecification,omitempty"`
	// Indicates whether an instance is enabled for stop protection. For more information,
	// see Stop protection (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection).
	DisableAPIStop *bool `json:"disableAPIStop,omitempty"`
	// If you set this parameter to true, you can't terminate the instance using
	// the Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute
	// after launch, use ModifyInstanceAttribute (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html).
	// Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate,
	// you can terminate the instance by running the shutdown command from the instance.
	//
	// Default: false
	DisableAPITermination *bool `json:"disableAPITermination,omitempty"`
	// Indicates whether the instance is optimized for Amazon EBS I/O. This optimization
	// provides dedicated throughput to Amazon EBS and an optimized configuration
	// stack to provide optimal Amazon EBS I/O performance. This optimization isn't
	// available with all instance types. Additional usage charges apply when using
	// an EBS-optimized instance.
	//
	// Default: false
	EBSOptimized *bool `json:"ebsOptimized,omitempty"`
	// An elastic GPU to associate with the instance.
	//
	// Amazon Elastic Graphics reached end of life on January 8, 2024.
	ElasticGPUSpecification []*ElasticGPUSpecification `json:"elasticGPUSpecification,omitempty"`
	// An elastic inference accelerator to associate with the instance.
	//
	// Amazon Elastic Inference is no longer available.
	ElasticInferenceAccelerators []*ElasticInferenceAccelerator `json:"elasticInferenceAccelerators,omitempty"`
	// Indicates whether the instance is enabled for Amazon Web Services Nitro Enclaves.
	// For more information, see What is Amazon Web Services Nitro Enclaves? (https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html)
	// in the Amazon Web Services Nitro Enclaves User Guide.
	//
	// You can't enable Amazon Web Services Nitro Enclaves and hibernation on the
	// same instance.
	EnclaveOptions *EnclaveOptionsRequest `json:"enclaveOptions,omitempty"`
	// Indicates whether an instance is enabled for hibernation. This parameter
	// is valid only if the instance meets the hibernation prerequisites (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html).
	// For more information, see Hibernate your Amazon EC2 instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html)
	// in the Amazon EC2 User Guide.
	//
	// You can't enable hibernation and Amazon Web Services Nitro Enclaves on the
	// same instance.
	HibernationOptions *HibernationOptionsRequest `json:"hibernationOptions,omitempty"`
	// The name or Amazon Resource Name (ARN) of an IAM instance profile.
	IAMInstanceProfile *IAMInstanceProfileSpecification `json:"iamInstanceProfile,omitempty"`
	// The ID of the AMI. An AMI ID is required to launch an instance and must be
	// specified here or in a launch template.
	ImageID *string `json:"imageID,omitempty"`
	// Indicates whether an instance stops or terminates when you initiate shutdown
	// from the instance (using the operating system command for system shutdown).
	//
	// Default: stop
	InstanceInitiatedShutdownBehavior *string `json:"instanceInitiatedShutdownBehavior,omitempty"`
	// The market (purchasing) option for the instances.
	//
	// For RunInstances, persistent Spot Instance requests are only supported when
	// InstanceInterruptionBehavior is set to either hibernate or stop.
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"instanceMarketOptions,omitempty"`
	// The instance type. For more information, see Amazon EC2 instance types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon EC2 User Guide.
	InstanceType *string `json:"instanceType,omitempty"`
	// The number of IPv6 addresses to associate with the primary network interface.
	// Amazon EC2 chooses the IPv6 addresses from the range of your subnet. You
	// cannot specify this option and the option to assign specific IPv6 addresses
	// in the same request. You can specify this option if you've specified a minimum
	// number of instances to launch.
	//
	// You cannot specify this option and the network interfaces option in the same
	// request.
	IPv6AddressCount *int64 `json:"ipv6AddressCount,omitempty"`
	// The IPv6 addresses from the range of the subnet to associate with the primary
	// network interface. You cannot specify this option and the option to assign
	// a number of IPv6 addresses in the same request. You cannot specify this option
	// if you've specified a minimum number of instances to launch.
	//
	// You cannot specify this option and the network interfaces option in the same
	// request.
	IPv6Addresses []*InstanceIPv6Address `json:"ipv6Addresses,omitempty"`
	// The ID of the kernel.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
	// information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html)
	// in the Amazon EC2 User Guide.
	KernelID *string `json:"kernelID,omitempty"`
	// The name of the key pair. You can create a key pair using CreateKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html)
	// or ImportKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html).
	//
	// If you do not specify a key pair, you can't connect to the instance unless
	// you choose an AMI that is configured to allow users another way to log in.
	KeyName *string `json:"keyName,omitempty"`
	// The launch template. Any additional parameters that you specify for the new
	// instance overwrite the corresponding parameters included in the launch template.
	LaunchTemplate *LaunchTemplateSpecification `json:"launchTemplate,omitempty"`
	// The license configurations.
	LicenseSpecifications []*LicenseConfigurationRequest `json:"licenseSpecifications,omitempty"`
	// The maintenance and recovery options for the instance.
	MaintenanceOptions *InstanceMaintenanceOptionsRequest `json:"maintenanceOptions,omitempty"`
	// The maximum number of instances to launch. If you specify a value that is
	// more capacity than Amazon EC2 can launch in the target Availability Zone,
	// Amazon EC2 launches the largest possible number of instances above the specified
	// minimum count.
	//
	// Constraints: Between 1 and the quota for the specified instance type for
	// your account for this Region. For more information, see Amazon EC2 instance
	// type quotas (https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-instance-quotas.html).
	MaxCount *int64 `json:"maxCount,omitempty"`
	// The metadata options for the instance. For more information, see Instance
	// metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).
	MetadataOptions *InstanceMetadataOptionsRequest `json:"metadataOptions,omitempty"`
	// The minimum number of instances to launch. If you specify a value that is
	// more capacity than Amazon EC2 can provide in the target Availability Zone,
	// Amazon EC2 does not launch any instances.
	//
	// Constraints: Between 1 and the quota for the specified instance type for
	// your account for this Region. For more information, see Amazon EC2 instance
	// type quotas (https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-instance-quotas.html).
	MinCount *int64 `json:"minCount,omitempty"`
	// Specifies whether detailed monitoring is enabled for the instance.
	Monitoring *RunInstancesMonitoringEnabled `json:"monitoring,omitempty"`
	// The network interfaces to associate with the instance.
	NetworkInterfaces []*InstanceNetworkInterfaceSpecification `json:"networkInterfaces,omitempty"`
	// The placement for the instance.
	Placement *Placement `json:"placement,omitempty"`
	// The options for the instance hostname. The default values are inherited from
	// the subnet. Applies only if creating a network interface, not attaching an
	// existing one.
	PrivateDNSNameOptions *PrivateDNSNameOptionsRequest `json:"privateDNSNameOptions,omitempty"`
	// The primary IPv4 address. You must specify a value from the IPv4 address
	// range of the subnet.
	//
	// Only one private IP address can be designated as primary. You can't specify
	// this option if you've specified the option to designate a private IP address
	// as the primary IP address in a network interface specification. You cannot
	// specify this option if you're launching more than one instance in the request.
	//
	// You cannot specify this option and the network interfaces option in the same
	// request.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`
	// The ID of the RAM disk to select. Some kernels require additional drivers
	// at launch. Check the kernel requirements for information about whether you
	// need to specify a RAM disk. To find kernel requirements, go to the Amazon
	// Web Services Resource Center and search for the kernel ID.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more
	// information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html)
	// in the Amazon EC2 User Guide.
	RAMDiskID *string `json:"ramDiskID,omitempty"`
	// The IDs of the security groups. You can create a security group using CreateSecurityGroup
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html).
	//
	// If you specify a network interface, you must specify any security groups
	// as part of the network interface instead of using this parameter.
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	// [Default VPC] The names of the security groups.
	//
	// If you specify a network interface, you must specify any security groups
	// as part of the network interface instead of using this parameter.
	//
	// Default: Amazon EC2 uses the default security group.
	SecurityGroups []*string `json:"securityGroups,omitempty"`
	// The ID of the subnet to launch the instance into.
	//
	// If you specify a network interface, you must specify any subnets as part
	// of the network interface instead of using this parameter.
	SubnetID  *string                                  `json:"subnetID,omitempty"`
	SubnetRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"subnetRef,omitempty"`
	// The tags. The value parameter is required, but if you don't want the tag
	// to have a value, specify the parameter with no value, and we set the value
	// to an empty string.
	Tags []*Tag `json:"tags,omitempty"`
	// The user data to make available to the instance. User data must be base64-encoded.
	// Depending on the tool or SDK that you're using, the base64-encoding might
	// be performed for you. For more information, see Work with instance user data
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-add-user-data.html).
	UserData *string `json:"userData,omitempty"`
}

// InstanceStatus defines the observed state of Instance
type InstanceStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The AMI launch index, which can be used to find this instance in the launch
	// group.
	// +kubebuilder:validation:Optional
	AMILaunchIndex *int64 `json:"amiLaunchIndex,omitempty"`
	// The architecture of the image.
	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty"`
	// The boot mode that was specified by the AMI. If the value is uefi-preferred,
	// the AMI supports both UEFI and Legacy BIOS. The currentInstanceBootMode parameter
	// is the boot mode that is used to boot the instance at launch or start.
	//
	// The operating system contained in the AMI must be configured to support the
	// specified boot mode.
	//
	// For more information, see Boot modes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html)
	// in the Amazon EC2 User Guide.
	// +kubebuilder:validation:Optional
	BootMode *string `json:"bootMode,omitempty"`
	// The ID of the Capacity Reservation.
	// +kubebuilder:validation:Optional
	CapacityReservationID *string `json:"capacityReservationID,omitempty"`
	// Deprecated.
	//
	// Amazon Elastic Graphics reached end of life on January 8, 2024.
	// +kubebuilder:validation:Optional
	ElasticGPUAssociations []*ElasticGPUAssociation `json:"elasticGPUAssociations,omitempty"`
	// Deprecated
	//
	// Amazon Elastic Inference is no longer available.
	// +kubebuilder:validation:Optional
	ElasticInferenceAcceleratorAssociations []*ElasticInferenceAcceleratorAssociation `json:"elasticInferenceAcceleratorAssociations,omitempty"`
	// Specifies whether enhanced networking with ENA is enabled.
	// +kubebuilder:validation:Optional
	ENASupport *bool `json:"enaSupport,omitempty"`
	// The hypervisor type of the instance. The value xen is used for both Xen and
	// Nitro hypervisors.
	// +kubebuilder:validation:Optional
	Hypervisor *string `json:"hypervisor,omitempty"`
	// The ID of the instance.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceID,omitempty"`
	// Indicates whether this is a Spot Instance or a Scheduled Instance.
	// +kubebuilder:validation:Optional
	InstanceLifecycle *string `json:"instanceLifecycle,omitempty"`
	// The IPv6 address assigned to the instance.
	// +kubebuilder:validation:Optional
	IPv6Address *string `json:"ipv6Address,omitempty"`
	// The time that the instance was last launched. To determine the time that
	// instance was first launched, see the attachment time for the primary network
	// interface.
	// +kubebuilder:validation:Optional
	LaunchTime *metav1.Time `json:"launchTime,omitempty"`
	// The license configurations for the instance.
	// +kubebuilder:validation:Optional
	Licenses []*LicenseConfiguration `json:"licenses,omitempty"`
	// The Amazon Resource Name (ARN) of the Outpost.
	// +kubebuilder:validation:Optional
	OutpostARN *string `json:"outpostARN,omitempty"`
	// The platform. This value is windows for Windows instances; otherwise, it
	// is empty.
	// +kubebuilder:validation:Optional
	Platform *string `json:"platform,omitempty"`
	// The platform details value for the instance. For more information, see AMI
	// billing information fields (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/billing-info-fields.html)
	// in the Amazon EC2 User Guide.
	// +kubebuilder:validation:Optional
	PlatformDetails *string `json:"platformDetails,omitempty"`
	// [IPv4 only] The private DNS hostname name assigned to the instance. This
	// DNS hostname can only be used inside the Amazon EC2 network. This name is
	// not available until the instance enters the running state.
	//
	// The Amazon-provided DNS server resolves Amazon-provided private DNS hostnames
	// if you've enabled DNS resolution and DNS hostnames in your VPC. If you are
	// not using the Amazon-provided DNS server in your VPC, your custom domain
	// name servers must resolve the hostname as appropriate.
	// +kubebuilder:validation:Optional
	PrivateDNSName *string `json:"privateDNSName,omitempty"`
	// The product codes attached to this instance, if applicable.
	// +kubebuilder:validation:Optional
	ProductCodes []*ProductCode `json:"productCodes,omitempty"`
	// [IPv4 only] The public DNS name assigned to the instance. This name is not
	// available until the instance enters the running state. This name is only
	// available if you've enabled DNS hostnames for your VPC.
	// +kubebuilder:validation:Optional
	PublicDNSName *string `json:"publicDNSName,omitempty"`
	// The public IPv4 address, or the Carrier IP address assigned to the instance,
	// if applicable.
	//
	// A Carrier IP address only applies to an instance launched in a subnet associated
	// with a Wavelength Zone.
	// +kubebuilder:validation:Optional
	PublicIPAddress *string `json:"publicIPAddress,omitempty"`
	// The device name of the root device volume (for example, /dev/sda1).
	// +kubebuilder:validation:Optional
	RootDeviceName *string `json:"rootDeviceName,omitempty"`
	// The root device type used by the AMI. The AMI can use an EBS volume or an
	// instance store volume.
	// +kubebuilder:validation:Optional
	RootDeviceType *string `json:"rootDeviceType,omitempty"`
	// Indicates whether source/destination checking is enabled.
	// +kubebuilder:validation:Optional
	SourceDestCheck *bool `json:"sourceDestCheck,omitempty"`
	// If the request is a Spot Instance request, the ID of the request.
	// +kubebuilder:validation:Optional
	SpotInstanceRequestID *string `json:"spotInstanceRequestID,omitempty"`
	// Specifies whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	// +kubebuilder:validation:Optional
	SRIOVNetSupport *string `json:"sriovNetSupport,omitempty"`
	// The current state of the instance.
	// +kubebuilder:validation:Optional
	State *InstanceState `json:"state,omitempty"`
	// The reason for the most recent state transition.
	// +kubebuilder:validation:Optional
	StateReason *StateReason `json:"stateReason,omitempty"`
	// The reason for the most recent state transition. This might be an empty string.
	// +kubebuilder:validation:Optional
	StateTransitionReason *string `json:"stateTransitionReason,omitempty"`
	// If the instance is configured for NitroTPM support, the value is v2.0. For
	// more information, see NitroTPM (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html)
	// in the Amazon EC2 User Guide.
	// +kubebuilder:validation:Optional
	TPMSupport *string `json:"tpmSupport,omitempty"`
	// The usage operation value for the instance. For more information, see AMI
	// billing information fields (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/billing-info-fields.html)
	// in the Amazon EC2 User Guide.
	// +kubebuilder:validation:Optional
	UsageOperation *string `json:"usageOperation,omitempty"`
	// The time that the usage operation was last updated.
	// +kubebuilder:validation:Optional
	UsageOperationUpdateTime *metav1.Time `json:"usageOperationUpdateTime,omitempty"`
	// The virtualization type of the instance.
	// +kubebuilder:validation:Optional
	VirtualizationType *string `json:"virtualizationType,omitempty"`
	// The ID of the VPC in which the instance is running.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcID,omitempty"`
}

// Instance is the Schema for the Instances API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ID",type=string,priority=0,JSONPath=`.status.instanceID`
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// InstanceList contains a list of Instance
// +kubebuilder:object:root=true
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
