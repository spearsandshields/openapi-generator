/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore
// EnumTest struct for EnumTest
type EnumTest struct {
	EnumString string `json:"enum_string,omitempty" xml:"enum_string"`
	EnumStringRequired string `json:"enum_string_required" xml:"enum_string_required"`
	EnumInteger int32 `json:"enum_integer,omitempty" xml:"enum_integer"`
	EnumNumber float64 `json:"enum_number,omitempty" xml:"enum_number"`
	OuterEnum OuterEnum `json:"outerEnum,omitempty" xml:"outerEnum"`
}