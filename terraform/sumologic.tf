provider "sumologic" {
}

resource "sumologic_collector" "base" {
  name        = "base"
  description = "Logs from the base app"
}

resource "sumologic_http_source" "http_source" {
  name         = "http-base-growatt-to-iot"
  description  = "Growatt To IOT Base App"
  category     = "prod/base/growatt-to-iot"
  collector_id = sumologic_collector.base.id
}

resource "sumologic_monitor" "log_monitor_base_growatt_to_iot_missing_data" {
  name         = "log-monitor-base-growatt-to-iot-missing-data"
  description  = "Ensure we are getting logs"
  monitor_type = "Logs"

  queries {
    row_id = "A"
    query  = "_collector=${sumologic_collector.base.id}"
  }

  trigger_conditions {
    logs_static_condition {
      field = "_count"
      critical {
        time_range = "15m"
        alert {
          threshold = 0
          threshold_type = "LessThanOrEqual"
        }
        resolution {
          threshold = "0"
          threshold_type = "GreaterThan"
        }
      }
    }
    logs_missing_data_condition {
      time_range = "15m"
    }
  }

  notifications {
    notification {
      connection_type = "Email"
      recipients      = ["${var.sumologic_trigger_email}"]
      subject         = "Monitor Alert: {{TriggerType}} on {{Name}}"
      time_zone       = var.sumologic_time_zone
      message_body    = "Triggered {{TriggerType}} Alert on {{Name}}: {{QueryURL}}"
    }
    run_for_trigger_types = ["Critical", "ResolvedCritical"]
  }
}