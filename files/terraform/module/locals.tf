locals {
  events = {for event in var.events : event["cognitoUserPoolID"] => event if var.cron_enabled}
}