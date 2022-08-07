locals {
  events = {for event in var.events : event["UserPoolId"] => event if var.cron_enabled}
}