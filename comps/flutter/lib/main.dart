import 'package:flutter/material.dart';
import 'package:single/apps/contas/Form.dart';
import 'package:single/pages/app.dart';
import 'package:single/pages/app2.dart';
// import 'apps/person/Form.dart';
import 'apps/person/Form.dart';
import 'forms/CrudFormDetail.dart';

void main() {
  List<Pages> pages = [];
  pages.add(Pages("Pessoa", "/pessoa", (context) => PersonForm()));
  pages.add(Pages("Contas Receber", "/teste2", (context) => ContasForm()));
  var app = App("crud", pages);

  runApp(app);
}


class FirstScreen extends StatelessWidget {
  const FirstScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('First Screen'),
      ),
      body: Center(
        child: ElevatedButton(
          onPressed: () {
            // Navigate to the second screen when tapped.
          },
          child: const Text('Launch screen'),
        ),
      ),
    );
  }
}

class SecondScreen extends StatelessWidget {
  const SecondScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Second Screen'),
      ),
      body: Center(
        child: ElevatedButton(
          onPressed: () {
            // Navigate back to first screen when tapped.
          },
          child: const Text('Go back!'),
        ),
      ),
    );
  }
}
//
// class MyApp extends StatelessWidget {
//   final Map<String, Widget> routes = {
//     '/cliente': PersonForm(), // Substitua PersonForm pelo seu formulário de cliente
//     // Defina outras rotas aqui, se necessário
//   };
//
//   @override
//   Widget build(BuildContext context) {
//     return MaterialApp(
//       theme: ThemeData.light(), // Definir tema claro como tema padrão
//       darkTheme: ThemeData.dark(), // Definir tema escuro
//       // home: MyForm,
//       routes:  {
//         // '/': (context) => PersonForm(),
//         '/person': (context) => PersonForm(),
//       },
//       home: MySystemWidget(
//         applicationName: "Meu Sistema",
//         applicationVersion: "1.0",
//         menus: [
//           {"rota": "/cliente", "nome": "Cliente"},
//           {
//             "rota": "/financeiro",
//             "nome": "Financeiro",
//             "menus": [
//               {"rota": "/financeiro/contas", "nome": "Contas"}
//             ]
//           }
//         ],
//         usuarios: [
//           {"nome": "Rodrigo", "avatar": "", "status": "online"},
//           {"nome": "José", "avatar": "", "status": "offline"},
//         ],
//       ),
//     );
//   }
// }
//
// class MySystemWidget extends StatelessWidget {
//   final String applicationName;
//   final String applicationVersion;
//   final List<Map<String, dynamic>> menus;
//   final List<Map<String, dynamic>> usuarios;
//
//   const MySystemWidget({
//     Key? key,
//     required this.applicationName,
//     required this.applicationVersion,
//     required this.menus,
//     required this.usuarios,
//   }) : super(key: key);
//
//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       appBar: AppBar(
//         title: Text(applicationName),
//         actions: [
//           IconButton(
//             icon: Icon(Icons.settings),
//             onPressed: () {
//               // Lógica para abrir a tela de configurações
//             },
//           ),
//         ],
//       ),
//       body: Column(
//         children: [
//           // Cabeçalho fixo ocupando 10% da tela
//           Container(
//             height: MediaQuery.of(context).size.height * 0.1,
//             alignment: Alignment.center,
//             child: SizedBox(
//               width: MediaQuery.of(context).size.width * 0.3,
//               child: TextField(
//                 textAlign: TextAlign.center,
//                 decoration: InputDecoration(
//                   hintText: 'Pesquisar...',
//                   border: OutlineInputBorder(),
//                 ),
//               ),
//             ),
//           ),
//           // Resto da tela dividida em duas colunas
//           Expanded(
//             child: Row(
//               children: [
//                 // Primeira coluna com menus e informações
//                 Container(
//                   width: MediaQuery.of(context).size.width * 0.15,
//                   color: Colors.grey[200], // Cor de fundo para o menu
//                   child: ListView.builder(
//                     itemCount: menus.length,
//                     itemBuilder: (context, index) {
//                       return MouseRegion(
//                         cursor: SystemMouseCursors.click, // Define o cursor como pointer
//                         child: InkWell(
//                           onTap: () {
//                             // Lógica para lidar com o clique do menu
//                             print("Clicou em ${menus[index]["nome"]}");
//                           },
//                           child: Container(
//                             color: Colors.grey[300], // Cor de fundo para hover mais escuro
//                             child: ListTile(
//                               title: Text(menus[index]["nome"]),
//                             ),
//                           ),
//                         ),
//                       );
//                     },
//                   ),
//                 ),
//                 // Segunda coluna para mostrar telas abertas
//                 Expanded(
//                   child: Container(
//                     color: Colors.grey[200], // Cor de fundo para mostrar telas abertas
//                     // Aqui você pode adicionar a lógica para exibir as telas abertas
//                     child: Center(
//                       child: Text("Telas abertas"),
//                     ),
//                   ),
//                 ),
//               ],
//             ),
//           ),
//         ],
//       ),
//     );
//   }
// }
