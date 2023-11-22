import 'package:flutter/material.dart';
import 'package:single/apps/contas/Form.dart';
import 'package:single/pages/app.dart';

import 'apps/person/Form.dart';

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

