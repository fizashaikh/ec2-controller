// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package capacity_reservation

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.AdditionalInfo, b.ko.Spec.AdditionalInfo) {
		delta.Add("Spec.AdditionalInfo", a.ko.Spec.AdditionalInfo, b.ko.Spec.AdditionalInfo)
	} else if a.ko.Spec.AdditionalInfo != nil && b.ko.Spec.AdditionalInfo != nil {
		if *a.ko.Spec.AdditionalInfo != *b.ko.Spec.AdditionalInfo {
			delta.Add("Spec.AdditionalInfo", a.ko.Spec.AdditionalInfo, b.ko.Spec.AdditionalInfo)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone) {
		delta.Add("Spec.AvailabilityZone", a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone)
	} else if a.ko.Spec.AvailabilityZone != nil && b.ko.Spec.AvailabilityZone != nil {
		if *a.ko.Spec.AvailabilityZone != *b.ko.Spec.AvailabilityZone {
			delta.Add("Spec.AvailabilityZone", a.ko.Spec.AvailabilityZone, b.ko.Spec.AvailabilityZone)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AvailabilityZoneID, b.ko.Spec.AvailabilityZoneID) {
		delta.Add("Spec.AvailabilityZoneID", a.ko.Spec.AvailabilityZoneID, b.ko.Spec.AvailabilityZoneID)
	} else if a.ko.Spec.AvailabilityZoneID != nil && b.ko.Spec.AvailabilityZoneID != nil {
		if *a.ko.Spec.AvailabilityZoneID != *b.ko.Spec.AvailabilityZoneID {
			delta.Add("Spec.AvailabilityZoneID", a.ko.Spec.AvailabilityZoneID, b.ko.Spec.AvailabilityZoneID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ClientToken, b.ko.Spec.ClientToken) {
		delta.Add("Spec.ClientToken", a.ko.Spec.ClientToken, b.ko.Spec.ClientToken)
	} else if a.ko.Spec.ClientToken != nil && b.ko.Spec.ClientToken != nil {
		if *a.ko.Spec.ClientToken != *b.ko.Spec.ClientToken {
			delta.Add("Spec.ClientToken", a.ko.Spec.ClientToken, b.ko.Spec.ClientToken)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DryRun, b.ko.Spec.DryRun) {
		delta.Add("Spec.DryRun", a.ko.Spec.DryRun, b.ko.Spec.DryRun)
	} else if a.ko.Spec.DryRun != nil && b.ko.Spec.DryRun != nil {
		if *a.ko.Spec.DryRun != *b.ko.Spec.DryRun {
			delta.Add("Spec.DryRun", a.ko.Spec.DryRun, b.ko.Spec.DryRun)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EBSOptimized, b.ko.Spec.EBSOptimized) {
		delta.Add("Spec.EBSOptimized", a.ko.Spec.EBSOptimized, b.ko.Spec.EBSOptimized)
	} else if a.ko.Spec.EBSOptimized != nil && b.ko.Spec.EBSOptimized != nil {
		if *a.ko.Spec.EBSOptimized != *b.ko.Spec.EBSOptimized {
			delta.Add("Spec.EBSOptimized", a.ko.Spec.EBSOptimized, b.ko.Spec.EBSOptimized)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EndDate, b.ko.Spec.EndDate) {
		delta.Add("Spec.EndDate", a.ko.Spec.EndDate, b.ko.Spec.EndDate)
	} else if a.ko.Spec.EndDate != nil && b.ko.Spec.EndDate != nil {
		if !a.ko.Spec.EndDate.Equal(b.ko.Spec.EndDate) {
			delta.Add("Spec.EndDate", a.ko.Spec.EndDate, b.ko.Spec.EndDate)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EndDateType, b.ko.Spec.EndDateType) {
		delta.Add("Spec.EndDateType", a.ko.Spec.EndDateType, b.ko.Spec.EndDateType)
	} else if a.ko.Spec.EndDateType != nil && b.ko.Spec.EndDateType != nil {
		if *a.ko.Spec.EndDateType != *b.ko.Spec.EndDateType {
			delta.Add("Spec.EndDateType", a.ko.Spec.EndDateType, b.ko.Spec.EndDateType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EphemeralStorage, b.ko.Spec.EphemeralStorage) {
		delta.Add("Spec.EphemeralStorage", a.ko.Spec.EphemeralStorage, b.ko.Spec.EphemeralStorage)
	} else if a.ko.Spec.EphemeralStorage != nil && b.ko.Spec.EphemeralStorage != nil {
		if *a.ko.Spec.EphemeralStorage != *b.ko.Spec.EphemeralStorage {
			delta.Add("Spec.EphemeralStorage", a.ko.Spec.EphemeralStorage, b.ko.Spec.EphemeralStorage)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.InstanceCount, b.ko.Spec.InstanceCount) {
		delta.Add("Spec.InstanceCount", a.ko.Spec.InstanceCount, b.ko.Spec.InstanceCount)
	} else if a.ko.Spec.InstanceCount != nil && b.ko.Spec.InstanceCount != nil {
		if *a.ko.Spec.InstanceCount != *b.ko.Spec.InstanceCount {
			delta.Add("Spec.InstanceCount", a.ko.Spec.InstanceCount, b.ko.Spec.InstanceCount)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.InstanceMatchCriteria, b.ko.Spec.InstanceMatchCriteria) {
		delta.Add("Spec.InstanceMatchCriteria", a.ko.Spec.InstanceMatchCriteria, b.ko.Spec.InstanceMatchCriteria)
	} else if a.ko.Spec.InstanceMatchCriteria != nil && b.ko.Spec.InstanceMatchCriteria != nil {
		if *a.ko.Spec.InstanceMatchCriteria != *b.ko.Spec.InstanceMatchCriteria {
			delta.Add("Spec.InstanceMatchCriteria", a.ko.Spec.InstanceMatchCriteria, b.ko.Spec.InstanceMatchCriteria)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.InstancePlatform, b.ko.Spec.InstancePlatform) {
		delta.Add("Spec.InstancePlatform", a.ko.Spec.InstancePlatform, b.ko.Spec.InstancePlatform)
	} else if a.ko.Spec.InstancePlatform != nil && b.ko.Spec.InstancePlatform != nil {
		if *a.ko.Spec.InstancePlatform != *b.ko.Spec.InstancePlatform {
			delta.Add("Spec.InstancePlatform", a.ko.Spec.InstancePlatform, b.ko.Spec.InstancePlatform)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.InstanceType, b.ko.Spec.InstanceType) {
		delta.Add("Spec.InstanceType", a.ko.Spec.InstanceType, b.ko.Spec.InstanceType)
	} else if a.ko.Spec.InstanceType != nil && b.ko.Spec.InstanceType != nil {
		if *a.ko.Spec.InstanceType != *b.ko.Spec.InstanceType {
			delta.Add("Spec.InstanceType", a.ko.Spec.InstanceType, b.ko.Spec.InstanceType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.OutpostARN, b.ko.Spec.OutpostARN) {
		delta.Add("Spec.OutpostARN", a.ko.Spec.OutpostARN, b.ko.Spec.OutpostARN)
	} else if a.ko.Spec.OutpostARN != nil && b.ko.Spec.OutpostARN != nil {
		if *a.ko.Spec.OutpostARN != *b.ko.Spec.OutpostARN {
			delta.Add("Spec.OutpostARN", a.ko.Spec.OutpostARN, b.ko.Spec.OutpostARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PlacementGroupARN, b.ko.Spec.PlacementGroupARN) {
		delta.Add("Spec.PlacementGroupARN", a.ko.Spec.PlacementGroupARN, b.ko.Spec.PlacementGroupARN)
	} else if a.ko.Spec.PlacementGroupARN != nil && b.ko.Spec.PlacementGroupARN != nil {
		if *a.ko.Spec.PlacementGroupARN != *b.ko.Spec.PlacementGroupARN {
			delta.Add("Spec.PlacementGroupARN", a.ko.Spec.PlacementGroupARN, b.ko.Spec.PlacementGroupARN)
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Tenancy, b.ko.Spec.Tenancy) {
		delta.Add("Spec.Tenancy", a.ko.Spec.Tenancy, b.ko.Spec.Tenancy)
	} else if a.ko.Spec.Tenancy != nil && b.ko.Spec.Tenancy != nil {
		if *a.ko.Spec.Tenancy != *b.ko.Spec.Tenancy {
			delta.Add("Spec.Tenancy", a.ko.Spec.Tenancy, b.ko.Spec.Tenancy)
		}
	}

	return delta
}
