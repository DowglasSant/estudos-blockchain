# Blockchain: Fundamentos Completos

Este documento consolida os conceitos fundamentais de blockchain, cobrindo **registros imutáveis**, **redes P2P distribuídas**, **Proof of Work** e **protocolo de consenso**, servindo como referência para estudo e revisão.

---

## Índice

### Parte 1: Estrutura e Criptografia
1. [O que é Blockchain](#o-que-é-blockchain)
2. [Estrutura de um Bloco](#estrutura-de-um-bloco)
3. [O Bloco Genesis](#o-bloco-genesis)
4. [Hash SHA256: O Coração da Segurança](#hash-sha256-o-coração-da-segurança)
5. [Propriedades Fundamentais do Hash Criptográfico](#propriedades-fundamentais-do-hash-criptográfico)
6. [Encadeamento de Blocos](#encadeamento-de-blocos)
7. [Imutabilidade: Por que Não se Corrige o Passado](#imutabilidade-por-que-não-se-corrige-o-passado)
8. [Custo de um Ataque (Cadeia Isolada)](#custo-de-um-ataque)

### Parte 2: Redes Distribuídas
9. [Redes P2P: Conceito Geral](#redes-p2p-conceito-geral)
10. [P2P Aplicado ao Blockchain](#p2p-aplicado-ao-blockchain)
11. [Consenso e Regra da Maioria](#consenso-e-regra-da-maioria)
    - [Tolerância a Falha Bizantina (BFT)](#a-base-teórica-tolerância-a-falha-bizantina-bft)
    - [O Problema dos Generais Bizantinos](#o-problema-original)
    - [BFT aplicado ao Blockchain](#traduzindo-para-blockchain)
12. [Ataque 51%: O Limite da Descentralização](#ataque-51-o-limite-da-descentralização)

### Parte 3: Proof of Work e Mineração
13. [O Problema: Criação Instantânea de Blocos](#o-problema-criação-instantânea-de-blocos)
14. [Nonce: O Número Mágico](#nonce-o-número-mágico)
15. [Proof of Work (PoW)](#proof-of-work-pow)
16. [Dificuldade da Rede](#dificuldade-da-rede)
17. [Por que PoW Funciona](#por-que-pow-funciona)

### Parte 4: Protocolo de Consenso em Ação
18. [Cadeias Concorrentes e Longest Chain Rule](#cadeias-concorrentes-e-longest-chain-rule)
19. [Hashing Power: Quem Decide a Corrida](#hashing-power-quem-decide-a-corrida)
20. [Blocos Órfãos](#blocos-órfãos)

### Resumo e Referências
21. [Resumo Visual](#resumo-visual)
22. [Próximos Passos](#próximos-passos)
23. [Referências Técnicas](#referências-técnicas)

---

## O que é Blockchain

Blockchain é uma **estrutura de dados** composta por blocos encadeados através de hashes criptográficos. Cada bloco contém informações e está matematicamente vinculado ao bloco anterior, formando uma corrente (chain) que cresce apenas por adição — nunca por modificação.

```
┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────┐
│ Bloco 0 │───▶│ Bloco 1 │───▶│ Bloco 2 │───▶│ Bloco 3 │
│ Genesis │    │         │    │         │    │         │
└─────────┘    └─────────┘    └─────────┘    └─────────┘
```

**Princípio central:** A integridade de toda a cadeia depende da integridade de cada bloco individual.

---

## Estrutura de um Bloco

Cada bloco na cadeia possui três componentes essenciais:

| Componente | Descrição |
|------------|-----------|
| **Dado** | A informação armazenada (transações, registros, etc.) |
| **Hash Anterior** | O hash do bloco que vem antes na cadeia |
| **Hash Próprio** | Calculado a partir do dado + hash anterior |

### Representação Visual

```
┌────────────────────────────────────┐
│            BLOCO N                 │
├────────────────────────────────────┤
│  Hash Anterior: 7a3f2b...          │  ← Vínculo com bloco N-1
├────────────────────────────────────┤
│  Dado: "Alice envia 10 para Bob"   │  ← Conteúdo do bloco
├────────────────────────────────────┤
│  Hash: 9c4e8d...                   │  ← SHA256(dado + hash anterior)
└────────────────────────────────────┘
```

### Fórmula do Hash

```
Hash do Bloco = SHA256(Dado do Bloco + Hash do Bloco Anterior)
```

Esta fórmula cria o **vínculo criptográfico** entre blocos consecutivos.

---

## O Bloco Genesis

O **bloco genesis** (bloco 0) é o primeiro bloco da cadeia e possui características únicas:

- **Não possui antecessor** — é a âncora de toda a cadeia
- **Hash anterior é zero** (ou string vazia/nula)
- **Seu hash é formado apenas pelo dado** que contém

```
┌────────────────────────────────────┐
│         BLOCO GENESIS              │
├────────────────────────────────────┤
│  Hash Anterior: 0000000000...      │  ← Valor nulo/zero
├────────────────────────────────────┤
│  Dado: "Bloco inicial da cadeia"   │
├────────────────────────────────────┤
│  Hash: a1b2c3d4e5...               │  ← SHA256(dado)
└────────────────────────────────────┘
```

**Importância:** O bloco genesis é imutável por definição. Qualquer alteração nele invalidaria toda a cadeia subsequente.

---

## Hash SHA256: O Coração da Segurança

### O que é SHA256?

SHA256 (Secure Hash Algorithm 256-bit) é uma função hash criptográfica que transforma qualquer entrada em uma saída de tamanho fixo.

### Especificações Técnicas

| Característica | Valor |
|----------------|-------|
| **Tamanho da saída** | 256 bits |
| **Representação** | 64 caracteres hexadecimais |
| **Caracteres válidos** | 0-9 e a-f |
| **Possibilidades** | 2²⁵⁶ ≈ 1.16 × 10⁷⁷ combinações |

### Por que 64 caracteres?

Cada caractere hexadecimal representa **4 bits**:

```
64 caracteres × 4 bits = 256 bits
```

### Exemplo Prático

```
Entrada: "Hello"
Saída:   185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969

Entrada: "hello" (apenas 'H' minúsculo)
Saída:   2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
```

Note como uma mudança mínima (H → h) gera um hash **completamente diferente**.

---

## Propriedades Fundamentais do Hash Criptográfico

Para que um hash seja considerado **criptograficamente seguro**, ele deve satisfazer cinco propriedades:

### 1. Função de Sentido Único (One-Way)

```
Dado ──────▶ Hash ✓ (fácil)
Hash ──────▶ Dado ✗ (impossível)
```

- A partir do hash, **não é possível reconstruir** o dado original
- Não é possível extrair informações relevantes sobre a entrada
- O hash serve para **verificação**, não para reversão

**Analogia:** É como um moedor de carne — você pode transformar carne em carne moída, mas não pode reverter carne moída em um bife.

### 2. Determinismo

```
"Blockchain" ──▶ SHA256 ──▶ sempre ef7797...
"Blockchain" ──▶ SHA256 ──▶ sempre ef7797...
"Blockchain" ──▶ SHA256 ──▶ sempre ef7797...
```

- O **mesmo dado** sempre produz o **mesmo hash**
- Não há aleatoriedade no processo
- Fundamental para **verificação de integridade**

**Sem determinismo:** Seria impossível verificar se um dado foi alterado.

### 3. Processamento Rápido

- O cálculo do hash deve ser **eficiente**
- Permite uso em **larga escala**: milhões de transações, arquivos, validações
- Viabiliza aplicações práticas em tempo real

**Contexto:** Uma rede blockchain processa milhares de transações por segundo. Cada uma requer múltiplos cálculos de hash.

### 4. Efeito Avalanche

```
Entrada A: "Transferir 100 reais"
Hash A:    7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069

Entrada B: "Transferir 101 reais"  (mudança de 1 dígito)
Hash B:    3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d
```

- Uma alteração **mínima** (1 bit) gera um hash **completamente diferente**
- Impossibilita "ajustes finos" maliciosos
- Não há como prever o hash de uma entrada similar

**Segurança:** Um atacante não consegue fazer pequenas alterações até "acertar" um hash desejado.

### 5. Resistência a Colisões

```
Dado A ──▶ Hash X
Dado B ──▶ Hash X   ← Colisão (mesmo hash para dados diferentes)
```

- Colisões são **matematicamente inevitáveis** (infinitas entradas, finitas saídas)
- Porém, é **computacionalmente inviável** encontrar uma colisão de propósito
- Com 2²⁵⁶ possibilidades, a probabilidade é desprezível

**Números:** Para ter 50% de chance de encontrar uma colisão no SHA256, seria necessário calcular aproximadamente 2¹²⁸ hashes — mais do que o número de átomos na Terra.

### Tabela Resumo das Propriedades

| Propriedade | O que Garante | Consequência se Falhar |
|-------------|---------------|------------------------|
| One-Way | Privacidade do dado original | Dados expostos |
| Determinismo | Verificação de integridade | Impossível validar |
| Velocidade | Escalabilidade | Sistema lento/inviável |
| Avalanche | Proteção contra manipulação | Ataques por aproximação |
| Resistência a Colisões | Unicidade do hash | Falsificação de dados |

---

## Encadeamento de Blocos

O encadeamento é o que transforma blocos isolados em uma **cadeia segura**.

### Como Funciona

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│    BLOCO 0      │     │    BLOCO 1      │     │    BLOCO 2      │
├─────────────────┤     ├─────────────────┤     ├─────────────────┤
│ Prev: 0000...   │     │ Prev: a1b2...   │────▶│ Prev: c3d4...   │
│ Data: "Genesis" │     │ Data: "Tx 1"    │     │ Data: "Tx 2"    │
│ Hash: a1b2...   │────▶│ Hash: c3d4...   │     │ Hash: e5f6...   │
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

### O Efeito Dominó

Se alterarmos o **Bloco 1**:

1. O dado do Bloco 1 muda
2. O hash do Bloco 1 muda (efeito avalanche)
3. O "hash anterior" do Bloco 2 não bate mais
4. O Bloco 2 se torna inválido
5. E assim por diante até o último bloco

```
Alteração no Bloco 1:

┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│    BLOCO 0      │     │    BLOCO 1      │     │    BLOCO 2      │
├─────────────────┤     ├─────────────────┤     ├─────────────────┤
│ Prev: 0000...   │     │ Prev: a1b2...   │  ✗  │ Prev: c3d4...   │
│ Data: "Genesis" │     │ Data: "ALTERADO"│     │ Data: "Tx 2"    │
│ Hash: a1b2...   │────▶│ Hash: XXXX...   │──╳──│ Hash: e5f6...   │
└─────────────────┘     └─────────────────┘     └─────────────────┘
                              │                        │
                              │    Hash não corresponde!
                              ▼                        ▼
                        INVÁLIDO                  INVÁLIDO
```

---

## Imutabilidade: Por que Não se Corrige o Passado

### O Princípio Fundamental

> **Em blockchain, não se corrige informações anteriores. Se um evento foi registrado erroneamente, ele deve ser compensado em um novo bloco.**

### Exemplo Prático

**Situação:** Alice transferiu 100 para Bob, mas deveria ter transferido 50.

❌ **Abordagem tradicional (banco de dados):**
```
UPDATE transacoes SET valor = 50 WHERE id = 123;
```

✅ **Abordagem blockchain:**
```
Bloco N:   "Alice transfere 100 para Bob"     ← Permanece inalterado
Bloco N+1: "Bob devolve 50 para Alice"        ← Compensação
```

### Por que Isso Importa?

| Aspecto | Benefício |
|---------|-----------|
| **Auditoria** | Histórico completo e inalterável |
| **Transparência** | Todos podem verificar o que aconteceu |
| **Confiança** | Ninguém pode "reescrever a história" |
| **Rastreabilidade** | Cada correção é documentada |

---

## Custo de um Ataque

### O Cenário de Ataque

Um atacante quer modificar um registro antigo na blockchain. O que ele precisa fazer?

```
Cadeia original (100 blocos):

[0]──[1]──[2]──[3]──...──[50]──[51]──...──[99]──[100]
                           │
                    Atacante quer alterar este bloco
```

### Passos Necessários para o Ataque

1. **Alterar o bloco 50** com a informação fraudulenta
2. **Recalcular o hash** do bloco 50 (novo dado = novo hash)
3. **Atualizar o bloco 51** com o novo hash anterior
4. **Recalcular o hash** do bloco 51
5. **Repetir para todos os blocos subsequentes** (51 até 100)

```
Trabalho necessário:

[50] ← Alterar + recalcular
  │
  ▼
[51] ← Recalcular (hash anterior mudou)
  │
  ▼
[52] ← Recalcular (hash anterior mudou)
  │
  ▼
 ...
  │
  ▼
[100] ← Recalcular (hash anterior mudou)

Total: 51 blocos precisam ser recalculados
```

### Por que Isso é Difícil?

Neste modelo simplificado (sem proof of work), o custo é proporcional ao número de blocos. Mas mesmo assim:

- **Tempo:** Quanto mais antiga a alteração, mais blocos para recalcular
- **Detecção:** Qualquer nó da rede pode verificar a integridade
- **Competição:** Enquanto o atacante recalcula, novos blocos legítimos são adicionados

### Visualização do Custo

```
Bloco alterado     Blocos a recalcular     Dificuldade relativa
─────────────────────────────────────────────────────────────────
Bloco 99           1 bloco                 Baixa
Bloco 90           10 blocos               Média
Bloco 50           50 blocos               Alta
Bloco 10           90 blocos               Muito alta
Bloco 0            100 blocos              Máxima (toda a cadeia)
```

---

# Parte 2: Redes Distribuídas

---

## Redes P2P: Conceito Geral

Antes de entender como blockchain usa redes distribuídas, é importante compreender o conceito de **P2P (Peer-to-Peer)** de forma geral.

### Modelo Tradicional: Cliente-Servidor

No modelo tradicional da internet, existe uma hierarquia clara:

```
              ┌──────────────┐
              │   SERVIDOR   │  ← Autoridade central
              │              │    (controla tudo)
              └──────┬───────┘
                     │
        ┌────────────┼────────────┐
        │            │            │
        ▼            ▼            ▼
   ┌────────┐   ┌────────┐   ┌────────┐
   │Cliente │   │Cliente │   │Cliente │  ← Apenas consomem
   └────────┘   └────────┘   └────────┘
```

**Características:**
- Um **servidor centralizado** detém toda a informação e controle
- Clientes apenas **solicitam e consomem** dados
- Se o servidor cai, **todo o sistema para**
- O servidor é um **ponto único de falha** e de controle

**Exemplos:** Netflix, bancos tradicionais, e-mail (Gmail, Outlook), redes sociais centralizadas.

### Modelo P2P (Peer-to-Peer)

No modelo P2P, **não existe hierarquia**. Todos os participantes são equivalentes:

```
       ┌──────┐                 ┌──────┐
       │ Peer │◀───────────────▶│ Peer │
       └──┬───┘                 └───┬──┘
          │                         │
          │       ┌──────┐          │
          └──────▶│ Peer │◀─────────┘
                  └──┬───┘
                     │
          ┌──────────┴──────────┐
          │                     │
          ▼                     ▼
      ┌──────┐              ┌──────┐
      │ Peer │◀────────────▶│ Peer │
      └──────┘              └──────┘

      Todos são iguais. Todos se conectam entre si.
```

**Características:**
- **Sem hierarquia** — todos os nós são equivalentes (peer = par, igual)
- **Sem ponto central de falha** — a rede sobrevive mesmo se vários nós caírem
- **Cada nó é cliente E servidor** simultaneamente
- **Recursos distribuídos** — dados, processamento e banda são compartilhados

### Comparação: Cliente-Servidor vs P2P

| Aspecto | Cliente-Servidor | P2P |
|---------|------------------|-----|
| **Controle** | Centralizado | Distribuído |
| **Ponto de falha** | Servidor é crítico | Nenhum nó é crítico |
| **Escalabilidade** | Servidor é gargalo | Cresce com mais nós |
| **Confiança** | Depositada no servidor | Depositada no protocolo |
| **Censura** | Fácil (desliga servidor) | Difícil (milhares de nós) |
| **Custo** | Infraestrutura cara | Distribuído entre participantes |

### Exemplos Históricos de P2P

| Sistema | Ano | O que Distribui | Status |
|---------|-----|-----------------|--------|
| **Napster** | 1999 | Músicas | Fechado (parcialmente P2P) |
| **BitTorrent** | 2001 | Arquivos em geral | Ativo |
| **Skype** (original) | 2003 | Chamadas de voz | Migrou para centralizado |
| **Bitcoin** | 2009 | Transações financeiras | Ativo |
| **IPFS** | 2015 | Arquivos (web distribuída) | Ativo |

### Definição Formal

> **Rede P2P é uma arquitetura onde não existe autoridade central. Os participantes (peers) colaboram diretamente entre si, compartilhando recursos e responsabilidades de forma igualitária, seguindo um protocolo comum.**

---

## P2P Aplicado ao Blockchain

Blockchain combina a **estrutura de dados encadeada** (blocos + hashes) com uma **rede P2P distribuída** para criar um sistema verdadeiramente descentralizado.

### Arquitetura da Rede Blockchain

```
                    REDE BLOCKCHAIN DISTRIBUÍDA
    ═══════════════════════════════════════════════════════

         ┌─────────┐                     ┌─────────┐
         │  NÓ A   │◀───────────────────▶│  NÓ B   │
         │ ┌─────┐ │                     │ ┌─────┐ │
         │ │Chain│ │                     │ │Chain│ │
         │ └─────┘ │                     │ └─────┘ │
         └────┬────┘                     └────┬────┘
              │                               │
              │         ┌─────────┐           │
              └────────▶│  NÓ C   │◀──────────┘
                        │ ┌─────┐ │
                        │ │Chain│ │
                        │ └─────┘ │
                        └────┬────┘
                             │
              ┌──────────────┴──────────────┐
              │                             │
              ▼                             ▼
         ┌─────────┐                   ┌─────────┐
         │  NÓ D   │◀─────────────────▶│  NÓ E   │
         │ ┌─────┐ │                   │ ┌─────┐ │
         │ │Chain│ │                   │ │Chain│ │
         │ └─────┘ │                   │ └─────┘ │
         └─────────┘                   └─────────┘

    Cada nó possui uma CÓPIA COMPLETA da blockchain
```

### Princípio Fundamental

> **Cada nó da rede mantém uma cópia completa e idêntica de toda a blockchain.**

Isso significa que:
- Não existe um "servidor" que guarda "a blockchain oficial"
- Cada participante tem sua própria cópia verificável
- A "verdade" emerge do **consenso** entre os nós

### O que Cada Nó Faz

| Função | Descrição |
|--------|-----------|
| **Armazena** | Mantém cópia completa da blockchain |
| **Valida** | Verifica integridade dos blocos recebidos |
| **Propaga** | Compartilha novos blocos com outros nós |
| **Sincroniza** | Mantém sua cópia atualizada com a rede |

### Fluxo de um Novo Bloco

Quando um novo bloco é criado:

```
1. Nó A cria/recebe novo bloco

       ┌─────────┐
       │  NÓ A   │  ← Bloco novo criado aqui
       │ [1][2][3]│
       └────┬────┘
            │
            ▼ Propaga para vizinhos

2. Vizinhos validam e propagam

       ┌─────────┐         ┌─────────┐
       │  NÓ B   │         │  NÓ C   │
       │ [1][2]  │   →     │ [1][2]  │
       │   ↓     │         │   ↓     │
       │  [3]✓   │         │  [3]✓   │  ← Validam e aceitam
       └────┬────┘         └────┬────┘
            │                   │
            ▼                   ▼

3. Toda a rede atualizada

    Todos os nós agora têm: [1][2][3]
```

---

## Consenso e Regra da Maioria

### A Base Teórica: Tolerância a Falha Bizantina (BFT)

O mecanismo de consenso do blockchain é baseado em um problema clássico da ciência da computação: o **Problema dos Generais Bizantinos** (1982, Leslie Lamport).

#### O Problema Original

Imagine vários generais de um exército cercando uma cidade inimiga. Eles precisam decidir em conjunto: **atacar ou recuar**. A comunicação é feita por mensageiros, mas:

- Alguns generais podem ser **traidores** (enviam mensagens falsas)
- Mensageiros podem ser **interceptados**
- Mesmo assim, os generais leais precisam chegar a um **consenso correto**

```
                    PROBLEMA DOS GENERAIS BIZANTINOS
    ═══════════════════════════════════════════════════════

              General A                General B
             (Leal ✓)                (Leal ✓)
            "Atacar!"               "Atacar!"
                 ╲                     ╱
                  ╲                   ╱
                   ▼                 ▼
              ┌────────────────────────┐
              │    CIDADE INIMIGA      │
              └────────────────────────┘
                   ▲                 ▲
                  ╱                   ╲
                 ╱                     ╲
            "Recuar!"               "Atacar!"
           General C                General D
          (Traidor ✗)              (Leal ✓)

    Problema: Como os generais leais (A, B, D) chegam
    a um consenso correto, mesmo com C sendo traidor?
```

#### Traduzindo para Blockchain

| Generais Bizantinos | Blockchain |
|---------------------|------------|
| Generais | Nós da rede |
| Generais leais | Nós honestos |
| Generais traidores | Nós maliciosos |
| Mensageiros | Protocolo de rede (P2P) |
| "Atacar ou recuar" | "Qual bloco é válido?" |
| Consenso dos leais | Consenso da maioria |

#### A Solução no Blockchain

O blockchain resolve o Problema dos Generais Bizantinos através de uma combinação:

```
┌───────────────────────────────────────────────────────────────┐
│              TOLERÂNCIA A FALHA BIZANTINA (BFT)               │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│  1. PROOF OF WORK                                             │
│     └─ Custa caro criar blocos → traidores pagam alto         │
│                                                               │
│  2. REGRA DA MAIORIA                                          │
│     └─ Maioria honesta sempre vence (>50%)                    │
│                                                               │
│  3. INCENTIVO ECONÔMICO                                       │
│     └─ Minerar honestamente é mais lucrativo que atacar       │
│                                                               │
│  4. VERIFICAÇÃO INDEPENDENTE                                  │
│     └─ Cada nó valida por si mesmo, sem confiar nos outros    │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

#### O Limite: Regra dos 2/3

A teoria da tolerância a falha bizantina prova matematicamente que:

> **Um sistema distribuído pode tolerar até 1/3 de nós maliciosos e ainda funcionar corretamente.**

```
Rede com 9 nós:

    Até 3 traidores (33%) → Sistema funciona corretamente ✓
    4+ traidores (44%+)   → Consenso pode ser comprometido ✗

Na prática, o blockchain (com PoW) eleva esse limite:
    O atacante precisa de >50% do poder computacional
    (não apenas >33% dos nós)
```

#### Por que o Blockchain é uma Solução Elegante

Antes do Bitcoin, as soluções para o Problema dos Generais Bizantinos exigiam:
- Que os participantes se conhecessem
- Comunicação direta entre todos os pares
- Rodadas de votação sincronizadas

O Bitcoin/blockchain resolveu isso para **participantes anônimos** em uma **rede aberta**, usando PoW como mecanismo de custo e a cadeia mais longa como critério de verdade.

---

### O Problema na Prática

Em uma rede distribuída, como decidir qual versão da blockchain é a "correta"?

```
Situação problemática:

    NÓ A: [0]──[1]──[2]──[3]
    NÓ B: [0]──[1]──[2]──[3]
    NÓ C: [0]──[1]──[2]──[3']  ← Bloco diferente!
    NÓ D: [0]──[1]──[2]──[3]
    NÓ E: [0]──[1]──[2]──[3]

    Qual é a versão correta? A de C ou a dos outros?
```

### A Solução: Regra da Maioria

> **A versão válida da blockchain é aquela aceita pela maioria dos nós da rede (50% + 1).**

```
Contagem:

    Versão com [3]:  A, B, D, E  →  4 nós (80%)  ✓ VÁLIDA
    Versão com [3']: C           →  1 nó  (20%)  ✗ REJEITADA

    Resultado: Nó C deve atualizar sua chain para a versão majoritária
```

### Verificação Contínua

A rede está **constantemente verificando** a consistência:

```
Loop de verificação (simplificado):

    ┌─────────────────────────────────────────────────────┐
    │                                                     │
    │   Para cada nó na rede:                             │
    │   │                                                 │
    │   ├──▶ Comparar minha chain com vizinhos            │
    │   │                                                 │
    │   ├──▶ Se minha chain ≠ maioria:                    │
    │   │        └──▶ Atualizar para versão majoritária   │
    │   │                                                 │
    │   └──▶ Repetir continuamente                        │
    │                                                     │
    └─────────────────────────────────────────────────────┘
```

### Correção Automática

Quando um nó está com dados diferentes da maioria, ele é **automaticamente corrigido**:

```
Antes da sincronização:

    NÓ A: [0]──[1]──[2]──[3]     ┐
    NÓ B: [0]──[1]──[2]──[3]     │ Maioria
    NÓ C: [0]──[1]──[2X]──[3X]   │ ← Divergente (malicioso ou defeituoso)
    NÓ D: [0]──[1]──[2]──[3]     │
    NÓ E: [0]──[1]──[2]──[3]     ┘

Após sincronização:

    NÓ A: [0]──[1]──[2]──[3]
    NÓ B: [0]──[1]──[2]──[3]
    NÓ C: [0]──[1]──[2]──[3]     ← Corrigido automaticamente
    NÓ D: [0]──[1]──[2]──[3]
    NÓ E: [0]──[1]──[2]──[3]

    ✓ Consenso restaurado
```

### Por que Isso Funciona

| Fator | Contribuição para Segurança |
|-------|----------------------------|
| **Redundância** | Milhares de cópias = impossível destruir dados |
| **Transparência** | Qualquer um pode verificar a integridade |
| **Automatização** | Correção acontece sem intervenção humana |
| **Descentralização** | Nenhum ponto único de controle ou falha |

---

## Ataque 51%: O Limite da Descentralização

### O Cenário de Ataque

Para manipular a blockchain em uma rede distribuída, um atacante precisaria:

```
Requisitos para ataque bem-sucedido:

    1. Controlar mais de 50% dos nós da rede
    2. Alterar a blockchain nesses nós
    3. Fazer isso SIMULTANEAMENTE
    4. Antes que a rede corrija os nós comprometidos

    ┌─────────────────────────────────────────────────────┐
    │                                                     │
    │   Rede com 10.000 nós                               │
    │                                                     │
    │   Para ataque: comprometer > 5.000 nós              │
    │                ao mesmo tempo                       │
    │                em questão de minutos                │
    │                                                     │
    └─────────────────────────────────────────────────────┘
```

### Por que é Praticamente Impossível

```
Obstáculos para o atacante:

    ┌──────────────────┐
    │ 1. ESCALA        │  Milhares de máquinas independentes
    └────────┬─────────┘  em diferentes países e jurisdições
             │
             ▼
    ┌──────────────────┐
    │ 2. COORDENAÇÃO   │  Ataque precisa ser simultâneo
    └────────┬─────────┘  (verificação é constante)
             │
             ▼
    ┌──────────────────┐
    │ 3. TEMPO         │  Enquanto ataca, rede continua
    └────────┬─────────┘  adicionando blocos legítimos
             │
             ▼
    ┌──────────────────┐
    │ 4. CUSTO         │  Recursos computacionais enormes
    └────────┬─────────┘  (especialmente com PoW)
             │
             ▼
    ┌──────────────────┐
    │ 5. DETECÇÃO      │  Anomalias são rapidamente
    └──────────────────┘  identificadas pela comunidade
```

### Exemplo Numérico

```
Bitcoin (exemplo aproximado):

    Nós na rede:        ~15.000 nós completos
    Para ataque 51%:    ~7.500 nós comprometidos

    Considerando:
    - Nós em ~100 países diferentes
    - Diferentes ISPs e infraestruturas
    - Operadores independentes e anônimos
    - Verificação contínua 24/7

    Probabilidade de sucesso: Praticamente zero
```

### Tabela de Dificuldade por Tamanho da Rede

| Tamanho da Rede | Nós para 51% | Dificuldade |
|-----------------|--------------|-------------|
| 10 nós | 6 nós | Trivial |
| 100 nós | 51 nós | Fácil |
| 1.000 nós | 501 nós | Moderada |
| 10.000 nós | 5.001 nós | Muito difícil |
| 100.000 nós | 50.001 nós | Praticamente impossível |

### Nota Importante

Este modelo considera apenas a **quantidade de nós**. Em redes reais com **Proof of Work** (como Bitcoin), o ataque 51% se refere ao **poder computacional** (hashrate), não apenas ao número de nós. Isso adiciona outra camada de dificuldade que veremos a seguir.

---

# Parte 3: Proof of Work e Mineração

---

## O Problema: Criação Instantânea de Blocos

Até agora, vimos que o hash de um bloco é calculado assim:

```
Hash = SHA256(dados + hash_anterior)
```

**O problema:** esse cálculo é **instantâneo**. Qualquer computador pode gerar milhões de hashes por segundo.

```
Sem custo computacional:

    Criar bloco legítimo:     0.001 segundos
    Criar bloco fraudulento:  0.001 segundos

    → Atacante pode criar blocos tão rápido quanto a rede honesta
    → Sem barreira de entrada
    → Sistema inseguro
```

**A solução:** adicionar um **custo computacional** à criação de blocos.

---

## Nonce: O Número Mágico

### O que é Nonce?

**Nonce** = "**N**umber used **ONCE**" (número usado uma vez)

É um campo adicional no bloco que o minerador pode **alterar livremente** para tentar gerar um hash que atenda a uma condição específica.

### Estrutura do Bloco com Nonce

```
┌─────────────────────────────────────────┐
│              BLOCO N                    │
├─────────────────────────────────────────┤
│  Hash Anterior: 0000abc123def456...     │
│  Dados: "Alice → Bob: 50 BTC"           │
│  Timestamp: 1699900800                  │
│  Nonce: 8294712                         │  ← Minerador ajusta isso
├─────────────────────────────────────────┤
│  Hash: 0000f7d2e1a9b8c7...              │  ← Deve atender à dificuldade
└─────────────────────────────────────────┘
```

### Nova Fórmula do Hash

```
Hash = SHA256(dados + hash_anterior + timestamp + nonce)
                                                   │
                                    Único campo que pode ser alterado
                                    para "procurar" um hash válido
```

### Por que o Nonce Existe?

O nonce permite que o minerador **tente diferentes valores** até encontrar um hash que satisfaça a condição de dificuldade, **sem alterar os dados reais do bloco**.

---

## Proof of Work (PoW)

### Mineração é Força Bruta

**Não existe estratégia.** Mineração é pura **tentativa e erro**:

```go
// Pseudocódigo da mineração
for {
    hash = SHA256(bloco + nonce)

    if hash.começa_com("0000...") {
        return "Bloco minerado!"
    }

    nonce++  // Tenta o próximo
}
```

### A Realidade da Mineração

| Aspecto | Realidade |
|---------|-----------|
| **Estratégia** | Não existe. É pura tentativa e erro |
| **Previsibilidade** | Zero. Impossível saber qual nonce vai funcionar |
| **Vantagem competitiva** | Quem tem mais poder computacional tenta mais nonces/segundo |
| **Fator sorte** | Um PC fraco pode encontrar antes de um forte (improvável, mas possível) |

### Por que Não Dá pra "Calcular" o Nonce Certo?

Por causa do **efeito avalanche** do SHA256 — não há padrão entre nonces consecutivos:

```
nonce = 1000000  →  hash = 7f3a8b2c1d...  ✗ (não começa com 0000)
nonce = 1000001  →  hash = 2e9c1da4f8...  ✗ (não começa com 0000)
nonce = 1000002  →  hash = 0000f7e2b1...  ✓ (ACHEI!)
```

- O hash do nonce `1000002` não tem **nenhuma relação** com o hash do `1000001`
- Não há como "se aproximar" do resultado
- Cada tentativa é independente — como jogar dados

### Consumo de Energia

É por isso que mineração consome tanta energia:

```
Rede Bitcoin (exemplo):

    ~500 EH/s (exahashes por segundo)
    = 500.000.000.000.000.000.000 tentativas por segundo
    = 500 quintilhões de hashes por segundo

    Tudo isso para "adivinhar" o próximo nonce válido
```

---

### O Conceito

**Proof of Work** (Prova de Trabalho) é um mecanismo que exige que o minerador **prove** que gastou recursos computacionais para criar um bloco.

A prova é simples: o hash do bloco deve começar com um **número específico de zeros**.

### Como Funciona na Prática

```
Dificuldade definida pela rede: hash deve começar com "0000"

O minerador tenta diferentes valores de nonce:

┌──────────────────────────────────────────────────────────────────┐
│ Tentativa │ Nonce │ Hash Resultante                    │ Válido │
├───────────┼───────┼────────────────────────────────────┼────────┤
│     1     │   0   │ 7a3f8b2c1d9e4f5a6b7c8d9e0f1a2b3c │   ✗    │
│     2     │   1   │ b2e9f1a4c8d7e6f5a4b3c2d1e0f9a8b7 │   ✗    │
│     3     │   2   │ 1c8d7e3f2a9b0c1d2e3f4a5b6c7d8e9f │   ✗    │
│    ...    │  ...  │ ...                                │   ✗    │
│ 8.294.712 │8294712│ 0000a8f3b2c1d4e5f6a7b8c9d0e1f2a3 │   ✓    │
└──────────────────────────────────────────────────────────────────┘

Após 8.294.712 tentativas, encontrou um hash válido!
```

### Visualização do Processo

```
                    PROCESSO DE MINERAÇÃO
    ═══════════════════════════════════════════════════════

    Dados do bloco (fixos):
    ┌─────────────────────────────────────────┐
    │ Hash Anterior: 0000abc...               │
    │ Transações: [tx1, tx2, tx3]             │
    │ Timestamp: 1699900800                   │
    └─────────────────────────────────────────┘
                        │
                        ▼
    ┌─────────────────────────────────────────┐
    │         LOOP DE MINERAÇÃO               │
    │                                         │
    │   nonce = 0                             │
    │   │                                     │
    │   ├─▶ hash = SHA256(bloco + nonce)      │
    │   │                                     │
    │   ├─▶ hash começa com "0000"?           │
    │   │       │                             │
    │   │       ├─ NÃO → nonce++ → repetir    │
    │   │       │                             │
    │   │       └─ SIM → BLOCO MINERADO! ✓    │
    │   │                                     │
    └───┴─────────────────────────────────────┘
                        │
                        ▼
    ┌─────────────────────────────────────────┐
    │ Bloco válido com nonce = 8294712        │
    │ Hash: 0000f7d2e1a9b8c7d6e5f4a3b2c1d0   │
    └─────────────────────────────────────────┘
```

---

## Dificuldade da Rede

### O que é Dificuldade?

A **dificuldade** define quantos zeros o hash deve ter no início. Mais zeros = mais difícil encontrar.

### Impacto da Dificuldade

```
Dificuldade    Prefixo exigido    Tentativas médias    Tempo estimado*
────────────────────────────────────────────────────────────────────────
    1          0...               16                   instantâneo
    2          00...              256                  instantâneo
    3          000...             4.096                milissegundos
    4          0000...            65.536               segundos
    5          00000...           1.048.576            minutos
    6          000000...          16.777.216           dezenas de min
    ...        ...                ...                  ...
    19         0000000000...      ~2⁷⁶                 anos/décadas

* Tempo varia conforme poder computacional
```

### Ajuste Dinâmico

No Bitcoin, a dificuldade é **ajustada automaticamente** a cada 2016 blocos (~2 semanas) para manter o tempo médio de mineração em **~10 minutos por bloco**.

```
Se blocos estão sendo minerados muito rápido:
    → Aumenta dificuldade (mais zeros)

Se blocos estão sendo minerados muito devagar:
    → Diminui dificuldade (menos zeros)
```

---

## Por que PoW Funciona

### Assimetria Fundamental

| Ação | Custo |
|------|-------|
| **Minerar** (encontrar nonce válido) | Milhões de tentativas |
| **Verificar** (conferir se hash é válido) | 1 única operação |

```
Minerador:
    hash = SHA256(bloco + 0)        ✗
    hash = SHA256(bloco + 1)        ✗
    hash = SHA256(bloco + 2)        ✗
    ... (milhões de tentativas) ...
    hash = SHA256(bloco + 8294712)  ✓  ← Encontrou!

Verificador:
    hash = SHA256(bloco + 8294712)
    hash começa com "0000"?  ✓  ← Verificado instantaneamente!
```

### Custo Real de um Ataque

Agora o atacante precisa:

1. **Alterar o bloco** com informação fraudulenta
2. **Re-minerar o bloco** (encontrar novo nonce) — **CARO**
3. **Re-minerar TODOS os blocos subsequentes** — **MUITO CARO**
4. **Fazer isso mais rápido que toda a rede honesta** — **PRATICAMENTE IMPOSSÍVEL**

```
Ataque em blockchain com PoW:

    Bloco alterado    Trabalho necessário
    ─────────────────────────────────────────────────
    Bloco 99          Re-minerar 1 bloco
    Bloco 90          Re-minerar 10 blocos
    Bloco 50          Re-minerar 50 blocos
    Bloco 0           Re-minerar TODA a cadeia

    E enquanto o atacante re-minera...
    A rede honesta continua adicionando blocos novos!

    → Atacante precisa de >50% do poder computacional TOTAL
    → Isso é o verdadeiro "Ataque 51%"
```

### Resumo: Camadas de Segurança com PoW

```
┌─────────────────────────────────────────────────────────────┐
│  CAMADA 4: PROOF OF WORK                                    │
│  └─ Cada bloco exige milhões de cálculos para ser criado    │
│                                                             │
│    ┌─────────────────────────────────────────────────────┐  │
│    │  CAMADA 3: REDE DISTRIBUÍDA                         │  │
│    │  └─ Ataque precisa de 51% do hashrate               │  │
│    │                                                     │  │
│    │    ┌─────────────────────────────────────────────┐  │  │
│    │    │  CAMADA 2: ENCADEAMENTO                     │  │  │
│    │    │  └─ Alterar 1 bloco exige re-minerar todos  │  │  │
│    │    │                                             │  │  │
│    │    │    ┌─────────────────────────────────────┐  │  │  │
│    │    │    │  CAMADA 1: HASH CRIPTOGRÁFICO       │  │  │  │
│    │    │    │  └─ SHA256: one-way, avalanche      │  │  │  │
│    │    │    └─────────────────────────────────────┘  │  │  │
│    │    └─────────────────────────────────────────────┘  │  │
│    └─────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘

Cada camada MULTIPLICA exponencialmente a dificuldade de ataque
```

---

# Parte 4: Protocolo de Consenso em Ação

---

## Cadeias Concorrentes e Longest Chain Rule

### Como Surgem Cadeias Concorrentes

Cadeias concorrentes não são apenas fruto de ataques — acontecem **naturalmente** na operação normal da rede. Quando dois mineradores encontram um nonce válido quase simultaneamente, a rede se divide temporariamente:

```
Situação: Dois mineradores encontram o bloco 101 ao mesmo tempo

    Minerador A (Brasil):  encontra bloco 101 às 14:00:00.000
    Minerador B (Japão):   encontra bloco 101 às 14:00:00.003

    Ambos são válidos! Ambos cumprem a dificuldade.
    Mas são blocos DIFERENTES (dados/nonce diferentes).
```

### A Divisão Temporária (Fork)

```
                    FORK NATURAL NA REDE
    ═══════════════════════════════════════════════════════

    Cadeia original (todos concordam até aqui):

    [0]──[1]──[2]──...──[100]
                            │
                 ┌──────────┴──────────┐
                 │                     │
                 ▼                     ▼
             [101a]                [101b]
          (Minerador A)        (Minerador B)

    Nós próximos de A: [0]──...──[100]──[101a]
    Nós próximos de B: [0]──...──[100]──[101b]

    A rede está temporariamente dividida em duas versões!
```

### A Resolução: Longest Chain Rule (Regra de Nakamoto)

> **A cadeia mais longa (com mais blocos) é considerada a versão verdadeira.**

A corrida é decidida por quem minerar o **próximo bloco** primeiro:

```
Cenário 1: Cadeia A minera o bloco 102 primeiro

    Cadeia A: [0]──...──[100]──[101a]──[102]   ← VENCEDORA (mais longa)
    Cadeia B: [0]──...──[100]──[101b]           ← PERDEDORA

    Resultado: Toda a rede adota a Cadeia A

─────────────────────────────────────────────────

Cenário 2: Cadeia B minera o bloco 102 primeiro

    Cadeia A: [0]──...──[100]──[101a]           ← PERDEDORA
    Cadeia B: [0]──...──[100]──[101b]──[102]   ← VENCEDORA

    Resultado: Toda a rede adota a Cadeia B
```

### Visualização Completa do Processo

```
TEMPO 0: Rede unificada
    Todos: [0]──...──[100]

TEMPO 1: Fork (dois blocos 101 minerados simultaneamente)
    Grupo A: [0]──...──[100]──[101a]
    Grupo B: [0]──...──[100]──[101b]

TEMPO 2: Corrida pelo bloco 102
    Grupo A: [0]──...──[100]──[101a]──[102] ✓ PRIMEIRO!
    Grupo B: [0]──...──[100]──[101b]        ✗ Perdeu

TEMPO 3: Rede reunificada
    Todos: [0]──...──[100]──[101a]──[102]
                              │
                          [101b] → BLOCO ÓRFÃO (descartado)
```

---

## Hashing Power: Quem Decide a Corrida

### Quantidade de Nós vs Poder Computacional

Um ponto crucial: **a corrida não é decidida por quantidade de nós, mas por poder de processamento (hashrate)**.

```
Cenário:

    Grupo A: 1.000 nós com 70% do hashrate total
    Grupo B: 5.000 nós com 30% do hashrate total

    Quem provavelmente minera o próximo bloco primeiro?
    → Grupo A (apesar de ter MENOS nós)

    Hashrate = tentativas de nonce por segundo
    Mais hashrate = mais chances de encontrar o nonce válido
```

### Hashrate e Probabilidade

| Hashrate do Grupo | Probabilidade de Minerar Próximo Bloco |
|--------------------|-----------------------------------------|
| 10% da rede | ~10% de chance |
| 30% da rede | ~30% de chance |
| 51% da rede | ~51% de chance (controle efetivo) |
| 70% da rede | ~70% de chance |

### Por que Hashrate Importa Mais que Nós

```
1 nó com ASIC dedicado          vs       100 nós com CPUs comuns
┌────────────────────────┐               ┌──────────────────────┐
│ 100 TH/s (terahashes)  │               │ 0.1 TH/s (total)    │
└────────────────────────┘               └──────────────────────┘
          │                                        │
          ▼                                        ▼
   1.000x mais poder                        Quase irrelevante
   de processamento                         na corrida

→ 1 nó potente > 100 nós fracos
→ O que conta é PODER DE PROCESSAMENTO, não democracia de nós
```

### Implicações para Segurança

Isso redefine o **Ataque 51%**:

```
Ataque 51% NÃO significa:
    ✗ Controlar 51% dos nós da rede

Ataque 51% SIGNIFICA:
    ✓ Controlar 51% do HASHRATE (poder computacional total)

Na prática (Bitcoin):
    Hashrate total: ~500 EH/s
    Para ataque: >250 EH/s
    Custo estimado em hardware: bilhões de dólares
    Custo em energia: milhões de dólares por dia
```

---

## Blocos Órfãos

### O que São

Blocos órfãos (orphan blocks) são blocos **válidos** que foram descartados porque pertenciam à cadeia que **perdeu a corrida**.

```
┌────────────────────────────────────────────────────────────┐
│                     BLOCO ÓRFÃO                            │
├────────────────────────────────────────────────────────────┤
│                                                            │
│  - Foi minerado corretamente (PoW válido)                  │
│  - Cumpriu a dificuldade                                   │
│  - Mas pertencia à cadeia mais curta                       │
│  - Foi descartado quando a cadeia mais longa venceu        │
│                                                            │
└────────────────────────────────────────────────────────────┘
```

### Ciclo de Vida de um Bloco Órfão

```
1. CRIAÇÃO
   Minerador B encontra nonce válido para bloco 101b   ✓

2. PROPAGAÇÃO
   Parte da rede aceita o bloco 101b                   ✓

3. COMPETIÇÃO
   Minerador A também encontra bloco 101a              ✓
   Rede se divide temporariamente

4. RESOLUÇÃO
   Cadeia com 101a minera o bloco 102 primeiro
   Cadeia A se torna a mais longa                      ✓

5. ORFANAMENTO
   Bloco 101b é descartado                             ✗
   └─ Vira bloco órfão
```

### O que Acontece com o Conteúdo

| Elemento | Destino |
|----------|---------|
| **Transações** | Voltam para o **pool de transações pendentes** (mempool) |
| **Recompensa do minerador** | **Perdida** — o minerador não recebe nada |
| **Trabalho computacional** | **Desperdiçado** — energia gasta sem retorno |

### As Transações Não se Perdem

```
Bloco 101b (órfão) continha:
    - Tx1: Alice → Bob: 10 BTC
    - Tx2: Carol → Dave: 5 BTC
    - Tx3: Eve → Frank: 2 BTC

Após orfanamento:

    Tx1, Tx2, Tx3 → voltam para a MEMPOOL
                     │
                     ▼
              Serão incluídas em um
              bloco futuro da cadeia
              vencedora

    → Nenhuma transação é perdida
    → Apenas atrasadas temporariamente
```

### Frequência de Blocos Órfãos

No Bitcoin, forks temporários e blocos órfãos são **raros mas normais**:

```
Taxa aproximada de blocos órfãos no Bitcoin:

    ~1-2 por semana (em ~1.008 blocos minerados/semana)
    ≈ 0.1% a 0.2% dos blocos

    Motivos:
    - Latência de rede (propagação não é instantânea)
    - Mineradores em regiões geográficas distantes
    - Dois mineradores encontram solução quase simultânea
```

### Resumo: Protocolo de Consenso Completo

```
┌─────────────────────────────────────────────────────────────────┐
│              PROTOCOLO DE CONSENSO EM AÇÃO                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  1. DEFESA (Regra da Maioria)                                   │
│     └─ Mudanças precisam de 50%+1 dos nós                       │
│     └─ Ataques exigem comprometer a maioria simultaneamente     │
│                                                                 │
│  2. COMPETIÇÃO (Cadeias Concorrentes)                           │
│     └─ Forks temporários são naturais                           │
│     └─ Longest Chain Rule resolve a disputa                     │
│     └─ Hashrate decide quem vence a corrida                     │
│                                                                 │
│  3. RESOLUÇÃO (Blocos Órfãos)                                   │
│     └─ Cadeia perdedora é descartada                            │
│     └─ Blocos descartados viram órfãos                          │
│     └─ Transações voltam ao pool (não se perdem)                │
│     └─ Minerador perdedor perde recompensa                      │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Resumo Visual

### Anatomia Completa de uma Blockchain

```
                    BLOCKCHAIN: REGISTROS IMUTÁVEIS
═══════════════════════════════════════════════════════════════════

    ┌─────────────────────────────────────────────────────────┐
    │                    BLOCO GENESIS                        │
    │  ┌───────────────────────────────────────────────────┐  │
    │  │ Hash Anterior: 000000000000000000000000000000000  │  │
    │  │ Dado: "Início da cadeia"                          │  │
    │  │ Hash: a1b2c3d4e5f6789...                          │  │
    │  └───────────────────────────────────────────────────┘  │
    └─────────────────────────┬───────────────────────────────┘
                              │
                              ▼
    ┌─────────────────────────────────────────────────────────┐
    │                      BLOCO 1                            │
    │  ┌───────────────────────────────────────────────────┐  │
    │  │ Hash Anterior: a1b2c3d4e5f6789...                 │◀─┤ Vínculo
    │  │ Dado: "Transação: Alice → Bob: 50"                │  │
    │  │ Hash: b2c3d4e5f6a7890...                          │  │
    │  └─────────────────────────────────────────────────��─┘  │
    └─────────────────────────┬───────────────────────────────┘
                              │
                              ▼
    ┌─────────────────────────────────────────────────────────┐
    │                      BLOCO 2                            │
    │  ┌───────────────────────────────────────────────────┐  │
    │  │ Hash Anterior: b2c3d4e5f6a7890...                 │◀─┤ Vínculo
    │  │ Dado: "Transação: Bob → Carol: 30"                │  │
    │  │ Hash: c3d4e5f6a7b8901...                          │  │
    │  └───────────────────────────────────────────────────┘  │
    └─────────────────────────┬───────────────────────────────┘
                              │
                              ▼
                            [...]


═══════════════════════════════════════════════════════════════════
                    PROPRIEDADES DO SHA256
═══════════════════════════════════════════════════════════════════

    ┌──────────────┐  ┌──────────────┐  ┌──────────────┐
    │   ONE-WAY    │  │ DETERMINISMO │  │  VELOCIDADE  │
    │              │  │              │  │              │
    │  Hash → Dado │  │ Mesmo input  │  │   Cálculo    │
    │  impossível  │  │ mesmo output │  │   rápido     │
    └──────────────┘  └──────────────┘  └──────────────┘

    ┌──────────────────────┐  ┌──────────────────────┐
    │   EFEITO AVALANCHE   │  │ RESISTÊNCIA COLISÃO  │
    │                      │  │                      │
    │  1 bit muda = hash   │  │ Encontrar 2 inputs   │
    │  totalmente diferente│  │ com mesmo hash é     │
    │                      │  │ computacionalmente   │
    │                      │  │ inviável             │
    └──────────────────────┘  └──────────────────────┘


═══════════════════════════════════════════════════════════════════
                    PRINCÍPIO DE IMUTABILIDADE
═══════════════════════════════════════════════════════════════════

    ERRADO ❌                        CORRETO ✓
    ──────────                       ─────────
    Alterar bloco antigo             Adicionar novo bloco
                                     com compensação

    [1]──[2]──[3]                    [1]──[2]──[3]──[4]
          │                                         │
          ▼                                         ▼
       "Corrigir"                              "Compensar"


═══════════════════════════════════════════════════════════════════
                    REDE P2P DISTRIBUÍDA
═══════════════════════════════════════════════════════════════════

    CLIENTE-SERVIDOR                 PEER-TO-PEER (P2P)
    ────────────────                 ──────────────────

         [SERVER]                     [P]──────[P]
            │                          │ ╲    ╱ │
        ┌───┼───┐                      │  ╲  ╱  │
        │   │   │                      │   ╲╱   │
       [C] [C] [C]                    [P]──────[P]

    Ponto único de falha             Sem ponto central
    Controle centralizado            Todos são iguais


═══════════════════════════════════════════════════════════════════
                    CONSENSO E REGRA DA MAIORIA
═══════════════════════════════════════════════════════════════════

    Nó divergente detectado:

    NÓ A: [1]─[2]─[3]  ───┐
    NÓ B: [1]─[2]─[3]     ├──▶ 80% = VERDADE
    NÓ C: [1]─[2]─[3]     │
    NÓ D: [1]─[2]─[3]  ───┘
    NÓ E: [1]─[2]─[X]  ───────▶ 20% = CORRIGIDO

    ✓ Nó E atualizado automaticamente para versão majoritária


═══════════════════════════════════════════════════════════════════
                    CAMADAS DE SEGURANÇA
═══════════════════════════════════════════════════════════════════

    ┌─────────────────────────────────────────────────────────┐
    │  CAMADA 3: REDE DISTRIBUÍDA                             │
    │  └─ Ataque precisa de 51% dos nós simultaneamente       │
    │                                                         │
    │    ┌─────────────────────────────────────────────────┐  │
    │    │  CAMADA 2: ENCADEAMENTO                         │  │
    │    │  └─ Alterar 1 bloco exige recalcular todos      │  │
    │    │                                                 │  │
    │    │    ┌─────────────────────────────────────────┐  │  │
    │    │    │  CAMADA 1: HASH CRIPTOGRÁFICO           │  │  │
    │    │    │  └─ SHA256: one-way, avalanche, etc.    │  │  │
    │    │    └─────────────────────────────────────────┘  │  │
    │    └─────────────────────────────────────────────────┘  │
    └─────────────────────────────────────────────────────────┘

    Cada camada MULTIPLICA a dificuldade de ataque

═══════════════════════════════════════════════════════════════════
```

---

## Próximos Passos

Este documento cobre os fundamentos de **registros imutáveis**, **redes P2P distribuídas** e **Proof of Work**. Os próximos tópicos a estudar incluem:

- [x] **Proof of Work (PoW)** — Mineração, dificuldade e o custo real de um ataque 51%
- [x] **Nonce** — O número que torna a mineração um desafio computacional
- [ ] **Proof of Stake (PoS)** — Alternativa ao PoW baseada em participação econômica
- [ ] **Merkle Trees** — Eficiência na verificação de transações
- [ ] **Transações e UTXO** — Como o dinheiro "flui" na blockchain
- [ ] **Contratos Inteligentes** — Código executável na blockchain
- [ ] **Wallets e Chaves** — Criptografia assimétrica e assinaturas digitais

---

## Referências Técnicas

- **SHA256:** [NIST FIPS 180-4](https://csrc.nist.gov/publications/detail/fips/180/4/final)
- **Bitcoin Whitepaper:** [Satoshi Nakamoto, 2008](https://bitcoin.org/bitcoin.pdf)

---

*Documento criado para consolidação de estudos sobre blockchain — Parte 1: Estrutura e Criptografia + Parte 2: Redes Distribuídas + Parte 3: Proof of Work e Mineração + Parte 4: Protocolo de Consenso em Ação.*
