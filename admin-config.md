# Admin Config

admin config 将定义admin 后台基础类型，  前端更具基础类型config 编写基础组件，  
通过基础组件实现渲染

### config 模块

不同页面 不同 admin config 进行排版

- 看板 大屏  (主要是 数据块) 
- 列表页面   (列表中什么什么的元素,  周围什么什么的组件)   
- 编辑页面   (文本编辑  数字编辑   等等)

``` 
type Page {
    id: String! # id
    name: String! # name
    pageType: PageType!
    pageConfig: PageConfig!
}

enum PageType {
    BigScreen # 大屏页面
    Manage # 管理页面
    Details # 详情页面
}

# PageConfig [ 大屏, 管理, 详情 ]
union  PageConfig = BigScreen | Manage | Details

# 大屏
type BigScreen {
    Items: [BigScreenItem!]!
}

# 大屏
type BigScreenItem {
    prefix: String! # 前缀
    suffix: String! # 后缀
    value: String!  # 值
}

# 管理页面
type Manage {
    line: [ConfigRow!]! # line 为有多少行
    readingMode: Boolean! # 预览模式
}

# 详情页
type Details {
    line: [ConfigRow!]! # line 为有多少行
    readingMode: Boolean! # 预览模式
}

# config row  每一行item
type ConfigRow {
    items: [ConfigItem!]!  # 每一行具体的
}

union ConfigItem = InputItem | TextButton | Action | Submit | Table | DataSelect

# 输入框
type InputItem {
    readOnly: Boolean!
    required: Boolean!
    key: String!
    name: String!
    type: InputType!
}

enum InputType {
    Text
    Number
    Bool
    File
}

# text
type TextButton {
    name: String!
    value: String!
    color: String!
}

# 跳转到新页面
type Action {
    id: String!
    name: String!
    pageType: PageType!
}

# 提交当前信息
type Submit {
    id: String!
    name: String!
    color: String!
}

# 时间选择
type DataSelect {
    readOnly: Boolean!
    required: Boolean!
    id: String!
    name: String!
}

# 枚举类型
type EnumInput {
    readOnly: Boolean!
    required: Boolean!
    id: String!
    name: String!
    enums: EnumItem!
}

type EnumItem {
    id: String!
    name: String!
}

# table
type Table {
    id: String!
    name: String!
    column: [TableColumn!]! # table header
}

type TableColumn {
    id: String!
    name: String!
}
```