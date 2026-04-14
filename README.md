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

Provide:

Access Key
Secret Key
Region (e.g., ap-south-1)
Output format: json
🐳 Docker Image

The application is containerized and available publicly on DockerHub:

https://hub.docker.com/r/<your-dockerhub-username>/simple-time-service
Build Locally
cd app
docker build -t simple-time-service .
Run Locally
docker run -p 8080:8080 simple-time-service

Test:

http://localhost:8080
🚀 Deployment Instructions

Clone the repository:

git clone <your-repo-url>
cd terraform

Initialize Terraform:

terraform init

Review the plan:

terraform plan

Deploy infrastructure:

terraform apply

Type yes when prompted.

🌐 Access the Application

After deployment:

terraform output alb_url

Open in browser:

http://<ALB_DNS_NAME>

Expected response:

{
  "timestamp": "2026-04-14T10:00:00Z",
  "ip": "x.x.x.x"
}
🧹 Cleanup (IMPORTANT)

To avoid unnecessary AWS charges:

terraform destroy
🔒 Security Considerations
No credentials are stored in the repository
ECS tasks run in private subnets
Security groups restrict access:
ALB → public HTTP access
ECS → only accessible from ALB
Container runs as a non-root user
📁 Project Structure
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
⚡ CI/CD Pipeline (Extra Credit)

A GitHub Actions workflow is included to automate:

Docker image build
Push to DockerHub
Workflow Location:
.github/workflows/docker.yml
Required GitHub Secrets:
DOCKER_USERNAME
DOCKER_PASSWORD