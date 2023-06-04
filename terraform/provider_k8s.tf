provider "kubernetes" {
  client_certificate     = k3d_cluster.rain-cloud-local[0].credentials.0.client_certificate
  client_key             = k3d_cluster.rain-cloud-local[0].credentials.0.client_key
  cluster_ca_certificate = k3d_cluster.rain-cloud-local[0].credentials.0.cluster_ca_certificate
  host                   = k3d_cluster.rain-cloud-local[0].credentials.0.host
}

provider "helm" {
  kubernetes {
    client_certificate     = k3d_cluster.rain-cloud-local[0].credentials.0.client_certificate
    client_key             = k3d_cluster.rain-cloud-local[0].credentials.0.client_key
    cluster_ca_certificate = k3d_cluster.rain-cloud-local[0].credentials.0.cluster_ca_certificate
    host                   = k3d_cluster.rain-cloud-local[0].credentials.0.host
  }
}

provider "kubectl" {
  client_certificate     = k3d_cluster.rain-cloud-local[0].credentials.0.client_certificate
  client_key             = k3d_cluster.rain-cloud-local[0].credentials.0.client_key
  cluster_ca_certificate = k3d_cluster.rain-cloud-local[0].credentials.0.cluster_ca_certificate
  host                   = k3d_cluster.rain-cloud-local[0].credentials.0.host
  load_config_file       = false # Otherwise it loads KUBECONFIG
}
