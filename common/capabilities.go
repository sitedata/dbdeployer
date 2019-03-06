// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2019 Giuseppe Maxia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package common

import "github.com/datacharmer/dbdeployer/globals"

type MinimumVersion []int
type Capability struct {
	Description string         `json:"description"`
	Since       MinimumVersion `json:"since"`
	Until       MinimumVersion `json:"until"`
}
type FeatureList map[string]Capability

type Capabilities struct {
	Flavor      string      `json:"flavor"`
	Description string      `json:"description"`
	Features    FeatureList `json:"features"`
}

type ElementPath struct {
	dir      string
	fileName string
}
type FlavorIndicator struct {
	elements  []ElementPath
	flavor    string
	AllNeeded bool
}

const (
	// Tarball flavors
	MySQLFlavor         = "mysql"
	PerconaServerFlavor = "percona"
	MariaDbFlavor       = "mariadb"
	NDBFlavor           = "ndb"
	PxcFlavor           = "pxc"
	TiDbFlavor          = "tidb"

	// Feature names
	InstallDb        = "installdb"
	DynVariables     = "dynVars"
	SemiSynch        = "semiSync"
	CrashSafe        = "crashSafe"
	GTID             = "GTID"
	EnhancedGTID     = "enhancedGTID"
	Initialize       = "initialize"
	CreateUser       = "createUser"
	SuperReadOnly    = "superReadOnly"
	MySQLX           = "mysqlx"
	MySQLXDefault    = "mysqlxDefault"
	MultiSource      = "multiSource"
	GroupReplication = "groupReplication"
	SetPersist       = "setPersist"
	Roles            = "roles"
	NativeAuth       = "nativeAuth"
	DataDict         = "datadict"
	XtradbCluster    = "xtradbCluster"
	RootAuth         = "rootAuth"
)

var MySQLCapabilities = Capabilities{
	Flavor:      MySQLFlavor,
	Description: "MySQL server",
	Features: FeatureList{
		InstallDb: {
			Description: "uses mysql_install_db",
			Since:       globals.MinimumMySQLInstallDb,
			Until:       globals.MaximumMySQLInstallDb,
		},
		DynVariables: {
			Description: "dynamic variables",
			Since:       globals.MinimumDynVariablesVersion,
		},
		SemiSynch: {
			Description: "semi-synchronous replication",
			Since:       globals.MinimumSemiSyncVersion,
		},
		CrashSafe: {
			Description: "crash-safe replication",
			Since:       globals.MinimumCrashSafeVersion,
		},
		GTID: {
			Description: "Global transaction identifiers",
			Since:       globals.MinimumGtidVersion,
		},
		EnhancedGTID: {
			Description: "Enhanced Global transaction identifiers",
			Since:       globals.MinimumEnhancedGtidVersion,
		},
		Initialize: {
			Description: "mysqld --initialize as default",
			Since:       globals.MinimumDefaultInitializeVersion,
		},
		CreateUser: {
			Description: "Create user mandatory",
			Since:       globals.MinimumCreateUserVersion,
		},
		SuperReadOnly: {
			Description: "super-read-only support",
			Since:       globals.MinimumSuperReadOnly,
		},
		MySQLX: {
			Description: "MySQLX supported",
			Since:       globals.MinimumMysqlxVersion,
		},
		MySQLXDefault: {
			Description: "MySQLX enabled by default",
			Since:       globals.MinimumMysqlxDefaultVersion,
		},
		MultiSource: {
			Description: "multi-source replication",
			Since:       globals.MinimumMultiSourceReplVersion,
		},
		GroupReplication: {
			Description: "group replication",
			Since:       globals.MinimumGroupReplVersion,
		},
		SetPersist: {
			Description: "Set persist supported",
			Since:       globals.MinimumPersistVersion,
		},
		Roles: {
			Description: "Roles supported",
			Since:       globals.MinimumRolesVersion,
		},
		NativeAuth: {
			Description: "Native Authentication plugin",
			Since:       globals.MinimumNativeAuthPluginVersion,
		},
		DataDict: {
			Description: "data dictionary",
			Since:       globals.MinimumDataDictionaryVersion,
		},
	},
}

