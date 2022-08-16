## Prerequisites
- Must have Go installed
- Must have Docker Desktop installed
- Must have Kubernetes + kubectl installed

## Run
`make up` from folder root to build docker image, run container, and deploy on k8's
`make forward` to port forward so you can run the endpoint
`curl http://localhost:8080/serve-customer/3` to hit endpoint on deployment
`make down` to tear down deployment, stop container, and remove image