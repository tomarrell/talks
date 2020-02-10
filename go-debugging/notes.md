% Notes: Debugging Go Programs
% Tom Arrell
% Wed 12th Feb – Golang Meetup – Berlin

---

# Overview

* Go at SumUp
* Go in Logistics
* Why we're moving to Go

* Primitive debugging
* Delve

---

# Go at SumUp

We use Go here at SumUp. Although we're still relatively early in our Go story.

For some context, Go has been used here in our SysOps team for a couple of
years. But it only really reached a production backend service early last year.

We run most of our services inside Kubernetes, on AWS, after a period of
migrating away from more primitive deployments. Some teams have included in
their migation breaking up monolithic Ruby services into smaller Go
microservices, we'll touch on why this decision was made in a bit.

In our payments domain, we have a custom environment bootstrapper we call
Theseus, which is written in Go. This was built in order to make it easier to
bootstrap our payments applications in ephemeral environments with all of their
dependencies.

---

# Go in Logistics

So I work on our Logistics team where we build software to make the delivery
process to our merchants more seamless for other teams within SumUp.

When I joined 6 months ago, almost all the Logistics processes within SumUp were
running through a single set of Python scripts, running on a single job server
which was hooked up to a very large Postgres database. Since then, we've
migrated our European operations to a new set of Golang services, which run in
Kubernetes, and provide API's for other teams to interact with.

Also during this time, we took 2 engineers with no background in Golang to being
comfortable contributing.

This has brought a massive improvement in productivity, reliability and
observability of our processes. As well as a steady platform to expand our
logistics operations easily into new markets.

Go in particular has made this easy to do with its simplicity. We've found that
engineers new to Go have had a easy time picking it up and becoming productive.

Another benefit I consider is its relative lack of abstractions. Logistics is
one of those domains with a large number of edge cases. Go intentionally doesn't
give you the abstract power that some other languages do. Heavy use of
abstractions puts the burden on the developer to be very careful that any
abstractions they do implement are in fact valid across many different markets.
Spoiler alert, I believe that's probably impossible.

---

# Why we're moving to Go

SumUp is in the stage of it's growth where scaling the engineering practices in
the organisation is important. Our engineering is rather scattered across the
globe, including in our offices in Sao Paulo Brazil, Sofia, Cologne and two here
in Berlin.

Also, at SumUp we really encourage people to be "T" shaped. In other words,
having a broad range of skills across many disciplines, while also specializing
where needed.

To do this, we support engineers who want to learn new things by helping them
change teams, form new teams, and do all of this as simply as possible. Part of
what makes this simple is having as much consistency across teams as makes
sense. This makes it much easier for someone to get up to speed with the
codebase of a new team in relatively little time. Adopting Go is a step towards
this possibility, putting aside the yak shaving, leaving more time for decisions
that bring value to the business.

We now have teams in all 4 of our locations with engineering who are writing Go.

// TODO ruby

---

# Debugging: Downsides of fmt.Println

Print debugging is something that probably the majority of programmers are
familiar with. It's a simple and easy to use tool, especially when you want to
inspect very specific parts of your program, and are running things locally on
your machine.

However if either of those things is not clearly defined, the challenge begins. What
if your bug occurs on the 237th iteration of this *for* loop? Or what if you
don't know which iteration it occurs on? Or what if your program is already
running, and you can't restart it in fear that you'll have to wait another 3
days for the bug to appear? Hence, sometimes you can save a lot of time by
picking up a debugger.

---

# Debugging: GDB

Debugging Go programs can be quite the non-trivial task. Especially when they
become highly concurrent. Go programs also confuse GDB with the way that they
handle stack management and threading for example.

A quick example is the defer statement. You can use the defer statement to
change the return value of the function, however this extra execution after a
return is non-standard, and can lead to execution of code which is not known to
GDB.

GDB is also not aware of the Go scheduler's context switching. It is possible
for a goroutine to be preempted, and scheduled on another processor, which can
cause the debugger to hang.

Other limitations of GDB include it struggling with types derived from strings,
and method qualifications from packages, causing it to treat identifiers
including a "." as unstructured literals. This is made even more difficult when
you have methods from other packages implementing interfaces defined locally.

GDB is extremely versatile, and the Go team have released extensions to make
using it more ergonomic. E.g. pretty printing strings, slices, maps, channels
and interfaces. You can even print directly the length or capacity of slices.

These all help to improve the usability of GDB, but we can do a little better.

---

# Debugging: Delve

Enter Delve. Now I'd expect a lot of people here have probably heard, if not
used Delve themselves to debug their programs. It's been around for quite a
while these days, having been started by Derek Parker back in early 2014.

Delve was purpose built for debugging Go, and deals with some of the shortfalls
that are present in GDB.

Probably the best way I can introduce the power of Delve is through a demo.

Just something to note, Delve works best on Linux, a few commands are only
available on Linux. If you're running a Mac, you can get most of the benefit
running within a Docker container.

---

# Race Detector:

In the words of someone much wiser than myself, "No race is a safe race." Don't
get into the habit of thinking that you have a

---

# Core Dumps

You can analyse a crashed program in more detail, including getting views of the
source as the program crashed.

---

# Memory Leaks


---

Notes:

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

---

# Delve TODO

- dlv run: start debugging, disables optimisations, compiles code, starts
    program, attaches to the process

- dlv test: used for programs without main, compiles test binary, starts binary,
    attaches to process

- dlv attach <pid>: careful, program could be optimised and could run into
    issues

- runtime.Breakpoint() to trigger a breakpoint in your debugger, being able to set
it in your program.

- Checkpoints, restart your program from a specific place

---

# Scenarios

Come up with scenarios to show debugging different situations

* Race condition
* Deadlock
* https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables









