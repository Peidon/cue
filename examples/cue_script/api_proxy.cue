#Item: {
    id:int
    name:string
    abc?:string
    price:{...}
}

#Price:{
    val:number
}

// test data, origin response
elems : [...#Item] & [
    {
        id:1
        name:"shop ee"
        price: #Price & {
            val: 100.00
        }
    }
]

// 取列表第一个作为返回
// 列表为空时, 会返回null
response : #Item & elems[0]


// 结构扁平化
flat_response : {
    id : elems[0].id
    price_val: elems[0].price.val
    name : *elems[0].name | "defalt" // 可以加默认值
}