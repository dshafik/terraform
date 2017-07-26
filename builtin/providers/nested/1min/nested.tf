resource "nested" "daveyshafikcom" {
  rule {
    comment = "The default rule applies to all requests"

    behavior {
      name = "allowPost"
      option {
        name = "enabled"
        flag = true
      }
      option {
        name = "allowWithoutContentLength"
        flag = true
      }
    }

    behavior {
      name = "realUserMonitoring"
      option {
        name = "enabled"
        flag = true
      }
    }

    rule {
      name = "Redirect to HTTPS"
      comment = "Redirect to the same URL on HTTPS protocol, issuing a 301 response code (Moved Permanently). You may change the response code to 302 if needed."

      criteria {
        name = "requestProtocol"
        option {
          name = "value"
          value = "HTTP"
        }
      }

      criteria {
        name = "hostname"
        option {
          name = "matchOperator"
          value = "IS_NOT_ONE_OF"
        }

        option {
          name = "values"
          values = ["staging.daveyshafik.com"]
        }
      }

      behavior {
        name = "redirect"

        option {
          name = "mobileDefaultChoice"
          value = "DEFAULT"
        }
        option {
          name = "destinationProtocol"
          value = "HTTPS"
        }
        option {
          name = "destinationHostname"
          value = "OTHER"
        }
        option {
          name = "destinationHostnameOther"
          value = "www.daveyshafik.com"
        }
        option {
          name = "destinationPath"
          value = "SAME_AS_REQUEST"
        }
        option {
          name = "queryString"
          value = "APPEND"
        }
        option {
          name = "responseCode"
          value = "301"
        }
      }
    }
  }
}