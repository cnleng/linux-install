Install Docker
-------------- 
```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

sudo apt-get update

sudo apt-cache policy docker-ce

sudo apt-get install -y docker-ce

sudo usermod -aG docker ${USER}

su ${USER}

id -nG
```

Install Docker-Compose
---------------------- 
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

Install Kubernetes Private Registry
----------------------------------- 
https://www.linuxtechi.com/setup-private-docker-registry-kubernetes/

Alternatively
------------------
```
sudo apt-get update && apt-get install -y curl apt-transport-https

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
cat <<EOF >/etc/apt/sources.list.d/docker.list
deb https://download.docker.com/linux/$(lsb_release -si | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) stable
EOF
sudo apt-get update && sudo apt-get install -y docker-ce=$(apt-cache madison docker-ce | grep 17.03 | head -1 | awk '{print $3}')
```
Install Kubernetes
------------------
Install Kubeadm
```
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubeadm=1.13\* kubectl=1.13\* kubelet=1.13\* kubernetes-cni=0.7\*
sudo apt-mark hold kubelet kubeadm kubectl
swapoff -a
# Comment out swap line in fstab so that it remains disabled after reboot
vi /etc/fstab
```
Tocreate your cluster
```
kubeadm init --kubernetes-version stable-1.13  --pod-network-cidr=10.244.0.0/16
```

To start using your cluster, you need to run the following as a regular user:
```
  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

Alternatively, if you are the root user, you can run:
```
export KUBECONFIG=/etc/kubernetes/admin.conf
```

You should now deploy a pod network to the cluster.
https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/#pod-network
```
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/62e44c867a2846fefb68bd5f178daf4da3095ccb/Documentation/kube-flannel.yml
```

You can now join any number of machines by running the following on each node
as root:
```
  kubeadm join 100.80.245.242:6443 --token b4l44r.babq3qyjgqcsgcph --discovery-token-ca-cert-hash sha256:ac1cb1bdca98c32db61e38a2a8953de66f333ba04525c8083faa322781ed3d22
```

Check all nodes in cluster (on master)
```
kubectl get nodes
```
