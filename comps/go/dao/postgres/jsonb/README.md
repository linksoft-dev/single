This package implements the interface dao.Crud for Postgres JSONB field.

This implementation have in mind multi tenant architecture where each tenant has its 
own table in the database, here are some key features of this implementation.

- Each tenant has their table in database, having completely isolation the data between tenants
- Tables are created automatically for each tenant if needed
- All the objects are stored as **doc** field, this make this implementation schemaless, no worries about schema changes
- All tenant space/table has a fixed structured and optimized
- 