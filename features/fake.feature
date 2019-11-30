Feature: Fake Moocro

  Describe the fake implementation of Moocro

  Scenario: Greeting action
    When the fake system sends a name of "Bob" to the greeting action
    Then the fake system should send back "Hello Bob"

  Scenario: Find a greeting and receive a found greeting
    When the fake system sends a name of "Bob" to the find greeting action
    Then the fake system should respond "Hello Bob" as a greeting found
