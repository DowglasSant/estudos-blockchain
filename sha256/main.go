package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// =============================================================================
// SHA256 em Go - Demonstração dos Conceitos Fundamentais
// =============================================================================
//
// Este programa demonstra três propriedades essenciais do SHA256:
//
// 1. SENTIDO ÚNICO (One-Way)
//    - Dado → Hash: fácil e rápido
//    - Hash → Dado: impossível (não existe função inversa)
//
// 2. DETERMINISMO
//    - Mesmo input SEMPRE gera o mesmo output
//    - Executar 1000x com "Hello" → sempre o mesmo hash
//
// 3. EFEITO AVALANCHE
//    - Mudança mínima no input → hash completamente diferente
//    - "Hello" vs "hello" → hashes totalmente distintos
//
// =============================================================================

// generateHash recebe um texto e retorna seu hash SHA256
// Esta é uma função de SENTIDO ÚNICO: não existe generateText(hash)
func generateHash(data string) string {
	// sha256.Sum256 retorna um array de 32 bytes (256 bits)
	// Cada byte tem 8 bits, então: 32 bytes × 8 bits = 256 bits
	hash := sha256.Sum256([]byte(data))

	// Convertemos os bytes para hexadecimal (string legível)
	// Cada byte vira 2 caracteres hex, então: 32 bytes × 2 = 64 caracteres
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("            SHA256 - DEMONSTRAÇÃO DE CONCEITOS")
	fmt.Println(strings.Repeat("=", 70))

	// =========================================================================
	// DEMONSTRAÇÃO 1: DETERMINISMO
	// =========================================================================
	// O mesmo input SEMPRE produz o mesmo output.
	// Isso é fundamental para verificação de integridade.

	fmt.Println("\n📌 CONCEITO 1: DETERMINISMO")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("O mesmo input sempre gera o mesmo hash.")
	fmt.Println()

	input := "Blockchain"

	fmt.Printf("Input: %q\n\n", input)

	// Geramos o hash 3 vezes para provar que é sempre igual
	for i := 1; i <= 3; i++ {
		hash := generateHash(input)
		fmt.Printf("Execução %d: %s\n", i, hash)
	}

	fmt.Println("\n✓ Resultado: Os 3 hashes são IDÊNTICOS")

	// =========================================================================
	// DEMONSTRAÇÃO 2: EFEITO AVALANCHE
	// =========================================================================
	// Uma mudança mínima no input gera um hash COMPLETAMENTE diferente.
	// Isso impede ataques por "aproximação" ou "ajuste fino".

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("\n📌 CONCEITO 2: EFEITO AVALANCHE")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("Mudança mínima no input → hash completamente diferente.")
	fmt.Println()

	// Teste 1: Maiúscula vs minúscula
	input1 := "Hello"
	input2 := "hello" // apenas 'H' → 'h'

	hash1 := generateHash(input1)
	hash2 := generateHash(input2)

	fmt.Println("Teste: Maiúscula vs Minúscula")
	fmt.Printf("Input 1: %q → %s\n", input1, hash1)
	fmt.Printf("Input 2: %q → %s\n", input2, hash2)
	fmt.Printf("\nDiferença no input: apenas 1 caractere (H → h)\n")
	fmt.Printf("Caracteres diferentes no hash: %d de 64\n", countDifferences(hash1, hash2))

	fmt.Println()

	// Teste 2: Número alterado
	input3 := "Transferir 100 reais"
	input4 := "Transferir 101 reais" // apenas 100 → 101

	hash3 := generateHash(input3)
	hash4 := generateHash(input4)

	fmt.Println("Teste: Alteração de valor")
	fmt.Printf("Input 1: %q\n         → %s\n", input3, hash3)
	fmt.Printf("Input 2: %q\n         → %s\n", input4, hash4)
	fmt.Printf("\nDiferença no input: apenas 1 dígito (100 → 101)\n")
	fmt.Printf("Caracteres diferentes no hash: %d de 64\n", countDifferences(hash3, hash4))

	fmt.Println("\n✓ Resultado: Mudanças mínimas geram hashes TOTALMENTE diferentes")

	// =========================================================================
	// DEMONSTRAÇÃO 3: SENTIDO ÚNICO (One-Way)
	// =========================================================================
	// Podemos facilmente ir de Dado → Hash
	// Mas é IMPOSSÍVEL ir de Hash → Dado

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("\n📌 CONCEITO 3: SENTIDO ÚNICO (One-Way)")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("Dado → Hash: trivial | Hash → Dado: impossível")
	fmt.Println()

	secreto := "minha senha secreta"
	hashSecreto := generateHash(secreto)

	fmt.Printf("Dado original: %q\n", secreto)
	fmt.Printf("Hash gerado:   %s\n", hashSecreto)
	fmt.Println()
	fmt.Println("Agora, dado apenas o hash acima, tente descobrir o texto original...")
	fmt.Println("→ NÃO EXISTE função inversa: reverseHash(hash) → dado")
	fmt.Println("→ A única forma seria testar TODAS as combinações possíveis (força bruta)")
	fmt.Println("→ Com 2²⁵⁶ possibilidades, isso é computacionalmente inviável")

	fmt.Println("\n✓ Resultado: SHA256 é uma função de MÃO ÚNICA")

	// =========================================================================
	// DEMONSTRAÇÃO 4: TAMANHO FIXO
	// =========================================================================
	// Independente do tamanho do input, o output sempre tem 64 caracteres hex

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("\n📌 BÔNUS: TAMANHO FIXO DA SAÍDA")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("Qualquer input → sempre 64 caracteres hexadecimais (256 bits)")
	fmt.Println()

	inputs := []string{
		"A",
		"Blockchain",
		"Este é um texto muito maior para demonstrar que não importa o tamanho",
	}

	for _, inp := range inputs {
		h := generateHash(inp)
		fmt.Printf("Input (%2d chars): %q\n", len(inp), truncate(inp, 50))
		fmt.Printf("Hash  (64 chars): %s\n\n", h)
	}

	fmt.Println("✓ Resultado: Output SEMPRE tem 64 caracteres (256 bits)")

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("                    FIM DA DEMONSTRAÇÃO")
	fmt.Println(strings.Repeat("=", 70))
}

// countDifferences conta quantos caracteres são diferentes entre dois hashes
func countDifferences(hash1, hash2 string) int {
	count := 0
	for i := 0; i < len(hash1) && i < len(hash2); i++ {
		if hash1[i] != hash2[i] {
			count++
		}
	}
	return count
}

// truncate trunca uma string se for maior que maxLen
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
