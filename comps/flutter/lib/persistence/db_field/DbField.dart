import 'package:flutter/material.dart';

// Classe Dataset fornecida para referência
class Dataset {
  // Implementação do Dataset
}

class DbField extends StatefulWidget {
  final Dataset dataset;
  final String fieldName;

  const DbField({
    Key key,
    @required this.dataset,
    @required this.fieldName,
  }) : super(key: key);

  @override
  _DbFieldState createState() => _DbFieldState();
}

class _DbFieldState extends State<DbField> {
  late TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController(text: getValueFromDataset());
    widget.dataset.onAfter.listen((_) {
      setState(() {
        _controller.text = getValueFromDataset();
      });
    });
  }

  String getValueFromDataset() {
    // Implemente a lógica para obter o valor do dataset com base no fieldName
    // Por enquanto, estamos apenas retornando um valor fixo para ilustração
    return widget.fieldName + " value";
  }

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: _controller,
      decoration: InputDecoration(
        labelText: widget.fieldName,
      ),
      onChanged: (newValue) {
        // Implemente a lógica para atualizar o dataset com o novo valor
      },
    );
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }
}
