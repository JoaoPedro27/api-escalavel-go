# 🚀 API Escalável em Go com Kubernetes

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-24.0+-2496ED?logo=docker)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-1.29+-326CE5?logo=kubernetes)](https://kubernetes.io/)

Uma API moderna em Go utilizando Gin Framework, containerizada com Docker e implantada em Kubernetes para alta escalabilidade.

![Diagrama de Arquitetura](https://i.imgur.com/mX9t3dD.png)  
*(Diagrama ilustrativo: Load Balancer, Kubernetes Cluster e Pods)*

## 📋 Pré-requisitos

- **Go 1.24+** ([Download](https://go.dev/dl/))
- **Docker Desktop** ([Windows](https://docs.docker.com/desktop/install/windows-install/))
- **Minikube** ([Instalação](https://minikube.sigs.k8s.io/docs/start/))
- **kubectl** ([Guia](https://kubernetes.io/docs/tasks/tools/))
- **k6** (Opcional para testes de carga: `choco install k6`)
