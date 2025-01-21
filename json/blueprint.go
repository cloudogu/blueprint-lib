package json

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
