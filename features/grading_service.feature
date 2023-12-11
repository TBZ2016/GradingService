Feature: Grading Service

  Background:
    Given there is a grading service

  Scenario: Retrieve grades by cursus ID
    Given a cursus ID
    When the user requests grades for the given cursus ID
    Then the service should return a list of grades for that cursus ID

  Scenario: Create a new grade
    Given grade details (student ID, teacher ID, assignment ID, course ID, grade, isPass)
    When the user creates a new grade with the provided details
    Then the service should confirm the successful creation of the grade

  Scenario: Retrieve grades by student ID
    Given a student ID
    When the user requests grades for the given student ID
    Then the service should return a list of grades for that student ID

  Scenario: Retrieve grades by class ID
    Given a class ID
    When the user requests grades for the given class ID
    Then the service should return a list of grades for that class ID

  Scenario: Retrieve a grade by ID
    Given a grade ID
    When the user requests a grade for the given grade ID
    Then the service should return the details of that grade

  Scenario: Update an existing grade
    Given an existing grade with updated details
    When the user updates the grade with the new details
    Then the service should confirm the successful update of the grade

  Scenario: Delete a grade by ID
    Given a grade ID
    When the user requests to delete the grade with the given ID
    Then the service should confirm the successful deletion of the grade
