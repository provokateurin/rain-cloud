terraform {
  required_providers {
    k3d = {
      source  = "pvotal-tech/k3d"
      version = "0.0.6"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.21.0"
    }

    helm = {
      source  = "hashicorp/helm"
      version = "2.10.0"
    }

    kubectl = {
      source  = "gavinbunney/kubectl"
      version = "1.14.0"
    }
  }
}
