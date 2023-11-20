import 'package:flutter/material.dart';

class MyForm extends StatefulWidget {
  @override
  _MyFormState createState() => _MyFormState();
}

class _MyFormState extends State<MyForm> {
  Widget _widgetToShow = Container(); // Initial empty widget

  void _include() {
    setState(() {
      _widgetToShow = Text('Include page'); // Include widget
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
            children: [
              ElevatedButton(onPressed: _include, child: Text('Include')),
              ElevatedButton(onPressed: _edit, child: Text('Edit')),
              ElevatedButton(onPressed: _save, child: Text('Save')),
              ElevatedButton(onPressed: _delete, child: Text('Delete')),
            ],
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