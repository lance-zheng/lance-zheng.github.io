<!-- customize-category:React -->

# Props

- [Props](#props)
  - [Props 验证](#props-验证)
  - [为 Props 指定默认值](#为-props-指定默认值)

## Props 验证

<https://www.freecodecamp.org/chinese/news/how-to-use-proptypes-in-react/>

## 为 Props 指定默认值

函数式组件可以利用 **ES6** 中的`解构赋值`定义默认值，也可以通过 `defaultProps` 指定默认值。

```jsx
// 方法一
const ListComment = ({ size = 10 }) => {
  return <div>{props.size}</div>;
};

// 方法二
const ListComment2 = (props) => {
  return <div>{props.size}</div>;
};
ListComment2.defaultProps = {
  size: 10,
};

export default function App() {
  return (
    <div className="App">
      <ListComment></ListComment>
    </div>
  );
}
```

类组件同样也可以通过定义静态属性 `defaultProps`

```jsx
import React from "react";

class ClassComponent extends React.Component {
  // 方法一
  static defaultProps = {
    size: 11,
  };

  render() {
    return <div>{this.props.size}</div>;
  }
}
// 方法二
ClassComponent.defaultProps = {
  size: 20,
};

export default function App() {
  return (
    <div className="App">
      <ClassComponent></ClassComponent>
    </div>
  );
}
```
