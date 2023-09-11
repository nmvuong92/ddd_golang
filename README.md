# DDD golang

# Domain
 - customer
    - repositories (interface)
    - in_memory-folder (implementation)
    - mysql_folder  (implementation)
    - mongo_db_folder (implementation)
    - ...
 -

# Subdomain

# Aggregates
- unique identifier by root Entity
- Multiple Entities / Value Objects combined
- Aggregate like a Container
- Aggregate is a business logic for customer should be inside inside a aggregates not entity
- use Factory for instantiated
- use Repository to store data (Db, disk, cloud, ram...)
- combination of entities or valueobject
- we store/manage them use repository
- repository store and manage aggregate
- repository hide implementation detail behind the interface
- repository allow to build moduler and changable software
- in-memory repository for unit test or mysql, mongodb repository... (we can change)
  # Entity
  - Unique Identifier
  - Mutable

  # Value Object
  - No Identifier
  - Immutable
  # Factory
  - create instance (complex stuff)
  - how to create a customer eg.
  - validate input
  - don't care about how to create instance
  - care about input, output


  # repository
  - to manage aggregates
  - one repository handle one aggregate
  - inside memory repository we cannot access to the aggregate => we need to fix that
  - => add Get/Set for Customer Aggregate
  - we cannot modify data in aggregate directly customer_aggregate.prop = value (wrong)
  - so we need to create some exposed function to handle that




