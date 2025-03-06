// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "go-mod-path/generated/go/core"
	time "time"
)

type CreateInstanceRequest struct {
	// The custom name of the instance.
	Name *string `json:"name,omitempty" url:"-"`
	// The GPU type of the instance.
	GpuType GpuType `json:"gpu_type" url:"-"`
	// The number of GPUs to attach to the instance.
	GpuCount *int `json:"gpu_count,omitempty" url:"-"`
	// The SSH key name to add to the instance. This SSH key is used to connect to the instance.
	SshKey string `json:"ssh_key" url:"-"`
	// The operating system label used to create the instance.
	OperatingSystemLabel *CreateInstanceRequestOperatingSystemLabel `json:"operating_system_label,omitempty" url:"-"`
	// The region in which to create the instance.
	Region *Region `json:"region,omitempty" url:"-"`
	// The volumes attached to the instance.
	Volumes []*VolumeInstanceResponse `json:"volumes,omitempty" url:"-"`
}

type InstancesListRequest struct {
	Page *int `json:"-" url:"page,omitempty"`
	// Include failed instances. Default is False.
	IncludeFailedInstances *bool `json:"-" url:"include_failed_instances,omitempty"`
}

type ConfigurationInstanceResponse struct {
	Id EntityId `json:"id" url:"id"`
	// The GPU model of the configuration.
	GpuModel *GpuModelResponse `json:"gpu_model,omitempty" url:"gpu_model,omitempty"`
	// The CPU model of the configuration.
	CpuModel *string `json:"cpu_model,omitempty" url:"cpu_model,omitempty"`
	// The number of GPUs in the configuration.
	GpuCount int `json:"gpu_count" url:"gpu_count"`
	// The number of CPUs in the configuration.
	CpuCount int `json:"cpu_count" url:"cpu_count"`
	// The size of NVMe in the configuration.
	NvmeStorageSizeGb int `json:"nvme_storage_size_gb" url:"nvme_storage_size_gb"`
	// The amount of RAM memory in the configuration.
	MemorySizeGb                     float64 `json:"memory_size_gb" url:"memory_size_gb"`
	EstimatedProvisioningTimeMinutes *int    `json:"estimated_provisioning_time_minutes,omitempty" url:"estimated_provisioning_time_minutes,omitempty"`
	Region                           Region  `json:"region" url:"region"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ConfigurationInstanceResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ConfigurationInstanceResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ConfigurationInstanceResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConfigurationInstanceResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConfigurationInstanceResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CreateInstanceResponse struct {
	Id                   string                                      `json:"id" url:"id"`
	Name                 string                                      `json:"name" url:"name"`
	GpuType              GpuType                                     `json:"gpu_type" url:"gpu_type"`
	OperatingSystemLabel *CreateInstanceResponseOperatingSystemLabel `json:"operating_system_label,omitempty" url:"operating_system_label,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateInstanceResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateInstanceResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateInstanceResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateInstanceResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateInstanceResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CreateInstanceResponseOperatingSystemLabel struct {
	SupportedOperatingSystem SupportedOperatingSystem
	String                   string
}

func NewCreateInstanceResponseOperatingSystemLabelFromSupportedOperatingSystem(value SupportedOperatingSystem) *CreateInstanceResponseOperatingSystemLabel {
	return &CreateInstanceResponseOperatingSystemLabel{SupportedOperatingSystem: value}
}

func NewCreateInstanceResponseOperatingSystemLabelFromString(value string) *CreateInstanceResponseOperatingSystemLabel {
	return &CreateInstanceResponseOperatingSystemLabel{String: value}
}

