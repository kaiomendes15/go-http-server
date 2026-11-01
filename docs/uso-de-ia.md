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

### 2026-10-31 ~17:05 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R3 (card R3.3)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** antes de começar o R3.3, pedi uma explicação sobre como ler o arquivo como bytes e saber o tamanho exato dele.
- **O que a IA respondeu ou fez:** explicou que `os.ReadFile` devolve o conteúdo inteiro em um `[]byte`, sem interpretar o conteúdo (serve igual para HTML e PNG), e que `len` do slice dá o tamanho exato em bytes, que é o valor do `Content-Length`. Lembrou que o `Content-Length` conta bytes e não caracteres (em UTF-8, "ç" ocupa 2 bytes). Citou o `os.Stat().Size()` como alternativa que dispensa ler o arquivo, desnecessária aqui porque o corpo será enviado de qualquer forma. Indicou também `errors.Is(err, fs.ErrNotExist)` para reconhecer o arquivo inexistente, que será útil no R4.1.
- **O que aprendi ou aproveitei:** usei `os.ReadFile` + `len` e já deixei o `errors.Is(err, fs.ErrNotExist)` separando o caso de arquivo inexistente dos outros erros.
- **O que NÃO foi usado:** o `os.Stat().Size()`. Nenhum código foi gerado.

### 2026-10-31 ~17:20 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R3 (card R3.3) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que tinha concluído o R3.3 e pedi que a ferramenta verificasse o que foi implementado, atualizasse os cards e registrasse as interações neste diário.
- **O que a IA respondeu ou fez:**
  - Leu o diff do meu `main.go` (sem alterá-lo), rodou o `go vet` (sem avisos), compilou o código numa pasta temporária, com uma cópia de `./public/` e arquivos de teste, e fez requisições com `curl`.
  - Os tamanhos lidos bateram com o `stat` dos arquivos: `index.html` com 224 bytes, um `.png` binário com 2048 bytes e um `.css` com "ç" com 30 bytes (29 caracteres, porque o "ç" ocupa 2 bytes). Um arquivo inexistente (`/nao.html`) caiu no ramo `[ERRO 404]`. Um diretório (`/dir.html`) e um arquivo sem permissão de leitura caíram no erro genérico de leitura.
  - Apontou para eu decidir no R4.1: como o Content-Type é verificado antes da leitura, um arquivo inexistente com extensão fora da tabela (ex.: `/favicon.ico`, que o navegador pede sozinho) retorna o erro de extensão e nunca chega ao ramo 404.
  - Moveu o R3.3 para Concluído e o marcou como completo. Moveu para A Fazer os cards desbloqueados por ele: R3.4 (o R3.2 já estava concluído) e R4.1.
- **O que aprendi ou aproveitei:** a confirmação de que o tamanho vem em bytes, e não em caracteres, e a pendência da ordem entre a verificação da extensão e a do 404, para resolver no R4.1.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~17:50 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R3 (card R3.4) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** antes de começar o R3.4, perguntei se podia montar a resposta inteira como uma única string, com uma função genérica que recebe os argumentos e devolve a resposta pronta. Depois pedi que a ferramenta commitasse o que eu já tinha feito.
- **O que a IA respondeu ou fez:**
  - Explicou que funciona: em Go, uma `string` é uma sequência imutável de bytes que não precisa ser UTF-8 válido, então `string(conteudo)` preserva os bytes binários de um PNG, e `len` de uma string conta bytes. O custo é uma cópia extra do arquivo na memória, irrelevante neste projeto.
  - Citou como alternativa montar como string só o cabeçalho (status line, headers e linha em branco) e enviar o body como `[]byte`, com dois `Write` ou com `append` num único slice.
  - Observou que a função genérica também serve para o R4 (404), mudando só o status e o conteúdo. Lembrou os cuidados: `\r\n` em todas as linhas, a linha em branco antes do body e o `Content-Length` igual ao `len` dos bytes do body.
  - Commitou o meu `main.go` (R3.3) sem coautoria da IA e, em outro commit, este diário.
- **O que aprendi ou aproveitei:** segui a alternativa: o cabeçalho é montado como string e o body continua `[]byte`, sem passar por string. A função de cabeçalho recebe status, mensagem, Content-Type e tamanho, para ser reaproveitada no R4.
- **O que NÃO foi usado:** a resposta inteira como string. Nenhum código foi gerado ou alterado.

