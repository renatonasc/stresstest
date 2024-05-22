# Stress Test

Stress Test é um sistema escrito em Go, projetado para executar testes de carga em serviços web. Ele fornece insights críticos sobre o desempenho e a robustez desses serviços, ajudando a identificar pontos fracos e a otimizar a capacidade de resposta sob carga pesada.

## Recursos

- **Testes de carga**: Gere uma grande quantidade de tráfego para o seu serviço web para testar sua capacidade de lidar com cargas pesadas.
- **Análise de desempenho**: Obtenha insights detalhados sobre o desempenho do seu serviço web sob diferentes níveis de carga.
- **Robustez**: Verifique a robustez do seu serviço web, identificando pontos de falha e áreas que precisam de melhorias.

## Instalação

Para instalar o Stress Test, você precisa ter o Go instalado em sua máquina. Você pode baixar o Go [aqui](https://golang.org/dl/).

Uma vez que o Go esteja instalado, você pode baixar e instalar o Stress Test com o seguinte comando:

```bash
go get github.com/renatonasc/stresstest
```

## Uso

Para executar um teste de carga em um serviço web, use o seguinte comando:

Com Docker:
```bash
docker build -t stresstest .

docker run stresstest3 --url=http://cameras.renatonasc.com --concurrency=20 --requests=100 
```

Este comando irá gerar 1000 solicitações para `http://seuservico.com` com uma concorrência de 20 solicitações simultâneas.

