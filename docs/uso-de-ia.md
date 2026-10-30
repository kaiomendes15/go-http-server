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
