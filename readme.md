# 1. Go Durations

[![Go Report Card](https://goreportcard.com/badge/github.com/ElecTwix/uploader)](https://goreportcard.com/badge/github.com/ElecTwix/uploader) [![PkgGoDev](#)](#)


## 1.1. What is this for

Uploader can upload files with in one line command and give back respond struct decoded.

## 1.2. Why

For people don't wanna rewrite upload code for them selfs and make things faster.

## 1.3. Usage

``` 	
resp ,err := uploader.Upload(uploader.BayFiles, file)
fmt.Println(resp.Data.File.Url) 

or

resp ,err := uploader.Upload(uploader.AnonFiles, file)
fmt.Println(resp.Data.File.Url) 
```

## 1.4. All Durations

|  Sites  |      URL      |  Package |
|----------     |:-------------:|------:|
| AnonFiles |  [AnonFiles](https://anonfiles.com) | [Uploader](https://github.com/ElecTwix/uploader)  |
| BayFiles |    [BayFiles](https://bayfiles.com/)  |   [Uploader](https://github.com/ElecTwix/uploader) |

## 1.5. Credit

[ElecTwix](https://github.com/ElecTwix)