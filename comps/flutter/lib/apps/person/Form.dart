import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:single/forms/CrudFormDetail.dart';

class PersonForm extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return CrudFormDetail(
      screenName: "Pessoa",
      body: const Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          Text('Conteúdo do formulário específico'),
          TextField(
            decoration: InputDecoration(labelText: 'Campo 1'),
          ),
          TextField(
            decoration: InputDecoration(labelText: 'Campo 2'),
          ),
        ],
      ),
      customButtons: getCustomButtons(context),
    );
  }

  // Você pode adicionar ou substituir métodos e propriedades aqui, se necessário
  // A classe PersonForm agora herda todas as funcionalidades de MyForm

  @override
  List<Widget> getCustomButtons(BuildContext context) {
    return [
      ElevatedButton(
        onPressed: () {
          // Ao pressionar o botão, exibir o diálogo
          showDialog(
            context: context,
            builder: (BuildContext context) {
              return AlertDialog(
                title: Text('Título do Diálogo'),
                content: Text('Este é um conteúdo de exemplo.'),
                actions: <Widget>[
                  TextButton(
                    onPressed: () {
                      // Fechar o diálogo ao pressionar o botão
                      Navigator.of(context).pop();
                    },
                    child: Text('Fechar'),
                  ),
                ],
              );
            },
          );
        },
        child: Text('Mostrar Popup'),
      ),
    ];
  }
}
