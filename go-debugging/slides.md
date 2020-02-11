% Debugging Go Programs
% Tom Arrell
% Wed 12th Feb – Golang Meetup – Berlin

---

# About me

## Tom Arrell 

- Senior Backend Engineer @ SumUp
- Logistics Squad
- twitter:  twitter.com/tom_arrell
- github:   github.com/tomarrell

---

# Agenda

:::incremental
- Go at SumUp
- Go in Logistics
- Why we're moving to Go
- Debugging...
  - Primitive
  - GDB
  - Delve
  - Scenarios...
:::

---

# Go at SumUp

- Adopted within the last 2 years
- Mainly used for tooling
- First services written ~1 year ago
- Now migrating to Go for new services
- Deployed to Kube

---

# Go in Logistics

- Joined around ~6 months ago
- Legacy Python scripts
  - Lack of monitoring
  - Email alerts
  - Git clone deployment
- Replaced with Go services
  - Deployed to Kubernetes
  - Prometheus, Sentry, OpsGenie
- 2 engineers with no prior Go experience brought up to speed

---

# Go in Logistics

![Alerting, the old way](./media/python-exception-email.png)

---

# Go in Logistics

![Alerting, the new way](./media/sentry-report.png){ width=350px }

---

# Go in Logistics

![Monitoring](./media/grafana-dashboard.png)

---

# Leaky Abstractions

> All non-trivial abstractions, to some degree, are leaky.
> 
> – Joel Spolsky

---

# Leaky Abstractions

![(Probably) leaky Rust Abstraction](./media/rust-dense.jpg){ width=350px }

---

# Ruby to Go

![Ruby to Go, why?](./media/ruby-to-go.jpg)

---

*fin*

**Questions?**