// Flavor indicators must be listed from the most complex ones to the
// simplest ones, because we want to catch the flavors that require
// multiple elements to be identified. If we put the simpler ones on top,
// we would miss the complex ones.
var FlavorCompositionList = []FlavorIndicator{
	{
		AllNeeded: true,
		elements: []ElementPath{
			{"bin", "garbd"},
			{"lib", "libgalera_smm.so"},
			{"lib", "libperconaserverclient.so"},
		},
		flavor: PxcFlavor,
	},
	{
		AllNeeded: true,
		elements: []ElementPath{
			{"bin", "garbd"},
			{"lib", "libgalera_smm.a"},
			{"lib", "libperconaserverclient.a"},
		},
		flavor: PxcFlavor,
	},
	{
		AllNeeded: true,
		elements: []ElementPath{
			{"bin", "garbd"},
			{"lib", "libgalera_smm.dylib"},
			{"lib", "libperconaserverclient.dylib"},
		},
		flavor: PxcFlavor,
	},
	//{
	//	AllNeeded: true,
	//	elements: []ElementPath{
	//		{"bin", "ndbd"},
	//		{"bin", "ndb_mgm"},
	//		{"bin", "ndb_mgmd"},
	//	},
	//	flavor: NDBFlavor,
	//},
	{
		AllNeeded: false,
		elements: []ElementPath{
			{"bin", "aria_chk"},
			{"lib", "libmariadbclient.a"},
			{"lib", "libmariadbclient.dylib"},
			{"lib", "libmariadb.a"},
			{"lib", "libmariadb.dylib"},
		},
		flavor: MariaDbFlavor,
	},

	{
		AllNeeded: false,
		elements: []ElementPath{
			{"lib", "libperconaserverclient.a"},
			{"lib", "libperconaserverclient.so"},
			{"lib", "libperconaserverclient.dylib"},
		},
		flavor: PerconaServerFlavor,
	},

	{
		AllNeeded: false,
		elements: []ElementPath{
			{"bin", "tidb-server"},
		},
		flavor: TiDbFlavor,
	},
	{
		AllNeeded: false,
		elements: []ElementPath{
			{"bin", "mysqld"},
			{"bin", "mysqld-debug"},
			{"lib", "libmysqlclient.a"},
		},
		flavor: MySQLFlavor,
	},
}

var PerconaCapabilities = Capabilities{
	Flavor:      PerconaServerFlavor,
	Description: "Percona Server",
	Features:    MySQLCapabilities.Features,
}

var TiDBCapabilities = Capabilities{
	Flavor:      TiDbFlavor,
	Description: "TiDB isolated server",
	Features:    FeatureList{
		// No capabilities so far
	},
}
var NDBCapabilities = Capabilities{
	Flavor:      NDBFlavor,
	Description: "MySQL NDB Cluster",
	Features:    FeatureList{
		// No capabilities so far
	},
}

var PxcCapabilities = Capabilities{
	Flavor:      PxcFlavor,
	Description: "Percona XtraDB Cluster",
	Features: addCapabilities(PerconaCapabilities.Features,
		FeatureList{
			XtradbCluster: {
				Description: "Xtradb Cluster creation",
				Since:       globals.MinimumXtradbClusterVersion,
			},
		}),
}

// NOTE: We only list the capabilities
// for which dbdeployer needs to take action
var MariadbCapabilities = Capabilities{
	Flavor: MariaDbFlavor,
	Features: FeatureList{
		InstallDb: {
			Description: "uses mysql_install_db",
			Since:       globals.MinimumMySQLInstallDb,
			Until:       nil,
		},
		RootAuth: {
			Description: "Root Authentication during install",
			Since:       globals.MinimumRootAuthVersion,
		},
		DynVariables: MySQLCapabilities.Features[DynVariables],
		SemiSynch:    MySQLCapabilities.Features[SemiSynch],
	},
}

var AllCapabilities = map[string]Capabilities{
	MySQLFlavor:         MySQLCapabilities,
	PerconaServerFlavor: PerconaCapabilities,
	MariaDbFlavor:       MariadbCapabilities,
	TiDbFlavor:          TiDBCapabilities,
	NDBFlavor:           NDBCapabilities,
	PxcFlavor:           PxcCapabilities,
}

// Returns a set of existing capabilities with custom ones
// added (or replaced) to the list
func addCapabilities(flavorFeatures, features FeatureList) FeatureList {
	var fList = make(FeatureList)
	for fName, feature := range flavorFeatures {
		fList[fName] = feature
	}
	for fName, feature := range features {
		fList[fName] = feature
	}
	return fList
}

// Returns a subset of a flavor capabilities
func copyCapabilities(flavor string, names []string) FeatureList {
	var fList = make(FeatureList)
	_, flavorExists := AllCapabilities[flavor]
	if !flavorExists {
		return fList
	}
	for fName, feature := range AllCapabilities[flavor].Features {
		for _, n := range names {
			if fName == n {
				fList[n] = feature
			}
		}
	}
	return fList
}

// Returns true if a given flavor and version support the wanted feature
func HasCapability(flavor, feature, version string) (bool, error) {
	versionList, err := VersionToList(version)
	if err != nil {
		return false, err
	}
	for flavorName, capabilities := range AllCapabilities {
		if flavorName == flavor {
			featureDefinition, ok := capabilities.Features[feature]
			if ok {
				overMinimum, err := GreaterOrEqualVersionList(versionList, featureDefinition.Since)
				if err != nil {
					return false, err
				}
				withinMaximum := true
				if featureDefinition.Until != nil {
					withinMaximum, err = GreaterOrEqualVersionList(featureDefinition.Until, versionList)
					if err != nil {
						return false, err
					}
				}
				return overMinimum && withinMaximum, nil
			}
		}
	}
	return false, nil
}