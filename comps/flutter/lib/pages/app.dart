import 'package:flutter/material.dart';

class Pages {
  String title, route, menu = '';
  WidgetBuilder widget;

  Pages(this.title, this.route, this.widget);
}

class App extends StatelessWidget {
  String title;
  late List<Pages> pages = [];

  App(this.title, this.pages);

  @override
  Widget build(BuildContext context) {
    Map<String, WidgetBuilder> routes = {};

    // add all routers of te pages to routers
    this.pages.forEach((element) {
      routes[element.route] = element.widget;
    });
    var materialApp = MaterialApp(
      title: this.title,
      routes: routes,
      home: HomePage(this.title, this.pages),
    );

    return materialApp;
  }
}

class HomePage extends StatefulWidget {
  late String title;
  late List<Pages> pages;

  HomePage(this.title, this.pages);

  @override
  _MyHomePageState createState() => _MyHomePageState(this.title, this.pages);
}

class _MyHomePageState extends State<HomePage> {
  String _selectedRouter = "testeee";
  late String title = "";
  late List<Pages> pages = [];
  late List<InkWell> menus = [];
  late Map<String, Pages> openedScreen= {};
  ListView menu = ListView();
  // get pages => this.pages;
  // String get title => this.title;
  _MyHomePageState(this.title, this.pages);

  Map<String, Pages> routes = {};

  selectRoute(Pages page) {
    setState(() {
      this._selectedRouter = page.route;
      this.openedScreen[page.route] = page;
    });
  }

  // store information if some resource was already added
  Map<String, bool> alreadyAdded = {};

  buildDrawer(){
    List<Widget> list = [
      const DrawerHeader(
        decoration: BoxDecoration(
          color: Colors.blue,
        ),
        child: Text('Usuário: ddRodrigo'),
      ),
    ];

    // list.add(menu);

    return Drawer(
      // Add a ListView to the drawer. This ensures the user can scroll
      // through the options in the drawer if there isn't enough vertical
      // space to fit everything.
      child: ListView(
        // Important: Remove any padding from the ListView.
          padding: EdgeInsets.zero,
          children: list),
    );
  }

  buildMenu() {
    // if (alreadyAdded[page.title] == true) {
    //   return;
    // }
    menu =  ListView.builder(
      itemCount: pages.length,
      itemBuilder: (context, index) {
        return InkWell(
          onTap: () {
            selectRoute(pages[index]);
            // Aqui você pode adicionar a lógica para ação ao clicar em um item do menu
            // print('Clicou em ${pages[index]}');
          },
          child: Padding(
            padding: EdgeInsets.symmetric(vertical: 16.0),
            child: Row(
              children: [
                SizedBox(width: 10),
                Icon(Icons.circle, size: 10, color: Colors.blue), // Ícone para destacar o item
                SizedBox(width: 10),
                Text(
                  pages[index].title,
                  style: TextStyle(fontSize: 14),
                ),
              ],
            ),
          ),
          hoverColor: Colors.amberAccent, // cor ao passar o mouse
          splashColor: Colors.transparent, // cor do splash ao clicar
          highlightColor: Colors.transparent, // cor de destaque ao clicar
          mouseCursor: SystemMouseCursors.click, // alterar cursor ao passar o mouse
        );
      },
    );
  }

  buildRouters() {
    this.pages.forEach((element) {
      routes[element.route] = element;
    });
  }