### 2026-10-31 ~18:05 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R3 (card R3.4)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** com o cabeçalho e o body já separados, perguntei como concatenar os dois e devolver como resposta.
- **O que a IA respondeu ou fez:** explicou duas formas: `append([]byte(header), body...)`, que resulta num `[]byte` pronto para o `conn.Write`, ou `header + string(body)`, que exige converter de volta para `[]byte` no `Write`. Explicou que o `...` passa cada elemento do slice como um argumento separado do `append` e que o `conn.Write` devolve `(n, err)`. Lembrou que o socket transmite bytes, por isso o `Write` recebe `[]byte`.
- **O que aprendi ou aproveitei:** usei o `append` com `...` numa função que junta cabeçalho e body.
- **O que NÃO foi usado:** a concatenação com `string(body)`. Nenhum código foi gerado.

### 2026-10-31 ~18:20 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R3 (card R3.4)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que tinha terminado o R3.4 e pedi que a ferramenta verificasse.
- **O que a IA respondeu ou fez:**
  - Leu o diff do meu `main.go` (sem alterá-lo), rodou o `go vet` (sem avisos), compilou o código numa pasta temporária e fez requisições com `curl`. `/`, um `.css` com "ç" e um `.png` binário voltaram com `HTTP/1.1 200 OK`, o Content-Type certo e o Content-Length igual ao tamanho do arquivo. O `cmp` confirmou que o body chegou idêntico ao arquivo.
  - Apontou para eu decidir: a status line repete a versão que vem no pedido (com `HTTP/1.0`, a resposta sai `HTTP/1.0 200 OK`), enquanto o enunciado mostra `HTTP/1.1` fixo; e o erro do `conn.Write` não era verificado.
  - Sugeriu rodar o `gofmt` e corrigir o comentário que chamava o `[]byte` de "array" (em Go, `[N]byte` e `[]byte` são tipos diferentes).
- **O que aprendi ou aproveitei:** passei a verificar o erro do `Write`, rodei o `gofmt` e corrigi o comentário para "slice".
- **O que NÃO foi usado:** mantive a versão vinda do pedido na status line. Nenhum código foi gerado ou alterado.

### 2026-10-31 ~18:35 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** R3 (card R3.4)
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** depois de mover a leitura do arquivo para uma função `getContent`, o código parou de compilar e perguntei qual era o problema. Em seguida, perguntei se as outras tratativas de erro estavam corretas.
- **O que a IA respondeu ou fez:**
  - Identificou, testando numa cópia do código, dois erros de compilação: `_, err := conn.Write(...)` com `:=` sem variável nova (o `err` já existia e o `_` não conta) e o `missing return` no fim da função. Apontou três avisos do `go vet`: `fmt.Errorf` com o `err` como argumento, mas sem `%v` na string. Também observou que quem chamava a função descartava o erro, então os erros sumiam do log.
  - Na revisão das tratativas, confirmou que a lógica estava correta e citou três detalhes: com `%v` o erro original se perde e o `errors.Is(err, fs.ErrNotExist)` deixa de funcionar fora da função (com `%w` ele continua funcionando), o que importa para o R4.1; mensagens terminadas em `\n` geravam uma linha em branco a mais no log; e a função devolvia a resposta e também a enviava, o que era redundante.
- **O que aprendi ou aproveitei:** a diferença entre `:=` e `=`, a regra de que toda função com retorno precisa de `return` em todos os caminhos e o uso de `%v` no `Errorf`. Tirei os `\n` das mensagens e deixei o `Write` só no handler, com o `getContent` apenas montando a resposta. A diferença entre `%v` e `%w` fica para o R4.1.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~18:55 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R3 (card R3.4) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi uma nova verificação do R3.4. Depois de corrigir o que foi apontado, pedi que a ferramenta atualizasse o card, registrasse as interações neste diário e commitasse.
- **O que a IA respondeu ou fez:**
  - Compilou o código numa pasta temporária e testou com `curl`: `/` (276 bytes) e `/neymar.jpg` (31294 bytes, `image/jpeg`) voltaram com `HTTP/1.1 200 OK` e o Content-Length correto, e o `cmp` confirmou que o JPG chegou idêntico.
  - Apontou que, no handler, faltava um `return` depois de imprimir o erro do `getContent`, e por isso o fluxo seguia para o `conn.Write` com a resposta `nil`. Confirmou que `/nao.html` e `/favicon.ico` aparecem no log e ainda não recebem resposta, o que é esperado até o R4.
  - Moveu o R3.4 para Concluído e o marcou como completo. Nenhum card foi desbloqueado: o R3.5 depende do A.1, o R4.2 do R4.1 e o R6.3 do R6.1, todos ainda abertos.
  - Commitou o meu `main.go` e os arquivos de `./public/` sem coautoria da IA e, em outro commit, este diário.
