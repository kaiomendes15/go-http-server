# 🌐 Go HTTP Server — Servidor HTTP com Sockets TCP

Atividade prática (T310) de **Redes de Computadores**: um servidor HTTP escrito
em Go usando **apenas sockets TCP** (pacote `net`), sem o pacote `net/http`.

O navegador abre uma conexão TCP com o servidor e envia uma requisição HTTP em
texto. O servidor lê os bytes crus, interpreta a request line e os headers,
localiza o arquivo pedido em `./public/` e devolve uma resposta HTTP completa
(status line, headers, linha em branco e body) pela mesma conexão. Cada cliente
é atendido em sua própria goroutine.

O servidor também recebe o formulário de `public/form.html` por POST: lê o body
pelo `Content-Length`, decodifica os campos `x-www-form-urlencoded` e responde
com uma página HTML montada em tempo de execução com os dados recebidos.

## 🛠️ Tecnologias

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![TCP Sockets](https://img.shields.io/badge/TCP-Sockets-4479A1?style=for-the-badge&logo=cloudflare&logoColor=white)
![HTTP/1.1](https://img.shields.io/badge/HTTP-1.1-E34F26?style=for-the-badge)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

- **Go (`net`)** — `net.Listen` e `Accept` para receber as conexões TCP
- **Goroutines** — uma conexão por goroutine, vários clientes ao mesmo tempo
- **`bufio.Reader`** — leitura da requisição linha a linha, independente de quantos pacotes TCP chegaram
- **`strings` / `strconv` / `fmt.Sscanf`** — parsing da request line e do `Content-Length`
- **`os.ReadFile` / `path/filepath`** — leitura dos arquivos de `./public/` como bytes
- **`io.ReadFull`** — leitura do body do POST com exatamente `Content-Length` bytes
- **`net/url` (`QueryUnescape`)** — decodificação dos campos do formulário (`+` e `%XX`)
- **`strings.Builder`** — montagem da página HTML dinâmica da resposta ao POST
- **`flag`** — porta configurável pela linha de comando

## 📋 Requisitos

| Req. | Descrição | Status |
|------|-----------|--------|
| R1 | Servidor TCP com conexões simultâneas (goroutines) | ✅ |
| R2 | Parsing da requisição (método, path, headers, `Content-Length`) | ✅ |
| R3 | Servir arquivos estáticos (HTML, CSS e imagens) | ✅ |
| R4 | Página de erro 404 | ✅ |
| R5 | Processamento de requisições POST (`x-www-form-urlencoded`) | ✅ |
| R6 | Resposta dinâmica ao POST | ✅ |

## 📡 Protocolo

O navegador envia uma requisição como esta:

```
GET /style.css HTTP/1.1\r\n
Host: localhost:9000\r\n
\r\n
```

E o servidor responde com:

```
HTTP/1.1 200 OK\r\n
Content-Type: text/css\r\n
Content-Length: 37\r\n
\r\n
[bytes do arquivo]
```

O path `/` é tratado como `/index.html`. O `Content-Type` é escolhido pela
extensão do arquivo:

| Extensão | Content-Type |
|----------|--------------|
| `.html`  | `text/html`  |
| `.css`   | `text/css`   |
| `.png`   | `image/png`  |
| `.jpg`   | `image/jpeg` |

### POST do formulário

Ao enviar o formulário, o navegador manda os campos no body, codificados como
`application/x-www-form-urlencoded`:

```
POST /submit HTTP/1.1\r\n
Content-Type: application/x-www-form-urlencoded\r\n
Content-Length: 48\r\n
\r\n
nome=Jo%C3%A3o+Silva&email=kaio%40email.com&msg=
```

O servidor lê exatamente `Content-Length` bytes do body, separa os pares por `&`,
divide cada par no primeiro `=` e só então decodifica chave e valor (`+` vira
espaço e `%XX` vira o byte correspondente). Separar antes de decodificar evita
que um `%26` decodificado vire um `&` e quebre os pares. Os campos ficam em um
slice, na ordem do formulário, e viram uma página com um `<li>` por campo:

```
HTTP/1.1 200 OK\r\n
Content-Type: text/html\r\n
Content-Length: 240\r\n
\r\n
<html lang='pt-br'>
<head>
<meta charset='UTF-8'>
</head>
<body>
<h1>Dados Recebidos</h1>
<ul>
<li><strong>nome:</strong> João Silva</li>
...
```

O `Content-Length` da resposta é o tamanho do HTML gerado em **bytes**, e não
em caracteres: cada letra acentuada ocupa 2 bytes em UTF-8.

### Tratamento de erros

Nos casos abaixo, o servidor registra o erro no terminal e fecha a conexão:

- **Request line malformada** → `[ERRO] linha inicial do HTTP malformada`
- **`Content-Length` inválido ou negativo** → `[ERRO] Content-Length inválido`
- **Extensão fora da tabela** → `[ERRO] formato de arquivo inválido.`
- **Body do POST menor que o `Content-Length`** → `[ERRO] Nenhum ou menos bytes que o Content-Length foram lidos.`
- **Sequência `%XX` inválida no body** (ex.: `%ZZ`) → `[ERRO] Formato inválido para decode no body: ...`

Quando o arquivo pedido não existe, o servidor responde `HTTP/1.1 404 Not Found` com a página `public/not_found.html` e registra `[404] Recurso não encontrado: /nao-existe.html`.

## 📁 Estrutura

```
go-http-server/
├── main.go          # servidor HTTP (porta padrão 9000)
├── public/          # arquivos estáticos servidos
│   ├── index.html
│   ├── about.html
│   ├── form.html        # formulário do POST /submit
│   ├── style.css
│   ├── not_found.html   # página do 404
│   ├── neymar.jpg
│   └── cris.jpg
└── docs/
    ├── T310 - Servidor HTTP (1).md   # enunciado
    └── uso-de-ia.md                  # diário de uso de IA
```

## ▶️ Como executar

Pré-requisito: **Go 1.22+** instalado.

### 1. Suba o servidor

Na raiz do repositório (o servidor procura os arquivos em `./public/`):

```bash
go run main.go
```

Para usar outra porta:

```bash
go run main.go -porta 8080
```

Uma porta fora do intervalo 1–65535 faz o servidor usar a porta 9000.

Saída esperada:

```
[SERVIDOR] Escutando na porta :9000
```

### 2. Acesse pelo navegador

Abra `http://localhost:9000/`. A página aparece com o CSS aplicado e a imagem
visível. Como o HTML referencia o CSS e a imagem, o navegador faz uma
requisição para cada arquivo, cada uma em sua própria conexão:

```
[CONEXÃO] Nova conexão de [::1]:51762
[REQUEST] Método: GET | Path: /index.html | Content-Length: 0
[ARQUIVO] Servindo /index.html (text/html, 336 bytes)
[CONEXÃO] Encerrada [::1]:51762
[CONEXÃO] Nova conexão de [::1]:51774
[REQUEST] Método: GET | Path: /style.css | Content-Length: 0
[ARQUIVO] Servindo /style.css (text/css, 37 bytes)
[CONEXÃO] Encerrada [::1]:51774
[CONEXÃO] Nova conexão de [::1]:51786
[REQUEST] Método: GET | Path: /neymar.jpg | Content-Length: 0
[ARQUIVO] Servindo /neymar.jpg (image/jpeg, 31294 bytes)
[CONEXÃO] Encerrada [::1]:51786
```

### 3. Envie o formulário

Abra `http://localhost:9000/form.html`, preencha e envie. O navegador mostra a
página "Dados Recebidos" com os campos, e o terminal registra:

```
[REQUEST] Método: POST | Path: /submit | Content-Length: 48
[POST] Body bruto (48 bytes): nome=Jo%C3%A3o+Silva&email=kaio%40email.com&msg=
[POST] Campo: nome = João Silva
[POST] Campo: email = kaio@email.com
[POST] Campo: msg = 
[RESPONSE] Resposta dinâmica enviada (text/html, 240 bytes)
```

O body bruto mostra os valores como o navegador os enviou (`+` e `%XX`), e cada
`[POST] Campo` mostra o valor já decodificado.

### 4. Ou teste pelo terminal

```bash
curl -v http://localhost:9000/style.css
curl -o out.jpg http://localhost:9000/neymar.jpg && cmp out.jpg public/neymar.jpg
curl -v -d 'nome=Jo%C3%A3o+Silva&curso=Programa%C3%A7%C3%A3o' http://localhost:9000/submit
```

O `cmp` não mostra diferença: a imagem chega byte a byte igual ao arquivo.
Você pode abrir várias abas ao mesmo tempo, e o servidor atende todas
simultaneamente.