  @override
  Widget build(BuildContext context) {
    buildMenu();
    buildDrawer();
    buildRouters();

    var scaffold = Scaffold(
      appBar: AppBar(title: Text(title)),
      body: Row(
        children: [
          // Menu da esquerda
          Container(
            width: 300,
            // color: Colors.grey[300],
            child: menu,
            // child:  ListView.builder(
            //   itemCount: menus.length,
            //   itemBuilder: (context, index) {
            //     return menus[index];
            //   },
            // ),
          ),
          // Formulário exibido no meio da tela
          Expanded(
            child: Padding(
              padding: EdgeInsets.all(16.0),
              // child: Navigator(
              //   onGenerateRoute: (settings) {
              //     return MaterialPageRoute(
              //       builder: (context) {
              //
              //         return  HomeScreen();
              //         // switch (settings.name) {
              //         //   case '/':
              //         //     return Screen1();
              //         //   case '/screen2':
              //         //     return Screen2();
              //         //   case '/screen3':
              //         //     return Screen3();
              //         //   default:
              //         //     return Screen1(); // Tela padrão
              //         // }
              //       },
              //     );
              //   },
              // ),
              child: Column(children: [
                // Primeira linha (cabeçalho fixo)
                Container(
                  decoration: BoxDecoration(
                    border: Border(
                      bottom: BorderSide(
                          width: 1, color: Colors.grey), // Borda inferior fina
                    ),
                  ),
                  height: 60, // Altura do cabeçalho
                  child: Container(
                    // color: Colors.blue,
                    decoration: BoxDecoration(
                      border: Border(
                        bottom: BorderSide(
                            width: 1,
                            color: Colors.grey), // Borda inferior fina
                      ),
                    ),
                    padding: EdgeInsets.all(16.0),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.start,
                      children: openedScreen.keys.map((String title) {
                        return Row(
                          mainAxisAlignment: MainAxisAlignment.start, // Alinhar à esquerda
                          children: [
                            ElevatedButton(
                              onPressed: () {
                                selectRoute(openedScreen[title]!);
                              },
                              style: ButtonStyle(
                                backgroundColor: MaterialStateProperty.resolveWith<Color?>(
                                      (Set<MaterialState> states) {
                                    if (openedScreen[title]?.route == _selectedRouter) {
                                      return Colors.blue; // Cor do botão pressionado
                                    } else {
                                      return Colors.grey; // Sem cor de pressionado
                                    }
                                  },
                                ),
                              ),
                              child: Text(openedScreen[title]!.title),
                            ),
                            SizedBox(width: 10), // Espaçamento entre os botões
                          ],
                        );
                      }).toList(),
                      // children: [
                      //   Icon(Icons.account_circle),
                      //   SizedBox(width: 10),
                      //   Text(
                      //     'Cabeçalho Fixo',
                      //     style: TextStyle(
                      //       fontSize: 18,
                      //       fontWeight: FontWeight.bold,
                      //     ),
                      //   ),
                      // ],
                    ),
                  ),
                ),
                // Segunda linha (expandida)
                Expanded(
                  child: Container(
                    color: Colors.green,
                    child: Center(
                      // child: Text('Home  $_selectedRouter'),

                      child: this.routes.length > 0
                          ? this.routes[_selectedRouter]?.widget(context)
                          : Text('Home' + _selectedRouter),
                    ),
                  ),
                ),
              ]),
            ),
          ),
        ],
      ),
      drawer: buildDrawer(),
    );
    return scaffold;
  }
}
//
// class RouteAwareWidget extends StatefulWidget {
//   @override
//   _RouteAwareWidgetState createState() => _RouteAwareWidgetState();
// }
//
// class _RouteAwareWidgetState extends State<RouteAwareWidget> {
//   late String _currentRoute;
//
//   @override
//   void didChangeDependencies() {
//     super.didChangeDependencies();
//     _currentRoute = ModalRoute.of(context)?.settings.name ?? '/';
//   }
//
//
//   @override
//   Widget build(BuildContext context) {
//     _currentRoute = getCurrentRoute(context);
//     Widget currentWidget;
//     switch (_currentRoute) {
//       case '/form1':
//         currentWidget = FirstScreen();
//         break;
//       case '/form2':
//         currentWidget = SecondScreen();
//         break;
//       default:
//         currentWidget = FirstScreen();
//     }
//
//     return currentWidget;
//   }
// }

String getCurrentRoute(BuildContext context) {
  return ModalRoute.of(context)?.settings.name ?? '/';
}

class HomeScreen extends StatelessWidget {
  const HomeScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Home Screen'),
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
