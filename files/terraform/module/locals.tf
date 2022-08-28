locals {
  events = {for event in var.events : event["cognitoUserPoolId"] => event if var.cron_enabled}
}