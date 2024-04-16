
// Interface para um repositório genérico
abstract class Dao<T> {
  Future<List<T>> getAll();
  Future<T> getById(String id);
  Future<void> insert(T item);
  Future<void> update(T item);
  Future<void> delete(String id);
}

// Exemplo de implementação de um repositório para uma entidade específica
class ExampleRepository<T> implements Dao<T> {
  // Implemente métodos da interface
  @override
  Future<List<T>> getAll() {
    // Lógica para recuperar todos os itens
    throw UnimplementedError();
  }

  @override
  Future<T> getById(String id) {
    // Lógica para recuperar um item por ID
    throw UnimplementedError();
  }

  @override
  Future<void> insert(T item) {
    // Lógica para inserir um item
    throw UnimplementedError();
  }

  @override
  Future<void> update(T item) {
    // Lógica para atualizar um item
    throw UnimplementedError();
  }

  @override
  Future<void> delete(String id) {
    // Lógica para excluir um item por ID
    throw UnimplementedError();
  }
}
