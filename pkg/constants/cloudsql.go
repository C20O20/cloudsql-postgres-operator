/*
Copyright 2019 The cloudsql-postgres-operator Authors.

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

package constants

const (
	// DatabaseInstanceActivationPolicyAlways is the activation policy of a running, healthy CSQLP instance.
	DatabaseInstanceActivationPolicyAlways = "ALWAYS"
	// DatabaseInstanceIPAddressTypePublic is the type associated with a CSQLP instance's public IP.
	DatabaseInstanceIPAddressTypePublic = "PRIMARY"
	// DatabaseInstanceIPAddressTypePrivate is the type associated with a CSQLP instance's private IP.
	DatabaseInstanceIPAddressTypePrivate = "PRIVATE"
	// DatabaseInstanceStateRunnable is the state of a running, healthy CSQLP instance.
	DatabaseInstanceStateRunnable = "RUNNABLE"
	// OperationStatusDone is the status of an operation that has terminated.
	OperationStatusDone = "DONE"
)
