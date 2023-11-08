# dev-provision
DevProGo is a Go-based application that empowers developers to swiftly provision development environments resembling production setups using Infrastructure as Code (IaC) principles. Simplify and accelerate the process of creating development environments with ease.

```
provision-cli/
│
├── cmd/
│   └── provision/
│       ├── main.go         # Ponto de entrada da aplicação CLI
│       └── commands/       # Definições dos comandos CLI
│           ├── create.go   # Comando para criar ambientes de desenvolvimento
│           ├── delete.go   # Comando para excluir ambientes de desenvolvimento
│           ├── list.go     # Comando para listar ambientes de desenvolvimento
│           └── <outros_comandos>.go
│
├── internal/
│   ├── provision/         # Lógica interna da aplicação
│   │   ├── provisioner.go # Módulo principal de provisionamento
│   │   ├── config.go     # Gerenciamento de configurações
│   │   ├── api.go        # Integração com a API REST (se aplicável)
│   │   └── <outros_módulos>.go
│   │
│   └── pkg/              # Pacotes e utilitários reutilizáveis
│
├── scripts/              # Scripts para automação, implantação e configuração
│
├── config/               # Arquivos de configuração da aplicação
│
├── tests/                # Testes unitários e de integração
│
├── README.md             # Documentação do projeto
│
├── LICENSE               # Licença do projeto
│
└── go.mod                # Arquivo de definição de módulo Go
```