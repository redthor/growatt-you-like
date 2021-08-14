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

  triggers {
    threshold_type   = "LessThanOrEqual"
    threshold        = 0
    time_range       = "15m"
    occurrence_type  = "ResultCount"
    trigger_source   = "AllResults"
    trigger_type     = "Critical"
    detection_method = "StaticCondition"
  }

  triggers {
    threshold_type   = "GreaterThan"
    threshold        = 0
    time_range       = "15m"
    occurrence_type  = "ResultCount"
    trigger_source   = "AllResults"
    trigger_type     = "ResolvedCritical"
    detection_method = "StaticCondition"
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

// I wanted to use these but the response from SL API was "get in contact with account representative".
// The "triggers" API is deprecated, but can't access "trigger_conditions"
//
//  trigger_conditions {
//    logs_static_condition {
//      field = "_count"
//      critical {
//        time_range = "15m"
//        alert {
//          threshold = 0
//          threshold_type = "LessThanOrEqual"
//        }
//        resolution {
//          threshold = "0"
//          threshold_type = "GreaterThan"
//        }
//      }
//    }
//    logs_outlier_condition {
//      direction = "Both"
//      critical {
//        consecutive = 1
//        window      = 10
//      }
//    }
//    logs_missing_data_condition {
//      time_range = "15m"
//    }
//  }
}