---
layout: page
title: Asm
---

Work in Progress
------------------

# These are just notes

# Root

Each tree has a root so katydid asm expects you to define a root type.

root = package.message

for example

root = main.Hello

root is a keyword

# Automaton

In katydid asm, automaton are defined just as described in the bottom up hedge automaton document

current input = next

# Start States

Each message is a parent node and requires a start state to an automaton.

package.message = start

for example in [Hello World](http://arborist.katydid.ws):

main.Hello = start

or

main.Hello = a

start is not a keyword

# Leaves

Input symbols for leaves are defined using an if statement.

for example

if contains(decString(main.Hello.World), "World") then world else noworld

The leave name is contained as a function parameter see main.Hello.World

or rather package.message.field

All fields are of type []byte, see Types and Builtin Functions Documents

After decoding the string with the decString function, it becomes type string

This can now be used in the contains function to check whether the decoded string contains the substring "World".

If the contains function returns true the input symbol for the field will be "world" otherwise it will be "noworld".

Nested ifs are also possible see Example [Index in Address List of Mover](http://arborist.katydid.ws/example/ListIndexAddressMover)