- **O que aprendi ou aproveitei:** coloquei o `return` de volta no tratamento do erro do `getContent`.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-10-31 ~19:10 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R3 (card R3.5), R4 (card R4.1), assets (card A.1) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que achava ter concluído o R3.5 e pedi que a ferramenta verificasse. Depois pedi que marcasse o R3.5 como concluído, revisasse o quadro do Trello conforme o que já foi implementado e commitasse tudo.
- **O que a IA respondeu ou fez:**
  - Compilou o código numa pasta temporária e simulou as requisições do navegador com `curl` (`/`, `/style.css`, `/neymar.jpg`, `/about.html` e `/cris.jpg`). Todas voltaram com 200 e o Content-Type correto, e o terminal mostrou um log `[ARQUIVO]` para cada arquivo, com o tipo e o tamanho.
  - Apontou que o `style.css` estava vazio (0 bytes), o que impediria o "CSS aplicado" do caso de teste. Eu salvei o arquivo com as regras, e a ferramenta confirmou o conteúdo antes de concluir o card.
  - Observou que o enunciado pede o log `[ARQUIVO]` depois de enviar o arquivo, mas o meu código imprime antes do `conn.Write`.
  - Moveu o R3.5 para Concluído. Na revisão do quadro, concluiu o A.1 (os arquivos de `./public/` existem e o `index.html` referencia o CSS e a imagem) e o R4.1 (o `errors.Is(err, fs.ErrNotExist)` já separa o arquivo inexistente, e o teste com `/nao.html` caiu no log `[ERRO 404]`). Com isso, moveu o R4.2 para A Fazer. Os cards A.2, R5.1 e E.1 continuam em A Fazer, porque ainda não existe o formulário nem a rota do POST.
  - Commitou os meus arquivos de `./public/` sem coautoria da IA e, em outro commit, este diário.
- **O que aprendi ou aproveitei:** a conferência dos logs `[ARQUIVO]` com o formato do enunciado e o lembrete do CSS vazio.
- **O que NÃO foi usado:** **Decisão:** mantive o log `[ARQUIVO]` antes do `conn.Write`. Nenhum código foi gerado ou alterado.

### 2026-10-31 ~19:25 (BRT) · Claude Code (modelo Claude Opus 5.5)
- **Requisito relacionado:** documentação do repositório
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi um `README.md` para o projeto, seguindo o mesmo esquema do README de outro repositório meu (go-calculator), e depois pedi que a ferramenta commitasse e registrasse neste diário.
- **O que a IA respondeu ou fez:**
  - Escreveu o `README.md` com as mesmas seções do go-calculator: introdução, tecnologias (badges e pacotes usados), protocolo (exemplo de requisição e resposta HTTP e tabela de Content-Type), tratamento de erros, estrutura de pastas e como executar (`go run main.go`, a flag `-porta`, o acesso pelo navegador e o teste com `curl` e `cmp`). Acrescentou uma tabela com o status de R1 a R6.
  - Rodou o servidor numa pasta temporária para conferir os exemplos. Corrigiu o exemplo de log, que não mostrava uma conexão por arquivo, e acrescentou uma nota sobre as linhas de depuração que o servidor ainda imprime.
  - Commitou o README e este diário com coautoria da IA, porque o texto do README foi escrito pela ferramenta.
- **O que aprendi ou aproveitei:** usei o README como documentação do repositório. A tabela de status deve ser atualizada a cada requisito concluído.
- **O que NÃO foi usado:** o README é só documentação. Nenhum código do servidor foi gerado ou alterado.

