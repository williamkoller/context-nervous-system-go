## Tipos de context disponíveis:

| tipo | quando usar | exemplo
| --------------------- | ------------------------------------ | ----------------------------|

| `context.Background()`  | Apenas no `main()` ou inicializacões         | `ctx := context.Background()`
| `context.TODO()`        | Placeholder temporário (não use em producão) | `ctx := context.TODO()`
| `context.WithCancel()`  | Cancelamento Manual                          | Shutdown graceful
| `context.WithTimeout()` | Timeout relativo                             | Request HTTP com limite de tempo
| `context.WithDeadline()`| Deadline absoluta                            | Job que deve terminar até X hora