## Aula 1 
### Code bank - estudo de caso
- O que iremos desenvolver?  
  - Um banco digital que disponibiliza cartões de crédito para clientes
  - O cliente realizará o checkout de um produto 
  - Processamento da transação
  - Caso seja aprovada, registrar no extrato do cliente 
  - Desenvolver um dashboard para as métricas das transações
- Pontos importantes
  - Tudo é uma simulação para usar microsserviços 
  - Não haverá validação de seguraça 
  - Os sistemas poderão se comunicar de forma síncrona ou assíncrona
- Desafios 
  - O banco precisa de alta performance na comunicação então no lugar do padrão REST, usaremos a solução com gRPC(permite que a comunicação seja veloz transitando o payload da requisição em binário utilizando o protocolo HTTP2)
  - Garantir resiliência entre a sicronização do processamento das transações com o serviço do extrato bancário 
  - Usaremos o Kafka Connect para consumir os dados de cada transação e inserir no Elasticsearch
- Dinâmica do sistema
  - TODO: colocar a imagem aqui depois
- Tecnologias utilizadas: 
  - Bank: Golang
  - Checkout e Extrato backend: Nestjs
  - Checkout e Extrato frontend: Nextjs
  - Kafka e Kafka Connect
  - Elasticsearch e Kibana
  - Docker e Kubernets
  - Istio, Kiali, Prometheus e Grafana