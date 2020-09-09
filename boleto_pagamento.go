package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
)

type SmartContract struct {

}

type Boleto struct {
	BoletoID       string  `json:"boletoId"`
	IfBeneficiario *Banco  `json:"ifBeneficiario"`
	Beneficiario   *Pessoa `json:"beneficiario"`
	IfPagadora     *Banco  `json:"ifPagadora"`
	Pagador        *Pessoa `json:"pagador"`
	CodigoBarra    string  `json:"codigoBarra"`
	DataVencimento string  `json:"datavencimento"`
	Status         string  `json:"status"`
	Valor          string  `json:"valor"`
}

type Pessoa struct {
	PessoaID string `json:"pessoaId"`
	Nome     string `json:"nome"`
	Cpf      string `json:"cpf"`
}

type Banco struct {
	BancoID string `json:"bancoId"`
	Nome    string `json:"nome"`
	Cnpj    string `json:"cnpj"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	
	function, args := APIstub.GetFunctionAndParameters()
	
	if function == "consultarBoleto" {
		return s.consultarBoleto(APIstub, args)
	} else if function == "criarBoleto" {
		return s.criarBoleto(APIstub, args)
	} else if function == "criarBanco" {
		return s.criarBanco(APIstub, args)
	} else if function == "consultarBanco" {
		return s.consultarBanco(APIstub, args)
	} else if function == "criarPessoa" {
		return s.criarPessoa(APIstub, args)
	} else if function == "consultarPessoa" {
		return s.consultarPessoa(APIstub, args)
	} else if function == "registrarBoleto" {
		return s.registrarBoleto(APIstub, args)
	} else if function == "pagarBoleto" {
		return s.pagarBoleto(APIstub, args)
	}

	return shim.Error("Invalid function name")
}

func (s *SmartContract) consultarPessoa(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	pessoaSerializado, err := APIstub.GetState(args[0])
	var objPessoa Pessoa

	err = json.Unmarshal(pessoaSerializado, &objPessoa)

	if err != nil {
		return shim.Error(err.Error())
	}

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")
	fmt.Println("Pessoa ID : ", objPessoa.PessoaID)
	buffer.WriteString("{\"Key\":")
	buffer.WriteString("\"")
	buffer.WriteString(objPessoa.PessoaID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Record\":")
	// Record is a JSON object, so we write as-is
	buffer.WriteString(string(pessoaSerializado))
	buffer.WriteString("}")
	buffer.WriteString("]")

	fmt.Printf("- consultarPessoa:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


func (s *SmartContract) criarBoleto(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	var boleto = Boleto{BoletoID: args[0]}
	boleto.Status = "PENDENTE"

	boletoAsBytes, _ := json.Marshal(boleto)
	APIstub.PutState(args[0], boletoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) consultarBoleto(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	boletoSerializado, err := APIstub.GetState(args[0])
	var objBoleto Boleto

	err = json.Unmarshal(boletoSerializado, &objBoleto)

	if err != nil {
		return shim.Error(err.Error())
	}

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")
	fmt.Println("Boleto ID : ", objBoleto.BoletoID)
	buffer.WriteString("{\"Key\":")
	buffer.WriteString("\"")
	buffer.WriteString(objBoleto.BoletoID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Record\":")
	// Record is a JSON object, so we write as-is
	buffer.WriteString(string(boletoSerializado))
	buffer.WriteString("}")
	buffer.WriteString("]")

	fmt.Printf("- consultarBoleto:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) criarPessoa(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var pessoa = Pessoa{PessoaID: args[0], Nome: args[1], Cpf: args[2]}

	pessoaAsBytes, _ := json.Marshal(pessoa)
	APIstub.PutState(args[0], pessoaAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) criarBanco(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var banco = Banco{BancoID: args[0], Nome: args[1], Cnpj: args[2]}

	bancoAsBytes, _ := json.Marshal(banco)
	APIstub.PutState(args[0], bancoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) consultarBanco(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	bancoSerializado, err := APIstub.GetState(args[0])
	var objBanco Banco

	err = json.Unmarshal(bancoSerializado, &objBanco)

	if err != nil {
		return shim.Error(err.Error())
	}

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")
	fmt.Println("Banco ID : ", objBanco.BancoID)
	buffer.WriteString("{\"Key\":")
	buffer.WriteString("\"")
	buffer.WriteString(objBanco.BancoID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Record\":")
	// Record is a JSON object, so we write as-is
	buffer.WriteString(string(bancoSerializado))
	buffer.WriteString("}")
	buffer.WriteString("]")

	fmt.Printf("- consultarBanco:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


func (s *SmartContract) registrarBoleto(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	boletoSerializado, err := APIstub.GetState(args[0])
	var objBoleto Boleto
	err = json.Unmarshal(boletoSerializado, &objBoleto)
	if err != nil {
		return shim.Error(err.Error())
	}

	objBoleto.CodigoBarra = strconv.Itoa((rand.Intn(5)+10000000 + rand.Intn(5)*rand.Intn(1040)*10000000))+strconv.Itoa((rand.Intn(5)+10000000 + rand.Intn(5)*rand.Intn(1040)*10000000))+strconv.Itoa((rand.Intn(5)+10000000 + rand.Intn(5)*rand.Intn(1040)*10000000))+"001000021"
	var notExpiredDate = time.Now()
	objBoleto.DataVencimento = notExpiredDate.Format("02/01/2006")
	objBoleto.Status = "PENDENTE"

	ifBeneficiarioSerializado, err1 := APIstub.GetState(args[1])
	var objIfBeneficiario Banco
	err1 = json.Unmarshal(ifBeneficiarioSerializado, &objIfBeneficiario)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	objBoleto.IfBeneficiario = &objIfBeneficiario

	beneficiarioSerializado, err2 := APIstub.GetState(args[2])
	var objBeneficiario Pessoa
	err2 = json.Unmarshal(beneficiarioSerializado, &objBeneficiario)
	if err2 != nil {
		return shim.Error(err2.Error())
	}
	objBoleto.Beneficiario = &objBeneficiario

	pagadorSerializado, err4 := APIstub.GetState(args[3])
	var objPagador Pessoa
	err4 = json.Unmarshal(pagadorSerializado, &objPagador)
	if err4 != nil {
		return shim.Error(err4.Error())
	}
	objBoleto.Pagador = &objPagador
	//setar o valor do boleto

	var valor = args[4]
	objBoleto.Valor = valor

	boletoAsBytes, _ := json.Marshal(objBoleto)
	APIstub.PutState(args[0], boletoAsBytes)

	return shim.Success(nil)
}


func (s *SmartContract) pagarBoleto(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	//carregar o boleto do ledger
	boletoSerializado, err := APIstub.GetState(args[0])
	var objBoleto Boleto
	err = json.Unmarshal(boletoSerializado, &objBoleto)
	if err != nil {
		return shim.Error(err.Error())
	}
	//verifcar se boleto esta vencido
	if objBoleto.IsExpired() {
		msgErro := "Boleto vencido"
		fmt.Println("PayBoleto: " + msgErro)
		return shim.Error(msgErro)
	}


	//mudar Status do boleto
	objBoleto.Status = "Pago"
	
	//adicionar ifPagadora
	ifPagadoraSerializado, err1 := APIstub.GetState(args[1])
	var objIfPagadora Banco
	err1 = json.Unmarshal(ifPagadoraSerializado, &objIfPagadora)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	objBoleto.IfPagadora = &objIfPagadora

	bancoAsBytes, _ := json.Marshal(objBoleto)
	APIstub.PutState(args[0], bancoAsBytes)


	return shim.Success(nil)
}



func (t *Boleto) IsExpired() bool {
	var vencimento time.Time

	vencimento, _ = time.Parse("02/01/2006", t.DataVencimento)

	var hoje = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())

	return vencimento.Sub(hoje).Hours() < 0
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
