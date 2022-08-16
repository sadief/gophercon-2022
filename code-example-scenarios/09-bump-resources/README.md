## Prerequisites
- Must have Go installed
- Must have Docker Desktop installed
- Must have Kubernetes + kubectl installed

## Run
`make run` from folder root to build docker container
`make forward` to port forward so you can run the endpoint
`curl http://localhost:8080/serve-customer/3000` to hit endpoint on deployment (and test the limits!)
`make down` to tear down deployment, stop container, and remove image