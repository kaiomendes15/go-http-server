# Diário de Uso de Inteligência Artificial

Projeto: T310 · Servidor HTTP em Go (somente pacote `net`)
Disciplina: Redes de Computadores
Autor: Kaio Mendes (trabalho individual)

Este diário registra todas as interações com ferramentas de IA durante o projeto e serve de base para a Seção 4 da Declaração de Uso de IA.

**Diretriz adotada:** a IA é usada apenas para consulta e organização (dúvidas pontuais sobre a linguagem Go e sobre comunicação em rede, além do planejamento das tarefas). Ela não gera o código-fonte do servidor, e só delego a ela o que já compreendo.

---

## Modelo de entrada

```
### AAAA-MM-DD HH:MM · <ferramenta>
- Requisito relacionado: R1–R6 / organização / entrega
- Nível de uso (guia institucional): <1 | 2 | 3>
- Pergunta ou solicitação:
- O que a IA respondeu ou fez:
- O que aprendi ou aproveitei:
- O que NÃO foi usado:
```

---

## Registros

### 2026-10-30 ~17:20 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** organização e planejamento (todos)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta lesse o enunciado do trabalho (`docs/T310 - Servidor HTTP (1).md`) para entender o contexto. Informei que ela seria usada só para consultas pontuais sobre Go e redes, sem geração de código.
- **O que a IA respondeu ou fez:** resumiu o enunciado: os seis requisitos (R1–R6), os dois ramos de dependência (R1→R2→R3→R4 e R1→R2→R5→R6), os itens obrigatórios da entrega (.zip com código, `./public/`, PDF de divisão de requisitos e Declaração de IA assinada) e as penalidades. Também sugeriu manter este diário desde o início, porque uma Seção 4 genérica zera a nota.
- **O que aprendi ou aproveitei:** a visão geral dos requisitos e das dependências entre eles, e a necessidade de registrar o uso de IA desde o primeiro dia.
- **O que NÃO foi usado:** nenhum código foi solicitado ou gerado.

### 2026-10-30 ~17:25 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** organização e planejamento (todos)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta criasse, pela integração com o Trello, um quadro kanban do projeto. As atividades deveriam ser divididas em entregáveis pequenos, com prazo em cada card e as dependências de bloqueio descritas no enunciado, considerando a entrega em 02/11. Informei que faria o trabalho sozinho e com esforço mínimo na parte de HTML/CSS/imagens.
- **O que a IA respondeu ou fez:**
  - Criou o quadro "T310 · Servidor HTTP em Go (net puro)" com as listas Guia, Bloqueado, A Fazer, Em Andamento, Validação e Concluído.
  - Criou 40 cards: setup (S), assets mínimos (A), subtarefas de cada requisito (R1.1 a R6.4), integração e ensaio da defesa (I) e entrega (E). Cada card tem prazo (horário de Brasília), critério de pronto e os campos "Bloqueado por" e "Desbloqueia", com links para os cards relacionados.
  - Nas descrições dos cards, incluiu uma seção "Para dominar na defesa" com perguntas conceituais para estudo. Elas citam nomes de recursos da biblioteca padrão do Go como pontos de pesquisa, por exemplo `net.Listen`, `bufio.Reader`, `io.ReadFull`, `strconv.Atoi`, `filepath.Clean`, `errors.Is`, `url.QueryUnescape`, `html.EscapeString` e `strings.Builder`. Também apontou cuidados conceituais: TCP é um fluxo de bytes, Content-Length conta bytes e não caracteres (UTF-8), e há riscos de path traversal e de XSS.
- **O que aprendi ou aproveitei:** a divisão do trabalho em tarefas pequenas na ordem de dependência, e uma lista de conceitos a estudar e saber explicar na defesa oral.
- **O que NÃO foi usado:** nenhum código foi solicitado ou gerado. Os cards descrevem objetivos e perguntas de estudo, não implementações.

### 2026-10-30 ~18:00 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** setup (card S0.1)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** "Como inicio o módulo Go?"
- **O que a IA respondeu ou fez:** explicou o comando `go mod init <caminho>` e como escolher o caminho do módulo (endereço do repositório x nome local), o conteúdo do `go.mod` gerado (diretivas `module` e `go`) e por que não há `require` nem `go.sum` quando só a biblioteca padrão é usada. Também explicou o papel do `package main` com `func main()`, o que o `go run .` compila e que o Git não versiona diretórios vazios.
- **O que aprendi ou aproveitei:** executei o `go mod init go-http-server` por conta própria. Entendi o que a diretiva `go` controla (versão mínima e semântica da linguagem).
- **O que NÃO foi usado:** nenhum código do projeto foi gerado.

