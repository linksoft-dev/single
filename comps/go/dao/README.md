# Db Crud interface Diagram (Simplified)

This is a simplified example of a class diagram using the Mermaid notation.

```mermaid
classDiagram
  class Crud {
    // Default interface for data access
    Create(obj T): (T, error)
    Update(obj T, fields UpdateField): error
    Delete(T): error
    DeleteById(id string): error
    Find(filter Query): ([]T, error)
    FindById(id string): (T, error)
  }

  class Jsonb {
    // Implementation of the Crud interface for Postgres JSONB
    Create(obj T): (T, error)
    Update(obj T, fields UpdateField): error
    Delete(T): error
    DeleteById(id string): error
    Find(filter Query): ([]T, error)
    FindById(id string): (T, error)
  }

  class Redis {
    // Implementation of the Crud interface for Redis
    Create(obj T): (T, error)
    Update(obj T, fields UpdateField): error
    Delete(T): error
    DeleteById(id string): error
    Find(filter Query): ([]T, error)
    FindById(id string): (T, error)
  }

  Crud --|> Jsonb
  Crud --|> Redis
