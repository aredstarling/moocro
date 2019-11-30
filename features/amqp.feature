Feature: AMQP Moocro

  Describe the AMQP implementation of Moocro

  Scenario: Greeting action
    When the AMQP system sends a name of "Bob" to the greeting action
    Then the AMQP system should send back "Hello Bob"

  Scenario: Find a greeting and receive a found greeting
    When the AMQP system sends a name of "Bob" to the find greeting action
    Then the AMQP system should respond "Hello Bob" as a greeting found
