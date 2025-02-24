// Example for https://github.com/seaube/tygo/issues/65
package genericany

type AnyStructField[T any] struct {
	Value     T
	SomeField string
}

type JsonArray[T any] []T
