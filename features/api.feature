# file: version.feature
Feature: get version
  In order to register 
  as a Cluber in Fiesta
  I need a email and password

  Scenario: Sign up as simple user 
    When I send "POST" request to "/v1/register" with json:
    """
      {
        "mail": "ahidoyatov@gmail.com",
        "password": "bugagaga"
      }
    """
    Then the response code should be 200
#    And the response should match json:
#      """
#      {
#        "error": "Method not allowed"
#      }
#      """

  Scenario: Get Clubs
    When I send "GET" request to "/v1/clubs/"
    Then the response code should be 404
