Scheduler preemptivo: define tempo por tarefa, orquestra as tarefas.
vantagem: não vai ficar travado, pq a tarefa sempre vai acabar.
desvantagem: mudança continua de contexto.
ex.: 5s para cada tarefas rodar

Scheduler cooperativo: Espera uma tarefa acabar para começar a outra.
vantagem: tem menos interrupção de tarefa
desvantagem: caso alguma tarefa seja longa, vai esperar até acabar.

GO

A linguagem com tem um runtime

O GO não faz syscall pra gerar thread, ele tem um próprio gerenciamento de threads.
As threads do Go não são verdadeiras, são greenthreads (threads in userlands)
GO cria pequenos processos no runtime.
GO tem seu próprio scheduler, que trabalha de forma cooperativa.
O GO perceber se uma green thread vai travar a execução, ele altera para modo preemptivo.

O GO usa 2kb inicialmente de memória, contra 1Mb das demais linguagens.

 

Para debugar se estamos em cenário de race condition
go run -race main.go
-> exibe no log os problemas conforme o processo está em execução


Channel:
- comunicação entre threads
- 