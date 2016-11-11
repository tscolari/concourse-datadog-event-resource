# datadog event resource

First of all, add the `datadog-event` to your resource types list:

```yaml
resource_types:
- name: datadog-event
  type: docker-image
  source:
    repository: tscolari/datadog-event-resource
```

## Source Configuration:

* `auth`: Required. Struct that holds api key and application key:
  * `api_key`: Required. Datadog api key.
  * `application_key`: Required. Datadog application key.
  ```yaml
  auth:
    api_key: "1234-5"
    application_key: "123444"
  ```
* `priority`: Optional. Default priority. It' used to filter checks and to create events.
* `tags`: Optional. Array of default tags. It's used to filter checks and to create events.
* `title_prefix`: Optional. String to be prepended to the events title. Used to filter checks and to create events.

### Example:

```yaml
resources:
- name: my-app-resource
  type: datadog-event
  source:
    auth:
      api_key: "12345"
      application_key: "123456"
    priority: "high"
    tags:
    - production
    - helps
    title_prefix: "my-app-1"
```

## Behavior

###`check`: check for new events.

**Parameters:**
* `priority`: same as from `source` but with higher precedence
* `tags`: same as from `source` but with higher precedence
* `title_prefix`: same as from `source` but with higher precedence

###`in`: fetches event and writes it as json.

Will create a `event.json` inside the resource path containing the event information.

**Parameters:**
* `title_prefix`: same as from `source` but with higher precedence

###`out`: check for new events.

Creates an event with the given information.

**Parameters:**
* `title_prefix`: same as from `source` but with higher precedence
* `event`; Required. Event structure to be created
  * `title`: Required. Event title
  * `text`: Required. Event text
  * `time`: Optional. Event time in UNIX format. Will default to time.Now() otherwise.
  * `alert_type`: Optional. Event alert type
  * `source_type`: Optional. Event source type
  * `host`: Optional. Event host
  * `Resource`: Optional. Event resource
  * `Url`: Optional. Event url
  * `tags`: same as from `source` but with higher precedence
  * `priority`: same as from `source` but with higher precedence

### Example:

```yaml
- put: my-app-event
  params:
    event:
      title: "panic"
      text: "all panic"
      tags: ["app"]
```
