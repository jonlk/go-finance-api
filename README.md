# Finance Api
## Go Version 1.20
### Simple finance calculations over http with the gin gonic framework
---
### Updates
- 3/1/2023 - Updated to work with AWS Lambda (fat) or standalone
- 2/28/2023 - Updated to work with Gin query binding
- 2/27/2023 - Updated to work with Kubernetes
---

Notes for local run:
- Export the environment variable RUN_AS_LAMBDA=`false` or `true`
- May need to adjust port as necessary
- Add `gin.SetMode(gin.ReleaseMode)` for release


Kubernetes Notes:
- Run in Minikube with Docker Env (adjust as necessary for your environment)
- Don't forget to enable registry and ingress
- `docker build -t vertwave/go-finance-api:latest .`
- In k8s folder
- `k apply -f ingress.yml`
- `k apply -f deploy.yml`
