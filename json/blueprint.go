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
	// TestEmpty is a non-production, test-only API identifier of Cloudogu EcoSystem blueprint mechanism.
	TestEmpty BlueprintApi = "test/empty"
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

// BlueprintV1 describes an abstraction of Cloudogu EcoSystem (CES) parts that should be absent or present within one or
// more CES instances. When the same Blueprint is applied to two different CES instances it is required to leave two
// equal instances in terms of the components.
//
// In general additions without changing the version are fine, as long as they don't change semantics. Removal or
// renaming are breaking changes and require a new blueprint API version.
type BlueprintV1 struct {
	GeneralBlueprint
	// ID is the unique name of the set over all parts. This blueprint ID should be used to distinguish from similar
	// blueprints between humans in an easy way. Must not be empty.
	ID string `json:"blueprintId"`
	// CesAppVersion defines the exact version of the cesapp that should be present in the CES instance after which this
	// blueprint was applied. Must not be empty.
	//
	// This field MUST NOT be MODIFIED or REMOVED because the cesapp is paramount for interpreting blueprint
	// implementations.
	CesAppVersion string `json:"cesappVersion"`
	// Dogus contains a set of exact dogu versions which should be present or absent in the CES instance after which this
	// blueprint was applied. Optional.
	Dogus []TargetDogu `json:"dogus,omitempty"`
	// Packages contains a set of exact package versions which should be present or absent in the CES instance after which
	// this blueprint was applied. The packages must correspond to the used operating system package manager. Optional.
	Packages []TargetPackage `json:"packages,omitempty"`
	// Used to configure registry globalRegistryEntries on blueprint upgrades
	RegistryConfig RegistryConfig `json:"registryConfig,omitempty"`
	// Used to remove registry globalRegistryEntries on blueprint upgrades
	RegistryConfigAbsent []string `json:"registryConfigAbsent,omitempty"`
	// Used to configure encrypted registry globalRegistryEntries on blueprint upgrades
	RegistryConfigEncrypted RegistryConfig `json:"registryConfigEncrypted,omitempty"`
}

type RegistryConfig map[string]map[string]interface{}

// TargetDogu defines a Dogu, its version, and the installation state in which it is supposed to be after a blueprint
// was applied.
type TargetDogu struct {
	// Name defines the name of the dogu including its namespace, f. i. "official/nginx". Must not be empty.
	Name string `json:"name"`
	// Version defines the version of the dogu that is to be installed. Must not be empty if the targetState is "present";
	// otherwise it is optional and is not going to be interpreted.
	Version string `json:"version"`
	// TargetState defines a state of installation of this dogu. Optional field, but defaults to "TargetStatePresent"
	TargetState TargetState `json:"targetState"`
}

// TargetPackage an operating system package, its version, and the installation state in which it is supposed to be
// after a blueprint was applied.
type TargetPackage struct {
	// Name defines the name of the package. Must not be empty.
	Name string `json:"name"`
	// Version defines the version of the package that is to be installed. Must not be empty if the targetState is
	// "present"; otherwise it is optional and is not going to be interpreted.
	Version string `json:"version"`
	// TargetState defines a state of installation of this package. Optional field, but defaults to "TargetStatePresent"
	TargetState TargetState `json:"targetState"`
}
