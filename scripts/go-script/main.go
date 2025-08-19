// fern-sdk-sandbox SDK Sandbox
//
// This is a test sandbox for the locally generated fern-sdk-sandbox SDK.
// Modify this file to test different SDK functionality and reproduce issues.

package main

import (
	"fmt"

	sdk "github.com/jsklan/fern-sdk-sandbox"
	"github.com/jsklan/fern-sdk-sandbox/client"
	"github.com/jsklan/fern-sdk-sandbox/option"
)

var (
	host = "localhost:3000"
)

func main() {
	fmt.Println("Starting fern-sdk-sandbox SDK sandbox...")

	sdkClient := client.NewClient(
		option.WithBaseURL(fmt.Sprintf("https://%s", host)),
	)
	fmt.Printf("SDK imported and client initialized successfully! Client: %T\n", sdkClient)

	st := sdk.String("myString")
	fmt.Printf("st: %+v\n", st)

	req := sdk.OptionalExamplesPostRequestBody{
		Primitive:              "primitiveValue",
		OptionalPrimitive:      sdk.String("optionalPrimitiveValue"),
		NamedPrimitive:         "NamedPrimitiveValue",
		NamedOptionalPrimitive: sdk.String("OptionalNamedPrimitiveValue"),
		Foo: &sdk.Foo{
			FooName:  "fooName",
			FooValue: "fooValue",
		},
		FooOptional: &sdk.Foo{
			FooName: "fooName",
		},
		FooNullable: &sdk.Foo{
			FooName: "fooName",
		},
	}
	fmt.Printf("req: %+v\n", req)

}
