# Vault Server Configuration
Steps to configurea Vault Server

## Set environement variables

```sh
export VAULT_ADDR='http://127.0.0.1:8200'
```

Initialize new vault server with one key shares:
```sh
vault operator init -key-shares=1 -key-threshold=1
```

Unseal vault server:
```sh
vault operator unseal <key> 
```

Check vault server status: 
```sh
vault status -format=json
```

Enable the secret kv engine:
```sh
vault secrets enable -version=1 -path=secret kv
```
