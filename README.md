## Tipos de context disponíveis

| Tipo                     | Quando usar                                  | Exemplo                          |
| ------------------------ | -------------------------------------------- | -------------------------------- |
| `context.Background()`   | Apenas no `main()` ou inicializações         | `ctx := context.Background()`    |
| `context.TODO()`         | Placeholder temporário (não use em produção) | `ctx := context.TODO()`          |
| `context.WithCancel()`   | Cancelamento manual                          | Shutdown graceful                |
| `context.WithTimeout()`  | Timeout relativo                             | Request HTTP com limite de tempo |
| `context.WithDeadline()` | Deadline absoluta                            | Job que deve terminar até X hora |
