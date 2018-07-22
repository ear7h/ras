# RAS

Random access storage, this uses a file for c-like storage calls. 

### Disclaimer
Due to the nature of the calls simple pointer mistakes can cause entire database corruption.

**todo**
* more alloc/free tests
* implement `Bsize` method in
* fuzzy testing operations
* get/set tests
* write `go generate` prog for runtime methods
* type system
    * implement `type Sizer interface{ Size() uint }` for fixed size vars
    * implement `type Pointer interface { Point() uint }` for dynamic size vars
    * test fixed size variables
    * array GetI methods
        * reuse the fixed types 
    * slices
        * fixed portion, 16 bytes
            * len, 64 bit uint
            * data, pointer to array
        * arrays with 64 bit length header
    * string
        * same immutable behavior as `go`
        * set always does a realloc
    * pointers
    * maps
        * how to handle typing?
            * check if `Sizer` or `Pointer`
        * hashing algo
    * structs?
        * `(Runtime).LoadStruct(v inteface{})` method which uses reflection
            * this would make things extremely simple for the user but has terrible performance
        * `type Structer interface{ Load(addr uint, a Allocator), Size() uint }`
            * complicated for end user but 
* ~~concurrency~~
    * add race tests
* ~~remove init~~
* change shrink operation
    * automatically remove links to thos
* acid compliance

**far ahead**
* optimize for speed
* maybe use mmap to circumvent io errors
    * mem only capability
* create a heap based allocator