### 2026-11-01 ~09:50 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R4 (cards R4.2 e R4.3) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** informei que tinha implementado o R4.2 e pedi que a ferramenta verificasse. Depois de trocar o texto da página, pedi que registrasse a consulta neste diário e atualizasse o quadro.
- **O que a IA respondeu ou fez:**
  - Compilou o código numa pasta temporária (o `go vet` não apontou nada) e testou com `curl`. `/nao-existe.html` voltou com `HTTP/1.1 404 Not Found`, `Content-Type: text/html` e um Content-Length igual ao tamanho de `public/not_found.html`, e `/` continuou voltando com 200.
  - Confirmou que a função `notFound` reaproveita o `buildHeader` e o `buildResponse` do R3.4.
  - Apontou que o texto da página (`404 - PAGE NOT FOUND`) era diferente do enunciado (`404 - Não Encontrado` e "O recurso solicitado não existe."), que é o texto conferido no teste do R4.3. Sem o "ã", também não dava para verificar o critério do acento.
  - Depois da minha troca, refez o teste: o arquivo tem 279 bytes e 277 caracteres (os dois "ã" ocupam 2 bytes cada em UTF-8), e o Content-Length enviado foi 279.
  - Lembrou a pendência do R4.1: como o `getContentType` roda antes do `ReadFile`, `/favicon.ico` e paths sem extensão não recebem resposta nenhuma, nem 404. O teste do enunciado usa `.html`, então isso não bloqueia o R4.
  - Moveu o R4.2 para Concluído e o marcou como completo. Moveu para A Fazer o R4.3, que ele desbloqueia.
- **O que aprendi ou aproveitei:** troquei o texto de `public/not_found.html` pelo do enunciado (e também `lang="pt-br"` e o `<title>`). Para a defesa: o `len` de um `[]byte` conta bytes, por isso o Content-Length bate mesmo com acento; e escolhi a página em arquivo, e não fixa no código, para poder editá-la sem recompilar e usar o mesmo caminho de leitura do R3.
- **O que NÃO foi usado:** a ordem entre a verificação da extensão e a do 404 continua como estava (é opcional para o enunciado). Nenhum código foi gerado ou alterado.

### 2026-11-01 ~10:05 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R4 (card R4.3), documentação e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta concluísse também o R4.3 e, depois de cada ajuste meu no log, que verificasse de novo.
- **O que a IA respondeu ou fez:**
  - Na primeira verificação, apontou que o log `[404]` ainda não existia no código. A única linha com "404" era o `[ERRO 404]` da função `notFound`, que só aparece quando falta o próprio `not_found.html`. Indicou que o `path` pedido pelo cliente está disponível no ramo `fs.ErrNotExist` do `getContent`, e não dentro da `notFound`.
  - Depois que eu escrevi o log, testou com `curl` e mostrou duas diferenças em relação ao enunciado: o log imprimia o `contentPath` (`public/nao-existe.html`) em vez do `path` (`/nao-existe.html`), e havia dois espaços depois de "encontrado:", porque o `fmt.Println` já põe um espaço entre os argumentos e a string terminava com outro.
  - Depois das minhas duas correções, confirmou o log `[404] Recurso não encontrado: /nao-existe.html`, a resposta 404 com Content-Length 279 e que `/` continua voltando 200.
  - Moveu o R4.3 para Concluído e o marcou como completo. O I.2 continua bloqueado, porque ainda depende do R5.5 e do R6.4.
  - Atualizou o `README.md`: o R4 passou para ✅, o arquivo inexistente saiu da lista de erros que fecham a conexão (agora o servidor responde com 404) e o `not_found.html` entrou na estrutura de pastas.
- **O que aprendi ou aproveitei:** troquei `contentPath` por `path` no log e tirei o espaço do fim da string. Aprendi que o `Println` separa os argumentos com espaço. Os critérios do navegador (texto com acento e status 404 na aba Network do DevTools) ficaram para eu conferir.
- **O que NÃO foi usado:** nenhum código do servidor foi gerado ou alterado pela ferramenta.

