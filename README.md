## How to Install

Do the following command
```
mkdir -p your_path/pre-test
cd your_path/pre-test
git clone https://github.com/iman-khaeruddin/pre-test.git
```

## Directory Structure

Get into the `pre-test` folder by :

```
cd pre-test/
```
And you should now find the source code under `pre-test` directory.

```
  + your_path/
  |
  +--+ src/
  |  |
  |  +--+ github.com/
  |     |
  |     +--+ pre-test/
  |         |
  |         +--+ main.go
  |            + app/
  |            + controller/
  |            + database/

```

## How to Run
```
go run main.go
```

## How to Test
There are 3 end points :
- Post a message
```
curl -X POST http://localhost:8080/post/**your_massage**
```
- Get all message
```
curl -X GET http://localhost:8080/get/all-message
```
- Get all message with long life connection

You can use this [tool](https://chrome.google.com/webstore/detail/websocket-test-client/fgponpodhbmadfljofbimhhlengambbn)
