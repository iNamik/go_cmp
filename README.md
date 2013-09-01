go_cmp
======

**Comparators in Go**


ABOUT
-----

Package `cmp` defines a set of comparator interfaces and functions, with default implementations for all standard comparable types.


Compare Results
---------------

comparing two values results in one of the following constants:

	const (
		cmp.LT int = -1 // Less-Than
		cmp.EQ     = 0  // Equal
		cmp.GT     = 1  // Greater-Than
	)


Type: `cmp.T`
-------------

Type `cmp.T` defines an interface that allows objects to export a comparator function:

	type T interface {
		Cmp(interface{}) int
	}


Type: `cmp.F`
-------------

Type `cmp.F` defines a stand-alone comparator function:

	type F func(a, b interface{}) int


Default Implementations
-----------------------

Default implementations are provided for all standard comparable types:

 * byte
 * float(32|64)
 * int & int(8|16|32|64)
 * rune
 * string
 * uint & uint(8|16|32|64)

Below are the definitions of the included `int` functions.  A complimentary set of functions exists for each of the standard comparable types (listed above):

	// cmp.Int is a cmp.T wrapper for basic type int.
	// Example:
	//   var comparable cmp.T = cmp.Int(3)
	type Int int

	// cmp.Int::Cmp fulfills the cmp.T interface.
	// b_ is expected to also be type cmp.Int.
	func (a Int) Cmp(b_ interface{}) int 

	// cmp.F_Int is a cmp.F for type cmp.Int
	// i.e. You can pass cmp.F_Int to functions
	//      requiring a cmp.F comparator
	// a_ and b_ are expected to be type cmp.Int.
	func F_Int(a_, b_ interface{}) int 

	// F_int is a cmp.F for basic type int
	// i.e. You can pass cmp.F_int to functions
	//      requiring a cmp.F comparator
	// a_ and b_ are expected to be basic type int.
	func F_int(a_, b_ interface{}) int 

	// cmp.Compare_int is a convenience function 
	// for explicitly comparing basic type int.
	func Compare_int(a, b int) int 


Extras
------

The package also contains constructs for treating functions and closures as `cmp.T` interfaces.  See the file "cmp.go" for details.


License
-------

This package is released under the MIT license.

See the file "LICENSE" for more details.


Contributors
------------

David Farrell <DavidPFarrell@yahoo.com>
