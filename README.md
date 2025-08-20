# Unit Converter Go

Uma ferramenta de linha de comando simples e eficiente para conversão de unidades, escrita em Go.

## ✨ Características

- 🌡️ Conversão de temperatura (Celsius, Fahrenheit, Kelvin)
- ⚖️ Conversão de peso (kg, g, lb, oz)
- 📏 Conversão de distância (m, km, ft, mile)
- 🎯 Interface de linha de comando intuitiva
- ✅ Validação robusta de entrada
- 🔧 Tratamento de erros detalhado
- ⚡ Rápido e preciso
- 🛠️ Fácil de usar e estender

## 🚀 Instalação

### Opção 1: Build local
```bash
git clone https://github.com/seu-usuario/unit-converter-go
cd unit-converter-go
go build -o converter
```

### Opção 2: Executar diretamente
```bash
git clone https://github.com/seu-usuario/unit-converter-go
cd unit-converter-go
go run main.go [opções]
```

### Opção 3: Instalar globalmente
```bash
go install github.com/seu-usuario/unit-converter-go@latest
```

## 📖 Uso

```bash
./converter -category [categoria] -from [unidade_origem] -to [unidade_destino] -value [valor]
```

## ⚙️ Opções

| Flag | Tipo | Descrição |
|------|------|-----------|
| `-category` | string | Categoria de conversão (temperature, weight, distance) |
| `-from` | string | Unidade de origem |
| `-to` | string | Unidade de destino |
| `-value` | float64 | Valor a ser convertido |

## 📊 Unidades Suportadas

### 🌡️ Temperatura
- `celsius` - Graus Celsius (°C)
- `fahrenheit` - Graus Fahrenheit (°F)
- `kelvin` - Kelvin (K)

### ⚖️ Peso
- `kg` - Quilogramas
- `g` - Gramas
- `lb` - Libras
- `oz` - Onças

### 📏 Distância
- `m` - Metros
- `km` - Quilômetros
- `ft` - Pés
- `mile` - Milhas

## 💡 Exemplos

### Conversões de Temperatura
```bash
# Celsius para Fahrenheit
./converter -category temperature -from celsius -to fahrenheit -value 25
# Saída: 25.00 celsius = 77.00 fahrenheit

# Fahrenheit para Celsius
./converter -category temperature -from fahrenheit -to celsius -value 100
# Saída: 100.00 fahrenheit = 37.78 celsius

# Celsius para Kelvin
./converter -category temperature -from celsius -to kelvin -value 0
# Saída: 0.00 celsius = 273.15 kelvin
```

### Conversões de Peso
```bash
# Quilogramas para Libras
./converter -category weight -from kg -to lb -value 70
# Saída: 70.00 kg = 154.32 lb

# Onças para Gramas
./converter -category weight -from oz -to g -value 16
# Saída: 16.00 oz = 453.59 g

# Libras para Quilogramas
./converter -category weight -from lb -to kg -value 150
# Saída: 150.00 lb = 68.04 kg
```

### Conversões de Distância
```bash
# Quilômetros para Milhas
./converter -category distance -from km -to mile -value 10
# Saída: 10.00 km = 6.21 mile

# Pés para Metros
./converter -category distance -from ft -to m -value 100
# Saída: 100.00 ft = 30.48 m

# Milhas para Quilômetros
./converter -category distance -from mile -to km -value 5
# Saída: 5.00 mile = 8.05 km
```

## 🧪 Casos de Teste

### Testes de Validação
```bash
# Erro: categoria obrigatória
./converter -from celsius -to fahrenheit -value 25

# Erro: unidade inválida
./converter -category temperature -from celsius -to xyz -value 25

# Conversão de valor zero (válido)
./converter -category temperature -from celsius -to fahrenheit -value 0
```

### Testes de Precisão
```bash
# Ponto de congelamento da água
./converter -category temperature -from celsius -to fahrenheit -value 0
# Esperado: 32.00 fahrenheit

# Ponto de ebulição da água
./converter -category temperature -from celsius -to fahrenheit -value 100
# Esperado: 212.00 fahrenheit

# Conversão métrica padrão
./converter -category weight -from kg -to g -value 1
# Esperado: 1000.00 g
```

## 🔧 Build

### Build simples
```bash
go build -o converter
```

### Build para múltiplas plataformas
```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o converter.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o converter

# MacOS
GOOS=darwin GOARCH=amd64 go build -o converter
```

### Build otimizado
```bash
go build -ldflags="-s -w" -o converter
```

## 📐 Fórmulas de Conversão

### Temperatura
- **Celsius → Fahrenheit**: `°F = (°C × 9/5) + 32`
- **Fahrenheit → Celsius**: `°C = (°F - 32) × 5/9`
- **Celsius → Kelvin**: `K = °C + 273.15`
- **Kelvin → Celsius**: `°C = K - 273.15`

### Peso (baseado em gramas)
- **1 kg** = 1000 g
- **1 lb** = 453.592 g
- **1 oz** = 28.3495 g

### Distância (baseado em metros)
- **1 km** = 1000 m
- **1 ft** = 0.3048 m
- **1 mile** = 1609.34 m

## 📋 Requisitos

- Go 1.16 ou superior

## 🛠️ Desenvolvimento

```bash
# Clonar repositório
git clone https://github.com/seu-usuario/unit-converter-go
cd unit-converter-go

# Executar em modo desenvolvimento
go run main.go [opções]

# Executar testes (quando implementados)
go test ./...

# Verificar formatação
go fmt ./...

# Verificar código
go vet ./...
```

## 🚨 Tratamento de Erros

O programa fornece mensagens de erro claras para:

- **Parâmetros obrigatórios faltando**
- **Categorias não suportadas**
- **Unidades inválidas**
- **Combinações de unidades incompatíveis**

Exemplo de saídas de erro:
```bash
Erro: categoria é obrigatória
Erro: unidade de origem 'xyz' inválida para temperatura
Error: Categoria 'invalid' não suportada
```

## 📄 Licença

MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🤝 Contribuição

Contribuições são bem-vindas! Por favor, abra uma issue ou envie um pull request.

## 🚧 Próximas Funcionalidades

- [ ] Conversão de área (m², ft², acre)
- [ ] Conversão de volume (L, mL, gallon, fl oz)
- [ ] Conversão de velocidade (km/h, mph, m/s)
- [ ] Conversão de energia (J, cal, kWh)
- [ ] Modo interativo
- [ ] Histórico de conversões
- [ ] Configuração de precisão decimal
- [ ] API REST
- [ ] Interface web

## 📚 Exemplos Avançados

### Conversões em Cadeia
```bash
# Converter 100°F para Celsius, depois usar o resultado para Kelvin
./converter -category temperature -from fahrenheit -to celsius -value 100
# Resultado: 37.78°C
./converter -category temperature -from celsius -to kelvin -value 37.78
# Resultado: 310.93K
```

### Casos de Uso Práticos
```bash
# Receita de culinária: 350°F para Celsius
./converter -category temperature -from fahrenheit -to celsius -value 350

# Peso corporal: 70kg para libras
./converter -category weight -from kg -to lb -value 70

# Corrida: 5 milhas para quilômetros
./converter -category distance -from mile -to km -value 5
```