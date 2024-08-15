# Prayer Schedule API

A simple api to provides prayer time written in Go.

## Endpoint

`GET /api/prayer-schedule`

### Parameters

| Parameter  | Type   | Required | Description                                                                                                                                    | Example              |
| ---------- | ------ | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | -------------------- |
| `lat`      | float  | Yes      | Latitude of the location.                                                                                                                      | `-7.275687449166608` |
| `lon`      | float  | Yes      | Longitude of the location.                                                                                                                     | `112.7936484102365`  |
| `year`     | int    | Yes      | The year for which the prayer schedule is requested.                                                                                           | `2024`               |
| `timezone` | string | No       | The timezone in which the prayer times should be returned. If not provided, it will automatically generate timezone based on given coordinate. | `Asia/Jakarta`       |

### Example Request

```bash
curl 'http://localhost:8080/api/prayer-schedule?lat=-7.275687449166608&lon=112.7936484102365&year=2024&timezone=Asia%2FJakarta'
```

### Example Response

```json
{
  "status": 200,
  "timezone": "Asia/Jakarta",
  "data": [
    {
      "date": "2024-01-01",
      "schedule": {
        "fajr": "2024-01-01T03:52:12+07:00",
        "sunrise": "2024-01-01T05:11:45+07:00",
        "zuhr": "2024-01-01T11:35:00+07:00",
        "asr": "2024-01-01T15:01:03+07:00",
        "maghrib": "2024-01-01T17:50:13+07:00",
        "isha": "2024-01-01T19:06:38+07:00"
      },
      "readableSchedule": {
        "fajr": "03:52 WIB",
        "sunrise": "05:11 WIB",
        "zuhr": "11:35 WIB",
        "asr": "15:01 WIB",
        "maghrib": "15:01 WIB",
        "isha": "19:06 WIB"
      }
    }
  ]
}
```
