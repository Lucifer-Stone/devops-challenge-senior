# SimpleTimeService — DevOps Challenge

## 📌 Overview

SimpleTimeService is a minimal web service that returns the current timestamp and the IP address of the client in JSON format.

This project demonstrates a complete end-to-end DevOps workflow:
- Application development (Go)
- Containerization (Docker)
- Infrastructure as Code (Terraform)
- Cloud deployment (AWS ECS Fargate)

---

## 🧱 Architecture


User (Internet)
↓
Application Load Balancer (Public Subnets)
↓
ECS Fargate Service (Private Subnets)
↓
Containerized Go Application


### 🔍 Design Decisions

- **Go**: Lightweight, fast startup time, ideal for microservices
- **Docker**: Ensures consistency across environments
- **ECS Fargate**:
  - No server management
  - Fully managed container runtime
- **Application Load Balancer**:
  - Provides public access
  - Handles routing and health checks
- **Private Subnets**:
  - Improves security by isolating compute resources
- **Terraform**:
  - Declarative infrastructure
  - Reproducible deployments

---

## ⚙️ Prerequisites

Ensure the following tools are installed:

- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
- [Terraform](https://developer.hashicorp.com/terraform/downloads)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Git](https://git-scm.com/)

---

## 🔐 AWS Configuration

Configure AWS credentials:

```bash
aws configure
```

Provide:

```bash
Access Key
Secret Key
Region (e.g., ap-south-1)
Output format: json
```

---

## 🐳 Docker Image

The application is containerized and available publicly on DockerHub:

`https://hub.docker.com/r/mdsn6719/simple-time-service`

Build Locally

```bash
cd app
docker build -t simple-time-service .
```

Run Locally

```bash
docker run -p 8080:8080 simple-time-service
```

Test:

```bash
http://localhost:8080
```

---

## 🚀 Deployment Instructions

Clone the repository:

```bash
git clone https://github.com/Lucifer-Stone/devops-challenge-senior.git
cd terraform
```

Initialize Terraform:

```bash
terraform init
```

Review the plan:

```bash
terraform plan
```

Deploy infrastructure:

```bash
terraform apply
```

Type yes when prompted.

---

## 🌐 Access the Application

After deployment:

```bash
terraform output alb_url
```
Open in browser:

```bash
http://<ALB_DNS_NAME>
```
Expected response:

```bash
{
  "timestamp": "2026-04-14T10:00:00Z",
  "ip": "x.x.x.x"
}
```
---

## 🧹 Cleanup (IMPORTANT)

To avoid unnecessary AWS charges:

```bash
terraform destroy
```

---

## 🔒 Security Considerations

 - No credentials are stored in the repository
 - ECS tasks run in private subnets
 - Security groups restrict access:
    - ALB → public HTTP access
    - ECS → only accessible from ALB
 - Container runs as a non-root user

---

## 📁 Project Structure

```bash
.
├── app
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
└── terraform
    ├── main.tf
    ├── variables.tf
    ├── outputs.tf
    └── terraform.tfvars
```

--- 

## ⚡ CI/CD Pipeline (Extra Credit)

A GitHub Actions workflow is included to automate:
 - Docker image build
 - Push to DockerHub

Workflow Location:
.github/workflows/docker.yml

Required GitHub Secrets:
 - `DOCKER_USERNAME`
 - `DOCKER_PASSWORD`