### 2026-11-01 ~10:35 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R5 (cards R5.1 e R5.2) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** depois de implementar o R5.1, mostrei o log do servidor e perguntei por que o terminal imprimia `[ERRO] Erro ao ler a request line: EOF` antes de cada POST que eu enviava. Depois pedi que a ferramenta registrasse a consulta e movesse o R5.1 para Concluído.
- **O que a IA respondeu ou fez:**
  - Explicou que o `EOF` aparece quando o cliente abre uma conexão TCP e a fecha sem enviar nenhum byte: o `reader.ReadString('\n')` chega ao fim do fluxo sem encontrar o `\n` e devolve o erro. Não é um problema do código do R5.1.
  - Comparou dois testes no servidor compilado numa pasta temporária: um `POST /submit` com `curl` gerou só a conexão do POST, e uma conexão aberta e fechada sem dados reproduziu a mesma linha de EOF do meu log. Concluiu que as conexões vazias vêm do cliente que usei para o POST (ferramentas como Postman ou Insomnia, ou o navegador, que pode abrir uma conexão antecipada e fechá-la sem usar).
  - Observou que o `curl` termina com `Empty reply from server` porque o ramo do POST ainda não envia resposta, o que é esperado até o R6.
  - Conferiu os critérios do R5.1: `GET /` e `GET /style.css` continuam com 200, `/nao-existe.html` com 404, e o `POST /submit` chega ao handler (log `POST para /submit`). Moveu o R5.1 para Concluído e o marcou como completo. Moveu para A Fazer o R5.2, que ele desbloqueia.
- **O que aprendi ou aproveitei:** o significado do `EOF` na leitura da request line e como reproduzi-lo com uma conexão vazia. O servidor já trata esse caso: registra o erro, fecha a conexão e continua aceitando outras.
- **O que NÃO foi usado:** nenhum código foi gerado ou alterado.

### 2026-11-01 ~10:45 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R5 (cards R5.2 e R5.3), R6 (visão geral) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** perguntei como ler o body a partir do Content-Length e se um `conn.Read` com um buffer desse tamanho bastava. Depois perguntei se o `io.ReadFull` devolve a quantidade de bytes lidos e como devolver uma resposta estruturada ao usuário. Por fim, pedi duas revisões da minha implementação do R5.2 e que a ferramenta o concluísse e registrasse neste diário.
- **O que a IA respondeu ou fez:**
  - Explicou que o buffer do tamanho do Content-Length está certo, mas o `conn.Read` tem dois problemas. Primeiro, o `bufio.Reader` dos headers pode já ter puxado parte ou todo o body para o buffer interno, então a leitura precisa usar o mesmo `reader` e não a `conn`. Segundo, um único `Read` pode devolver menos bytes do que o pedido, porque TCP é um fluxo. Indicou o `io.ReadFull`, que repete o `Read` até preencher o buffer ou devolve erro (`io.ErrUnexpectedEOF`).
  - Esclareceu que o `int` do `io.ReadFull` é só a contagem de bytes: os dados ficam no próprio buffer, e `string(buf)` dá o texto para o R5.3. Explicou que a resposta estruturada é o R6: o HTML é gerado pelo código a partir dos campos e enviado com o `buildHeader` e o `buildResponse` que já existem (200, `text/html`, Content-Length em bytes), mais o log `[RESPONSE]`. Para isso, o ramo do POST também precisa da `version`.
  - Na primeira revisão, testou com `curl` e apontou que o handler descartava o retorno do `postContent`: o body lido não aparecia no log e, com um body curto (anunciando 20 bytes e mandando 5), o erro do `io.ReadFull` não era impresso. Indicou também que o `gofmt` acusava o `main.go` (falta um espaço em `([]byte, error){`), que a mensagem de erro fixa perde o erro original (`EOF` x `unexpected EOF`) e, fora do escopo, que sem um timeout o `ReadFull` espera para sempre se o cliente não fechar a conexão.
  - Na segunda revisão, confirmou que o `curl` do card teve o body de 36 bytes lido inteiro, igual ao Content-Length, que o body curto agora registra o erro e encerra, e que o `GET /` continua com 200.
  - Avisou que o critério "envio pelo formulário do navegador" não tinha como ser testado, porque o formulário (A.2) ainda não existe. Com a minha autorização, moveu o R5.2 para Concluído, marcou como completo, deixou um comentário no card sobre esse teste pendente (será conferido no R5.5) e moveu o R5.3 para A Fazer.
- **O que aprendi ou aproveitei:** passei o `reader` para o `postContent` e li o body com `io.ReadFull`. No handler, passei a receber o `buf` e o `err`, imprimir o erro com `return` e mostrar `string(buf)` como log temporário. Para a defesa: por que o mesmo `bufio.Reader` resolve o body adiantado e a diferença entre `io.ReadFull` e um único `Read`.
- **O que NÃO foi usado:** a mensagem de erro continua fixa, e o timeout de leitura não foi implementado. Nenhum código foi gerado ou alterado pela ferramenta.

