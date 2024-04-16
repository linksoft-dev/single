import 'package:flutter/material.dart';

class CalculaFatorR extends StatefulWidget {
  @override
  _CalculaFatorRState createState() => _CalculaFatorRState();
}

class _CalculaFatorRState extends State<CalculaFatorR> {
  TextEditingController totalFaturadoController = TextEditingController();
  TextEditingController totalFolhaPagamentoController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Meu Formulário'),
      ),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            TextField(
              controller: totalFaturadoController,
              keyboardType: TextInputType.numberWithOptions(decimal: true),
              decoration: InputDecoration(labelText: 'Total Faturado Últimos Meses'),
            ),
            SizedBox(height: 16.0),
            TextField(
              controller: totalFolhaPagamentoController,
              keyboardType: TextInputType.numberWithOptions(decimal: true),
              decoration: InputDecoration(labelText: 'Total Folha Pagamento'),
            ),
            SizedBox(height: 16.0),
            ElevatedButton(
              onPressed: () {
                // Aqui você pode acessar os valores inseridos nos campos
                double totalFaturado = double.parse(totalFaturadoController.text);
                double totalFolhaPagamento = double.parse(totalFolhaPagamentoController.text);

                // Faça o que precisar com os valores
                // Por exemplo, pode imprimir no console por enquanto
                print('Total Faturado: $totalFaturado');
                print('Total Folha Pagamento: $totalFolhaPagamento');
              },
              child: Text('Calcular'),
            ),
          ],
        ),
      ),
    );
  }
}