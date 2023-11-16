// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ServersDatabasesAuditingSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabasesAuditingSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForServersDatabasesAuditingSetting, ServersDatabasesAuditingSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForServersDatabasesAuditingSetting tests if a specific instance of ServersDatabasesAuditingSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForServersDatabasesAuditingSetting(subject ServersDatabasesAuditingSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.ServersDatabasesAuditingSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ServersDatabasesAuditingSetting
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersDatabasesAuditingSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ServersDatabasesAuditingSetting to ServersDatabasesAuditingSetting via AssignProperties_To_ServersDatabasesAuditingSetting & AssignProperties_From_ServersDatabasesAuditingSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForServersDatabasesAuditingSetting, ServersDatabasesAuditingSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServersDatabasesAuditingSetting tests if a specific instance of ServersDatabasesAuditingSetting can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServersDatabasesAuditingSetting(subject ServersDatabasesAuditingSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ServersDatabasesAuditingSetting
	err := copied.AssignProperties_To_ServersDatabasesAuditingSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ServersDatabasesAuditingSetting
	err = actual.AssignProperties_From_ServersDatabasesAuditingSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ServersDatabasesAuditingSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesAuditingSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesAuditingSetting, ServersDatabasesAuditingSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesAuditingSetting runs a test to see if a specific instance of ServersDatabasesAuditingSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesAuditingSetting(subject ServersDatabasesAuditingSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesAuditingSetting
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ServersDatabasesAuditingSetting instances for property testing - lazily instantiated by
// ServersDatabasesAuditingSettingGenerator()
var serversDatabasesAuditingSettingGenerator gopter.Gen

// ServersDatabasesAuditingSettingGenerator returns a generator of ServersDatabasesAuditingSetting instances for property testing.
func ServersDatabasesAuditingSettingGenerator() gopter.Gen {
	if serversDatabasesAuditingSettingGenerator != nil {
		return serversDatabasesAuditingSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting(generators)
	serversDatabasesAuditingSettingGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesAuditingSetting{}), generators)

	return serversDatabasesAuditingSettingGenerator
}

// AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesAuditingSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = Servers_Databases_AuditingSetting_SpecGenerator()
	gens["Status"] = Servers_Databases_AuditingSetting_STATUSGenerator()
}

func Test_Servers_Databases_AuditingSetting_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Databases_AuditingSetting_Spec to Servers_Databases_AuditingSetting_Spec via AssignProperties_To_Servers_Databases_AuditingSetting_Spec & AssignProperties_From_Servers_Databases_AuditingSetting_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Databases_AuditingSetting_Spec, Servers_Databases_AuditingSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Databases_AuditingSetting_Spec tests if a specific instance of Servers_Databases_AuditingSetting_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServers_Databases_AuditingSetting_Spec(subject Servers_Databases_AuditingSetting_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Databases_AuditingSetting_Spec
	err := copied.AssignProperties_To_Servers_Databases_AuditingSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Databases_AuditingSetting_Spec
	err = actual.AssignProperties_From_Servers_Databases_AuditingSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_Databases_AuditingSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_AuditingSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_AuditingSetting_Spec, Servers_Databases_AuditingSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_AuditingSetting_Spec runs a test to see if a specific instance of Servers_Databases_AuditingSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_AuditingSetting_Spec(subject Servers_Databases_AuditingSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_AuditingSetting_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Servers_Databases_AuditingSetting_Spec instances for property testing - lazily instantiated by
// Servers_Databases_AuditingSetting_SpecGenerator()
var servers_Databases_AuditingSetting_SpecGenerator gopter.Gen

// Servers_Databases_AuditingSetting_SpecGenerator returns a generator of Servers_Databases_AuditingSetting_Spec instances for property testing.
func Servers_Databases_AuditingSetting_SpecGenerator() gopter.Gen {
	if servers_Databases_AuditingSetting_SpecGenerator != nil {
		return servers_Databases_AuditingSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_AuditingSetting_Spec(generators)
	servers_Databases_AuditingSetting_SpecGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_AuditingSetting_Spec{}), generators)

	return servers_Databases_AuditingSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_AuditingSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_AuditingSetting_Spec(gens map[string]gopter.Gen) {
	gens["AuditActionsAndGroups"] = gen.SliceOf(gen.AlphaString())
	gens["IsAzureMonitorTargetEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsManagedIdentityInUse"] = gen.PtrOf(gen.Bool())
	gens["IsStorageSecondaryKeyInUse"] = gen.PtrOf(gen.Bool())
	gens["QueueDelayMs"] = gen.PtrOf(gen.Int())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(DatabaseBlobAuditingPolicyProperties_State_Disabled, DatabaseBlobAuditingPolicyProperties_State_Enabled))
	gens["StorageAccountSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

func Test_Servers_Databases_AuditingSetting_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Servers_Databases_AuditingSetting_STATUS to Servers_Databases_AuditingSetting_STATUS via AssignProperties_To_Servers_Databases_AuditingSetting_STATUS & AssignProperties_From_Servers_Databases_AuditingSetting_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForServers_Databases_AuditingSetting_STATUS, Servers_Databases_AuditingSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForServers_Databases_AuditingSetting_STATUS tests if a specific instance of Servers_Databases_AuditingSetting_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForServers_Databases_AuditingSetting_STATUS(subject Servers_Databases_AuditingSetting_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.Servers_Databases_AuditingSetting_STATUS
	err := copied.AssignProperties_To_Servers_Databases_AuditingSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Servers_Databases_AuditingSetting_STATUS
	err = actual.AssignProperties_From_Servers_Databases_AuditingSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Servers_Databases_AuditingSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Servers_Databases_AuditingSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServers_Databases_AuditingSetting_STATUS, Servers_Databases_AuditingSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServers_Databases_AuditingSetting_STATUS runs a test to see if a specific instance of Servers_Databases_AuditingSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServers_Databases_AuditingSetting_STATUS(subject Servers_Databases_AuditingSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Servers_Databases_AuditingSetting_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Servers_Databases_AuditingSetting_STATUS instances for property testing - lazily instantiated by
// Servers_Databases_AuditingSetting_STATUSGenerator()
var servers_Databases_AuditingSetting_STATUSGenerator gopter.Gen

// Servers_Databases_AuditingSetting_STATUSGenerator returns a generator of Servers_Databases_AuditingSetting_STATUS instances for property testing.
func Servers_Databases_AuditingSetting_STATUSGenerator() gopter.Gen {
	if servers_Databases_AuditingSetting_STATUSGenerator != nil {
		return servers_Databases_AuditingSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServers_Databases_AuditingSetting_STATUS(generators)
	servers_Databases_AuditingSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(Servers_Databases_AuditingSetting_STATUS{}), generators)

	return servers_Databases_AuditingSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServers_Databases_AuditingSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServers_Databases_AuditingSetting_STATUS(gens map[string]gopter.Gen) {
	gens["AuditActionsAndGroups"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAzureMonitorTargetEnabled"] = gen.PtrOf(gen.Bool())
	gens["IsManagedIdentityInUse"] = gen.PtrOf(gen.Bool())
	gens["IsStorageSecondaryKeyInUse"] = gen.PtrOf(gen.Bool())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["QueueDelayMs"] = gen.PtrOf(gen.Int())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(DatabaseBlobAuditingPolicyProperties_State_STATUS_Disabled, DatabaseBlobAuditingPolicyProperties_State_STATUS_Enabled))
	gens["StorageAccountSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
