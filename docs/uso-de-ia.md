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
