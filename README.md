# RAS

Random access storage, this uses a file for c-like storage calls. 

### Disclaimer
Due to the nature of the calls simple pointer mistakes can cause entire database corruption.

**todo**
* more alloc/free tests
* fuzzy testing operations
* get/set tests
* write `go generate` prog for runtime methods 
* finish fixed size variable mechanism
    * test variable mechanism
* come up with dynamic size variable mechanism
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