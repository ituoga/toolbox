# datastar Hot reload

register route: 

```go
dsRouter.HandleFunc("/hotreload", hotreload.Handler)
```

and add this line to your templ file

```templ
@templ.Raw(hotreload.HTML)
```

or  just

```html
<div data-on-load="@get('/hotreload', {retryMaxCount: 1000,retryInterval:20, retryMaxWaitMs:200})"></div>
```