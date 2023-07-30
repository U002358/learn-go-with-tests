### Dicas

* Use channels quando for passar a propriedade de um dado
* Use mutexes para gerenciar estados

    [A wiki do Go tem uma página dedicada para esse tópico: Mutex ou Channel?](https://github.com/golang/go/wiki/MutexOrChannel)

* Context: se uma função necessita de alguns valores, coloque-os como parâmetros tipados em vez de tentar obtê-los a partir de **context.Value**, isto torna-o estaticamente verificado e documentado para todos o vejam