### 2026-10-30 ~18:05 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** setup (card S0.1)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi um `.gitignore` simples para projetos Go.
- **O que a IA respondeu ou fez:** criou o arquivo `.gitignore` (binário do `go build`, artefatos de teste, `go.work`, `.env`, arquivos de editor e do SO, `*.zip` da entrega) e explicou cada grupo de regras.
- **O que aprendi ou aproveitei:** o `.gitignore` foi usado como gerado pela IA. É um arquivo de configuração do repositório, não código do servidor.
- **O que NÃO foi usado:** nenhum código do servidor foi gerado.

### 2026-10-30 ~18:20 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** organização do repositório
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta fizesse os commits do que existia no repositório, com mensagens específicas ao contexto da entrega.
- **O que a IA respondeu ou fez:** conferiu o `main.go` que eu havia escrito (rodou `go vet` para verificar a compilação, sem alterar o arquivo) e criou 4 commits separados por assunto: setup (`go.mod` e `.gitignore`), enunciado, diário de IA e a primeira versão do servidor TCP (R1). Marcou com `Co-Authored-By: Claude` só os commits com conteúdo escrito pela IA (`.gitignore` e este diário), e não o commit do `main.go`, que é de minha autoria.
- **O que aprendi ou aproveitei:** organização do histórico de commits por assunto.
- **O que NÃO foi usado:** o `main.go` não foi alterado pela IA.

### 2026-10-30 ~18:22 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta movesse o card S0.1 para Concluído e o R1.1 para A Fazer.
- **O que a IA respondeu ou fez:** moveu os cards. Apontou que o critério "`./public/` existe" do S0.1 ainda estava pendente (ficou para o card A.1).
- **O que NÃO foi usado:** nenhum código foi gerado.

### 2026-10-30 ~18:25 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R1 (card R1.1)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** argumentei que a porta configurável por linha de comando não constava no enunciado e que eu poderia deixá-la fixa no código, considerando o R1.1 concluído.
- **O que a IA respondeu ou fez:** mostrou que o requisito consta no enunciado, no resumo do R1 ("Escutar em uma porta configurável via argumento de linha de comando"), e que a seção detalhada cita 9000 apenas como exemplo. O card R1.1 permaneceu em A Fazer.
- **O que aprendi ou aproveitei:** corrigi minha leitura do enunciado. A porta deve ser configurável, com 9000 como valor padrão.
- **O que NÃO foi usado:** nenhum código foi gerado.

### 2026-10-30 ~18:30 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R1 (card R1.1)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** "Explique a diferença entre `os.Args` e `flag`."
- **O que a IA respondeu ou fez:** explicou que `os.Args` é um `[]string` cru e posicional (o índice 0 é o executável, risco de panic sem checar `len`, conversão e valor padrão manuais) e que o pacote `flag` declara opções nomeadas com tipo, valor padrão e descrição, preenchidas por `flag.Parse()`, com validação de tipo e ajuda `-h` automáticas. Explicou também por que `flag.Int` retorna um ponteiro e que a checagem da faixa de porta é responsabilidade do programador nos dois casos. Mostrou um trecho **genérico e ilustrativo** do padrão de uso (`n := flag.Int("nome", 10, "descrição")`, `flag.Parse()`, `fmt.Println(*n)`), com comentários explicativos.
- **O que aprendi ou aproveitei:** escolhi o pacote `flag` para o R1.1. **Adaptei no `main.go` o padrão de declaração do trecho ilustrativo** (`flag.Int` + `flag.Parse()`), mantendo os comentários explicativos da resposta. O nome da flag, o valor padrão e a lógica de validação da porta são meus.
- **O que NÃO foi usado:** a IA não escreveu a lógica do R1.1 (validação, fallback, montagem do endereço).

