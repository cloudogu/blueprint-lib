package json

type BlueprintMaskApi string

const (
	// BlueprintMaskAPIV1 contains the API version number of the Blueprint Mask mechanism.
	BlueprintMaskAPIV1 BlueprintMaskApi = "v1"
)

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

// BlueprintMaskV1 describes an abstraction of CES components that should alter a blueprint definition before
// applying it to a CES system via a blueprint upgrade. The blueprint mask should not change the blueprint JSON file
// itself, but is applied to the information in it to generate a new, effective blueprint.
//
// In general additions without changing the version are fine, as long as they don't change semantics. Removal or
// renaming are breaking changes and require a new blueprint mask API version.
type BlueprintMaskV1 struct {
	GeneralBlueprintMask
	// ID is the unique name of the set over all components. This blueprint mask ID should be used to distinguish
	// from similar blueprint masks between humans in an easy way. Must not be empty.
	ID string `json:"blueprintMaskId"`
	// Dogus contains a set of dogus which alters the states of the dogus in the blueprint this mask is applied on.
	// The names and target states of all dogus must not be empty.
	Dogus []TargetDogu `json:"dogus"`
}
