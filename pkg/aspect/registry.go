// Copyright 2017 Google Inc.
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

package aspect

import (
	"istio.io/mixer/pkg/config"
)

// ManagerFinder find an aspect Manager given aspect kind.
type ManagerFinder interface {
	// ByImpl queries the registry by adapter name.
	FindManager(kind string) (Manager, bool)
}

// ProcessBindings returns a fully constructed aspect registry given APIBinding.
func ProcessBindings(bnds []APIBinding) (map[string]Manager, map[config.APIMethod]config.AspectSet) {
	r := make(map[string]Manager)
	as := make(map[config.APIMethod]config.AspectSet)

	// setup aspect sets for all methods
	for _, am := range config.APIMethods() {
		as[am] = config.AspectSet{}
	}
	for _, bnd := range bnds {
		r[bnd.aspect.Kind()] = bnd.aspect
		as[bnd.method][bnd.aspect.Kind()] = true
	}
	return r, as

}

// APIBinding associates an aspect with an API method
type APIBinding struct {
	aspect Manager
	method config.APIMethod
}

// Inventory returns a manager inventory that contains
// all available aspect managers
func Inventory() []APIBinding {
	// Update the following list to add a new Aspect manager
	return []APIBinding{
		{NewDenyCheckerManager(), config.CheckMethod},
		{NewListCheckerManager(), config.CheckMethod},
		{NewLoggerManager(), config.ReportMethod},
	}
}
