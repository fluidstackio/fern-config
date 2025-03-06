// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "go-mod-path/generated/go/core"
)

type OperatingSystemResponse struct {
	// The friendly name of the operating system.
	Name string `json:"name" url:"name"`
	// The description of the operating system, detailing the pre-installed packages and customisations, if any.
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	// The unique slug identifier of the operating system.
	Label SupportedOperatingSystem `json:"label" url:"label"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (o *OperatingSystemResponse) GetExtraProperties() map[string]interface{} {
	return o.extraProperties
}

func (o *OperatingSystemResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler OperatingSystemResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OperatingSystemResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties

	o._rawJSON = json.RawMessage(data)
	return nil
}

func (o *OperatingSystemResponse) String() string {
	if len(o._rawJSON) > 0 {
		if value, err := core.StringifyJSON(o._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}
