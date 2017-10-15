package nonewlines

import (
	"fmt"
	"testing"
)

const (
	funcWithNewline = `
package p

func f() {

	return
}
`

	funcWithNewlineComment = `
package p

func f() {

	// comment

	return // trailing comment
}
`
	funcWithTrailingNewline = `
package p

func f() {
	assert.True()


	// trailing comment
}
`

	funcWithTrailingNewlineAndMultilineStatement = `
package p

func f() Thing {

	return Thing {
		X: ok,
		Y: false,
	}

}
`

	multlineDecl = `
package p

func f(
	argC int,
	args []string) error {

	return nil
}
`
)

func TestNewlineRemoval(t *testing.T) {
	out, err := Process("test.go", []byte(funcWithNewline))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", string(out))
}

func TestNewlineRemovalComment(t *testing.T) {
	out, err := Process("test.go", []byte(funcWithNewlineComment))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", string(out))
}

func TestTrailingNewlineRemoval(t *testing.T) {
	out, err := Process("test.go", []byte(funcWithTrailingNewline))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", string(out))
}

func TestTrailingNewlineRemovalMultilineStatement(t *testing.T) {
	out, err := Process("test.go", []byte(funcWithTrailingNewlineAndMultilineStatement))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", string(out))
}

func TestMultilineFuncDecl(t *testing.T) {
	out, err := Process("test.go", []byte(multlineDecl))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v\n", string(out))
}
