package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// BlueprintApi is a string that contains a Blueprint API version identifier.
type BlueprintApi string

const (
	// V1 is the classic version 1 API identifier of Cloudogu EcoSystem blueprint mechanism inside VMs.
	V1 BlueprintApi = "v1"
	// V2 is the API version of the BlueprintV2 json format used in the MultiNode-CES inside kubernetes, e.g. for Blueprint-CRs.
	V2 BlueprintApi = "v2"
	// TestEmpty is a non-production, test-only API identifier of Cloudogu EcoSystem blueprint mechanism.
	TestEmpty BlueprintApi = "test/empty"
)

type BlueprintMaskApi string

const (
	// MaskV1 is the classic version 1 API identifier of Cloudogu EcoSystem blueprint mask mechanism which hide or show
	// certain parts from a blueprint.
	MaskV1 BlueprintMaskApi = "v1"
)

// GeneralBlueprint defines the minimum set to parse the blueprint API version string in order to select the right
// blueprint handling strategy. This is necessary in order to accommodate maximal changes in different blueprint API
// versions.
type GeneralBlueprint struct {
	// API is used to distinguish between different versions of the used API and impacts directly the interpretation of
	// this blueprint. Must not be empty.
	//
	// This field MUST NOT be MODIFIED or REMOVED because the API is paramount for distinguishing between different
	// blueprint version implementations.
	API BlueprintApi `json:"blueprintApi"`
}

// GeneralBlueprintMask defines the minimum set to parse the blueprint mask API version string in order to select the
// right blueprint mask handling strategy. This is necessary in order to accommodate maximal changes in different
// blueprint mask API versions.
type GeneralBlueprintMask struct {
	// API is used to distinguish between different versions of the used API and impacts directly the interpretation of
	// this blueprint mask. Must not be empty.
	//
	// This field MUST NOT be MODIFIED or REMOVED because the API is paramount for distinguishing between different
	// blueprint mask version implementations.
	API BlueprintMaskApi `json:"blueprintMaskApi"`
}

// TargetState defines an enum of values that determines a state of installation.
type TargetState int

const (
	// TargetStatePresent is the default state. If selected the chosen item must be present after the blueprint was
	// applied.
	TargetStatePresent TargetState = iota
	// TargetStateAbsent sets the state of the item to absent. If selected the chosen item must be absent after the
	// blueprint was applied.
	TargetStateAbsent
	// TargetStateIgnore is currently only internally used to mark items that are present in the CES instance at hand
	// but not mentioned in the blueprint.
	TargetStateIgnore
)

// String returns a string representation of the given TargetState enum value.
func (state TargetState) String() string {
	return toString[state]
}

var toString = map[TargetState]string{
	TargetStatePresent: "present",
	TargetStateAbsent:  "absent",
}

var toID = map[string]TargetState{
	"present": TargetStatePresent,
	"absent":  TargetStateAbsent,
}

// MarshalJSON marshals the enum as a quoted json string
func (state TargetState) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[state])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmarshals a quoted json string to the enum value. Use it with usual json unmarshalling:
//
//	 jsonBlob := []byte("\"present\"")
//		var state TargetState
//		err := json.Unmarshal(jsonBlob, &state)
func (state *TargetState) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return fmt.Errorf("cannot unmarshal value %s to a TargetState: %w", string(b), err)
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*state = toID[j]
	return nil
}
