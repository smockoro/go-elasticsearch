/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package action

import "github.com/elastic/go-elasticsearch/generator/api"

// Var is a variable instantiated by an action.
type Var struct {
	Name    string
	Package string
	Type    string
	Pointer bool
}

var errorVar = &Var{
	Name: "err",
	Type: "error",
}

var valueVar = &Var{
	Name: "v",
	Type: "interface{}",
}

var intVar = &Var{
	Name: "i",
	Type: "int",
}

var boolVar = &Var{
	Name: "b",
	Type: "bool",
}

var bodyVar = &Var{
	Name:    "body",
	Package: "util",
	Type:    "MapStr",
}

func newResponseVar(m *api.Method) *Var {
	return &Var{
		Name:    m.ResponseName,
		Package: m.PackageName,
		Type:    m.Name + "Response",
		Pointer: true,
	}
}

// Context is the context of an action, it includes the methods invoked and the variables the action instantiates.
type Context struct {
	Methods []*api.Method
	Vars    []*Var
}
