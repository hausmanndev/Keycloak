# Keycloak + Go (OIDC) — Exemplo prático

Este repositório mostra como integrar uma aplicação Go com **Keycloak** utilizando **OpenID Connect (OIDC)** e **OAuth2**. Inclui um `docker-compose.yml` para subir o Keycloak (com PostgreSQL) e instruções passo a passo para configurar realm, client e usuário.
<br>

## Estrutura do repositório
keycloak-app/<br>
│<br>
├── cmd/<br>
│ ├── .env.sample<br>
│ └── main.go # Exemplo de app Go usando OIDC (fornecido)<br>
│<br>
├── config/<br>
│ └── config.go # Loader de variáveis de ambiente (ex.: .env)<br>
│<br>
├── docker-compose.yml # Compose para Keycloak + Postgres<br>
│<br>
├── .env # Variáveis de ambiente da aplicação Go<br>
│<br>
└── README.md # Este arquivo<br>
<br>
<br>

## Configuração passo a passo no Keycloak

Abaixo está o passo a passo para configurar um **Realm**, **Client** e **Usuário de teste** no Keycloak para uso com a aplicação Go do repositório.

> **Pré-requisitos:** Keycloak rodando (ex.: `http://localhost:8080`) e acesso ao Admin Console.

---

### 1. Acessar o Admin Console
1. Abra `http://localhost:8080` no navegador.
2. Clique em **Administration Console** e faça login com o usuário admin (ex.: `admin` / `admin` conforme `docker-compose.yml`).

---

### 2. Criar um Realm
1. No canto superior esquerdo, clique no seletor de realm (por padrão `Master`) → **Add realm**.
2. **Name:** `demo`  
3. Clique em **Create**.

---

### 3. Criar um Client (aplicação)
1. No menu lateral do realm `demo`, clique em **Clients** → **Create**.
2. Preencha:
   - **Client ID:** `go-client`
   - **Client Protocol:** `openid-connect`
   - **Root URL:** `http://localhost:8081`
3. Clique em **Save**.

#### Ajustes do Client (após salvar)
1. Em **Settings**:
   - **Access Type:** `confidential`
   - **Standard Flow Enabled:** `ON` (habilita Authorization Code Flow)
   - **Direct Access Grants Enabled:** `ON` (opcional; habilita Resource Owner Password Credentials)
   - **Service Accounts Enabled:** `OFF` (ou `ON` se precisar usar client credentials)
   - **Valid Redirect URIs:** `http://localhost:8081/auth/callback`
   - **Web Origins:** `http://localhost:8081` (ou `*` apenas para desenvolvimento local)
   - **Root URL:** `http://localhost:8081` (se não preenchido antes)
2. Clique em **Save**.

#### Obter Client Secret
1. Vá até a aba **Credentials** do client.
2. Copie o **Secret** (Client Secret) e coloque no seu `.env` como `CLIENT_SECRET`.

---

### 4. Incluir roles no token (opcional — quando precisar enviar roles)
Se sua aplicação precisa que roles apareçam no ID/Access Token, siga um destes caminhos:

**Opção A — Usar mappers do client**
1. Dentro do client `go-client`, vá em **Mappers** → **Create**.
2. Exemplo de mapper para incluir **realm roles**:
   - **Name:** `realm-roles-mapper`
   - **Mapper Type:** `User Realm Role` (mapeia papéis do realm)
   - **Token Claim Name:** `realm_access.roles` (ou `roles`, conforme sua preferência)
   - Marque: **Add to ID token** e **Add to access token**
   - Salve.

**Opção B — Usar Client Scopes**  
1. Em **Client Scopes** (menu lateral do realm), crie/edite um scope que inclua os mappers desejados e depois associe esse scope ao client em **Client → Client Scopes**.

> Observação: o Keycloak também pode incluir roles automatica
