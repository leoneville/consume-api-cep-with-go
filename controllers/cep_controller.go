package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
}

func formatCep(cep string) string {
	cepFormatado := strings.Replace(cep, ".", "", -1)
	cepFormatado = strings.Replace(cepFormatado, "-", "", -1)
	return cepFormatado
}

func Cep(c *gin.Context) {
	cep := c.Param("cep")

	cepFormatado := formatCep(cep)

	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", cepFormatado)
	client := &http.Client{}
	requisicao, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "erro interno no servidor: " + err.Error(),
		})
		return
	}

	resp, err := client.Do(requisicao)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "erro ao fazer a requisição: " + err.Error(),
		})
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "erro interno no servidor: " + err.Error(),
		})
		return
	}

	var responseObject Response
	json.Unmarshal(data, &responseObject)

	c.JSON(200, responseObject)
}