func (c *CreateInstanceResponseOperatingSystemLabel) UnmarshalJSON(data []byte) error {
	var valueSupportedOperatingSystem SupportedOperatingSystem
	if err := json.Unmarshal(data, &valueSupportedOperatingSystem); err == nil {
		c.SupportedOperatingSystem = valueSupportedOperatingSystem
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateInstanceResponseOperatingSystemLabel) MarshalJSON() ([]byte, error) {
	if c.SupportedOperatingSystem != "" {
		return json.Marshal(c.SupportedOperatingSystem)
	}
	if c.String != "" {
		return json.Marshal(c.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CreateInstanceResponseOperatingSystemLabelVisitor interface {
	VisitSupportedOperatingSystem(SupportedOperatingSystem) error
	VisitString(string) error
}

func (c *CreateInstanceResponseOperatingSystemLabel) Accept(visitor CreateInstanceResponseOperatingSystemLabelVisitor) error {
	if c.SupportedOperatingSystem != "" {
		return visitor.VisitSupportedOperatingSystem(c.SupportedOperatingSystem)
	}
	if c.String != "" {
		return visitor.VisitString(c.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

type EntityId = string

type GpuModelResponse struct {
	// The FluidStack unique name of the GPU model.
	Name string `json:"name" url:"name"`
	// Memory capacity of the GPU in gigabytes.
	MemorySizeGb *int `json:"memory_size_gb,omitempty" url:"memory_size_gb,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GpuModelResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GpuModelResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GpuModelResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GpuModelResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GpuModelResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type InstanceResponse struct {
	// The unique identifier of the instance.
	Id string `json:"id" url:"id"`
	// The current status of the instance.
	Status *InstanceStatus `json:"status,omitempty" url:"status,omitempty"`
	// The username used to connect to the instance. For example, to connect to the instance via SSH, use: "ssh -i <path/to/private/key> <username>@<ip_address>".
	Username *string `json:"username,omitempty" url:"username,omitempty"`
	// The SSH port used to connect to the instance.
	SshPort *string `json:"ssh_port,omitempty" url:"ssh_port,omitempty"`
	// The names of the SSH keys used to login to the instance.
	SshKeys []string `json:"ssh_keys,omitempty" url:"ssh_keys,omitempty"`
	// The IP address of the instance.
	IpAddress *string `json:"ip_address,omitempty" url:"ip_address,omitempty"`
	// The name provided when the instance was created.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The current hourly price of the instance per processor based on its current status.
	CurrentGpuHrCost *float64 `json:"current_gpu_hr_cost,omitempty" url:"current_gpu_hr_cost,omitempty"`
	// The configuration used to create the instance.
	Configuration *ConfigurationInstanceResponse `json:"configuration,omitempty" url:"configuration,omitempty"`
	// The creation date and time of the instance.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The volumes attached to the instance.
	Volumes []*VolumeInstanceResponse `json:"volumes,omitempty" url:"volumes,omitempty"`
	// The operating system used to create the instance.
	OperatingSystemLabel *InstanceResponseOperatingSystemLabel `json:"operating_system_label,omitempty" url:"operating_system_label,omitempty"`
	// The email of the user that owns the instance.
	UserEmail *string `json:"user_email,omitempty" url:"user_email,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (i *InstanceResponse) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *InstanceResponse) UnmarshalJSON(data []byte) error {
	type embed InstanceResponse
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*i),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*i = InstanceResponse(unmarshaler.embed)
	i.CreatedAt = unmarshaler.CreatedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties

	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *InstanceResponse) MarshalJSON() ([]byte, error) {
	type embed InstanceResponse
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*i),
		CreatedAt: core.NewOptionalDateTime(i.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (i *InstanceResponse) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

// The operating system used to create the instance.
type InstanceResponseOperatingSystemLabel struct {
	SupportedOperatingSystem SupportedOperatingSystem
	String                   string
}

func NewInstanceResponseOperatingSystemLabelFromSupportedOperatingSystem(value SupportedOperatingSystem) *InstanceResponseOperatingSystemLabel {
	return &InstanceResponseOperatingSystemLabel{SupportedOperatingSystem: value}
}

func NewInstanceResponseOperatingSystemLabelFromString(value string) *InstanceResponseOperatingSystemLabel {
	return &InstanceResponseOperatingSystemLabel{String: value}
}

func (i *InstanceResponseOperatingSystemLabel) UnmarshalJSON(data []byte) error {
	var valueSupportedOperatingSystem SupportedOperatingSystem
	if err := json.Unmarshal(data, &valueSupportedOperatingSystem); err == nil {
		i.SupportedOperatingSystem = valueSupportedOperatingSystem
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		i.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, i)
}

func (i InstanceResponseOperatingSystemLabel) MarshalJSON() ([]byte, error) {
	if i.SupportedOperatingSystem != "" {
		return json.Marshal(i.SupportedOperatingSystem)
	}
	if i.String != "" {
		return json.Marshal(i.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", i)
}

type InstanceResponseOperatingSystemLabelVisitor interface {
	VisitSupportedOperatingSystem(SupportedOperatingSystem) error
	VisitString(string) error
}

func (i *InstanceResponseOperatingSystemLabel) Accept(visitor InstanceResponseOperatingSystemLabelVisitor) error {
	if i.SupportedOperatingSystem != "" {
		return visitor.VisitSupportedOperatingSystem(i.SupportedOperatingSystem)
	}
	if i.String != "" {
		return visitor.VisitString(i.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", i)
}

type InstanceStatus string

const (
	InstanceStatusRunning      InstanceStatus = "running"
	InstanceStatusStarting     InstanceStatus = "starting"
	InstanceStatusPending      InstanceStatus = "pending"
	InstanceStatusFailed       InstanceStatus = "failed"
	InstanceStatusTerminated   InstanceStatus = "terminated"
	InstanceStatusStopped      InstanceStatus = "stopped"
	InstanceStatusStopping     InstanceStatus = "stopping"
	InstanceStatusProvisioning InstanceStatus = "provisioning"
)

func NewInstanceStatusFromString(s string) (InstanceStatus, error) {
	switch s {
	case "running":
		return InstanceStatusRunning, nil
	case "starting":
		return InstanceStatusStarting, nil
	case "pending":
		return InstanceStatusPending, nil
	case "failed":
		return InstanceStatusFailed, nil
	case "terminated":
		return InstanceStatusTerminated, nil
	case "stopped":
		return InstanceStatusStopped, nil
	case "stopping":
		return InstanceStatusStopping, nil
	case "provisioning":
		return InstanceStatusProvisioning, nil
	}
	var t InstanceStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (i InstanceStatus) Ptr() *InstanceStatus {
	return &i
}

type ListInstanceResponse struct {
	// The unique identifier of the instance.
	Id string `json:"id" url:"id"`
	// The current status of the instance.
	Status *InstanceStatus `json:"status,omitempty" url:"status,omitempty"`
	// The username used to connect to the instance. For example, to connect to the instance via SSH, use: "ssh -i <path/to/private/key> <username>@<ip_address>".
	Username *string `json:"username,omitempty" url:"username,omitempty"`
	// The SSH port used to connect to the instance.
	SshPort *string `json:"ssh_port,omitempty" url:"ssh_port,omitempty"`
	// The names of the SSH keys used to login to the instance.
	SshKeys []string `json:"ssh_keys,omitempty" url:"ssh_keys,omitempty"`
	// The IP address of the instance.
	IpAddress *string `json:"ip_address,omitempty" url:"ip_address,omitempty"`
	// The name provided when the instance was created.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The current hourly price of the instance per processor based on its current status.
	CurrentGpuHrCost *float64 `json:"current_gpu_hr_cost,omitempty" url:"current_gpu_hr_cost,omitempty"`
	// The configuration used to create the instance.
	Configuration *ConfigurationInstanceResponse `json:"configuration,omitempty" url:"configuration,omitempty"`
	// The creation date and time of the instance.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The volumes attached to the instance.
	Volumes []*VolumeInstanceResponse `json:"volumes,omitempty" url:"volumes,omitempty"`
	// The operating system used to create the instance.
	OperatingSystemLabel *ListInstanceResponseOperatingSystemLabel `json:"operating_system_label,omitempty" url:"operating_system_label,omitempty"`
	// The email of the user that owns the instance.
	UserEmail *string `json:"user_email,omitempty" url:"user_email,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListInstanceResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListInstanceResponse) UnmarshalJSON(data []byte) error {
	type embed ListInstanceResponse
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListInstanceResponse(unmarshaler.embed)
	l.CreatedAt = unmarshaler.CreatedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListInstanceResponse) MarshalJSON() ([]byte, error) {
	type embed ListInstanceResponse
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*l),
		CreatedAt: core.NewOptionalDateTime(l.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (l *ListInstanceResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The operating system used to create the instance.
type ListInstanceResponseOperatingSystemLabel struct {
	SupportedOperatingSystem SupportedOperatingSystem
	String                   string
}

func NewListInstanceResponseOperatingSystemLabelFromSupportedOperatingSystem(value SupportedOperatingSystem) *ListInstanceResponseOperatingSystemLabel {
	return &ListInstanceResponseOperatingSystemLabel{SupportedOperatingSystem: value}
}

func NewListInstanceResponseOperatingSystemLabelFromString(value string) *ListInstanceResponseOperatingSystemLabel {
	return &ListInstanceResponseOperatingSystemLabel{String: value}
}

func (l *ListInstanceResponseOperatingSystemLabel) UnmarshalJSON(data []byte) error {
	var valueSupportedOperatingSystem SupportedOperatingSystem
	if err := json.Unmarshal(data, &valueSupportedOperatingSystem); err == nil {
		l.SupportedOperatingSystem = valueSupportedOperatingSystem
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		l.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l ListInstanceResponseOperatingSystemLabel) MarshalJSON() ([]byte, error) {
	if l.SupportedOperatingSystem != "" {
		return json.Marshal(l.SupportedOperatingSystem)
	}
	if l.String != "" {
		return json.Marshal(l.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type ListInstanceResponseOperatingSystemLabelVisitor interface {
	VisitSupportedOperatingSystem(SupportedOperatingSystem) error
	VisitString(string) error
}

func (l *ListInstanceResponseOperatingSystemLabel) Accept(visitor ListInstanceResponseOperatingSystemLabelVisitor) error {
	if l.SupportedOperatingSystem != "" {
		return visitor.VisitSupportedOperatingSystem(l.SupportedOperatingSystem)
	}
	if l.String != "" {
		return visitor.VisitString(l.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}

type VolumeInstanceResponse struct {
	// The ID of the volume.
	Id EntityId `json:"id" url:"id"`
	// The name of the volume.
	Name string `json:"name" url:"name"`
	// The size of the volume in GB.
	SizeGb int `json:"size_gb" url:"size_gb"`
	// The status of the volume.
	Status *VolumeStatus `json:"status,omitempty" url:"status,omitempty"`
	// The current hourly rate of the volume.
	CostGbHr string `json:"cost_gb_hr" url:"cost_gb_hr"`
	// The creation time of the volume.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The update time of the volume.
	UpdatedAt *time.Time `json:"updated_at,omitempty" url:"updated_at,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (v *VolumeInstanceResponse) GetExtraProperties() map[string]interface{} {
	return v.extraProperties
}

func (v *VolumeInstanceResponse) UnmarshalJSON(data []byte) error {
	type embed VolumeInstanceResponse
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
		UpdatedAt *core.DateTime `json:"updated_at,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = VolumeInstanceResponse(unmarshaler.embed)
	v.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	v.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()

	extraProperties, err := core.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties

	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *VolumeInstanceResponse) MarshalJSON() ([]byte, error) {
	type embed VolumeInstanceResponse
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
		UpdatedAt *core.DateTime `json:"updated_at,omitempty"`
	}{
		embed:     embed(*v),
		CreatedAt: core.NewOptionalDateTime(v.CreatedAt),
		UpdatedAt: core.NewOptionalDateTime(v.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (v *VolumeInstanceResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type VolumeStatus string

const (
	VolumeStatusCreating  VolumeStatus = "creating"
	VolumeStatusAttached  VolumeStatus = "attached"
	VolumeStatusAttaching VolumeStatus = "attaching"
	VolumeStatusReady     VolumeStatus = "ready"
	VolumeStatusUnknown   VolumeStatus = "unknown"
	VolumeStatusError     VolumeStatus = "error"
	VolumeStatusDeleted   VolumeStatus = "deleted"
)

func NewVolumeStatusFromString(s string) (VolumeStatus, error) {
	switch s {
	case "creating":
		return VolumeStatusCreating, nil
	case "attached":
		return VolumeStatusAttached, nil
	case "attaching":
		return VolumeStatusAttaching, nil
	case "ready":
		return VolumeStatusReady, nil
	case "unknown":
		return VolumeStatusUnknown, nil
	case "error":
		return VolumeStatusError, nil
	case "deleted":
		return VolumeStatusDeleted, nil
	}
	var t VolumeStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (v VolumeStatus) Ptr() *VolumeStatus {
	return &v
}

// The operating system label used to create the instance.
type CreateInstanceRequestOperatingSystemLabel struct {
	SupportedOperatingSystem SupportedOperatingSystem
	String                   string
}

func NewCreateInstanceRequestOperatingSystemLabelFromSupportedOperatingSystem(value SupportedOperatingSystem) *CreateInstanceRequestOperatingSystemLabel {
	return &CreateInstanceRequestOperatingSystemLabel{SupportedOperatingSystem: value}
}

func NewCreateInstanceRequestOperatingSystemLabelFromString(value string) *CreateInstanceRequestOperatingSystemLabel {
	return &CreateInstanceRequestOperatingSystemLabel{String: value}
}

func (c *CreateInstanceRequestOperatingSystemLabel) UnmarshalJSON(data []byte) error {
	var valueSupportedOperatingSystem SupportedOperatingSystem
	if err := json.Unmarshal(data, &valueSupportedOperatingSystem); err == nil {
		c.SupportedOperatingSystem = valueSupportedOperatingSystem
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateInstanceRequestOperatingSystemLabel) MarshalJSON() ([]byte, error) {
	if c.SupportedOperatingSystem != "" {
		return json.Marshal(c.SupportedOperatingSystem)
	}
	if c.String != "" {
		return json.Marshal(c.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CreateInstanceRequestOperatingSystemLabelVisitor interface {
	VisitSupportedOperatingSystem(SupportedOperatingSystem) error
	VisitString(string) error
}

func (c *CreateInstanceRequestOperatingSystemLabel) Accept(visitor CreateInstanceRequestOperatingSystemLabelVisitor) error {
	if c.SupportedOperatingSystem != "" {
		return visitor.VisitSupportedOperatingSystem(c.SupportedOperatingSystem)
	}
	if c.String != "" {
		return visitor.VisitString(c.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
