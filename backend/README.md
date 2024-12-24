# Sequence Diagram of Application

```mermaid
sequenceDiagram;
    create participant Client
    create participant Authentication Server
    create participant Authentication Database

    Client->>Authentication Server: POST /auth — Login/Authentication request
    activate Authentication Server
    Authentication Server->>Authentication Database: Check Username and Password
    activate Authentication Database
    Authentication Database-->>Authentication Server: Username and Password Validation
    deactivate Authentication Database
    Authentication Server-->>Client: Returns JSON Web Token (JWT)
    deactivate Authentication Server

    destroy Authentication Server
    destroy Authentication Database

    create participant Reverse Proxy
    create participant API Server
    create participant API Database

    Client->>Reverse Proxy: POST /getData and /ws — Request containing JWT
    activate Reverse Proxy
    Reverse Proxy-->>Client: If JWT rejected
    Reverse Proxy->>API Server: If JWT approved
    deactivate Reverse Proxy

    activate API Server
    API Server->>API Database: GET data
    activate API Database
    API Database-->>API Server: Data
    activate API Database
    API Server-->>Client: Data
    deactivate API Server

    destroy Client
    destroy Reverse Proxy
    destroy API Server
    destroy API Database
```
