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

We run most of our services inside Kubernetes, on AWS. Many teams have included in their migations breaking up monolithic Ruby services into smaller microservices, we'll touch on why this decision was made in a bit.

In our payments domain, we have a custom environment bootstrapper we call Theseus, which is written in Go. This was built in order to make it easier to bootstrap our payments applications in ephemeral environments with all of their dependencies.

--- 

# Go in Logistics

So I work on our Logistics team, building software to make the delivery process to our merchants more seamless for other teams within SumUp.

Since I joined 6 months ago, we began migrating the existing processes to two main services within my team. What may or may not surprise you is that global logistics processes are choc-full of edge cases and rather complicated business logic. One thing I particularly like about using Go for such a project is the explicitness of each case. 

There are no fancy abstractions, we have a set of models


---

# Why we're moving to Go

SumUp is in the stage of it's growth where scaling the engineering practices in the organisation is important. Our engineering is rather scattered across the globe, including in our offices in Sao Paulo Brazil, Sofia, Cologne and two here in Berlin.

Also, at SumUp we really encourage people to be "T" shaped. In other words, having a broad range of skills across many disciplines, while also specializing where needed.

To do this, we support engineers who want to learn new things by helping them change teams, form new teams, and do all of this as simply as possible.

Previously most of the auxiallary software at SumUp was written in Ruby. 

Developers are naturally a pretty opinionated bunch.  
