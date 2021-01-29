package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/hex"
	"strings"
	"errors"
	//"encoding/base64"
	"regexp"
	"log"
)

func main() {
    
	Arquivo01 := os.Args[1] //texto01
	Arquivo02 := os.Args[2] //texto02
	var name string

	texto01 := lerArquivo(Arquivo01)
	texto02 := lerArquivo(Arquivo02)

	inp1, _:= decodeHexBytes([]byte(strings.Join(texto01, "")))   
    inp2, _ := decodeHexBytes([]byte(strings.Join(texto02, "")))

    decoded, _ := fixedXorDecrypt(inp1, inp2) //xor das duas strings encriptadas

    for {
    	fmt.Println("Digite a entrada desejada: ")
	    fmt.Scanf("%s", &name)

	    nameHex:= encodeHexBytes([]byte(name))

	    i := verificaPalavra([]byte(nameHex),[]byte(decoded),0)
	    if i != 0 {
    		fmt.Println("Erro na verifica Palavra")
    	}
    }     
}

func verificaPalavra(palavra, texto []byte, indice int) (int){
	indice += 1 
	if len(palavra) <= len(texto){
		tam := len(palavra)
		
		var txt, novoTexto []byte

		for i:=0; i < tam; i++{
			txt = append(txt, texto[i])
		}
		res,_:= fixedXorDecrypt(txt, palavra)

		palavraE, _ := regexp.MatchString(`^[a-zA-Z.,:;?! ]+$`, string(res))
		
		if(palavraE == true){
			
			fmt.Printf("[%d]mensagem: %s \n",indice,string(res))
		}

		for i:=tam; i <= len(texto)-1; i++{
			novoTexto = append(novoTexto, texto[i])
		}
		
		if tam > len(texto) {
			return 0
		}

		return verificaPalavra(palavra, novoTexto,indice)
	}
	return 0
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

func decodeHexBytes(hexBytes []byte) ([]byte, error) {
   ret := make([]byte, hex.DecodedLen(len(hexBytes)))
   _, err := hex.Decode(ret, hexBytes)
   if err != nil {
      return nil, err
   }
   return ret, nil
}
func encodeHexBytes(input []byte) []byte {
    ret := make([]byte, hex.EncodedLen(len(input)))
    hex.Encode(ret, input)
    return ret
}
func fixedXorDecrypt(input1, input2 []byte) ([]byte, error) {
    if len(input1) != len(input2) {
        return nil, errors.New("As entradas devem ter o mesmo tamanho.")
    }
    ret := make([]byte, len(input1))
    for i := 0; i < len(input1); i++ {
         ret[i] = input1[i] ^ input2[i]
    }
    return ret, nil
}


func salvarTexto(texto string, NomeArquivo string) (error){

	//var linhas []string
	arquivo, err := os.Create(NomeArquivo)
	if err != nil {//
		log.Fatalf("Erro:", err)
	}

	defer arquivo.Close()
	escritor := bufio.NewWriter(arquivo)
	//for _, linha := range linhas {
		fmt.Fprint(escritor, texto) 
	//}
	fmt.Fprintln(escritor,"")
	return escritor.Flush()
}
