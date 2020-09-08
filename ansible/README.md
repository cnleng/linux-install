# Data Cohenrece Challenge - Ansible

#### - [Purpose](README.md#purpose)
#### - [Requirements](README.md#requirements)
#### - [Getting Started](README.md#getting-started)

</br>

## [Purpose](#purpose)
-----

This playbook and roles were made to install Halyard and some tools which are necessary to make Spinnaker and AWS integration to work properly. This version only supports Halyard installation, but Spinnaker Configuration and Deployment will be added in the future.

By running this playbook, you will install:
- [docker](roles/docker/README.md)
- [awscli](roles/k8s/README.md)
- [docker](roles/spinnaker/README.md)
- [halyard](roles/halyard/README.md)
- [kubectl](roles/minio/README.md)

</br>

## [Requirements](#requirements)
[Click here](README.md#getting-started) if you already have all requirements

-----
#### Desired host
- Ubuntu == 18.04
- Python >= 2.7

#### Your host
- Python >= 2.7
- PIP
- Ansible >= 2.6

</br>

#### Installing Requirements

1. Install Python and PIP

```sh
# Debian/Ubuntu
sudo apt-get install python python-pip

# CentOS/RHEL
sudo yum install yum-utils epel-release
sudo yum-config-manager --enable epel
sudo yum install python python-pip
```

2. Install Ansible using the PIP distribution

```sh
sudo -i pip install --upgrade ansible
```

</br>

## [Getting Started](#getting-started)
-----
1. Clone the repository
```sh
git clone https://github.com/tioxy/spinnaker-ansible.git
```

2. Edit the *sample-inventory* file and configure the desired host connection following [Ansible Docs](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html)

3. Run the file *main.yml*
```sh
ansible-playbook -i inventory/hosts master.yml
```
