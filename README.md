## Reto

[🔗Url](https://roadmap.sh/projects/expense-tracker)

Cree una aplicación sencilla de seguimiento de gastos para administrar sus finanzas. La aplicación debe permitir a los usuarios agregar, eliminar y ver sus gastos. La solicitud también debe proporcionar un resumen de los gastos.

### Requisitos
La aplicación debe ejecutarse desde la línea de comando y debe tener las siguientes características:

- Los usuarios pueden agregar un gasto con una descripción y un monto.
- Los usuarios pueden actualizar un gasto.
- Los usuarios pueden eliminar un gasto.
- Los usuarios pueden ver todos los gastos.
- Los usuarios pueden ver un resumen de todos los gastos.
- Los usuarios pueden ver un resumen de los gastos de un mes específico (del año en curso).
Aquí hay algunas características adicionales que puedes agregar a la aplicación:

- Agregue categorías de gastos y permita a los usuarios filtrar gastos por categoría.
- Permitir a los usuarios establecer un presupuesto para cada mes y mostrar una advertencia cuando el usuario excede el presupuesto.
- Permitir a los usuarios exportar gastos a un archivo CSV.
La lista de comandos y su resultado esperado se muestra a continuación:

```bash
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

$ expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ expense-tracker summary
# Total expenses: $30

$ expense-tracker delete --id 2
# Expense deleted successfully

$ expense-tracker summary
# Total expenses: $20

$ expense-tracker summary --month 8
# Total expenses for August: $20
```

Comandos y flags

// add  --description --amount
// update --id --description --amount
// list
// summary --month
// delete --id
// help


// Conn bd Servicios CRUD
//Posgrest

## Orden de las carpetas
```
ExpenseTracker/
├── cmd/                    # Entrada principal de la CLI
│   └── root.go            )
│   └── add.go              # Subcomandos
│   └── version.go
│
├── internal/               # Lógica interna, dominio y casos de uso
│   ├── domain/             # Entidades, interfaces del negocio
│   │   └── user.go         # Ejemplo: entidad User
│   │   └── user_service.go # Interfaz de dominio
│   │
│   ├── usecase/            # Casos de uso (lógica de aplicación)
│   │   └── create_user.go
│   │   └── list_users.go
│   │
│   ├── infra/              # Implementaciones concretas
│   │   └── repo/           # Acceso a archivos, DB, etc.
│   │   │   └── file_user_repo.go
│   │   └── logger/         # Logger, config, etc.
│   │       └── zap_logger.go
│   │
│   └── cli/                # Lógica de entrada/salida CLI
│       └── user_presenter.go
│
├── pkg/                    # Código reutilizable fuera del core
│   └── utils/              # Validaciones, parsers, etc.
│       └── file_helper.go
│
├── test/                   # Tests de integración y e2e
│   └── cli_test.go
│
├── mocks/                 # Mocks generados (gomock/mockery)
│   └── user_service_mock.go
│
├── go.mod
├── go.sum
└── main.go                # Punto de entrada de la app
```

## Comandos

```bash
make build   # Compila el binario
make run     # Lo compila y lo ejecuta
make fmt     # Formatea el código
make test    # Ejecuta los tests
make clean   # Borra los binarios generados
```
