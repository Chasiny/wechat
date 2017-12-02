## 基于微信公众号的支付接口

### 实现的接口

> - 统一下单
> - 查询订单
> - 生成前端支付的参数

### 使用步骤

>- ```go
>  //配置公众号
>  	wp := &WePay.WePay{}
>  	wp.Init("填写公众号APPID","公众号对应的商户号id","商户API密钥","支付通知地址（后端接口）")
>  ```
>
>- ```go
>  //生成订单号
>  	outTradeNO:=WePay.CreateOutTradeNO()
>
>  	//获取前端支付需要的参数，接口参数二为支付金额，单位分
>  	res, err := wp.WebRequest(outTradeNO, 1, "127.0.0.1", 10, "用户openid")
>  	if (err != nil) {
>  		fmt.Println(err.Error())
>  		return
>  	}
>  	fmt.Println(res.Package)
>  ```
>
>- ```go
>  //查询订单
>  	ans,err:=wp.Query(outTradeNO)
>  	if (err != nil) {
>  		fmt.Println(err.Error())
>  		return
>  	}
>  ```