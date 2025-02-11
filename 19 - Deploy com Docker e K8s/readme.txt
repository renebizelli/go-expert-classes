Para buildar:

>> GOOS=windwos|linux go build CGO_ENABLED=0 -o NOME_DO_BIN_SERA_GERADO . 
 -CGO_ENABLED=0 = desliga a possibilidade de usar libs em C.

Otimizar arquivo gerado

>> GOOS=windwos|linux go build CGO_ENABLED=0 -ldflags="-w -s" -o NOME_DO_BIN_SERA_GERADO . 
-w -s = remove os simbolos e DWARF da compilação
DWARF = Debugging with arbitrary record format
- recurso para produção
- CGO_ENABLED explicado abaixo.

>> Gerar imagem docker
- docker build -t renebizelli_aula_deploy:latest -f Dockerfile .
-- -f só se o Dockerfile tiver um nome diferente, tipo Dockerfile.prod

>> rodar container 
- docker run --rm  -p 8080:8080 NOMEDAIMAGEM
-- -rm: matar container quando parar de rodar NOMEDAIMAGEM

FROM scratch ->> imagem docker que não tem nada, é vazia.

****************************************
erro ao rodar docker otimizado abaixo:
exec ./server: no such file or directory
****************************************
FROM golang:latest AS builder
WORKDIR /app
COPY . . 
RUN GOOS=linux  go build -ldflags="-w -s" -o server .

FROM scratch 
COPY --from=builder /app/server .
CMD ["./server"]
*****************************************
um dos recursos do Go é poder importar libs da linguagem C.
esse recurso é chamado de "C GO", ativado por padrão.

No scratch acima, não existe a biblioteca C instalada.
Para parar o erro acima, devemos desativar o C GO com  CGO_ENABLED=0 no build do GO
- GOOS=windwos|linux CGO_ENABLED=0 go build -ldflags="-w -s" -o NOME_DO_BIN_SERA_GERADO
*****************************************

>> criar a imagem:
docker build -t renebizelli/renebizelli_aula_deploy:latest .

>> publicar a image do docker registry
docker push renebizelli/renebizelli_aula_deploy:latest

## K8S

kind: Para rodar k8s na máquina local
https://kind.sigs.k8s.io/

1 - criar cluster: kind create cluster --name=goexpert
2 - entrar no cluster: kubectl cluster-info --context kind-goexpert
3 - listar nodes do cluster: kubectl get nodes
4 - Criar pasta k8s
5 - criar arquivo deployment.yaml - nesse arquivo é descrito tudo o que é preciso para o container rodar
6 - rodar deployment: kubectl apply -f k8s/deployment.yaml
7 - listar pods: kubectl get pods
8 - adicinar prop replicas no deployment.yaml
9 - listar os pods novamente: kubectl get pods 
10 - deletar um pod: kubectl delete pod NOME_DO_POD

Service: é responsável apontar quais os pods serão usados, e ele faz o loadbalance.
Ou seja, nunca se acessa um Pod, é o service que o acessa.
Em spec/selector/app/server  é indicado o nome do Pod (server) que ele vai gerenciar
Nas portas:
      port: 80 >> porta do service
      targetPort: 8080 >> porta onde o service vai bater nos pods criados

11 - rodar deployment: kubectl apply -f k8s/service.yaml
12 - kubectl get services

13 - Para testar, configurar: kubectl port-forward svc/servicesvc 80:8080
- quando acessar 80 local, vai direcionar para 8080 do svc/servicesvc


Probes:
- São verificações para garantir que o POD está pronto, e o container está pronto para receber requisições.
e para verificar se o container está no ar.

        startupProbe:  # para iniciar
          httpGet: 
            path: /health
            port: 8080
          periodSeconds: 10  # verifica a cada 10s
          failureThreshold: 10 # 10 tentativas antes de considerar falha

