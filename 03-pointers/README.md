# Pointers

## Introduction

Pointers are variables that store the memory address instead of values. To retrieve the memory address of a variable, we use the `&` operator. To retrieve the value of a memory address, we use the `*` operator.

In Go pointers are used to pass the reference of a variable to a function, instead of passing the value. This is useful when we want to modify the value of a variable inside a function. Another use case is when we want to avoid copying large data structures.
