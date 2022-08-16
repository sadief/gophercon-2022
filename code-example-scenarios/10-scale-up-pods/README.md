## Prerequisites
- Must have Go installed
- Must have Docker Desktop installed
- Must have Kubernetes + kubectl installed
- Must have Ingress controller installed

## Run
`make run` from folder root to build docker container
`curl http://localhost/serve-customer/30000` several times to hit endpoint on deployment and see the pod traffic shared across pods
`make down` to tear down deployment, stop container, and remove image