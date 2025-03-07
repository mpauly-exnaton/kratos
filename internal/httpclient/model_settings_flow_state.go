/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.
 *
 * API version:
 * Contact: office@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SettingsFlowState The experimental state represents the state of a settings flow. This field is EXPERIMENTAL and subject to change!
type SettingsFlowState string

// List of settingsFlowState
const (
	SETTINGSFLOWSTATE_SHOW_FORM SettingsFlowState = "show_form"
	SETTINGSFLOWSTATE_SUCCESS   SettingsFlowState = "success"
)

func (v *SettingsFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SettingsFlowState(value)
	for _, existing := range []SettingsFlowState{"show_form", "success"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SettingsFlowState", value)
}

// Ptr returns reference to settingsFlowState value
func (v SettingsFlowState) Ptr() *SettingsFlowState {
	return &v
}

type NullableSettingsFlowState struct {
	value *SettingsFlowState
	isSet bool
}

func (v NullableSettingsFlowState) Get() *SettingsFlowState {
	return v.value
}

func (v *NullableSettingsFlowState) Set(val *SettingsFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsFlowState(val *SettingsFlowState) *NullableSettingsFlowState {
	return &NullableSettingsFlowState{value: val, isSet: true}
}

func (v NullableSettingsFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
