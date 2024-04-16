import 'package:flutter/material.dart';


class MySystemWidget extends StatelessWidget {
  final String applicationName;
  final String applicationVersion;
  final List<Map<String, dynamic>> menus;
  final List<Map<String, dynamic>> usuarios;

  const MySystemWidget({
    Key? key,
    required this.applicationName,
    required this.applicationVersion,
    required this.menus,
    required this.usuarios,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Cabeçalho fixo ocupando 10% da tela
        Container(
          height: MediaQuery.of(context).size.height * 0.1,
          // Aqui você pode adicionar o campo de texto de pesquisa
          child: TextField(
            decoration: InputDecoration(
              hintText: 'Pesquisar...',
            ),
          ),
        ),
        // Resto da tela dividida em duas colunas
        Expanded(
          child: Row(
            children: [
              // Primeira coluna com menus e informações
              Container(
                width: MediaQuery.of(context).size.width * 0.15,
                // Aqui você pode criar a estrutura dos menus e usuários
                child: ListView.builder(
                  itemCount: menus.length + 2, // 2 para o nome da empresa e espaço vazio
                  itemBuilder: (context, index) {
                    if (index == 0) {
                      // Nome da empresa
                      return Container(
                        height: MediaQuery.of(context).size.height * 0.1 * 0.1, // 10% da coluna
                        alignment: Alignment.center,
                        child: Text(
                          applicationName,
                          style: TextStyle(fontWeight: FontWeight.bold),
                        ),
                      );
                    } else if (index == 1) {
                      // Espaço vazio entre o nome da empresa e os menus
                      return SizedBox(height: 10);
                    } else {
                      // Construa aqui os itens do menu e usuários conectados
                      return ListTile(
                        title: Text(menus[index - 2]["nome"]),
                        // Você pode adicionar outras informações dos menus/usuarios
                      );
                    }
                  },
                ),
              ),
              // Segunda coluna para mostrar telas abertas
              Expanded(
                child: Container(
                  color: Colors.grey[200], // Cor de fundo para mostrar telas abertas
                  // Aqui você pode adicionar a lógica para exibir as telas abertas
                  child: Center(
                    child: Text("Telas abertas"),
                  ),
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }
}
