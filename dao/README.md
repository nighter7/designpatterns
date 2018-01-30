# DAO - Data Access Object

Data Access Object is a design pattern that help hiding the implementation of data access from the rest of the application.

### Situation

> You are designing an application that is prepared to persist and get its data from different sources (databases, files, legacy systems, APIs)

A not-so-easy-maintainable-sadly-common approach is to have your database access code right in the business logic and do all the data manipulation right there.

### Problem

Every time it is needed to change for maintenance or bug fix the data manipulation code, it would be needed to change all the implementations where the code is as the code is not centralized

### Solution

Using DAO DP we achieve centralization, maintainability and layer approach

### Implementation

- ***DTO***: We will need to use a Plain Old X Object (replace X for Java, Go, PHP, Python whatever you want) that will be the mean used to hold the data coming from or to our DAO.  
`In this context let's call this POXO: DTO - Data Transport Object`

- ***Actions Interface***: As any interface, this one will hold all the allowed operations for the modularization
    > Add
    > Delete
    > Update
    > List or get

- ***Concreate Implementation***: Class or classes that will implement the actions interface.  Many implementations could exists, let´s say, you want to save to a database or files or API, it does not matter the concrete implementation for the rest of the application as its methods will be consumed using the actions interface.
