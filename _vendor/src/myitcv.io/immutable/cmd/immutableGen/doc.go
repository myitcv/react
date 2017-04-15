/*

immutableGen is a go generate generator that creates immutable struct, map and
slice type declarations from template type declarations.

All such generated types "implement" a common immutable "interface" as well as
providing functions and methods specific to the immutable data structure
(struct, map or slice).

For more information see https://myitcv.io/immutable/wiki/immutableGen

*/
package main