### 2026-11-01 ~11:15 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R5 (cards R5.3 e R5.4) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi que a ferramenta validasse se a minha forma de separar o body do POST estava no caminho certo. Depois pedi mais duas verificações, à medida que eu ajustava o código, e por fim que movesse o R5.3, registrasse neste diário e commitasse.
- **O que a IA respondeu ou fez:**
  - Na primeira versão (vários `strings.ReplaceAll` trocando `=` por `: `, `&` por quebra de linha e `%40` por `@`), testou com `curl` e apontou três problemas. O resultado era uma string só, e o R6 precisa percorrer os campos um a um (um slice de pares, porque o `map` não garante a ordem). A versão misturava separar com decodificar, e o `%40` era só um caso do que o R5.4 precisa tratar (`+` e `%C3%A3` continuavam codificados). E o `ReplaceAll` do `=` trocava todas as ocorrências, e não só a primeira de cada par. Indicou `strings.Split` por `&`, `strings.Cut` para o primeiro `=`, um slice de `struct` com chave e valor, e o `url.QueryUnescape` para o R5.4.
  - Na segunda versão (slice de `struct` com `strings.Cut`), testou o exemplo do enunciado, `msg=`, `&&` seguidos, um par sem `=` e o body vazio: nenhum caso deu panic, e os critérios do card passaram. Apontou que a função não devolvia o slice, que o tipo estava declarado dentro da função e que o `gofmt` acusava o alinhamento da `struct`. Como opcionais, sugeriu um nome que não sugira decodificação e a troca de um comentário com palavrão, já que o código será entregue.
  - Na terceira versão, confirmou que o handler recebe `[{nome Jo%C3%A3o+Silva} {email joao%40email.com} {msg Ol%C3%A1}]` na ordem original, que `msg=` vira um valor vazio, que o `gofmt` e o `go vet` passam e que os valores continuam codificados, o que é esperado até o R5.4.
  - Moveu o R5.3 para Concluído e o marcou como completo. Moveu para A Fazer o R5.4, que ele desbloqueia.
- **O que aprendi ou aproveitei:** troquei o `map` por um slice da `struct` `Fields` (declarada no nível do pacote) e usei `strings.Cut` para separar cada par no primeiro `=`. Renomeei a função para `sliceFields` e fiz com que ela retornasse o slice. Para a defesa: por que separar antes de decodificar (um `%26` decodificado viraria `&`) e por que o `map` não serve para manter a ordem.
- **O que NÃO foi usado:** o comentário com palavrão continua no código. Nenhum código foi gerado ou alterado pela ferramenta.

### 2026-11-01 ~17:30 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R5 (cards R5.4, R5.5 e R6.1) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** perguntei qual era a melhor forma de aplicar o decode do R5.4. Depois de implementar, pedi que a ferramenta verificasse o código e, por fim, que movesse o card, registrasse neste diário e commitasse.
- **O que a IA respondeu ou fez:**
  - Explicou que o `url.QueryUnescape` deve ser aplicado em cada pedaço depois do `strings.Cut`, e não no body inteiro antes do `Split`, porque um `%26` decodificado viraria `&` e quebraria os pares. Sugeriu decodificar também a chave. Explicou que só o `QueryUnescape` converte `+` em espaço (o `PathUnescape` não converte), que a função devolve erro quando o `%` não vem seguido de dois dígitos hexadecimais (`100%`, `%ZZ`) e que esse erro não deve ser ignorado. Observou que o `%C3%A3` vira os dois bytes UTF-8 do `ã`.
  - Na revisão, testou numa cópia do código com `curl`: `Jo%C3%A3o+Silva` → `João Silva`, `joao%40email.com` → `joao@email.com`, `Programa%C3%A7%C3%A3o` → `Programação` e `a%26b=c%3Dd` → `{a&b c=d}`. Com `100%` e `%ZZ`, o servidor registra o `[ERRO]`, fecha a conexão e continua atendendo. O `gofmt` e o `go vet` passam.
  - Apontou três observações, nenhuma bloqueando o card: um par malformado descarta o body inteiro e o cliente fica sem resposta (decidir no R6); o `Errorf` sem `%w` perde a mensagem original da `QueryUnescape`; e o `Printf` do laço ainda mostra os valores codificados, o que fica para o formato de log do R5.5.
  - Moveu o R5.4 para Concluído e o marcou como completo. Como o A.2 já estava concluído, moveu o R5.5 e o R6.1 para A Fazer.
