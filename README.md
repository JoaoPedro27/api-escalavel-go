# 🚀 API Escalável em Go com Kubernetes
![RESPOSTA](https://github.com/user-attachments/assets/e6168ad7-991a-4e83-a4b7-a565f0fced04)

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://golang.org/)

[![Docker](https://img.shields.io/badge/Docker-24.0+-2496ED?logo=docker)](https://www.docker.com/)

[![Kubernetes](https://img.shields.io/badge/Kubernetes-1.29+-326CE5?logo=kubernetes)](https://kubernetes.io/)

Uma API teste em Go utilizando Gin Framework, containerizada com Docker e implantada em Kubernetes para alta escalabilidade.


## 📋 Pré-requisitos

- **Go 1.24+** ([Download](https://go.dev/dl/))
- **Docker Desktop** ([Windows](https://docs.docker.com/desktop/install/windows-install/))
- **Minikube** ([Instalação](https://minikube.sigs.k8s.io/docs/start/))
- **kubectl** ([Guia](https://kubernetes.io/docs/tasks/tools/))
- **k6** (Opcional para testes de carga: `choco install k6`)

---

## 🛠️ Como Usar
### 🦫 Go
#### 1. Clone o Repositório
```bash
git clone https://github.com/JoaoPedro27/api-escalavel-go.git
cd api-escalavel-go
```
#### 2. Instale as Dependências
```bash
go mod download
```
#### 3. Execute Localmente
```bash
go run main.go
```
#### Teste a API:
```bash
http://localhost:8080/health
http://localhost:8080/hello
```
### 🐳 Docker
#### 1. Construa a Imagem
```bash
docker build -t go-api:latest .
```
#### 2. Execute o Container
```bash
docker run -p 8080:8080 go-api
```
### ☸️ Implantação no Kubernetes
#### 1. Inicie o Cluster
```bash
minikube start --driver=docker
```
#### 2. Carregue a Imagem no Minikube
```bash
minikube image load go-api:latest
```
#### 3. Implante a Aplicação
```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```
#### 4. Acesse a API
```bash
minikube service go-api-service --url
```
### 📊 Teste de Carga com k6
#### Execute o Teste (100 usuários por 1 minuto)
```bash
k6 run load-test.js
```
#### Exemplo de Resultados:
```bash
http_req_duration........: avg=12.27ms min=503.6µs med=1.81ms max=87.79ms
http_req_failed..........: 0.00% ✓ 0 ✗ 5976
```

[![Linkedin](https://img.shields.io/badge/LinkedIn-326CE5?logo=Linkedin)](https://www.linkedin.com/in/joaopmroberto/)
