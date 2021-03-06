package features

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ChangeType enumerates the values for change type.
type ChangeType string

const (
	// Create The resource does not exist in the current state but is present in the desired state. The
	// resource will be created when the deployment is executed.
	Create ChangeType = "Create"
	// Delete The resource exists in the current state and is missing from the desired state. The resource will
	// be deleted when the deployment is executed.
	Delete ChangeType = "Delete"
	// Deploy The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource may or may not change.
	Deploy ChangeType = "Deploy"
	// Ignore The resource exists in the current state and is missing from the desired state. The resource will
	// not be deployed or modified when the deployment is executed.
	Ignore ChangeType = "Ignore"
	// Modify The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource will change.
	Modify ChangeType = "Modify"
	// NoChange The resource exists in the current state and the desired state and will be redeployed when the
	// deployment is executed. The properties of the resource will not change.
	NoChange ChangeType = "NoChange"
)

// PossibleChangeTypeValues returns an array of possible values for the ChangeType const type.
func PossibleChangeTypeValues() []ChangeType {
	return []ChangeType{Create, Delete, Deploy, Ignore, Modify, NoChange}
}

// DeploymentMode enumerates the values for deployment mode.
type DeploymentMode string

const (
	// Complete ...
	Complete DeploymentMode = "Complete"
	// Incremental ...
	Incremental DeploymentMode = "Incremental"
)

// PossibleDeploymentModeValues returns an array of possible values for the DeploymentMode const type.
func PossibleDeploymentModeValues() []DeploymentMode {
	return []DeploymentMode{Complete, Incremental}
}

// OnErrorDeploymentType enumerates the values for on error deployment type.
type OnErrorDeploymentType string

const (
	// LastSuccessful ...
	LastSuccessful OnErrorDeploymentType = "LastSuccessful"
	// SpecificDeployment ...
	SpecificDeployment OnErrorDeploymentType = "SpecificDeployment"
)

// PossibleOnErrorDeploymentTypeValues returns an array of possible values for the OnErrorDeploymentType const type.
func PossibleOnErrorDeploymentTypeValues() []OnErrorDeploymentType {
	return []OnErrorDeploymentType{LastSuccessful, SpecificDeployment}
}

// PropertyChangeType enumerates the values for property change type.
type PropertyChangeType string

const (
	// PropertyChangeTypeArray The property is an array and contains nested changes.
	PropertyChangeTypeArray PropertyChangeType = "Array"
	// PropertyChangeTypeCreate The property does not exist in the current state but is present in the desired
	// state. The property will be created when the deployment is executed.
	PropertyChangeTypeCreate PropertyChangeType = "Create"
	// PropertyChangeTypeDelete The property exists in the current state and is missing from the desired state.
	// It will be deleted when the deployment is executed.
	PropertyChangeTypeDelete PropertyChangeType = "Delete"
	// PropertyChangeTypeModify The property exists in both current and desired state and is different. The
	// value of the property will change when the deployment is executed.
	PropertyChangeTypeModify PropertyChangeType = "Modify"
)

// PossiblePropertyChangeTypeValues returns an array of possible values for the PropertyChangeType const type.
func PossiblePropertyChangeTypeValues() []PropertyChangeType {
	return []PropertyChangeType{PropertyChangeTypeArray, PropertyChangeTypeCreate, PropertyChangeTypeDelete, PropertyChangeTypeModify}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
	// SystemAssignedUserAssigned ...
	SystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// UserAssigned ...
	UserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned, SystemAssignedUserAssigned, UserAssigned}
}

// WhatIfResultFormat enumerates the values for what if result format.
type WhatIfResultFormat string

const (
	// FullResourcePayloads ...
	FullResourcePayloads WhatIfResultFormat = "FullResourcePayloads"
	// ResourceIDOnly ...
	ResourceIDOnly WhatIfResultFormat = "ResourceIdOnly"
)

// PossibleWhatIfResultFormatValues returns an array of possible values for the WhatIfResultFormat const type.
func PossibleWhatIfResultFormatValues() []WhatIfResultFormat {
	return []WhatIfResultFormat{FullResourcePayloads, ResourceIDOnly}
}