### 2026-10-30 ~18:38 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R1 (card R1.1)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** "As portas que eu posso ter, tanto no Linux quanto no Windows, são de 1 a 65535?"
- **O que a IA respondeu ou fez:** explicou que o campo de porta do TCP tem 16 bits (0 a 65535) e que a porta 0 significa "o SO escolhe". Explicou as diferenças entre os sistemas: no Linux, portas abaixo de 1024 exigem privilégio e a faixa efêmera é 32768–60999; no Windows não há portas privilegiadas, mas existem faixas reservadas pelo Hyper-V/WSL. Apresentou a classificação da IANA (well-known, registered, dynamic) e o comportamento no WSL2.
- **O que aprendi ou aproveitei:** usei a faixa 1–65535 como critério de validação da porta. Entendi que uma porta válida ainda pode falhar no `Listen` por permissão ou por já estar em uso.
- **O que NÃO foi usado:** nenhum código foi gerado.

### 2026-10-30 ~18:45 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** organização do repositório
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta registrasse as consultas neste diário e commitasse as alterações.
- **O que a IA respondeu ou fez:** escreveu estas entradas e commitou o diário e o meu `main.go` em andamento. Antes do commit, apontou que a condição de validação da porta no meu código estava invertida (portas válidas caíam no fallback). A correção ficou a meu cargo.
- **O que NÃO foi usado:** a IA não corrigiu o código.

### 2026-10-31 ~11:00 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R1 (cards R1.1 a R1.5) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta lesse os cards do quadro no Trello, comparasse com o que já estava implementado no `main.go` e definisse o que já estava concluído.
- **O que a IA respondeu ou fez:**
  - Leu os 40 cards e comparou o critério de pronto de cada um com o código e o histórico de commits.
  - Compilou o meu `main.go` em uma pasta temporária, fora do repositório, e rodou testes de verificação sem alterar o código. Testou: a flag `-porta 8080`, a porta fora da faixa (70000, que mostra a mensagem e usa a 9000), um valor não numérico (`abc`, recusado pelo pacote `flag`), porta já em uso (`bind: address already in use`), a porta em LISTEN no `ss -ltn` e 3 conexões TCP simultâneas com logs de abertura e encerramento.
  - Concluiu que S0.2, R1.1, R1.2, R1.3 e R1.4 atendem ao critério de pronto. O R1.5 não atende: um `curl` esperou 2 s sem resposta, porque o servidor ainda não envia nenhuma resposta HTTP, e o handler lê até o cliente fechar a conexão. Também apontou que a pasta `./public/` (card A.1) ainda não existe.
  - Apontou três pontos para eu revisar: um `fmt.Println(*n)` que sobrou de depuração, o endereço IPv6 `[::1]` que aparece no log quando o cliente usa `localhost`, e uma alteração só de espaços em branco ainda não commitada no `main.go`.
- **O que aprendi ou aproveitei:** a confirmação de que o R1 funciona, exceto o teste no navegador, e de que o navegador só "recebe resposta" quando o servidor escreve uma resposta HTTP na conexão. Os pontos de revisão ficaram para eu decidir e corrigir.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado. Os testes só executaram o binário compilado do meu código.

### 2026-10-31 ~11:08 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** a primeira tentativa de mover os cards foi parcialmente bloqueada pela permissão do Claude Code. Perguntei se R1.1, R1.4 e R1.5 deveriam ir para Concluído e como liberar a ferramenta de escrita do Trello.
- **O que a IA respondeu ou fez:** confirmou R1.1 e R1.4 e explicou que o R1.5 não cumpre o critério "3 abas carregam", porque o servidor ainda não responde. Explicou como liberar a ferramenta pelo comando `/permissions` ou pelo arquivo `.claude/settings.local.json`. Depois que liberei a ferramenta e decidi que o R1.5 iria para A Fazer, moveu S0.2, R1.1, R1.2, R1.3 e R1.4 para Concluído e os marcou como completos. Moveu R1.5 e R2.1 para A Fazer.
- **O que aprendi ou aproveitei:** a diferença entre validar só os logs e cumprir o caso de teste do enunciado. O R1.5 será fechado quando o R3 servir o `index.html`, ou na regressão do I.2.
- **O que NÃO foi usado:** nenhum código foi gerado.

