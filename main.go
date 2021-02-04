package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/hex"
	"strings"
	"errors"
	"regexp"
	"log"
)

func main() {
    
	Arquivo01 := os.Args[1] //texto01
	Arquivo02 := os.Args[2] //texto02
	var opc int
	var nameByte, textoClaro1, textoClaro2 []byte

	texto01 := lerArquivo(Arquivo01)
	texto02 := lerArquivo(Arquivo02)

	textoEncriptado1, _:= HexForBytes([]byte(strings.Join(texto01, "")))   
    textoEncriptado2, _ := HexForBytes([]byte(strings.Join(texto02, "")))

    decoded, _ := XorBytes(textoEncriptado1, textoEncriptado2) //xor das duas strings encriptadas
	for len(textoClaro1) != 76{

    	fmt.Println("Digite sua entrada: ")
    	scanner := bufio.NewScanner(os.Stdin)
	    scanner.Scan()
	    line := scanner.Text()

	    nameHex := BytesForHex([]byte(line))
	    nameByte,_ = HexForBytes([]byte(nameHex))   
	
		textoClaro1,textoClaro2 = verificaPalavra(nameByte,[]byte(decoded),0)
		
	    if len(textoClaro1) == 76 {
			fmt.Printf("\nDeseja gerar chave dessa frase? Se sim digite 1, se nao digite 2.\n")
			fmt.Scanf("%d", &opc)
			if opc == 1{ //GERAÇÃO DE CHAVE

				chave1,_ := XorBytes(textoEncriptado1,textoClaro1)
				chave2,_ := XorBytes(textoEncriptado2,textoClaro2)

				if Equal(chave1,chave2) {

					final := BytesForHex([]byte(chave1))
					fmt.Printf("\nCHAVE: %s\n",string(final))
					break
				}else{
					chave1,_ := XorBytes(textoEncriptado1,textoClaro2)
					chave2,_ := XorBytes(textoEncriptado2,textoClaro1)

					if Equal(chave1,chave2) == false{
						fmt.Printf("Erro em gerar chave!")
						break
					}

					final := BytesForHex([]byte(chave1))
					fmt.Printf("\nCHAVE: %s\n",string(final))
					break
				}
			}
		}else{
			textoClaro1 = []byte("")
		}
	}	
}

func Equal(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func verificaPalavra(palavra, texto []byte, indice int) ([]byte,[]byte){
	indice += 1 
	var frase []string
	var frs []byte
	if len(palavra) <= len(texto){
		tam := len(palavra)
		
		var txt, novoTexto []byte
		var i int 

		for i:=0; i < tam; i++{
			txt = append(txt, texto[i])
		}
		res,_:= XorBytes(txt, palavra)

		palavraE, _ := regexp.MatchString(`^[a-zA-Z.,:;?! ]+$`, string(res))
		j := indice;
		if(palavraE == true){
			for i = 1; i <= 76; i++ {
				if i == j {
					frase = append(frase, string(res))
				}
				frs = []byte(strings.Join(frase, ""))
			}
			fmt.Printf("\n--------Possiveis resultados para sua entrada--------\n")
			fmt.Printf("\n[%d]frase:%s\n",indice,strings.Join(frase, ""))
			return frs, palavra		
		}

		for i:=1; i < len(texto); i++{
			novoTexto = append(novoTexto, texto[i])
		}
		
		if tam > len(texto) {
			return frs, palavra
		}

		return verificaPalavra(palavra, novoTexto,indice)
	}
	return frs, palavra
}

func lerArquivo(arquivo string) ([]string){
	var linhasArquivo []string

	textoArquivo, err := os.Open(arquivo)
	if err != nil {
		fmt.Printf(" %v \n", err)
	}

	txtArquivo := bufio.NewScanner(textoArquivo)
	for txtArquivo.Scan() {
		linhasArquivo = append(linhasArquivo, txtArquivo.Text())
	}
	textoArquivo.Close()

	return linhasArquivo
}

func HexForBytes(hexBytes []byte) ([]byte, error) {
   ret := make([]byte, hex.DecodedLen(len(hexBytes)))
   _, err := hex.Decode(ret, hexBytes)
   if err != nil {
      return nil, err
   }
   return ret, nil
}
func BytesForHex(input []byte) []byte {
    ret := make([]byte, hex.EncodedLen(len(input)))
    hex.Encode(ret, input)
    return ret
}
func XorBytes(input1, input2 []byte) ([]byte, error) {
    if len(input1) != len(input2) {
        return nil, errors.New("As entradas devem ter o mesmo tamanho.")
    }
    ret := make([]byte, len(input1))
    for i := 0; i < len(input1); i++ {
         ret[i] = input1[i] ^ input2[i]
    }
    return ret, nil
}
