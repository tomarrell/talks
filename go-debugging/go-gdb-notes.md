# Notes:

Why does go confuse existing debuggers?

- Go's execution model, very different from C, C++
    - Defer statement: going to be used for something relatively
        inconsequential, but sometimes you may do something that may change the
        return value fo your function. GDB doesn't know about defer statements,
        they will not follow flow to the defer statement.
    - Threads vs goroutines: GDB plans to deal with relatively standard process.
        Standard process will probably have many threads. It expects this to be
        the smallest unit of execution. Go has threads, but also has goroutines,
        which are scheduled on to the underlying threads.
    - Go scheduler: Processors, os threads, goroutines
    - Context switches: Goroutines can be paused, can be rescheduled onto
        another processor. Debuggers may not expect this. GDB doesn't know about
        the context switch, can cause the debugger to hang.
        - Can manually schedule a context switch
        - Blocking syscalls can cause
        - Channel operations
        - Garbage collection
        - Go statement, the running goroutine will be swapped out so the new one
            can be run immediately
- Stack management
    - Runtime stack inspection
        - Goroutines initialized with 4k stack
        - Check for stack growth
        - Confuses debuggers
    - GDB will crash when you try to call a function, as it tries to allocated a
        new stack, runtime stack inspection fails
- Compiler optimisations
    - Function inlining
    - Registerizing variables
        - Storing information in a register, not the stack
