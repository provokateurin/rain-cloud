resource "k3d_cluster" "rain-cloud-local" {
  name  = "rain-cloud-local"
  image = "rancher/k3s:v1.26.5-k3s1"

  kube_api {
    host_ip   = "127.0.0.1"
    host_port = 6445
  }

  kubeconfig {
    update_default_kubeconfig = true
  }

  port {
    host_port      = 80
    container_port = 80
    node_filters   = ["loadbalancer"]
  }

  port {
    host_port      = 443
    container_port = 443
    node_filters   = ["loadbalancer"]
  }

  registries {
    create {
      name = "rain-cloud-local-registry.localhost"
      host = "0.0.0.0"
      host_port = 5000
    }
  }
}
