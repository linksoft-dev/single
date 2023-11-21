
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:single/forms/CrudFormDetail.dart';

class PersonForm extends CrudFormDetail {
  PersonForm({String screenName = 'Person'}) : super(screenName: screenName);

  set currentScreen(String currentScreen) {}


  // Você pode adicionar ou substituir métodos e propriedades aqui, se necessário
  // A classe PersonForm agora herda todas as funcionalidades de MyForm

  @override
  List<Widget> customButtons(BuildContext context) {
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