- **O que aprendi ou aproveitei:** criei a função `decodeValue` com o `url.QueryUnescape`, decodifiquei chave e valor dentro do laço do `sliceFields` e fiz a função devolver `([]Fields, error)`, tratando o erro no handler. Para a defesa: por que `QueryUnescape` e não `PathUnescape`, o que significa `%40` (byte 0x40 = `@`) e por que o `ã` ocupa duas sequências `%XX`.
- **O que NÃO foi usado:** o erro continua com mensagem fixa (sem `%w`), e o log do laço ainda mostra os valores codificados. Nenhum código foi gerado ou alterado pela ferramenta.

### 2026-11-01 ~18:00 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R5 (card R5.5, com o teste pendente do R5.2) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi duas revisões da minha implementação dos logs `[POST]` do R5.5. Depois de fazer o teste pelo navegador, pedi que a ferramenta movesse o card, registrasse neste diário e commitasse.
- **O que a IA respondeu ou fez:**
  - Na primeira revisão, testou numa cópia do código com `curl`. O log do body bruto já mostrava o `+`, e os 56 bytes batiam com o Content-Length. Apontou três problemas: o `Printf` do campo não tinha `\n` e os três campos saíam na mesma linha; o `gofmt` acusava um espaço depois do `}`; e o log ficava antes do `if found`, então registrava pares que não entravam no slice (com `a=1&&semigual`, aparecia `[POST] Campo:  = `). Como opcional, citou os logs de depuração que não estão no formato do enunciado.
  - Na segunda revisão, confirmou que os logs seguiam o formato do enunciado, que o log e o slice tinham os mesmos pares, que o erro de decode passou a mostrar a causa (`invalid URL escape "%ZZ"`) e que o `gofmt` e o `go vet` passavam. Lembrou que o teste pelo navegador só eu podia fazer.
  - Com o meu log do navegador (`Body bruto (48 bytes)` com `Content-Length: 48` e `nome = João Silva`), moveu o R5.5 para Concluído, marcou como completo e deixou um comentário no card com o resultado, que também fecha o teste pendente do R5.2. O I.2 continua bloqueado pelos cards do R6.
- **O que aprendi ou aproveitei:** troquei os logs de depuração pelos `[POST] Body bruto` e `[POST] Campo` do enunciado, coloquei o log do campo dentro do `if found` e acrescentei o `\n`. Também passei a incluir o erro original com `%w` no `decodeValue`. Para a defesa: por que o `Printf` precisa do `\n` e o `Println` não, e por que, pelo navegador, o body bruto chega com os acentos codificados (`Jo%C3%A3o`), diferente do exemplo do enunciado.
- **O que NÃO foi usado:** o `fmt.Println(fields)` e os logs de depuração do GET continuam no código, para limpar no I.2. Nenhum código foi gerado ou alterado pela ferramenta.

### 2026-11-01 ~18:30 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R6 (cards R6.1 e R6.3) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** pedi três revisões da minha implementação da página dinâmica do R6.1. Na primeira, discordei de um critério do card. Na última, pedi que a ferramenta movesse o R6.1 e também o R6.3, registrasse neste diário e commitasse.
- **O que a IA respondeu ou fez:**
  - Na primeira revisão, testou numa cópia do código com `curl`: a página lista 3 campos com 3 pares e 5 com 5, na ordem do formulário, e o body vazio não dá panic. Apontou que nem o HTML nem o header declaravam UTF-8, então o navegador poderia mostrar `JoÃ£o` (o caso de teste do enunciado pede acentos corretos no browser). Apontou também que o log de erro do `Write` verificava o `err2` mas imprimia o `err` (sempre `nil` naquele ponto) e que o `gofmt` acusava os parênteses em `(string)`.
  - Eu apontei que o enunciado não pede para imprimir o HTML no terminal. A ferramenta conferiu o enunciado, confirmou que o único log do R6 é o `[RESPONSE]` e que esse critério tinha sido criado por ela no card. Manteve as outras observações.
  - Na segunda revisão, confirmou o `<meta charset>` e a correção do `err`, mas viu com `curl -i` que o header saía como `Content-Type: charset=utf-8`, sem o tipo de mídia. Explicou que o navegador só exibiu a página certa por *content sniffing* e indicou voltar para `text/html`, já que o meta tag resolve os acentos. Observou que o Content-Length (207) batia com os bytes reais do body e que o log `[RESPONSE]` já cumpria o R6.3, mas era impresso antes do `conn.Write`.
  - Na terceira revisão, confirmou `Content-Type: text/html`, o critério do R6.3 (`Content-Length: 154` igual ao `wc -c` do body), o `gofmt`, o `go vet` e o `GET /` com 200. Moveu o R6.1 e o R6.3 para Concluído, marcou os dois como completos e moveu o R6.2 para A Fazer.
