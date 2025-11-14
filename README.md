# Keycloak + Go (OIDC) — Exemplo prático

Este repositório mostra como integrar uma aplicação Go com **Keycloak** utilizando **OpenID Connect (OIDC)** e **OAuth2**. Inclui um `docker-compose.yml` para subir o Keycloak (com PostgreSQL) e instruções passo a passo para configurar realm, client e usuário.
<br>

## Iniciar o Keycloak com Docker
Rode o comando abaixo em seu terminal e espere o Keycloak subir:
```bash
docker run -p 127.0.0.1:8080:8080 -e KC_BOOTSTRAP_ADMIN_USERNAME=admin -e KC_BOOTSTRAP_ADMIN_PASSWORD=admin quay.io/keycloak/keycloak:26.4.5 start-dev
```

## Configuração do ambiente
Crie um arquivo .env dentro da pasta cmd/ com base no exemplo fornecido:
```bash
cp cmd/.env.sample cmd/.env
```

## App
```bash
cd cmd/
go run main.go
```

## Configuração passo a passo no Keycloak

Abaixo está o passo a passo para configurar um **Realm**, **Client** e **Usuário de teste** no Keycloak para uso com a aplicação Go do repositório.

> **Pré-requisitos:** Keycloak rodando (ex.: `http://localhost:8080`) e acesso ao Admin Console.

---

### 1. Acessar o Admin Console
1. Abra [http://localhost:8080](http://localhost:8080) no navegador.
2. Faça login com o usuário e senha admin (ex.: `admin` / `admin`).

---

### 2. Criar um Realm
1. No canto superior esquerdo, clique em **Manage realms** (por padrão `Master`) → **Create realm**.
2. **Realm name:** `demo`  
3. Clique em **Create**.

---

### 3. Criar um Client (aplicação)
1. No menu lateral do realm `demo`, clique em **Clients** → **Create client**.
2. Preencha:
   - **Client ID:** `go-client` → **Next**
   - **Authentication flow:** ✅`Standard flow` | ✅`Direct access grants`
   - **Client authentication** `On`
   - **Authorization** `On`
   - **Root URL:** `http://localhost:8081`
   - **Valid redirect URIs** `http://localhost:8081/*`
3. Clique em **Save**.

#### Obter Client Secret
1. Vá até a aba **Credentials** do client.
2. Copie o **Client secret** e coloque no seu `.env` como `CLIENT_SECRET`.

---

### 4. Criar um User
1. No menu lateral do realm `demo`, clique em **Users** → **Create new user**.
2. Preencha:
   - **Email verified** `On`
   - **Username:** go-user
   - **Email:** foobar@gmail.com
   - **First name:** Foo
   - **Last name:** Bar
3. Clique em **Create**.

#### Definir password
1. Vá até a aba **Credentials** do user.
2. Preencha:
   - **Password:** ex: 123456
   - **Password confirmation:** 123456
   - **Temporary** `Off`
3. Clique em **Save**.

---

### Saiba mais
1. Documentação: [keycloak.org](https://www.keycloak.org/documentation)
2. API: [keycloak API](https://www.keycloak.org/docs-api/latest/rest-api/index.html)
3. OAuth 2.0 Authorization Framework: [OAuth2](https://auth0.com/docs/authenticate/protocols/oauth)
4. OpenID Connect Protocol: [OpenID Connect](https://auth0.com/docs/authenticate/protocols/openid-connect-protocol)
5. Mini curso de Keycloak: [YouTube - Marco Seabra](https://www.youtube.com/watch?v=ZiN3NLBro5U&list=PL-XEb3GK7JUrzMDzRAiRG-AV8iP73FF8T&index=2)