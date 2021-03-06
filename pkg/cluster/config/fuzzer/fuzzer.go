/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fuzzer

import (
	fuzz "github.com/google/gofuzz"

	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"

	"sigs.k8s.io/kind/pkg/cluster/config"
)

// Funcs returns custom fuzzer functions for the `kind` Config.
func Funcs(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		fuzzConfig,
	}
}

func fuzzConfig(obj *config.Config, c fuzz.Continue) {
	c.FuzzNoCustom(obj)

	// Pinning values for fields that get defaults if fuzz value is empty string or nil (thus making the round trip test fail)
	obj.Image = "fuzzimage:latest"
}