- **O que aprendi ou aproveitei:** criei o `buildHtmlPage` com `strings.Builder` e um `<li>` por campo, e fiz o `postContent` montar a resposta com o `buildHeader` e o `buildResponse` do R3.4, calculando o Content-Length com o `len` dos bytes do HTML. Coloquei o `<meta charset='UTF-8'>` e o log `[RESPONSE]`, e testei os acentos no navegador. Para a defesa: a diferença entre o tipo de mídia e o parâmetro `charset` no `Content-Type`, o que é *content sniffing*, e por que o Content-Length conta bytes, e não caracteres.
- **O que NÃO foi usado:** o critério de imprimir o HTML no terminal (não está no enunciado). O log `[RESPONSE]` continua antes do `conn.Write`, e o comentário `// A variável 'campo'` continua desatualizado. Nenhum código foi gerado ou alterado pela ferramenta.

### 2026-11-01 ~18:55 (BRT) · Claude Code (modelo Claude Opus 5.5) + integração com o Trello
- **Requisito relacionado:** R1 a R6 (cards R6.4 e I.2) e organização
- **Nível de uso:** _<preencher conforme o guia institucional>_
- **Pergunta ou solicitação:** depois de testar o R6.4 no navegador, pedi que a ferramenta rodasse a regressão completa do I.2. Depois de remover os prints de depuração, pedi que rodasse de novo e, por fim, que commitasse, movesse os cards e registrasse neste diário.
- **O que a IA respondeu ou fez:**
  - Rodou numa cópia compilada do código, numa única execução do servidor, todos os casos de teste do enunciado com `curl` e conexões TCP abertas pelo bash: o `gofmt` e o `go vet`; o R1 com 3 conexões abertas e um GET atendido ao mesmo tempo, além de 3 GETs em paralelo; o R2 com `/about.html` e `/`; o R3 com HTML, CSS e as duas imagens (tipo, Content-Length e bytes idênticos aos originais, conferidos com `cmp`); o R4 com `/nao-existe.html`; e o R5/R6 com 5 campos, espaços e acentos (Content-Length igual ao body e ao log `[RESPONSE]`).
  - Na primeira rodada, todos os casos passaram, mas os logs não batiam com o enunciado: a porta impressa numa linha solta no início, o Content-Type impresso antes do `[ARQUIVO]` e do `[404]` e os headers da resposta impressos depois de cada GET. Também viu que as conexões abertas com `nc` não registravam o `Encerrada`, repetiu o teste com conexões que fecham de verdade e confirmou que o problema era do `nc`, e não do servidor.
  - Na segunda rodada, depois da minha limpeza, confirmou que todas as linhas do log seguiam o formato do enunciado. A única fora do enunciado é a opcional `[ERRO] Erro ao ler a request line: EOF`, de conexões que fecham sem enviar nada.
  - Moveu o R6.4 e o I.2 para Concluído, marcou os dois como completos e deixou um comentário em cada um (o R6.2 ficou como opcional, sem bloquear o R6.4). Moveu o I.3 e o E.2 para A Fazer.
- **O que aprendi ou aproveitei:** comentei os três prints de depuração (a porta, o Content-Type e os headers da resposta) para deixar os logs iguais aos do enunciado. Testei pelo navegador as partes que a ferramenta não podia testar: as 3 abas, o CSS com a imagem e a página de 404.
- **O que NÃO foi usado:** o log de `EOF` continua. O R6.2 (escape do HTML) não foi implementado. Nenhum código foi gerado ou alterado pela ferramenta.
