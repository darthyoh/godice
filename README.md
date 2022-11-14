# godice
Simple implementation of Dice Coefficient with go.

## Introduction

The [Dice Coefficient](https://en.wikipedia.org/wiki/S%C3%B8rensen%E2%80%93Dice_coefficient) is used to gauge the similarity of two samples. This version uses it to return the similarity of two strings.

This implementation doesn't care with spaces or special characters, and is case unsensitive :

Foo <=> FOO <=> F O O <=> ###F'%o$£O

## Installation

`go get github.com/darthyoh/godice`

## Usage

As a simple form, you can use it to simply compare two strings :

```
package main

import (
    "fmt"
    "github.com/darthyoh/godice
)

func main() {
    string a := "foo"
    string b := "bar"

    fmt.Printf("Coef is %v", dice.Compare(&a, &b)) // returns 0.0

    a = "nicht"
    b = "nacht"

    fmt.Printf("Coef is %v", dice.Compare(&a, &b)) // returns 0.5
}
```

This package comes with a simple struct called `DiceString`. It's just a wrapper around a string value. The `NewDiceString()` function generate a DiceString. You can then use the `CompareDiceString()` method of a `DiceString` to get the similarity between the two strings. It can be useful for massive operations of comparing :

```
valueA := "nicht"
valueB := "nacht"

NewDiceString(&valueA).CompareDiceString(NewDiceString(&valueB)) // returns 0.5
```