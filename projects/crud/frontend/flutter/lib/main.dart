import 'package:flutter/material.dart';
import '../../../forms/CrudFormDetail.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData.light(), // Definir tema claro como tema padrão
      darkTheme: ThemeData.dark(), // Definir tema escuro
      home: MyForm,
      // home: MySystemWidget(
      //   applicationName: "Meu Sistema",
      //   applicationVersion: "1.0",
      //   menus: [
      //     {"rota": "/cliente", "nome": "Cliente"},
      //     {
      //       "rota": "/financeiro",
      //       "nome": "Financeiro",
      //       "menus": [
      //         {"rota": "/financeiro/contas", "nome": "Contas"}
      //       ]
      //     }
      //   ],
      //   usuarios: [
      //     {"nome": "Rodrigo", "avatar": "", "status": "online"},
      //     {"nome": "José", "avatar": "", "status": "offline"},
      //   ],
      // ),
    );
  }
}

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
    return Scaffold(
      appBar: AppBar(
        title: Text(applicationName),
        actions: [
          IconButton(
            icon: Icon(Icons.settings),
            onPressed: () {
              // Lógica para abrir a tela de configurações
            },
          ),
        ],
      ),
      body: Column(
        children: [
          // Cabeçalho fixo ocupando 10% da tela
          Container(
            height: MediaQuery.of(context).size.height * 0.1,
            alignment: Alignment.center,
            child: SizedBox(
              width: MediaQuery.of(context).size.width * 0.3,
              child: TextField(
                textAlign: TextAlign.center,
                decoration: InputDecoration(
                  hintText: 'Pesquisar...',
                  border: OutlineInputBorder(),
                ),
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
                  color: Colors.grey[200], // Cor de fundo para o menu
                  child: ListView.builder(
                    itemCount: menus.length,
                    itemBuilder: (context, index) {
                      return MouseRegion(
                        cursor: SystemMouseCursors.click, // Define o cursor como pointer
                        child: InkWell(
                          onTap: () {
                            // Lógica para lidar com o clique do menu
                            print("Clicou em ${menus[index]["nome"]}");
                          },
                          child: Container(
                            color: Colors.grey[300], // Cor de fundo para hover mais escuro
                            child: ListTile(
                              title: Text(menus[index]["nome"]),
                            ),
                          ),
                        ),
                      );
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
      ),
    );
  }
}
