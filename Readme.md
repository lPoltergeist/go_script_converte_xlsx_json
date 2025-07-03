Script simples para converter `.xls` para `.json`.

Foi feito na necessidade de ter algo rápido e fácil de utilizar.
As ferramentas disponiveis na internet tem limite de tamanho e são mais lentas.

---

## 🚀 Como usar

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/xls-to-json.git
cd xls-to-json
```

2. Coloque o arquivo .xls desejado na pasta do projeto.

3. Altere o caminho do arquivo de entrada e o nome do JSON de saída no main.go:

```bash
xl, err := xlsxreader.OpenFile("./nome do arquivo.xlsx")
...
file, err := os.Create("./nome de saida.json")
```

4. Rode o script:

```bash
go run main.go
```


📦 Requisitos
Go instalado (versão 1.18 ou superior recomendada)

Pacote github.com/thedatashed/xlsxreader

Instale com:

```bash
go get github.com/thedatashed/xlsxreader