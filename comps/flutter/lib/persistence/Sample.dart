void main() {
  // Criando um dataset fictício
  final dataset = Dataset([
    {'name': 'John', 'age': 30},
    {'name': 'Alice', 'age': 25},
  ]);

  runApp(MyApp(dataset: dataset));
}

class MyApp extends StatelessWidget {
  final Dataset dataset;

  const MyApp({Key key, @required this.dataset}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: Text('Dataset Example'),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              DbField(dataset: dataset, fieldName: 'name'),
              DbField(dataset: dataset, fieldName: 'age'),
            ],
          ),
        ),
      ),
    );
  }
}