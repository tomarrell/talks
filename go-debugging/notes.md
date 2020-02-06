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
* 

---

# Go at SumUp

We use Go here at SumUp. Although we're still relatively early in our Go story. 

For some context, Go has been used here in our SysOps team for a couple of years. But it only really reached a production backend service early last year.

We run most of our services inside Kubernetes, on AWS, after a period of migrating away from more primitive deployments. Some teams have included in their migation breaking up monolithic Ruby services into smaller Go microservices, we'll touch on why this decision was made in a bit.

In our payments domain, we have a custom environment bootstrapper we call Theseus, which is written in Go. This was built in order to make it easier to bootstrap our payments applications in ephemeral environments with all of their dependencies.

--- 

# Go in Logistics

So I work on our Logistics team where we build software to make the delivery process to our merchants more seamless for other teams within SumUp.

When I joined 6 months ago, almost all the Logistics processes within SumUp were running through a single set of Python scripts, running on a single job server which was hooked up to a very large Postgres database. Since then, we've migrated our European operations to a new set of Golang services, which run in Kubernetes, and provide API's for other teams to interact with.

Also during this time, we took 2 engineers with no background in Golang to being comfortable contributing.

This has brought a massive improvement in productivity, reliability and observability of our processes. As well as a steady platform to expand our logistics operations easily into new markets.

Go in particular has made this easy to do with its simplicity. We've found that engineers new to Go have had a easy time picking it up and becoming productive.

Another benefit I consider is its relative lack of abstractions. Logistics is one of those domains with a large number of edge cases. Go intentionally doesn't give you the abstract power that some other languages do. Heavy use of abstractions put the burden on the developer to be very careful that any abstractions they do implement are in fact valid across many different markets. Spoiler alert, I believe that's nearly impossible.

---

# Why we're moving to Go

SumUp is in the stage of it's growth where scaling the engineering practices in the organisation is important. Our engineering is rather scattered across the globe, including in our offices in Sao Paulo Brazil, Sofia, Cologne and two here in Berlin.

Also, at SumUp we really encourage people to be "T" shaped. In other words, having a broad range of skills across many disciplines, while also specializing where needed.

To do this, we support engineers who want to learn new things by helping them change teams, form new teams, and do all of this as simply as possible.

Previously most of the auxiallary software at SumUp was written in Ruby. 

Developers are naturally a pretty opinionated bunch.  
