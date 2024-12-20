import 'package:flutter/material.dart';

class CrudFormDetail extends StatefulWidget {
  final String screenName;
  final Widget body;
  List<Widget> customButtons = [];

  CrudFormDetail({
    required this.screenName,
    required this.body,
    required this.customButtons ,
  });

  @override
  _CrudFormDetail createState() => _CrudFormDetail(
      this.screenName,
      this.body,
      this.customButtons,
  );
}

class _CrudFormDetail extends State<CrudFormDetail> {
  final Widget body;
  List<Widget>? customButtons = [];

  _CrudFormDetail(
      this._currentScreen,
      this.body,
      this.customButtons,
      );

  Widget _widgetToShow = Container(); // Initial empty widget
  String _currentScreen = "base"; // Changed to private
  String get currentScreen => _currentScreen;

  void _include() {
    setState(() {
      _widgetToShow = Text('Including in $currentScreen'); // Include widget
    });
  }

  void _edit() {
    setState(() {
      _widgetToShow = Text('Edit page'); // Edit widget
    });
  }

  void _save() {
    setState(() {
      _widgetToShow = Text('Saving...'); // Simulated save action
    });
  }

  void _delete() {
    setState(() {
      _widgetToShow = Text('Delete page'); // Delete widget
    });
  }

  buildButtons() {
    List<Widget> buttons = [
      ElevatedButton(onPressed: _include, child: Text('Inlcuir')),
      ElevatedButton(onPressed: _edit, child: Text('Editar')),
      ElevatedButton(onPressed: _save, child: Text('salvar')),
      ElevatedButton(onPressed: _delete, child: Text('Delete')),
    ];
    if (customButtons == null) {
      return buttons;
    }
    return buttons + customButtons!;
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Header with buttons
        Container(
          padding: EdgeInsets.all(16.0),
          color: Colors.grey[300],
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: buildButtons(),
          ),
        ),
        // Area to display dynamic widgets
        Expanded(
          child: SingleChildScrollView(
            padding: EdgeInsets.all(16.0),
            child: _widgetToShow,
          ),
        ),
      ],
    );
  }
}
