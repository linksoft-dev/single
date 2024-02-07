import 'dart:async';

// Interface para operações de CRUD
abstract class Dao<T> {
  Future<void> create(T object);
  Future<void> update(T object);
  Future<void> delete(T object);
  Future<List<T>> list();
  Future<T> get(int id);
}

class Dataset<T> {
  final Dao<T> dao;
  final List<T> _data;
  int _currentIndex = 0;

  // Eventos
  final StreamController<void> _onNextController = StreamController<void>.broadcast();
  final StreamController<void> _onPriorController = StreamController<void>.broadcast();
  final StreamController<void> _onFirstController = StreamController<void>.broadcast();
  final StreamController<void> _onLastController = StreamController<void>.broadcast();
  final StreamController<void> _onBeforeController = StreamController<void>.broadcast();
  final StreamController<void> _onAfterController = StreamController<void>.broadcast();

  Dataset(this.dao, this._data);

  // Métodos de navegação
  T get current => _data[_currentIndex];

  void next() {
    _emitEvent(_onNextController);
    if (_currentIndex < _data.length - 1) {
      _currentIndex++;
    }
  }

  void prior() {
    _emitEvent(_onPriorController);
    if (_currentIndex > 0) {
      _currentIndex--;
    }
  }

  void first() {
    _emitEvent(_onFirstController);
    _currentIndex = 0;
  }

  void last() {
    _emitEvent(_onLastController);
    _currentIndex = _data.length - 1;
  }

  // Métodos CRUD
  Future<void> create(T object) async {
    await _emitBeforeEvent();
    await dao.create(object);
    await _emitAfterEvent();
  }

  Future<void> update(T object) async {
    await _emitBeforeEvent();
    await dao.update(object);
    await _emitAfterEvent();
  }

  Future<void> delete(T object) async {
    await _emitBeforeEvent();
    await dao.delete(object);
    await _emitAfterEvent();
  }

  Future<List<T>> list() async {
    return await dao.list();
  }

  Future<T> get(int id) async {
    return await dao.get(id);
  }

  // Método privado para emitir eventos
  Future<void> _emitBeforeEvent() async {
    _emitEvent(_onBeforeController);
  }

  Future<void> _emitAfterEvent() async {
    _emitEvent(_onAfterController);
  }

  void _emitEvent(StreamController<void> controller) {
    if (!controller.isClosed) {
      controller.add(null);
    }
  }

  // Métodos para inscrever-se nos eventos
  Stream<void> get onNext => _onNextController.stream;
  Stream<void> get onPrior => _onPriorController.stream;
  Stream<void> get onFirst => _onFirstController.stream;
  Stream<void> get onLast => _onLastController.stream;
  Stream<void> get onBefore => _onBeforeController.stream;
  Stream<void> get onAfter => _onAfterController.stream;
}

// Exemplo de uso:

// Implementação de Dao para uma classe específica
class ExampleDao<T> implements Dao<T> {
  @override
  Future<void> create(T object) async {
    // Implementação de create
  }

  @override
  Future<void> update(T object) async {
    // Implementação de update
  }

  @override
  Future<void> delete(T object) async {
    // Implementação de delete
  }

  @override
  Future<List<T>> list() async {
    // Implementação de list
    return <T>[];
  }

  @override
  Future<T> get(int id) async {
    // Implementação de get
    return null;
  }
}

void main() {
  // Exemplo de uso
  final dao = ExampleDao<int>();
  final dataset = Dataset<int>(dao, [1, 2, 3, 4, 5]);

  // Inscrever-se nos eventos
  dataset.onNext.listen((_) => print('Next'));
  dataset.onPrior.listen((_) => print('Prior'));
  dataset.onFirst.listen((_) => print('First'));
  dataset.onLast.listen((_) => print('Last'));
  dataset.onBefore.listen((_) => print('Before CRUD'));
  dataset.onAfter.listen((_) => print('After CRUD'));

  // Operações de navegação
  dataset.next(); // Next
  dataset.prior(); // Prior
  dataset.first(); // First
  dataset.last(); // Last

  // Operações CRUD
  dataset.create(6); // Before CRUD, After CRUD
  dataset.update(7); // Before CRUD, After CRUD
  dataset.delete(8); // Before CRUD, After CRUD
}
