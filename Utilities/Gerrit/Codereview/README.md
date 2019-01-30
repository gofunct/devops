- General
  [ ] The code works
  [ ] The code is easy to understand
  [ ] Follows coding conventions
  [ ] Names are simple and if possible short
  [ ] Names are spelt correctly
  [ ] Names contain units where applicable
  [ ] Enums are used instead of int constants where applicable
  [ ] There are no usages of 'magic numbers'
  [ ] All variables are in the smallest scope possible
  [ ] All class, variable, and method modifiers are correct.
  [ ] There is no commented out code
  [ ] There is no dead code (inaccessible at Runtime)
  [ ] No code can be replaced with library functions
  [ ] Required logs are present
  [ ] Frivolous logs are absent
  [ ] Debugging code is absent
  [ ] No System.out.println or similar calls exist
  [ ] No stack traces are printed
  [ ] Variables are not accidentally used with null values
  [ ] Variables are immutable where possible
  [ ] Code is not repeated or duplicated
  [ ] There is an else block for every if clause even if it is empty
  [ ] No complex/long boolean expressions
  [ ] No negatively named boolean variables
  [ ] No empty blocks of code
  [ ] Ideal data structures are used
  [ ] Constructors do not accept null/none values
  [ ] Collections are initialised with a specific estimated capacity
  [ ] Arrays are checked for out of bound conditions
  [ ] Catch clauses are fine grained and catch specific exceptions
  [ ] Exceptions are not eaten if caught, unless explicitly documented otherwise
  [ ] APIs and other public contracts check input values and fail fast
  [ ] Files/Sockets/Cursors and other resources are properly closed even when an exception occurs in using them
  [ ] StringBuilder is used to concatenate strings
  [ ] Null/None are not returned from any method
  [ ] Floating point numbers are not compared for equality
  [ ] Loops have a set length and correct termination conditions
  [ ] Blocks of code inside loops are as small as possible
  [ ] Order/index of a collection is not modified when it is being looped over
  [ ] No methods with boolean parameters
  [ ] No object exists longer than necessary
  [ ] Design patterns if used are correctly applied
  [ ] No memory leaks
  [ ] Law of Demeter is not violated
  [ ] Methods return early without compromising code readability
  - Java only
    [ ] Appropriate JCIP annotations are used
    [ ] No use of Object class, use generics instead
    [ ] Uses final modifier to prevent mistaken assignments


- Documentation
  [ ] All methods are commented in clear language.
  [ ] Comments exist and describe rationale or reasons for decisions in code
  [ ] All public methods/interfaces/contracts are commented describing usage
  [ ] All edge cases are described in comments
  [ ] All unusual behaviour or edge case handling is commented
  [ ] Data structures and units of measurement are explained

- Threading
  [ ] Objects accessed by multiple threads are accessed only through a lock, or synchronized methods.
  [ ] Race conditions have been handled
  [ ] Locks are acquired and released in the right order to prevent deadlocks, even in error handling code.
  [ ] StringBuffer is used to concatenate strings in multi-threaded code

- Security
  [ ] All data inputs are checked (for the correct type, length/size, format, and range)
  [ ] Invalid parameter values handled such that exceptions are not thrown
  [ ] No sensitive information is logged or visible in a stacktrace