### 2026-10-31 ~11:15 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R1 (card R1.5)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** executei o teste do R1.5 no navegador e compartilhei o log do terminal. Perguntei por que só consegui uma conexão: as outras abas normais na mesma URL não conectavam, e só uma aba anônima abria conexão nova. Perguntei também por que o navegador às vezes abre mais de uma conexão por aba.
- **O que a IA respondeu ou fez:**
  - Explicou que quem segurou as abas foi o navegador, não o servidor (o teste anterior já tinha aceitado 3 conexões simultâneas). Navegadores baseados no Chromium têm um **cache lock**: enquanto um GET para uma URL ainda não recebeu resposta, os pedidos idênticos para a mesma URL no mesmo perfil ficam em fila. Como o meu servidor ainda não responde, o 1º pedido nunca termina e as outras abas nem abrem conexão TCP. A aba anônima usa outro perfil, com cache e pool de conexões próprios, e por isso conectou na hora.
  - Sugeriu confirmar pelo DevTools (guia Network: pedido *pending*, tempo em *Stalled*/*Queueing*) e testar antes do R3 com URLs diferentes por aba (`/?1`, `/?2`, `/?3`).
  - Listou os motivos para mais de uma conexão por aba: pré-conexão especulativa ao digitar na barra de endereço, o pedido automático de `/favicon.ico` (no HTTP/1.1 uma conexão atende um pedido por vez), recursos em paralelo (até ~6 conexões por host), `localhost` resolvendo para `::1` e `127.0.0.1` (*Happy Eyeballs*, o que explica o `[::1]` no log) e o keep-alive, que atrasa o log "Encerrada".
  - Observou que o "Encerrada" só aparece ao fechar a aba porque o handler fica no laço de `Read` até receber `EOF`, e que no R2/R3 terei de decidir quando a conexão termina (fechar após a resposta ou manter keep-alive).
- **O que aprendi ou aproveitei:** o comportamento das abas é do navegador (cache lock e perfis separados), e não um limite do meu servidor. Também entendi as origens das conexões extras, que preciso saber explicar na defesa do R1.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~11:30 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R2 (card R2.1)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** "Só do navegador tentar acessar a URL do meu servidor, eu já consigo capturar essa requisição no lado do servidor?"
- **O que a IA respondeu ou fez:**
  - Explicou que sim: ao acessar a URL, o navegador abre a conexão TCP (é quando o `Accept` retorna) e escreve nela a requisição HTTP em texto, que chega ao `buf` pelo `conn.Read`. O HTTP é só texto trafegando sobre a conexão TCP do R1, sem configuração extra.
  - Apontou que o meu handler descarta o primeiro retorno do `Read` (a quantidade de bytes lidos) e nunca examina o `buf`. Sugeriu, como passo de estudo, exibir só a parte preenchida do buffer para ver o formato real da requisição antes de fazer o parsing.
  - Avisou o que pode confundir no teste: o pedido automático de `/favicon.ico`, conexões especulativas sem requisição, o navegador "carregando" indefinidamente (o servidor ainda não responde e a conexão fica em keep-alive) e a quantidade de headers além do exemplo do enunciado.
  - Explicou que o TCP é um fluxo de bytes, então um único `Read` não garante a requisição inteira. O fim do cabeçalho deve ser identificado pelo `\r\n\r\n`, ponto que volta no R5 com o `Content-Length`.
  - Sugeriu testar com `curl -v` ou `nc` para controlar exatamente o que é enviado.
- **O que aprendi ou aproveitei:** a requisição já chega ao servidor desde o R1. O trabalho do R2 é ler e interpretar esse texto, sem depender de um `Read` corresponder a uma requisição.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~11:35 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R2 (card R2.1)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** "Com o `Read`, eu leio toda a requisição, mas não tenho garantia de que o formato está correto. Podem haver partes faltando ou quebradas na leitura, correto?"
- **O que a IA respondeu ou fez:** separou o problema em três casos:
  - **Bytes corrompidos, perdidos ou fora de ordem:** não acontece. O TCP garante a integridade e a ordem (checksum, número de sequência, ACK e retransmissão) antes de entregar os dados à aplicação. Essa é a diferença em relação ao UDP.
  - **Requisição incompleta:** pode acontecer. O `Read` devolve o que já chegou, que pode ser só uma parte. O resto vem nos próximos `Read`s. O tamanho do buffer (1024 bytes) também limita cada leitura. O critério de cabeçalho completo é encontrar o `\r\n\r\n`.
  - **Formato inválido:** pode acontecer e é responsabilidade da aplicação. O TCP entrega exatamente o que o cliente enviou, mas não sabe nada de HTTP. Exemplos: request line sem as três partes, `nc` sem `-C` enviando só `\n`, método não suportado. A resposta HTTP para esse caso é `400 Bad Request`, e o parser não deve travar nem derrubar a goroutine com uma entrada inesperada.
- **O que aprendi ou aproveitei:** o TCP garante a integridade dos bytes. Cabe ao meu código garantir a completude (juntar as leituras até o `\r\n\r\n`) e a validade (conferir se o texto é HTTP bem formado).
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~11:45 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R2 (card R2.1), com impacto no R5
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** "Considerando o parser, usar o `Read` para preencher o buffer e iterar sobre os itens no buffer para verificar a sua estrutura é realmente a forma mais eficiente, ou existe alguma outra alternativa?"
- **O que a IA respondeu ou fez:**
  - Explicou que percorrer o buffer manualmente não é lento (qualquer abordagem passa pelos bytes, e o gargalo é a rede), mas obriga a tratar na mão o acúmulo de leituras até o `\r\n\r\n`, as linhas cortadas entre dois `Read`s e a "sobra" de bytes do corpo lida junto com o cabeçalho.
  - Apresentou como alternativa o `bufio.Reader` em volta do `conn`: `ReadString('\n')` devolve linhas completas independentemente da fragmentação TCP, a leitura em blocos reduz as chamadas de sistema, e `io.ReadFull` lê exatamente `Content-Length` bytes do corpo no R5. Destacou que, depois de criado o `bufio.Reader`, toda leitura deve passar por ele, porque ele lê adiantado e parte do corpo pode estar no buffer interno.
  - Explicou por que não usar o `bufio.Scanner` (não lê "exatamente N bytes" com naturalidade e também lê adiantado) nem o `net/textproto` (`ReadMIMEHeader` é a peça do `net/http` que trata o protocolo, fora do espírito do enunciado). Avaliou que o `bufio` é permitido pelo mesmo critério que liberou o `net/url`: não trata do protocolo HTTP.
  - Indicou funções do pacote `strings` para quebrar as linhas (`Fields`/`Split` para a request line, `Cut` para separar `Chave: Valor` na primeira ocorrência de `:`, `TrimSpace`/`TrimSuffix` para o fim de linha) e sugeriu pesquisar o retorno de `ReadString` quando a conexão fecha no meio de uma linha.
- **O que aprendi ou aproveitei:** a escolha entre leitura manual e leitura bufferizada é de simplicidade e correção, não de desempenho. Também entendi o risco de perder bytes do corpo ao misturar leitores diferentes sobre a mesma conexão.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~12:00 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R2 (card R2.1), com impacto no R5
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** depois de implementar a leitura da request line, perguntei se precisava ler todas as informações do cabeçalho ou apenas o `Content-Length`.
- **O que a IA respondeu ou fez:**
  - Diferenciou **ler** de **interpretar**: todas as linhas do cabeçalho precisam ser consumidas até a linha vazia, porque elas ficam no fluxo entre a request line e o corpo. Se o laço parar no `Content-Length`, as linhas restantes seriam lidas como corpo no R5.
  - Mostrou que a seção detalhada do R2 pede só método, path e `Content-Length` (quando presente), e que o resumo do enunciado fala em "headers". Guardar todos num `map` ou extrair só o `Content-Length` são escolhas defensáveis.
  - Listou cuidados com o `Content-Length`: o nome do header não diferencia maiúsculas e minúsculas (`strings.EqualFold` ou `strings.ToLower`), o espaço depois dos dois-pontos, o valor padrão 0 quando o header não vem (GET) e o valor inválido como caso de `400 Bad Request`.
  - Observou que a linha vazia chega como `"\r\n"` (ou `"\n"` pelo `nc`), então a condição de parada deve ser testada depois de limpar o fim da linha, e que o laço deve tratar a conexão fechando antes da linha vazia.
- **O que aprendi ou aproveitei:** implementei o laço de leitura dos headers com parada na linha vazia, depois de limpar o fim da linha.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~12:30 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R2 (card R2.1)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** compartilhei o log do terminal: a request line aparecia, mas o `Content-Length` não era impresso. Perguntei se, mesmo no GET, o `Content-Length` vem preenchido.
- **O que a IA respondeu ou fez:**
  - Corrigiu a premissa: num GET o navegador normalmente **não envia** o header `Content-Length`, porque não há corpo. O 0 deve vir do valor padrão.
  - Leu o meu `main.go` (sem alterá-lo) e identificou a causa real. Eu passava `*reader` para funções que recebiam `bufio.Reader` por valor, então cada função trabalhava sobre uma **cópia** da struct. A cópia de `getRequestLine` leu a requisição inteira do `conn` para o buffer, mas o avanço ficou só nela. A cópia em `getContentLenght` via o buffer vazio e ficava bloqueada num novo `Read` no `conn`, porque o navegador já tinha enviado tudo e esperava a resposta. O sinal disso era que nenhum header era impresso.
  - Sugeriu confirmar fechando a aba (o `Read` bloqueado retorna EOF) e apontou outros pontos para eu revisar: o `Sscanf` com o formato literal diferenciando maiúsculas e minúsculas e com o erro ignorado, a falta de `return` depois do erro da request line, a falta de `\n` no `Printf`, o formato do log `[REQUEST]` do enunciado e o redirecionamento de `/` para `/index.html`.
- **O que aprendi ou aproveitei:** num GET o `Content-Length` costuma não existir. Copiar um `bufio.Reader` divide o estado de leitura sem dar erro de compilação.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~13:30 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R2 (card R2.1), com impacto no R5
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** testei um POST pelo Postman com corpo JSON (*raw*), e o log mostrava `Content-Length: 0`. Perguntei se o problema era a estrutura do corpo ou o meu parse.
- **O que a IA respondeu ou fez:**
  - Leu o meu `main.go` (sem alterá-lo) e mostrou que o problema era o parse. Eu tinha aplicado `strings.ToLower` ao **formato** do `Sscanf` (`"content-length: %d"`), e não à linha recebida (`Content-Length: 72`). O `Sscanf` falhava no primeiro caractere e, como o erro era ignorado, a variável ficava em 0 sem aviso. O corpo JSON não tinha relação, porque o Postman calcula e envia o `Content-Length` automaticamente.
  - Ofereceu duas soluções: a mínima (aplicar o `ToLower` à linha e verificar o erro do `Sscanf`) e a recomendada (`strings.Cut` + `TrimSpace` + `strconv.Atoi` com verificação de erro).
  - Explicou que o `Erro ao ler a request line: EOF` vinha de uma conexão aberta e fechada pelo cliente sem enviar dados, e que o Postman acusaria erro porque o servidor fecha a conexão sem ler o corpo nem responder (R5, R3 e R6). Lembrou que o formulário do R5 usa `application/x-www-form-urlencoded`, e não JSON.
- **O que aprendi ou aproveitei:** apliquei a solução mínima por conta própria (`ToLower` na linha, com o formato em minúsculas). Também aprendi que um erro ignorado vira um zero silencioso, e que o retorno `(n, err)` do `Sscanf` teria mostrado a falha.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado. A solução recomendada (`Cut` + `Atoi`) não foi adotada por enquanto.

### 2026-10-31 ~13:55 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R2 (cards R2.1 a R2.4) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta lesse os cards do quadro no Trello e dissesse o que já estava concluído, considerando o que já estava implementado.
- **O que a IA respondeu ou fez:**
  - Leu os 40 cards e o meu `main.go` (sem alterá-lo). Compilou o código numa pasta temporária, fora do repositório, e testou: `curl` em `/about.html`, `/` e um POST com `-d "a=1"`, e `nc` com uma request line malformada (`LIXO`), com `content-length: -5` e com `Content-Length:7` (sem espaço).
  - Confirmou que S0.1, S0.2 e R1.1 a R1.5 estão concluídos (a pasta `./public/` do S0.1 ainda não existe). Avaliou que o R2.1 cumpre o critério de pronto, que o R2.2 e o R2.3 estavam parciais e que o R2.4 (normalização de `/` para `/index.html`) não estava feito.
  - Apontou que um `Content-Length` negativo era aceito, que `Content-Length:7` era rejeitado e que o log `[REQUEST]` era impresso em dois `Printf`, o que deixava a mensagem de erro colada ao log.
- **O que aprendi ou aproveitei:** o estado real de cada card do R2 em relação ao critério de pronto, comprovado por testes, e não só pela leitura do código.
- **O que NÃO foi usado:** nenhum código do projeto foi gerado ou alterado. Os testes só executaram o binário compilado do meu código.

### 2026-10-31 ~14:15 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R2 (cards R2.2, R2.3 e R2.4)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que corrigi o R2.2 por conta própria, que vou deixar a validação de métodos para quando for necessária e que não encontrei no enunciado nenhuma exigência de mapa de headers.
- **O que a IA respondeu ou fez:**
  - Leu o meu `main.go` (sem alterá-lo) e confirmou que a troca para `strings.Fields` com `len == 3` resolveu os dois casos do R2.2, que a versão agora é extraída e que a checagem de negativo passou a ser feita sobre o valor do `Content-Length`.
  - Apontou que o log agora inclui `| Version: ...`, o que difere do formato exato exigido pelo enunciado e precisa ser removido até o R2.4.
  - Concordou que a seção detalhada do R2 pede só método, path e o valor do `Content-Length`, e que o mapa de headers e a struct foram acréscimos dela ao montar os cards. Observou que o resumo do enunciado menciona "headers" e sugeriu o argumento para a defesa: todos os headers são **lidos** até a linha em branco, mas só o `Content-Length` é **interpretado**.
  - Concordou em deixar a validação de método para o roteamento do R5.1.
- **O que aprendi ou aproveitei:** a distinção entre ler e interpretar os headers, como justificativa para não construir o mapa.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado. A sugestão de mapa de headers e de struct não foi adotada.

### 2026-10-31 ~14:20 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R2 (card R2.3) e documentação
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que não vou tratar os dois casos restantes do `Content-Length` e pedi que a ferramenta registrasse as interações neste diário.
- **O que a IA respondeu ou fez:** escreveu estas entradas no diário.
- **Decisão registrada:** `Content-Length:7` (sem espaço, rejeitado) e `Content-Length: 7abc` (aceito como 7) ficam como **limitações conhecidas** do parser. Nenhum dos dois casos aparece com navegador, curl ou Postman, que sempre enviam o espaço e um valor numérico válido.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~14:25 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** organização (cards R2.2 e R2.3)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta movesse o R2.2 para Concluído e ajustasse os cards R2.2 e R2.3 às minhas decisões de escopo.
- **O que a IA respondeu ou fez:** moveu o R2.2 para Concluído e o marcou como completo. Moveu o R2.3 (desbloqueado) para A Fazer. Reescreveu as descrições: tirou a struct do R2.2 e o mapa de headers do R2.3, renomeou o R2.3 para "Ler todos os headers + extrair Content-Length", acrescentou os pontos de defesa desta sessão (`Split` x `Fields`, retorno `n` do `Sscanf`, panic com tamanho negativo, ler x interpretar) e registrou no R2.3 as duas limitações conhecidas do `Content-Length`.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~14:30 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** organização e versionamento (R2)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que testei o R2.3 (incluindo o `Content-Length` negativo) e pedi que a ferramenta atualizasse os cards e commitasse tudo, sem colocar a IA como coautora nos arquivos que ela não escreveu.
- **O que a IA respondeu ou fez:** moveu o R2.3 para Concluído e o marcou como completo. Moveu o R2.4 (desbloqueado) para A Fazer. Separou o commit em dois: um com o meu `main.go` (parsing do R2), **sem** coautoria da IA, e outro com este diário, que ela escreveu, com a linha de coautoria.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~14:45 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R2 (cards R2.4 e R2.5) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que achava ter concluído o R2.4 (normalização de `/` para `/index.html`) e pedi que a ferramenta verificasse. Depois pedi que atualizasse os cards, commitasse o código e registrasse a verificação.
- **O que a IA respondeu ou fez:**
  - Leu o diff do meu `main.go` (sem alterá-lo), rodou o `go vet` (sem avisos), compilou o código numa pasta temporária e testou com `curl`: `/about.html`, `/` (o log mostrou `Path: /index.html`), um POST com `-d "a=1"` (`Content-Length: 3`), `/?x=1` e `/img/`. Confirmou que o R2.4 cumpre o critério de pronto e que os três cenários do R2.5 já passam pelo `curl`, faltando o teste no navegador.
  - Apontou para eu decidir: a query string (`/?x=1`) não é normalizada e, no R3, viraria a busca de um arquivo inexistente; e há sobras de depuração (o `fmt.Println(*n)`, que imprime a porta antes do log `[SERVIDOR]`, e um comentário com o comando `nc` no início do `main`).
  - Moveu o R2.4 para Concluído e o marcou como completo. Seguindo a regra do quadro, moveu para A Fazer os cards desbloqueados por ele: R2.5, R3.1, R3.2 e R5.1. Commitou o meu `main.go` sem coautoria da IA.
- **O que aprendi ou aproveitei:** a confirmação de que o parsing do R2 está completo, e a pendência da query string para decidir no R3.1.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~15:00 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R2 (card R2.5) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que executei no navegador o caso de teste do R2 (`/about.html` e `/`) e pedi que a ferramenta atualizasse os cards.
- **O que a IA respondeu ou fez:** moveu o R2.5 para Concluído e o marcou como completo. O I.2 continua bloqueado, porque depende dos testes de R3 a R6.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~15:05 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R3 (card R3.1) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** antes de começar o R3.1, pedi uma explicação breve sobre como mapear o path da requisição para um arquivo dentro de `./public/`.
- **O que a IA respondeu ou fez:**
  - Explicou o uso de `filepath.Join` em vez de concatenar strings e a diferença entre os pacotes `path` (URL) e `path/filepath` (disco). Explicou também tópicos fora do enunciado: path traversal (`/../main.go`) e formas de impedi-lo, decodificação de `%20` com `url.PathUnescape` e corte da query string.
  - Respondi que a maior parte disso não era necessária: a meta é só receber o path no GET, buscar em `./public/` e responder. A IA concordou e resumiu o mínimo.
  - A meu pedido, tirou do card R3.1 o critério de proteção contra `../` (e o trecho correspondente do título), deixou o path traversal como ponto de defesa fora do escopo e moveu o card para Em Andamento.
- **O que aprendi ou aproveitei:** usei apenas a indicação do `filepath.Join`. **Decisão de escopo:** a proteção contra path traversal fica fora do projeto, por não ser exigida pelo enunciado.
- **O que NÃO foi usado:** as técnicas contra path traversal, a decodificação de `%20` e o corte da query string. Nenhum código foi gerado ou alterado.

### 2026-10-31 ~15:15 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R3 (cards R3.1 e R3.3) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que achava ter terminado o mapeamento do path (R3.1) e pedi que a ferramenta verificasse. Depois pedi que atualizasse os cards e registrasse a sessão neste diário.
- **O que a IA respondeu ou fez:**
  - Leu o diff do meu `main.go` (sem alterá-lo), compilou o código numa pasta temporária e testou com `curl`: `/style.css` → `public/style.css`, `/` → `public/index.html` e `/img/logo.png` → `public/img/logo.png`. Confirmou que o R3.1 cumpre o critério de pronto.
  - Moveu o R3.1 para Concluído e o marcou como completo. Moveu o R3.3 (desbloqueado) para A Fazer.
- **O que aprendi ou aproveitei:** -
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~17:00 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R3 (card R3.2) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta verificasse se eu tinha concluído o R3.2 (detecção do Content-Type pela extensão). Depois da correção, pedi que atualizasse o card, registrasse a sessão e commitasse.
- **O que a IA respondeu ou fez:**
  - Leu o diff do meu `main.go` (sem alterá-lo), rodou o `go vet` (sem avisos), compilou o código numa pasta temporária e testou com `curl` os paths `/`, `/style.css`, `/img/a.png`, `/a.jpg`, `/a.jpeg` e `/a.txt`.
  - Na primeira verificação, apontou que o `switch` tratava `.jpeg`, mas não o `.jpg` exigido pela tabela do enunciado. Apontou também uma variável global `extensions` sem uso, e que o critério do card previa um fallback para extensão desconhecida, enquanto eu retorno um erro.
  - Corrigi o `.jpg` e removi a variável. Na segunda verificação, as quatro extensões da tabela retornaram o tipo certo, e as desconhecidas retornaram erro.
  - A meu pedido, trocou o critério do card para "uma extensão desconhecida retorna erro", ajustou o ponto de defesa correspondente, moveu o R3.2 para Concluído e o marcou como completo. O R3.4 continua bloqueado pelo R3.3.
- **O que aprendi ou aproveitei:** a conferência do `switch` com a tabela do enunciado. **Decisão de escopo:** uma extensão desconhecida retorna erro em vez de usar `application/octet-stream`.
- **O que NÃO foi usado:** o fallback `application/octet-stream`. Nenhum código foi gerado ou alterado.
